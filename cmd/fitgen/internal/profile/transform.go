package profile

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/tormoder/fit/internal/base"
)

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
		types[t.CCName] = &t
	}

	return types, nil
}

func (t *Type) transform() (skip bool, err error) {
	_, found := timestampTypes[t.data.Header[tNAME]]
	if found {
		return true, nil
	}

	t.Name = t.data.Header[tNAME]
	if len(t.Name) == 0 {
		return false, fmt.Errorf(
			"found empty type name in header %q",
			t.data.Header)
	}
	t.OrigName = t.Name
	t.CCName = toCamelCase(t.Name)

	t.PkgName = strings.ToLower(t.CCName)

	t.BaseType, err = base.TypeFromString(t.data.Header[tBTYPE])
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

	if renamed, found := typeQuirks[t.Name]; found {
		t.Name = renamed
		t.CCName = toCamelCase(t.Name)
		t.PkgName = strings.ToLower(t.CCName)
	}

	return false, nil
}

func TransformMsgs(pmsgs []*PMsg, types map[string]*Type) ([]*Msg, error) {
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
		debugln("transforming message", msg.CCName)

		for _, pfield := range pmsg.Fields {
			f := &Field{data: pfield.Field}
			skip, err := f.transform(false, types)
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
				skip, err := sf.transform(true, types)
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

func (f *Field) transform(subfield bool, types map[string]*Type) (skip bool, err error) {
	if f.data[mEXAMPLE] == "" {
		return true, nil
	}

	f.DefNum = f.data[mFDEFN]
	f.Name = f.data[mFNAME]
	f.CCName = toCamelCase(f.Name)

	err = f.parseBaseType(types)
	if err != nil {
		return false, err
	}

	err = f.parseType(types)
	if err != nil {
		return false, err
	}

	f.parseArray()

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

	return false, f.parseComponents(types)
}

func (f *Field) parseType(types map[string]*Type) error {
	switch {
	case strings.HasSuffix(f.Name, "_lat"):
		f.Type = "Latitude"
		f.GoType = "lat"
		f.GoInvalid = "NewLatitudeInvalid()"
	case strings.HasSuffix(f.Name, "_long"):
		f.Type = "Longitude"
		f.GoType = "lng"
		f.GoInvalid = "NewLongitudeInvalid()"
	default:
		ft := f.data[mFTYPE]
		_, found := timestampTypes[ft]
		if found {
			f.Type = "time.Time"
			f.GoInvalid = "timeBase"
			if strings.HasPrefix(ft, "local") {
				f.GoType = "timelocal"
			} else {
				f.GoType = "timeutc"
			}
			return nil
		}

		f.GoType = "fit"

		if renamed, shouldRename := typeQuirks[ft]; shouldRename {
			f.Type = toCamelCase(renamed)
		} else {
			if btype, isBaseType := baseTypeToGoType[ft]; isBaseType {
				f.Type = btype
			} else {
				f.Type = toCamelCase(ft)
			}
		}

		tdef, tfound := types[f.Type]
		if tfound {
			f.GoInvalid = tdef.BaseType.GoInvalidValue()
			return nil
		}

		if f.Type == "Bool" { // Special case for now.
			f.GoInvalid = base.Enum.GoInvalidValue()
			return nil
		}

		// Assume base type.
		val, vfound := goBaseTypeToInvalidValue[f.Type]
		if !vfound {
			return fmt.Errorf(
				"base type for type %q not found",
				f.Type,
			)

		}
		f.GoInvalid = val
		f.BTInvalid = val // GoInvalid may be overwritten in parseArray.
	}

	return nil
}

func (f *Field) parseBaseType(types map[string]*Type) error {
	typeName := f.data[mFTYPE]
	if rewrite, tfound := typeQuirks[typeName]; tfound {
		typeName = rewrite
	}
	typeDef, found := types[toCamelCase(typeName)]
	if found {
		f.BaseType = typeDef.BaseType
	} else {
		_, found = timestampTypes[typeName]
		if found {
			f.BaseType = base.Uint32
		} else if typeName == "bool" {
			f.BaseType = base.Enum
		} else {
			var err error
			f.BaseType, err = base.TypeFromString(typeName)
			if err != nil {
				return err
			}
		}
	}
	return nil
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
		f.GoInvalid = "nil"
	default:
		f.Array = arrayStr
		f.GoInvalid = "nil"
	}
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

func (f *Field) parseComponents(types map[string]*Type) error {
	if f.data[mCOMPS] == "" {
		return nil
	}

	debugln("parsing components for field", f.CCName)

	switch f.BaseType {
	case base.Uint8, base.Uint16, base.Uint32:
	case base.Byte:
		if !f.IsArray() {
			return fmt.Errorf("parseComponents: base type was byte but not an array")
		}
	default:
		return fmt.Errorf(
			"parseComponents: unhandled base type (%s) for field %s",
			f.BaseType, f.CCName)
	}

	components := strings.Split(f.data[mCOMPS], ",")
	if len(components) == 0 {
		return fmt.Errorf("parseComponents: zero components after string split")
	}

	bits := strings.Split(f.data[mBITS], ",")
	if len(components) != len(bits) {
		return fmt.Errorf(
			"parseComponents: number of components (%d) and bits (%d) differ",
			len(components), len(bits))
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

	if bitsTotal > 32 || bitsTotal < 0 {
		return fmt.Errorf("parseComponents: illegal size for total number of bits: %d", bitsTotal)
	}

	if len(components) == 1 {
		// Set any scale on the "main" field.
		// TODO(tormoder): Verify that this is correct.
		f.parseScaleOffset()
		return nil
	}

	cscale := strings.Split(f.data[mSCALE], ",")
	coffset := strings.Split(f.data[mOFFSET], ",")

	if len(coffset) == 1 && coffset[0] == "" {
		coffset = nil
	}
	if len(cscale) != len(components) {
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
		f.Components[i].Scale = strings.TrimSpace(cscale[i])
		if len(coffset) == 0 {
			continue
		}
		f.Components[i].Offset = strings.TrimSpace(coffset[i])
	}

	return nil
}
