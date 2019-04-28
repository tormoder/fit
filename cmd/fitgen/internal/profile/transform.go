package profile

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	"github.com/tormoder/fit/internal/types"
)

var camelRegex = regexp.MustCompile("[0-9A-Za-z]+")

func toCamelCase(s string) string {
	chunks := camelRegex.FindAllString(s, -1)
	for i, val := range chunks {
		chunks[i] = strings.Title(val)
	}
	return strings.Join(chunks, "")
}

var typeQuirks = map[string]string{
	"activity": "activity_mode",
	"file":     "file_type",
}

func isTimestamp(name string) (types.Kind, bool) {
	if name == "date_time" {
		return types.TimeUTC, true
	}
	if name == "local_date_time" {
		return types.TimeLocal, true
	}
	return 0, false
}

func isCoordinate(name string) (types.Kind, bool) {
	if strings.HasSuffix(name, "_lat") {
		return types.Lat, true
	}
	if strings.HasSuffix(name, "_long") {
		return types.Lng, true
	}
	return 0, false
}

func TransformTypes(ptypes []*PType) (map[string]*Type, error) {
	types := make(map[string]*Type)
	for _, pt := range ptypes {
		t := Type{data: pt}
		skip, err := t.transform()
		if err != nil {
			return nil, err
		}
		if skip {
			continue
		}
		types[t.Name] = &t
	}

	return types, nil
}

func (t *Type) transform() (skip bool, err error) {
	_, isTS := isTimestamp(t.data.Header[tNAME])
	if isTS {
		return true, nil
	}

	name := t.data.Header[tNAME]
	if name == "" {
		return false, fmt.Errorf(
			"found empty type name in header %q",
			t.data.Header)
	}
	t.OrigName = name
	t.Name = toCamelCase(name)

	t.BaseType, err = types.BaseFromString(t.data.Header[tBTYPE])
	if err != nil {
		return false, err
	}

	for _, f := range t.data.Fields {
		vt := ValueTriple{
			Name:    toCamelCase(f[tVALNAME]),
			Value:   f[tVAL],
			Comment: f[tCOMMENT],
		}
		t.Values = append(t.Values, vt)
	}

	if renamed, found := typeQuirks[name]; found {
		t.Name = toCamelCase(renamed)
	}

	return false, nil
}

func TransformMsgs(pmsgs []*PMsg, ftypes map[string]*Type, logger *log.Logger) ([]*Msg, error) {
	var msgs []*Msg
	for _, pmsg := range pmsgs {
		msg := Msg{
			Name:        pmsg.Header[mMSGNAME],
			FieldByName: make(map[string]*Field),
		}

		if len(msg.Name) == 0 {
			return nil, fmt.Errorf(
				"found empty message name in header %q",
				pmsg.Header)
		}
		msg.CCName = toCamelCase(msg.Name)
		logger.Println("transforming message:", msg.CCName)

		for _, pfield := range pmsg.Fields {
			f := &Field{data: pfield.Field}
			skip, err := f.transform(false, ftypes, logger)
			if err != nil {
				return nil, err
			}
			if skip {
				continue
			}
			msg.Fields = append(msg.Fields, f)
			msg.FieldByName[f.CCName] = f
			if len(pfield.Subfields) == 0 {
				continue
			}

			for _, sfield := range pfield.Subfields {
				sf := &Field{data: sfield}
				skip, err := sf.transform(true, ftypes, logger)
				if err != nil {
					return nil, fmt.Errorf("error parsing subfield: %v", err)
				}
				if skip {
					continue
				}
				f.Subfields = append(f.Subfields, sf)
			}
		}
		msgs = append(msgs, &msg)
	}

	return msgs, nil
}

func (f *Field) transform(subfield bool, ftypes map[string]*Type, logger *log.Logger) (skip bool, err error) {
	if f.data[mEXAMPLE] == "" || f.data[mEXAMPLE] == "0" {
		return true, nil
	}

	f.DefNum = f.data[mFDEFN]
	f.Name = f.data[mFNAME]
	f.CCName = toCamelCase(f.Name)

	f.parseArray()

	err = f.parseType(ftypes)
	if err != nil {
		return false, err
	}

	f.Units = f.data[mUNITS]
	f.Comment = f.data[mCOMMENT]
	f.Example = f.data[mEXAMPLE]

	if subfield {
		f.parseRefFields()
	}

	if f.data[mCOMPS] == "" {
		f.parseScaleOffset()
		return false, nil
	}

	return false, f.parseComponents(logger)
}

func (f *Field) parseArray() {
	arrayStr := strings.TrimFunc(
		f.data[mARRAY], func(r rune) bool {
			if r == '[' || r == ']' {
				return true
			}
			return false
		})
	switch arrayStr {
	case "":
		f.Array = "0"
	case "N":
		f.Array = "255"
	default:
		f.Array = arrayStr
	}
}

func (f *Field) parseType(ftypes map[string]*Type) error {
	array := f.Array != "0"

	coordKind, isCoord := isCoordinate(f.Name)
	if isCoord {
		f.FType = types.Make(coordKind, array)
		f.TypeName = f.FType.GoType()
		return nil
	}

	originalTypeName := f.data[mFTYPE]
	if rewritten, tfound := typeQuirks[originalTypeName]; tfound {
		f.TypeName = toCamelCase(rewritten)
	} else {
		f.TypeName = toCamelCase(originalTypeName)
	}

	tsKind, isTimestamp := isTimestamp(originalTypeName)
	if isTimestamp {
		f.FType = types.Make(tsKind, array)
		f.TypeName = f.FType.GoType()
		return nil
	}

	if f.TypeName == "Bool" {
		f.FType = types.MakeNative(types.BaseEnum, array)
		return nil
	}

	typeDef, found := ftypes[f.TypeName]
	if found {
		f.FType = types.MakeNative(typeDef.BaseType, array)
		if array {
			f.TypeName = "[]" + f.TypeName
		}
		return nil
	}

	// Assume base type.
	baseType, err := types.BaseFromString(originalTypeName)
	if err != nil {
		return err
	}
	f.FType = types.MakeNative(baseType, array)
	f.TypeName = f.FType.GoType()

	return nil
}

