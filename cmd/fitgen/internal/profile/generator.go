package profile

import (
	"bufio"
	"bytes"
	"fmt"
	"go/parser"
	"go/printer"
	"go/token"
	"log"
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

func (g *Generator) GenerateTypes(types map[string]*Type) ([]byte, error) {
	g.Reset()
	g.genHeader()
	g.genTypes(types)
	err := g.formatCode()
	if err != nil {
		return nil, err
	}
	return g.Bytes(), nil
}

func (g *Generator) GenerateMsgs(msgs []*Msg) ([]byte, error) {
	g.Reset()
	g.genHeader()
	g.genMsgs(msgs)
	err := g.formatCode()
	if err != nil {
		return nil, err
	}
	return g.Bytes(), nil
}

func (g *Generator) GenerateProfile(types map[string]*Type, msgs []*Msg, jmptable bool) ([]byte, error) {
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
}

func (g *Generator) genTypes(types map[string]*Type) {
	// sort for determinstic print order
	tkeys := make([]string, 0, len(types))
	for tkey := range types {
		tkeys = append(tkeys, tkey)
	}
	sort.Strings(tkeys)

	g.p()
	for _, tkey := range tkeys {
		t := types[tkey]
		g.p("// ", t.CCName, " represents the ", t.OrigName, " FIT type.")
		g.p("type ", t.CCName, " ", t.GoBaseType)
		g.p()
		g.p("const (")
		for _, v := range t.Values {
			if v.Comment == "" {
				g.p(t.CCName, v.Name, " ", t.CCName, " = ", v.Value)
			} else {
				g.p(t.CCName, v.Name, " ", t.CCName, " = ", v.Value, " // ", v.Comment)
			}
		}
		g.p(t.CCName, "Invalid ", t.CCName, " = ", t.InvalidValue)
		g.p(")")
	}
}

func (g *Generator) genMsgs(msgs []*Msg) {
	g.p("import (")
	g.p("\"math\"")
	g.p("\"time\"")
	g.p(")")
	for _, msg := range msgs {
		g.p()
		g.p("// ", msg.CCName, "Msg represents the ", msg.Name, " FIT message type.")
		g.p("type ", msg.CCName, "Msg", " struct {")
		scaledfs, dynfs, compfs, dyncompfs := g.genFields(msg)
		for _, scaledfi := range scaledfs {
			g.genScaledGetter(msg, scaledfi)
		}
		for _, dynfi := range dynfs {
			g.genDynamicGetter(msg, dynfi)
		}
		if len(compfs) > 0 {
			g.genComponentsRelated(msg, compfs, dyncompfs)
		}
	}
}

func (g *Generator) genFields(msg *Msg) (scaledfi, dynfi, compfi []int, dyncompfi map[int][]int) {
	dyncompfi = make(map[int][]int)
	g.p()
	for i, f := range msg.Fields {
		switch {
		case f.Comment == "" && f.Array == "0":
			g.p(f.CCName, " ", f.Type)
		case f.Comment != "" && f.Array == "0":
			g.p(f.CCName, " ", f.Type, " // ", f.Comment)
		case f.Comment == "" && f.Array != "0":
			g.p(f.CCName, " []", f.Type)
		case f.Comment != "" && f.Array != "0":
			g.p(f.CCName, " []", f.Type, " // ", f.Comment)
		default:
			panic("genFields: unreachable")
		}
		if len(f.Scale) > 0 {
			scaledfi = append(scaledfi, i)
		}
		if len(f.Subfields) > 0 {
			dynfi = append(dynfi, i)
			var compsubi []int
			for j, sf := range f.Subfields {
				if len(sf.Components) == 0 {
					continue
				}
				compsubi = append(compsubi, j)
			}
			if len(compsubi) > 0 {
				dyncompfi[i] = compsubi
			}
		}
		if len(f.Components) > 0 {
			compfi = append(compfi, i)
		}
	}
	g.p("}")
	return
}

func (g *Generator) genScaledGetter(msg *Msg, fieldIndex int) {
	f := msg.Fields[fieldIndex]
	g.p()
	if f.Array == "0" {
		g.genScaledGetterReg(msg, f)
	} else {
		g.genScaledGetterArray(msg, f)
	}
	g.p("}")
}

func (g *Generator) genScaledGetterReg(msg *Msg, f *Field) {
	g.p("// Get", f.CCName, "Scaled returns ", f.CCName)
	g.p("// with scale and any offset applied. NaN is returned if the")
	g.p("// field has an invalid value (i.e. has not been set).")
	if f.Units != "" {
		g.p("// Units: ", f.Units)
	}

	g.p("func (x *", msg.CCName, "Msg) Get", f.CCName, "Scaled() float64 {")

	g.p("if x.", f.CCName, " == ", f.GoInvalid, " {")
	g.p("return math.NaN()")
	g.p("}")

	var out bytes.Buffer
	out.WriteString("return float64(x.")
	out.WriteString(f.CCName)
	out.WriteString(") / ")
	out.WriteString(f.Scale)
	if f.Offset != "" {
		out.WriteString(" - ")
		out.WriteString(f.Offset)
	}
	g.p(out.String())
}

func (g *Generator) genScaledGetterArray(msg *Msg, f *Field) {
	g.p("// Get", f.CCName, "Scaled returns ", f.CCName)
	g.p("// as a slice with scale and any offset applied to every element.")
	if f.Units != "" {
		g.p("// Units: ", f.Units)
	}

	g.p("func (x *", msg.CCName, "Msg) Get", f.CCName, "Scaled() []float64 {")
	g.p("if len(x.", f.CCName, ") == 0 {")
	g.p("return nil")
	g.p("}")
	g.p("s := make([]float64, len(x.", f.CCName, "))")
	g.p("for i, v := range x.", f.CCName, " {")
	if f.Offset == "0" || f.Offset == "" {
		g.p("s[i] = float64(v) / ", f.Scale)
	} else {
		g.p("s[i] = float64(v) / ", f.Scale, " - ", f.Offset)
	}
	g.p("}")
	g.p("return s")
}

func (g *Generator) genDynamicGetter(msg *Msg, fieldIndex int) {
	field := msg.Fields[fieldIndex]

	refFieldNamesSet := make(map[string]bool)
	for _, subf := range field.Subfields {
		for _, reffn := range subf.RefFieldName {
			refFieldNamesSet[reffn] = true
		}
	}

	refFieldNameToType := make(map[string]string)
	for rfn := range refFieldNamesSet {
		for _, f := range msg.Fields {
			if f.CCName == rfn {
				refFieldNameToType[rfn] = f.Type
				break
			}
		}
		if refFieldNameToType[rfn] == "" {
			panic("genMsgs: could not find type for ref field name")
		}
	}

	nrefs := len(refFieldNameToType)

	g.p()
	g.p("// Get", field.CCName, " returns the appropriate ", field.CCName)
	g.p("// subfield if a matching reference field/value combination is found.")
	g.p("// If none of the reference field/value combinations are true")
	g.p("// then the main field is returned.")
	g.p("func (", "x", " *", msg.CCName, "Msg) Get", field.CCName, "() interface{} {")

	if nrefs == 1 {
		var refField, refType string
		for rf, ty := range refFieldNameToType {
			refField = rf
			refType = ty
			break
		}

		g.p("switch ", "x", ".", refField, " {")
		for _, sf := range field.Subfields {
			var scase bytes.Buffer
			scase.WriteString("case ")
			for i := range sf.RefFieldName {
				scase.WriteString(refType)
				scase.WriteString(sf.RefFieldValue[i])
				if i < len(sf.RefFieldName)-1 {
					scase.WriteByte(',')
				}
			}
			scase.WriteString(":")
			g.p(scase.String())
			if sf.Scale != "" {
				// We need to scale and apply offset.
				var out bytes.Buffer
				out.WriteString("return float64(x.")
				out.WriteString(field.CCName)
				out.WriteString(") / ")
				out.WriteString(sf.Scale)
				if sf.Offset != "" {
					out.WriteString(" - ")
					out.WriteString(sf.Offset)
				}
				g.p(out.String())
			} else {
				g.p("return ", sf.Type, "(x.", field.CCName, ")")
			}
		}
	} else {
		// We can't switch on one message field. Currently only 1 case
		// in the SDK. See field "target_value" in "workout_step"
		// message.
		g.p("switch {")
		for _, sf := range field.Subfields {
			for i, rfn := range sf.RefFieldName {
				rtype, found := refFieldNameToType[rfn]
				if !found {
					panic("genMsgs: could not get type for ref field name")
				}
				g.p("case x.", rfn, " == ", rtype, sf.RefFieldValue[i], ":")
				if sf.Scale != "" {
					var out bytes.Buffer
					out.WriteString("return float64(x.")
					out.WriteString(field.CCName)
					out.WriteString(") / ")
					out.WriteString(sf.Scale)
					if sf.Offset != "" {
						out.WriteString(" - ")
						out.WriteString(sf.Offset)
					}
					g.p(out.String())
				} else {
					g.p("return ", sf.Type, "(x.", field.CCName, ")")
				}
			}
		}
	}
	g.p("default:")
	g.p("return ", "x", ".", field.CCName)
	g.p("}")
	g.p("}")
}

func (g *Generator) genComponentsRelated(msg *Msg, compFieldIndices []int, dynCompFieldIndices map[int][]int) {
	g.genGetterForComponents(msg, compFieldIndices)
	g.genExpandComponents(msg, compFieldIndices, dynCompFieldIndices)
}

func (g *Generator) genGetterForComponents(msg *Msg, compFieldIndices []int) {
	// Add getter for target field if scale for component differs. Only
	// relevant if # components > 1.

	// TODO(tormoder): Verify if this is correct.

	// TODO(tormoder): Add getter for subfields. There are no such cases in
	// the SDK now, expect 'gear_change_data' in the 'event' message, but
	// scale there is '1' for every component and non for target fields,
	// meaning getters are not neccecary.

	for _, i := range compFieldIndices {
		f := msg.Fields[i]
		for _, comp := range f.Components {
			targetf, found := msg.FieldByName[comp.Name]
			if !found {
				panic("target field for component not found")
			}
			if comp.Scale == "" || comp.Scale == "1" {
				continue
			}
			if comp.Scale != targetf.Scale {
				g.p()
				g.p("// Get", targetf.CCName, "From", f.CCName, " returns ")
				g.p("// ", targetf.CCName, " with the scale and offset defined by the \"", comp.Name, "\"")
				g.p("// component in the ", f.CCName, " field. NaN is")
				g.p("// if the field has an invalid value (i.e. has not been set).")

				g.p("func (x *", msg.CCName, "Msg) Get", targetf.CCName, "From", f.CCName, "() float64 {")

				g.p("if x.", targetf.CCName, " == ", targetf.GoInvalid, " {")
				g.p("return math.NaN()")
				g.p("}")

				var out bytes.Buffer
				out.WriteString("return float64(x.")
				out.WriteString(targetf.CCName)
				out.WriteString(") / ")
				out.WriteString(comp.Scale)
				if comp.Offset != "" {
					out.WriteString(" - ")
					out.WriteString(comp.Offset)
				}
				g.p(out.String())
				g.p("}")
			}
		}
	}
}

func (g *Generator) genExpandComponents(msg *Msg, compFieldIndices []int, dynCompFieldIndices map[int][]int) {
	if len(compFieldIndices) == 0 && len(dynCompFieldIndices) == 0 {
		return
	}

	log.Println("message should call expandComponents() on add in file_types.go:", msg.CCName)

	g.p()
	g.p("func (", "x", " *", msg.CCName, "Msg) expandComponents() {")

	for _, cfi := range compFieldIndices {
		field := msg.Fields[cfi]
		switch field.BaseType {
		case "fitByte", "fitUint8", "fitUint16", "fitUint32":
		default:
			panic("genExpandComponents: unhandled base type")
		}

		debugln("expand components: msg:", msg.CCName, "- field:", field.CCName)

		if field.Array == "0" {
			g.genExpandComponentsReg(msg, field)
		} else {
			g.genExpandComponentsArray(field)
		}
	}

	for dcfi, dcsfis := range dynCompFieldIndices {
		field := msg.Fields[dcfi]
		switch field.BaseType {
		case "fitUint8", "fitUint16", "fitUint32":
		case "fitByte":
			panic("genExpandComponentsDyn: unhandled base type when array")
		default:
			panic("genExpandComponentsDyn: unhandled base type")
		}
		g.genExpandComponentsDyn(msg, field, dcsfis)
	}

	g.p("}")
}

func (g *Generator) genExpandComponentsReg(msg *Msg, field *Field) {
	g.p("if x.", field.CCName, " != ", field.GoInvalid, " {")
	g.genExpandComponentsMaskShift(msg, field)
	g.p("}")
}

func (g *Generator) genExpandComponentsArray(field *Field) {
	// Handle this as a special case for now.
	if field.CCName != "CompressedSpeedDistance" {
		panic("genExpandComponents: only specific field " +
			"'CompressedSpeedDistance' in Record message" +
			"currently handled for byte arrays")
	}
	g.p("expand := false")
	g.p("if len(x.", field.CCName, ") == 3 {")
	g.p("for _, v := range x.", field.CCName, " {")
	g.p("if v != ", field.BTInvalid, "{")
	g.p("expand = true")
	g.p("break")
	g.p("}")
	g.p("}")
	g.p("}")
	g.p("if expand {")
	g.p("x.Speed = uint16(x.", field.CCName, "[0] | ((x.", field.CCName, "[1]", "& 0x0F) << 8))")
	g.p("x.Distance = uint32((x.", field.CCName, "[1] >> 4) | (x.", field.CCName, "[2] << 4))")
	g.p("}")
}

func (g *Generator) genExpandComponentsDyn(msg *Msg, field *Field, dcsfis []int) {
	refFieldNamesSet := make(map[string]bool)
	for _, subfi := range dcsfis {
		subf := field.Subfields[subfi]
		debugln("expand components: msg:", msg.CCName, "- field:", field.CCName, "- subfield:", subf.CCName)
		for _, reffn := range subf.RefFieldName {
			refFieldNamesSet[reffn] = true
		}
	}

	refFieldNameToType := make(map[string]string)
	for rfn := range refFieldNamesSet {
		for _, f := range msg.Fields {
			if f.CCName == rfn {
				refFieldNameToType[rfn] = f.Type
				break
			}
		}
		if refFieldNameToType[rfn] == "" {
			panic("genExpandComponentsDyn: could not find type for ref field name")
		}
	}

	if len(refFieldNameToType) > 1 {
		panic("genExpandComponentsDyn: unhandled case, more than one reference field name")
	}

	g.p("if x.", field.CCName, " != ", field.GoInvalid, " {")

	var refField, refType string
	for rf, ty := range refFieldNameToType {
		refField = rf
		refType = ty
		break
	}

	g.p("switch ", "x", ".", refField, " {")
	for _, subfi := range dcsfis {
		sf := field.Subfields[subfi]
		var scase bytes.Buffer
		scase.WriteString("case ")
		for i := range sf.RefFieldName {
			scase.WriteString(refType)
			scase.WriteString(sf.RefFieldValue[i])
			if i < len(sf.RefFieldName)-1 {
				scase.WriteByte(',')
			}
		}
		scase.WriteString(":")
		g.p(scase.String())
		g.genExpandComponentsMaskShiftDyn(msg, sf, field)
	}
	g.p("}")
	g.p("}")

}

func (g *Generator) genExpandComponentsMaskShift(msg *Msg, field *Field) {
	bits := 0
	for _, comp := range field.Components {
		tfield, tfound := msg.FieldByName[comp.Name]
		if !tfound {
			panic("genExpandComponentsMaskShift: target field not found")
		}
		g.p("x.", comp.Name, " = ", tfield.Type, "((x.", field.CCName, " >> ", bits, ") & ((1 << ", comp.Bits, ") - 1))")
		bits += comp.BitsInt
	}
}

func (g *Generator) genExpandComponentsMaskShiftDyn(msg *Msg, sfield *Field, mfield *Field) {
	bits := 0
	for _, comp := range sfield.Components {
		tfield, tfound := msg.FieldByName[comp.Name]
		if !tfound {
			panic("genExpandComponentsMaskShift: target field not found")
		}
		g.p("x.", comp.Name, " = ", tfield.Type, "((x.", mfield.CCName, " >> ", bits, ") & ((1 << ", comp.Bits, ") - 1))")
		bits += comp.BitsInt
	}
}

func (g *Generator) genProfile(types map[string]*Type, msgs []*Msg, jmptable bool) {
	g.p("import (")
	g.p("\"fmt\"")
	g.p("\"reflect\"")
	g.p(")")

	g.genFieldDef()
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

func (g *Generator) genFieldDef() {
	g.p()
	g.p("// field represents a fit message field in the profile field lookup table.")
	g.p("type field struct {")
	g.p("sindex int")
	g.p("array  uint8")
	g.p("t      gotype")
	g.p("num    byte")
	g.p("btype  fitBaseType")
	g.p("}")
	g.p()
	g.p("func (f field) String() string {")
	g.p("return fmt.Sprintf(")
	g.p("\"num: %d | btype: %v | sindex: %d | array: %d\",")
	g.p("f.num, f.btype, f.sindex, f.array,")
	g.p(")")
	g.p("}")
	g.p()
	g.p("// gotype is used in the profile field lookup table to represent the data type")
	g.p("// (or type category) for a field when decoded into a Go message struct.")
	g.p("type gotype uint8")
	g.p()
	g.p("const (")
	g.p("fit gotype = iota // Standard -> Fit base type/alias")
	g.p()
	g.p("// Special (non-profile types)")
	g.p("	timeutc   // Time UTC 	-> time.Time")
	g.p("	timelocal // Time Local -> time.Time with Location")
	g.p("	lat       // Latitude 	-> fit.Latitude")
	g.p("	lng       // Longitude 	-> fit.Longitude")
	g.p(")")
	g.p()
	g.p("func (g gotype) String() string {")
	g.p("	if int(g) > len(gotypeString) {")
	g.p("		return fmt.Sprintf(\"gotype(%d)\", g)")
	g.p("	}")
	g.p("	return gotypeString[g]")
	g.p("}")
	g.p()
	g.p("var gotypeString = [...]string{")
	g.p("	\"fit\",")
	g.p("	\"timeutc\",")
	g.p("	\"timelocal\",")
	g.p("	\"lat\",")
	g.p("	\"lng\",")
	g.p("}")
}

func (g *Generator) genKnownMsgs(types map[string]*Type) {
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

func (g *Generator) genFieldsVarsAndMap(msgs []*Msg) {
	g.p()
	for _, msg := range msgs {
		g.p("var ", unexport(msg.CCName), "Fields = map[byte]*field{")
		for i := 0; i < len(msg.Fields); i++ {
			f := msg.Fields[i]
			g.p(f.DefNum, ": {", i, ", ", f.Array, ", ", f.GoType, ", ", f.DefNum, ", ", f.BaseType, "},")
		}
		g.p("}")
		g.p()
	}
}

func (g *Generator) genFieldsArray(msgs []*Msg) {
	g.p()
	g.p("// Set length to 256, so that lookup for any")
	g.p("// field 255 (localMesgNumInvalid) will return nil.")
	g.p("var _fields = [...][256]*field{")
	for _, msg := range msgs {
		g.p("MesgNum", msg.CCName, ": {")
		for i := 0; i < len(msg.Fields); i++ {
			f := msg.Fields[i]
			g.p(f.DefNum, ": {", i, ", ", f.Array, ", ", f.GoType, ", ", f.DefNum, ", ", f.BaseType, "},")
		}
		g.p("},")
		g.p()
	}
	g.p("}")
}

func (g *Generator) genGetFieldSwitch(msgs []*Msg) {
	g.p()
	g.p("func getField(gmn MesgNum, fdn byte) (*field, bool) {")
	g.p("var f *field")
	g.p("var ok bool")
	g.p("switch gmn {")
	for _, msg := range msgs {
		g.p("case ", "MesgNum", msg.CCName, ":")
		g.p("f, ok = ", unexport(msg.CCName), "Fields[fdn]")
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

func (g *Generator) genGetFieldArrayLookup(msgs []*Msg) {
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

func (g *Generator) genZeroValueMsgsVars(msgs []*Msg) {
	g.p()
	g.p("var (")
	for _, msg := range msgs {
		g.p(unexport(msg.CCName), "AI = reflect.ValueOf(&", msg.CCName, "Msg{")
		for _, f := range msg.Fields {
			g.p(f.GoInvalid, ",")
		}
		g.p("})")
		g.p()
	}
	g.p(")")
}

func (g *Generator) genZeroValueMsgsArray(msgs []*Msg) {
	g.p()
	g.p("var msgsAllInvalid = [...]reflect.Value{")
	for _, msg := range msgs {

		g.p("MesgNum", msg.CCName, ": reflect.ValueOf(&", msg.CCName, "Msg{")
		for _, f := range msg.Fields {
			g.p(f.GoInvalid, ",")
		}
		g.p("}),")
	}
	g.p("}")
}

func (g *Generator) genGetZeroValueMsgsSwitch(msgs []*Msg) {
	g.p()
	g.p("func getMesgAllInvalid(mn MesgNum) reflect.Value {")
	g.p("switch mn {")
	for _, msg := range msgs {
		g.p("case MesgNum", msg.CCName, ":")
		g.p("return reflect.ValueOf(", unexport(msg.CCName), "AI.Interface()).Elem()")
	}
	g.p("default:")
	g.p("panic(\"getMesgAllInvalid: unknown message number\")")
	g.p("}")
	g.p("}")
}

func (g *Generator) genGetZeroValueMsgsArrayLookup(msgs []*Msg) {
	g.p()
	g.p("func getMesgAllInvalid(mn MesgNum) reflect.Value {")
	g.p("return reflect.ValueOf(msgsAllInvalid[mn].Interface()).Elem()")
	g.p("}")
}
