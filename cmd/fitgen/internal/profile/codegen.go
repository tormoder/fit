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
	"time"

	"github.com/tormoder/fit/internal/types"
)

type codeGenerator struct {
	*bytes.Buffer
	sdkFullVer           string
	sdkMajVer, sdkMinVer int
	addGenTime           bool
	genTime              time.Time
	logger               *log.Logger
}

func newCodeGenerator(sdkMajVer, sdkMinVer int, addGenerationTime bool, logger *log.Logger) *codeGenerator {
	g := new(codeGenerator)
	g.sdkMajVer = sdkMajVer
	g.sdkMinVer = sdkMinVer
	g.sdkFullVer = fmt.Sprintf("%d.%d", sdkMajVer, sdkMinVer)
	g.addGenTime = addGenerationTime
	g.genTime = time.Now()
	g.logger = logger
	return g
}

func (g *codeGenerator) generateTypes(types map[string]*Type) ([]byte, error) {
	g.Buffer = new(bytes.Buffer)
	g.genHeader()
	g.genTypes(types)
	err := g.formatCode()
	if err != nil {
		return nil, err
	}
	return g.Bytes(), nil
}

func (g *codeGenerator) generateMsgs(msgs []*Msg) ([]byte, error) {
	g.Buffer = new(bytes.Buffer)
	g.genHeader()
	g.genMsgs(msgs)
	err := g.formatCode()
	if err != nil {
		return nil, err
	}
	return g.Bytes(), nil
}

