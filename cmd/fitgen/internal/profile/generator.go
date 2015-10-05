package profile

import (
	"bufio"
	"bytes"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"sort"
	"strings"
	"time"
)

type Generator struct {
	*bytes.Buffer
	sdkVersion string
	genTime    time.Time
}

func NewGenerator(sdkVersion string) *Generator {
	g := new(Generator)
	g.Buffer = new(bytes.Buffer)
	g.sdkVersion = sdkVersion
	g.genTime = time.Now()
	return g
}

func (g *Generator) GenerateTypes(types map[string]*GoType) ([]byte, error) {
	g.Reset()
	g.genHeader()
	g.genTypes(types)
	err := g.formatCode()
	if err != nil {
		return nil, err
	}
	return g.Bytes(), nil
}

func (g *Generator) GenerateMsgs(msgs []*GoMsg) ([]byte, error) {
	g.Reset()
	g.genHeader()
	g.genMsgs(msgs)
	err := g.formatCode()
	if err != nil {
		return nil, err
	}
	return g.Bytes(), nil
}

func (g *Generator) GenerateProfile(types map[string]*GoType, msgs []*GoMsg, jmptable bool) ([]byte, error) {
	g.Reset()
	g.genHeader()
	g.genProfile(types, msgs, jmptable)
	err := g.formatCode()
	if err != nil {
		return nil, err
	}
	return g.Bytes(), nil
}

// Large parts of this method is taken from
// github.com/golang/protobuf/proto-gen-go/generator/generator.go
// Copyright 2010 The Go Authors.
// https://raw.githubusercontent.com/golang/protobuf/master/LICENSE
func (g *Generator) formatCode() error {
	fset := token.NewFileSet()
	raw := g.Bytes()
	ast, err := parser.ParseFile(fset, "", g, parser.ParseComments)
	if err != nil {
		// Print out the bad code with line numbers.
		var src bytes.Buffer
		s := bufio.NewScanner(bytes.NewReader(raw))
		for line := 1; s.Scan(); line++ {
			fmt.Fprintf(&src, "%5d\t%s\n", line, s.Bytes())
		}
		return fmt.Errorf("bad Go source code was generated: %v\n%v", err, src.String())
	}
	g.Reset()
	err = (&printer.Config{Mode: printer.TabIndent | printer.UseSpaces, Tabwidth: 8}).Fprint(g, fset, ast)
	if err != nil {
		return fmt.Errorf("generated Go source code could not be reformatted: %v", err)
	}
	return nil
}

func (g *Generator) p(str ...interface{}) {
	for _, v := range str {
		switch s := v.(type) {
		case string:
			g.WriteString(s)
		case int:
			g.WriteString(fmt.Sprintf("%d", s))
		default:
			panic("unknown type in generator printer")
		}
	}
	g.WriteByte('\n')
}

func unexport(s string) string { return strings.ToLower(s[:1]) + s[1:] }

func (g *Generator) genHeader() {
	g.p("// This file is auto-generated using the")
	g.p("// program found in 'cmd/fitgen/main.go'")
	g.p("// DO NOT EDIT.")
	g.p("// SDK Version: ", g.sdkVersion)
	g.p("// Generation time: ", g.genTime.UTC().Format(time.UnixDate))
	g.p()
	g.p("package fit")
	g.p()
}

func (g *Generator) genTypes(types map[string]*GoType) {
	// sort for determinstic print order
	tkeys := make([]string, 0, len(types))
	for tkey := range types {
		tkeys = append(tkeys, tkey)
	}
	sort.Strings(tkeys)

	for _, tkey := range tkeys {
		t := types[tkey]
		g.p("// ", t.CamelCaseName, " represents the ", t.OrigName, " FIT type.")
		g.p("type ", t.CamelCaseName, " ", t.GoBaseType)
		g.p()
		g.p("const (")
		for _, v := range t.Values {
			if v.Comment == "" {
				g.p(t.CamelCaseName, v.Name, " ", t.CamelCaseName, " = ", v.Value)
			} else {
				g.p(t.CamelCaseName, v.Name, " ", t.CamelCaseName, " = ", v.Value, " // ", v.Comment)
			}
		}
		g.p(t.CamelCaseName, "Invalid ", t.CamelCaseName, " = ", t.InvalidValue)
		g.p(")")
	}
}

func (g *Generator) genMsgs(msgs []*GoMsg) {
	g.p("import \"time\"")
	g.p()
	for _, msg := range msgs {
		g.p("// ", msg.CamelCaseName, " represents the ", msg.Name, " FIT message type.")
		g.p("type ", msg.CamelCaseName, "Msg", " struct {")
		dynFieldsIdx := g.genFields(msg)
		if len(dynFieldsIdx) == 0 {
			continue
		}
		for _, fieldIdx := range dynFieldsIdx {
			g.genDynamicGetter(msg, fieldIdx)
		}
	}
}