func (f *Field) parseRefFields() {
	f.RefFieldName = strings.Split(f.data[mRFNAME], ",")
	if len(f.RefFieldName) > 0 {
		for i, rfn := range f.RefFieldName {
			tmp := strings.TrimSpace(rfn)
			f.RefFieldName[i] = toCamelCase(tmp)
		}
	}

	f.RefFieldValue = strings.Split(f.data[mRFVAL], ",")
	if len(f.RefFieldValue) > 0 {
		for i, rfv := range f.RefFieldValue {
			tmp := strings.TrimSpace(rfv)
			f.RefFieldValue[i] = toCamelCase(tmp)
		}
	}
}

func (f *Field) parseScaleOffset() {
	if f.data[mSCALE] == "" {
		return
	}
	f.Scale = f.data[mSCALE]
	if f.data[mOFFSET] != "" {
		f.Offset = f.data[mOFFSET]
	}
}

func (f *Field) parseComponents(logger *log.Logger) error {
	if f.data[mCOMPS] == "" {
		return nil
	}

	logger.Println("parsing components for field:", f.CCName)

	switch f.FType.BaseType() {
	case types.BaseUint8, types.BaseUint16, types.BaseUint32, types.BaseByte:
	default:
		return fmt.Errorf(
			"parseComponents: unhandled base type (%s) for field %s",
			f.FType.BaseType(), f.CCName)
	}

	components := strings.Split(f.data[mCOMPS], ",")
	if len(components) == 0 {
		return fmt.Errorf("parseComponents: zero components after string split")
	}

	bitsFull := f.data[mBITS]
	if new, rewrite := bitsRewrite[bitsFull]; rewrite {
		bitsFull = new
	}
	bits := strings.Split(bitsFull, ",")

	if len(components) != len(bits) {
		return fmt.Errorf("parseComponents: number of components (%d) and bits (%d) differ", len(components), len(bits))
	}

	accumulate := strings.Split(f.data[mACCUMU], ",")
	if len(accumulate) == 1 && accumulate[0] == "" {
		accumulate = nil
	}

	if len(accumulate) > 0 && (len(accumulate) != len(components)) {
		return fmt.Errorf(
			"parseComponents: number of components (%d) and accumulate flags (%d) differ",
			len(components), len(accumulate))
	}

	f.Components = make([]Component, len(components))

	var (
		err       error
		bitsTotal int
	)

	for i, comp := range components {
		f.Components[i].Name = strings.TrimSpace(comp)
		f.Components[i].Name = toCamelCase(f.Components[i].Name)
		f.Components[i].Bits = strings.TrimSpace(bits[i])
		f.Components[i].BitsInt, err = strconv.Atoi(f.Components[i].Bits)
		if err != nil {
			return fmt.Errorf("parseComponents: error converting bit to integer: %v", err)
		}
		bitsTotal += f.Components[i].BitsInt
		if len(accumulate) == len(components) {
			tmp := strings.TrimSpace(accumulate[i])
			f.Components[i].Accumulate, err = strconv.ParseBool(tmp)
			if err != nil {
				return fmt.Errorf("parseComponents: %v", err)
			}
		}
	}

	if (bitsTotal > 32 && !f.FType.Array()) || bitsTotal < 0 {
		return fmt.Errorf("parseComponents: illegal size for total number of bits: %d", bitsTotal)
	}

	if len(components) == 1 {
		// Set any scale on the "main" field.
		// TODO(tormoder): Verify that this is correct.
		f.parseScaleOffset()
		return nil
	}

	cscaleFull := f.data[mSCALE]
	if new, rewrite := scaleRewrite[cscaleFull]; rewrite {
		cscaleFull = new
	}
	cscale := strings.Split(cscaleFull, ",")
	if len(cscale) == 1 && cscale[0] == "" {
		cscale = nil
	}

	coffset := strings.Split(f.data[mOFFSET], ",")
	if len(coffset) == 1 && coffset[0] == "" {
		coffset = nil
	}

	if len(cscale) != 0 && len(cscale) != len(components) {
		return fmt.Errorf(
			"parseComponents: number of components (%d) and scales (%d) differ",
			len(components), len(cscale))
	}
	if len(coffset) != 0 && len(coffset) != len(components) {
		return fmt.Errorf(
			"parseComponents: #offset != 0 and number of components (%d) and offsets (%d) differ",
			len(components), len(coffset))
	}

	for i := range f.Components {
		if len(cscale) == 0 {
			continue
		}
		f.Components[i].Scale = strings.TrimSpace(cscale[i])
		if len(coffset) == 0 {
			continue
		}
		f.Components[i].Offset = strings.TrimSpace(coffset[i])
	}

	return nil
}

// Rewrite maps.
// In SDK ~16.20 all bits and scales were separated by commas.
// In SDK 20.14 some bits and scales are just concatenated together,
// making them hard to parse. Use two maps for now.
var bitsRewrite = map[string]string{
	"1616":      "16,16",
	"88888888":  "8,8,8,8,8,8,8,8",
	"888888888": "8,8,8,8,8,8,8,8,8",
	"53":        "5,3",
	"44":        "4,4",
}

var scaleRewrite = map[string]string{
	"11":   "1,1",
	"1111": "1,1,1,1",
}