func (g *codeGenerator) generateProfile(types map[string]*Type, msgs []*Msg) ([]byte, error) {
	g.Buffer = new(bytes.Buffer)
	g.genHeader()
	g.genProfile(types, msgs)
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
func (g *codeGenerator) formatCode() error {
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

func (g *codeGenerator) p(str ...interface{}) {
	for _, v := range str {
		switch s := v.(type) {
		case string:
			g.WriteString(s)
		case int:
			g.WriteString(fmt.Sprintf("%d", s))
		default:
			err := fmt.Sprintf("unknown type in generator printer: %T", s)
			panic(err)
		}
	}
	g.WriteByte('\n')
}

func (g *codeGenerator) genHeader() {
	g.p("// Code generated using the program found in 'cmd/fitgen/main.go'. DO NOT EDIT.")
	g.p()
	g.p("// SDK Version: ", g.sdkFullVer)
	if g.addGenTime {
		g.p("// Generation time: ", g.genTime.UTC().Format(time.UnixDate))
	}
	g.p()
	g.p("package fit")
}

func (g *codeGenerator) genTypes(types map[string]*Type) {
	// Sort for determinstic print order.
	tkeys := make([]string, 0, len(types))
	for tkey := range types {
		tkeys = append(tkeys, tkey)
	}
	sort.Strings(tkeys)

	g.p()
	for _, tkey := range tkeys {
		t := types[tkey]
		g.p("// ", t.Name, " represents the ", t.OrigName, " FIT type.")
		g.p("type ", t.Name, " ", t.BaseType.GoType())
		g.p()
		g.p("const (")
		for _, v := range t.Values {
			if v.Comment == "" {
				g.p(t.Name, v.Name, " ", t.Name, " = ", v.Value)
			} else {
				g.p(t.Name, v.Name, " ", t.Name, " = ", v.Value, " // ", v.Comment)
			}
		}
		g.p(t.Name, "Invalid ", t.Name, " = ", t.BaseType.GoInvalidValue())
		g.p(")")
	}
}

func (g *codeGenerator) genMsgs(msgs []*Msg) {
	g.p("import (")
	g.p("\"math\"")
	g.p("\"time\"")
	g.p(")")
	for _, msg := range msgs {
		g.p()
		g.p("// ", msg.CCName, "Msg represents the ", msg.Name, " FIT message type.")
		g.p("type ", msg.CCName, "Msg", " struct {")
		scaledfs, dynfs, compfs, dyncompfs := g.genFields(msg)
		g.p("}")
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

func (g *codeGenerator) genFields(msg *Msg) (scaledfi, dynfi, compfi []int, dyncompfi map[int][]int) {
	dyncompfi = make(map[int][]int)
	for i, f := range msg.Fields {
		if !f.HasComment() {
			g.p(f.CCName, " ", f.TypeName)
		} else {
			g.p(f.CCName, " ", f.TypeName, " // ", f.Comment)
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
	return
}

func (g *codeGenerator) genScaledGetter(msg *Msg, fieldIndex int) {
	f := msg.Fields[fieldIndex]
	g.p()
	if !f.FType.Array() {
		g.genScaledGetterReg(msg, f)
	} else {
		g.genScaledGetterArray(msg, f)
	}
	g.p("}")
}

func (g *codeGenerator) genScaledGetterReg(msg *Msg, f *Field) {
	g.p("// Get", f.CCName, "Scaled returns ", f.CCName)
	g.p("// with scale and any offset applied. NaN is returned if the")
	g.p("// field has an invalid value (i.e. has not been set).")
	if f.Units != "" {
		g.p("// Units: ", f.Units)
	}

	g.p("func (x *", msg.CCName, "Msg) Get", f.CCName, "Scaled() float64 {")

	g.p("if x.", f.CCName, " == ", f.FType.GoInvalidValue(), " {")
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

func (g *codeGenerator) genScaledGetterArray(msg *Msg, f *Field) {
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
	if !f.HasOffset() {
		g.p("s[i] = float64(v) / ", f.Scale)
	} else {
		g.p("s[i] = float64(v) / ", f.Scale, " - ", f.Offset)
	}
	g.p("}")
	g.p("return s")
}

func (g *codeGenerator) genDynamicGetter(msg *Msg, fieldIndex int) {
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
				refFieldNameToType[rfn] = f.TypeName
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
				g.p("return ", sf.TypeName, "(x.", field.CCName, ")")
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
					g.p("return ", sf.TypeName, "(x.", field.CCName, ")")
				}
			}
		}
	}
	g.p("default:")
	g.p("return ", "x", ".", field.CCName)
	g.p("}")
	g.p("}")
}

func (g *codeGenerator) genComponentsRelated(msg *Msg, compFieldIndices []int, dynCompFieldIndices map[int][]int) {
	g.genGetterForComponents(msg, compFieldIndices)
	g.genExpandComponents(msg, compFieldIndices, dynCompFieldIndices)
}

func (g *codeGenerator) genGetterForComponents(msg *Msg, compFieldIndices []int) {
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
				panic("genGetterForComponents: target field for component not found")
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

				g.p("if x.", targetf.CCName, " == ", targetf.FType.GoInvalidValue(), " {")
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

func (g *codeGenerator) genExpandComponents(msg *Msg, compFieldIndices []int, dynCompFieldIndices map[int][]int) {
	if len(compFieldIndices) == 0 && len(dynCompFieldIndices) == 0 {
		return
	}

	g.logger.Println("msggen:", msg.CCName, "should call expandComponents() on add in file_types.go")

	g.p()
	g.p("func (", "x", " *", msg.CCName, "Msg) expandComponents() {")

	for _, cfi := range compFieldIndices {
		field := msg.Fields[cfi]
		switch field.FType.BaseType() {
		case types.BaseByte, types.BaseUint8, types.BaseUint16, types.BaseUint32:
		default:
			panic("genExpandComponents: unhandled base type")
		}

		g.logger.Println("expand components: msg:", msg.CCName, "- field:", field.CCName)

		if !field.FType.Array() {
			g.genExpandComponentsReg(msg, field)
		} else {
			g.genExpandComponentsArray(field)
		}
	}

	for dcfi, dcsfis := range dynCompFieldIndices {
		field := msg.Fields[dcfi]
		switch field.FType.BaseType() {
		case types.BaseUint8, types.BaseUint16, types.BaseUint32:
		case types.BaseByte:
			panic("genExpandComponentsDyn: unhandled base type when array")
		default:
			panic("genExpandComponentsDyn: unhandled base type")
		}
		g.genExpandComponentsDyn(msg, field, dcsfis)
	}

	g.p("}")
}

func (g *codeGenerator) genExpandComponentsReg(msg *Msg, field *Field) {
	g.p("if x.", field.CCName, " != ", field.FType.GoInvalidValue(), " {")
	g.genExpandComponentsMaskShift(msg, field)
	g.p("}")
}

func (g *codeGenerator) genExpandComponentsArray(field *Field) {
	// Handle every byte array manually for now.
	// One case in the SDK per version 16.10.
	switch field.CCName {
	case "CompressedSpeedDistance":
		g.p("expand := false")
		g.p("if len(x.", field.CCName, ") == 3 {")
		g.p("for _, v := range x.", field.CCName, " {")
		g.p("if v != ", field.FType.BaseType().GoInvalidValue(), "{")
		g.p("expand = true")
		g.p("break")
		g.p("}")
		g.p("}")
		g.p("}")
		g.p("if expand {")
		g.p("x.Speed = uint16(x.", field.CCName, "[0]) | uint16(x.", field.CCName, "[1]", "&0x0F) << 8")
		g.p("if accumuDistance == nil {")
		g.p("accumuDistance = uint32NewAccumulator(12)")
		g.p("}")
		g.p("x.Distance = accumuDistance.accumulate(")
		g.p("uint32(x.", field.CCName, "[1]>>4) | uint32(x.", field.CCName, "[2]<< 4),")
		g.p(")")
		g.p("}")
	case "EventTimestamp12":
		g.p("// TODO")
	case "MesgData":
		g.p("if len(x.", field.CCName, ") != 0 {")
		g.p("x.Data = make([]byte, len(x.", field.CCName, ")-1)")
		g.p("for i, v := range x.", field.CCName, " {")
		g.p("if v == ", field.FType.BaseType().GoInvalidValue(), "{")
		g.p("break")
		g.p("}")
		g.p("if i == 0 {")
		g.p("x.ChannelNumber = v")
		g.p("} else {")
		g.p("x.Data[i-1] = v")
		g.p("}")
		g.p("}")
		g.p("}")
	default:
		fatalErr := fmt.Sprintf("genExpandComponentsArray: unhandled case for field %q", field.CCName)
		panic(fatalErr)
	}
}

func (g *codeGenerator) genExpandComponentsDyn(msg *Msg, field *Field, dcsfis []int) {
	refFieldNamesSet := make(map[string]bool)
	for _, subfi := range dcsfis {
		subf := field.Subfields[subfi]
		g.logger.Println("expand components: msg:", msg.CCName, "- field:", field.CCName, "- subfield:", subf.CCName)
		for _, reffn := range subf.RefFieldName {
			refFieldNamesSet[reffn] = true
		}
	}

	refFieldNameToType := make(map[string]string)
	for rfn := range refFieldNamesSet {
		for _, f := range msg.Fields {
			if f.CCName == rfn {
				refFieldNameToType[rfn] = f.TypeName
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

	g.p("if x.", field.CCName, " != ", field.FType.GoInvalidValue(), " {")

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

func (g *codeGenerator) genExpandComponentsMaskShift(msg *Msg, field *Field) {
	if msg.CCName == "Hr" {
		return
	}
	bits := 0
	for _, comp := range field.Components {
		tfield, tfound := msg.FieldByName[comp.Name]
		if !tfound {
			panic("genExpandComponentsMaskShift: target field not found")
		}
		if comp.Accumulate {
			accumulator := "accumu" + comp.Name
			g.p("if ", accumulator, " == nil {")
			g.p(accumulator, " = new(", tfield.FType.GoType(), "Accumulator)")
			g.p("}")
			g.p("x.", comp.Name, " = ", accumulator, ".accumulate(")
			g.p(tfield.FType.GoType(), "(")
			g.p("(x.", field.CCName, " >> ", bits, ") & ((1 << ", comp.Bits, ") - 1),")
			g.p("),")
			g.p(")")
			continue
		}
		g.p("x.", comp.Name, " = ", tfield.FType.GoType(), "(")
		g.p("(x.", field.CCName, " >> ", bits, ") & ((1 << ", comp.Bits, ") - 1),")
		g.p(")")
		bits += comp.BitsInt
	}
}

func (g *codeGenerator) genExpandComponentsMaskShiftDyn(msg *Msg, sfield *Field, mfield *Field) {
	bits := 0
	for _, comp := range sfield.Components {
		tfield, tfound := msg.FieldByName[comp.Name]
		if !tfound {
			panic("genExpandComponentsMaskShift: target field not found")
		}
		g.p("x.", comp.Name, " = ", tfield.FType.GoType(), "(")
		g.p("(x.", mfield.CCName, " >> ", bits, ") & ((1 << ", comp.Bits, ") - 1),")
		g.p(")")
		bits += comp.BitsInt
	}
}

func (g *codeGenerator) genProfile(types map[string]*Type, msgs []*Msg) {
	g.p("import (")
	g.p("\"reflect\"")
	g.p()
	g.p("\"github.com/tormoder/fit/internal/types\"")
	g.p(")")

	g.genVersionConsts()
	g.genKnownMsgs(types)
	g.genAccumulators(msgs)
	g.genFieldsArray(msgs)
	g.genGetFieldArrayLookup()
	g.genMsgTypesArray(msgs)
	g.genZeroValueMsgsArray(msgs)
	g.genGetZeroValueMsgsArrayLookup()
}

func (g *codeGenerator) genVersionConsts() {
	g.p("const (")
	g.p("// ProfileMajorVersion is the current supported profile major version of the FIT SDK.")
	g.p("ProfileMajorVersion =", g.sdkMajVer)
	g.p()
	g.p("// ProfileMinorVersion is the current supported profile minor version of the FIT SDK.")
	g.p("ProfileMinorVersion =", g.sdkMinVer)
	g.p(")")
	g.p()
}

func (g *codeGenerator) genKnownMsgs(types map[string]*Type) {
	mesgNums, found := types["MesgNum"]
	if !found {
		panic("genKnownMsgs: can't find MesgNum type")
	}
	mnvals := mesgNums.Values
	g.p()
	g.p("var knownMsgNums = map[MesgNum]bool{")
	for i := 0; i < len(mnvals)-2; i++ { // -2: Skip the last two: RangeMin/Max
		if knownMesgNumButNoMsg(g.sdkFullVer, mnvals[i].Name) {
			continue
		}
		g.p("MesgNum", mnvals[i].Name, ": true,")
	}
	g.p("}")
}

func (g *codeGenerator) genAccumulators(msgs []*Msg) {
	g.p()
	g.p("var (")
	// For-loop hell.
	for _, msg := range msgs {
		if msg.CCName == "Hr" {
			continue
		}
		for _, field := range msg.Fields {
			for _, comp := range field.Components {
				if comp.Accumulate {
					g.genAccumulator(comp, msg)
				}
			}
			for _, sfield := range field.Subfields {
				for _, comp := range sfield.Components {
					if comp.Accumulate {
						g.genAccumulator(comp, msg)
					}
				}
			}
		}
	}
	g.p(")")
}

func (g *codeGenerator) genAccumulator(comp Component, msg *Msg) {
	targetf, found := msg.FieldByName[comp.Name]
	if !found {
		panic("genAccumulator: target field for component not found")
	}
	g.p("accumu", comp.Name, " *", targetf.FType.GoType(), "Accumulator")
}

func (g *codeGenerator) genFieldsArray(msgs []*Msg) {
	g.p()
	g.p("// Set length to 256, so that lookup for any")
	g.p("// field 255 (localMesgNumInvalid) will return nil.")
	g.p("var _fields = [...][256]*field{")
	for _, msg := range msgs {
		g.p("MesgNum", msg.CCName, ": {")
		for i := 0; i < len(msg.Fields); i++ {
			f := msg.Fields[i]
			g.p(f.DefNum, ": {", i, ", ", f.DefNum, ", ", f.FType.ValueString(), "},")
		}
		g.p("},")
		g.p()
	}
	g.p("}")
}

func (g *codeGenerator) genGetFieldArrayLookup() {
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

func (g *codeGenerator) genZeroValueMsgsArray(msgs []*Msg) {
	g.p()
	g.p("var msgsAllInvalid = [...]reflect.Value{")
	for _, msg := range msgs {

		g.p("MesgNum", msg.CCName, ": reflect.ValueOf(", msg.CCName, "Msg{")
		for _, f := range msg.Fields {
			g.p(f.FType.GoInvalidValue(), ",")
		}
		g.p("}),")
	}
	g.p("}")
}

func (g *codeGenerator) genMsgTypesArray(msgs []*Msg) {
	g.p()
	g.p("var msgsTypes = [...]reflect.Type{")
	for _, msg := range msgs {
		g.p("MesgNum", msg.CCName, ": reflect.TypeOf(", msg.CCName, "Msg{}),")
	}
	g.p("}")
}

func (g *codeGenerator) genGetZeroValueMsgsArrayLookup() {
	g.p()
	g.p("func getMesgAllInvalid(mn MesgNum) reflect.Value {")
	g.p("val := reflect.New(msgsTypes[mn]).Elem()")
	g.p("val.Set(msgsAllInvalid[mn])")
	g.p("return val")
	g.p("}")
}