func (g *Generator) genFields(msg *GoMsg) []int {
	var dynFieldsIdx []int
	for i, f := range msg.Fields {
		switch {
		case f.Comment == "" && f.Array == "0":
			g.p(f.CamelCaseName, " ", f.Type)
		case f.Comment != "" && f.Array == "0":
			g.p(f.CamelCaseName, " ", f.Type, " // ", f.Comment)
		case f.Comment == "" && f.Array != "0":
			g.p(f.CamelCaseName, " []", f.Type)
		case f.Comment != "" && f.Array != "0":
			g.p(f.CamelCaseName, " []", f.Type, " // ", f.Comment)
		default:
			panic("genFields: unreachable")
		}
		if len(f.DynFields) > 0 {
			dynFieldsIdx = append(dynFieldsIdx, i)
		}
	}
	g.p("}")
	g.p()
	return dynFieldsIdx
}

func (g *Generator) genDynamicGetter(msg *GoMsg, fieldIdx int) {
	field := msg.Fields[fieldIdx]

	refFieldNamesSet := make(map[string]bool)
	for _, dynf := range field.DynFields {
		for _, reffn := range dynf.RefFieldName {
			refFieldNamesSet[reffn] = true
		}
	}

	refFieldNameToType := make(map[string]string)
	for rfn := range refFieldNamesSet {
		for _, f := range msg.Fields {
			if f.CamelCaseName == rfn {
				refFieldNameToType[rfn] = f.Type
				break
			}
		}
		if refFieldNameToType[rfn] == "" {
			panic("genMsgs: could not find type for ref field name")
		}
	}

	nRefs := len(refFieldNameToType)

	g.p("func (", "x", " *", msg.CamelCaseName, "Msg) Get", field.CamelCaseName, "() interface{} {")

	if nRefs == 1 {
		var refField, refType string
		for rf, ty := range refFieldNameToType {
			refField = rf
			refType = ty
			break
		}

		g.p("switch ", "x", ".", refField, " {")
		for _, df := range field.DynFields {
			var dcase bytes.Buffer
			dcase.WriteString("case ")
			for i := range df.RefFieldName {
				dcase.WriteString(refType)
				dcase.WriteString(df.RefFieldValue[i])
				if i < len(df.RefFieldName)-1 {
					dcase.WriteByte(',')
				}
			}
			dcase.WriteString(":")
			g.p(dcase.String())
			if df.Type == "float64" {
				// We need to scale and apply offset
				g.p(
					"return ", df.Type, "(float32(x.",
					field.CamelCaseName, ") / ", df.Scale,
					" - float32(", df.Offset, "))",
				)
			} else {
				g.p("return ", df.Type, "(x.", field.CamelCaseName, ")")
			}
		}
	} else {
		// We can't switch on one message field...
		// Currently only 1 case in the SDK.
		// See field "target_value" in "workout_step" message.
		g.p("switch {")
		for _, df := range field.DynFields {
			for i, rfn := range df.RefFieldName {
				rtype, found := refFieldNameToType[rfn]
				if !found {
					panic("genMsgs: could not get type for ref field name")
				}
				g.p("case x.", rfn, " == ", rtype, df.RefFieldValue[i], ":")
				if df.Type == "float64" {
					g.p(
						"return ", df.Type, "(float32(x.",
						field.CamelCaseName, ") / ", df.Scale,
						" - float32(", df.Offset, "))",
					)
				} else {
					g.p("return ", df.Type, "(x.", field.CamelCaseName, ")")
				}
			}
		}
	}
	g.p("default:")
	g.p("return ", "x", ".", field.CamelCaseName)
	g.p("}")
	g.p("}")
	g.p()
}

func (g *Generator) genProfile(types map[string]*GoType, msgs []*GoMsg, jmptable bool) {
	g.p("import \"reflect\"")
	g.genKnownMsgs(types)

	if jmptable {
		g.genFieldsArray(msgs)
		g.genGetFieldArrayLookup(msgs)
		g.genZeroValueMsgsArray(msgs)
		g.genGetZeroValueMsgsArrayLookup(msgs)
		return
	}

	g.genFieldsVarsAndMap(msgs)
	g.genGetFieldSwitch(msgs)
	g.genZeroValueMsgsVars(msgs)
	g.genGetZeroValueMsgsSwitch(msgs)
}

func (g *Generator) genKnownMsgs(types map[string]*GoType) {
	mesgNums, found := types["MesgNum"]
	if !found {
		panic("genKnownMsgs: can't find MesgNum type")
	}
	mnvals := mesgNums.Values
	g.p()
	g.p("var knownMsgNums = map[MesgNum]bool{")
	for i := 0; i < len(mnvals)-2; i++ { // -2: Skip the last two: RangeMin/Max
		if knownMesgNumButNoMsg[mnvals[i].Name] {
			continue
		}
		g.p("MesgNum", mnvals[i].Name, ": true,")
	}
	g.p("}")
}

func (g *Generator) genFieldsVarsAndMap(msgs []*GoMsg) {
	g.p()
	for _, msg := range msgs {
		g.p("var ", unexport(msg.CamelCaseName), "Fields = map[byte]*field{")
		for i := 0; i < len(msg.Fields); i++ {
			f := msg.Fields[i]
			g.p(f.DefNum, ": {", i, ", ", f.Scale, ", ", f.Offset, ", ", f.Array, ", ", f.GoType, ", ", f.DefNum, ", ", f.BaseType, "},")
		}
		g.p("}")
		g.p()
	}
}

func (g *Generator) genFieldsArray(msgs []*GoMsg) {
	g.p()
	g.p("// Set length to 256, so that lookup for any")
	g.p("// field 255 (localMesgNumInvalid) will return nil.")
	g.p("var _fields = [...][256]*field{")
	for _, msg := range msgs {
		g.p("MesgNum", msg.CamelCaseName, ": {")
		for i := 0; i < len(msg.Fields); i++ {
			f := msg.Fields[i]
			g.p(f.DefNum, ": {", i, ", ", f.Scale, ", ", f.Offset, ", ", f.Array, ", ", f.GoType, ", ", f.DefNum, ", ", f.BaseType, "},")
		}
		g.p("},")
		g.p()
	}
	g.p("}")
}

func (g *Generator) genGetFieldSwitch(msgs []*GoMsg) {
	g.p()
	g.p("func getField(gmn MesgNum, fdn byte) (*field, bool) {")
	g.p("var f *field")
	g.p("var ok bool")
	g.p("switch gmn {")
	for _, msg := range msgs {
		g.p("case ", "MesgNum", msg.CamelCaseName, ":")
		g.p("f, ok = ", unexport(msg.CamelCaseName), "Fields[fdn]")
		g.p("if !ok {")
		g.p("return nil, false")
		g.p("}")
	}
	g.p("default:")
	g.p("panic(\"getField: global message not found\")")
	g.p("}")
	g.p("return f, true")
	g.p("}")
}

func (g *Generator) genGetFieldArrayLookup(msgs []*GoMsg) {
	g.p()
	g.p("func getField(gmn MesgNum, fdn byte) (*field, bool) {")
	g.p("if int(gmn) >= len(_fields) {")
	g.p("return nil, false")
	g.p("}")
	g.p("f := _fields[gmn][fdn]")
	g.p("if f == nil {")
	g.p("return nil, false")
	g.p("}")
	g.p("return f, true")
	g.p("}")
}

func (g *Generator) genZeroValueMsgsVars(msgs []*GoMsg) {
	g.p()
	g.p("var (")
	for _, msg := range msgs {
		g.p(unexport(msg.CamelCaseName), "AI = reflect.ValueOf(&", msg.CamelCaseName, "Msg{")
		for _, f := range msg.Fields {
			g.p(f.GoInvalid, ",")
		}
		g.p("})")
		g.p()
	}
	g.p(")")
}

func (g *Generator) genZeroValueMsgsArray(msgs []*GoMsg) {
	g.p()
	g.p("var msgsAllInvalid = [...]reflect.Value{")
	for _, msg := range msgs {

		g.p("MesgNum", msg.CamelCaseName, ": reflect.ValueOf(&", msg.CamelCaseName, "Msg{")
		for _, f := range msg.Fields {
			g.p(f.GoInvalid, ",")
		}
		g.p("}),")
	}
	g.p("}")
}

func (g *Generator) genGetZeroValueMsgsSwitch(msgs []*GoMsg) {
	g.p()
	g.p("func getMesgAllInvalid(mn MesgNum) reflect.Value {")
	g.p("switch mn {")
	for _, msg := range msgs {
		g.p("case MesgNum", msg.CamelCaseName, ":")
		g.p("return reflect.ValueOf(", unexport(msg.CamelCaseName), "AI.Interface()).Elem()")
	}
	g.p("default:")
	g.p("panic(\"getMesgAllInvalid: unknown message number\")")
	g.p("}")
	g.p("}")
}

func (g *Generator) genGetZeroValueMsgsArrayLookup(msgs []*GoMsg) {
	g.p()
	g.p("func getMesgAllInvalid(mn MesgNum) reflect.Value {")
	g.p("return reflect.ValueOf(msgsAllInvalid[mn].Interface()).Elem()")
	g.p("}")
}
