package profile

import (
	"fmt"
	"strings"
)

var typeQuirks = map[string]string{
	"activity": "activity_mode",
}

type GoType struct {
	Name          string
	CamelCaseName string
	PkgName       string
	BaseType      string
	GoBaseType    string
	InvalidValue  string
	Values        []ValueTriple
}

type ValueTriple struct {
	Name, Value, Comment string
}

type GoMsg struct {
	Name          string
	CamelCaseName string
	Fields        []*GoField
}

type GoField struct {
	DefNum        string
	Name          string
	CamelCaseName string
	Type          string
	Array         string // 255=N
	Components    []string
	Scale         string
	Offset        string
	Units         string
	Bits          []string
	RefFieldName  []string
	RefFieldValue []string
	Comment       string
	Example       string

	BaseType  string
	GoType    string
	GoInvalid string

	DynFields []*GoField
}

func TransformTypes(types []*Type) (map[string]*GoType, error) {
	gts := make(map[string]*GoType)
	for _, t := range types {
		_, found := timestampTypes[t.Header[tname]]
		if found {
			continue
		}
		var gt GoType
		gt.Name = t.Header[tname]
		if len(gt.Name) == 0 {
			return nil, fmt.Errorf("found empty type name in header %q", t.Header)
		}
		gt.CamelCaseName = toCamelCase(t.Header[tname])
		gt.PkgName = strings.ToLower(gt.CamelCaseName)
		gt.BaseType = t.Header[tbtype]
		gt.GoBaseType, found = baseTypeToGoType[gt.BaseType]
		if !found {
			return nil, fmt.Errorf("no go type found for base type: %q", t.Header[tbtype])
		}
		gt.InvalidValue, found = baseTypeToInvalidValue[gt.BaseType]
		if !found {
			return nil, fmt.Errorf("no invalid value found for base type: %q", gt.BaseType)
		}

		for _, f := range t.Fields {
			vt := ValueTriple{
				Name:    toCamelCase(f[tvalname]),
				Value:   trimFloat(f[tval]),
				Comment: f[tcomment],
			}
			gt.Values = append(gt.Values, vt)
		}

		if renamed, found := typeQuirks[gt.Name]; found {
			gt.Name = renamed
			gt.CamelCaseName = toCamelCase(gt.Name)
			gt.PkgName = strings.ToLower(gt.CamelCaseName)
		}
		gts[gt.CamelCaseName] = &gt
	}
	return gts, nil
}

func TransformMsgs(msgs []*Msg, types map[string]*GoType) ([]*GoMsg, error) {
	var goMsgs []*GoMsg
	for _, msg := range msgs {
		var goMsg GoMsg
		goMsg.Name = msg.Header[mmsgname]
		if len(goMsg.Name) == 0 {
			return nil, fmt.Errorf(
				"found empty message name in header %q",
				msg.Header,
			)
		}
		goMsg.CamelCaseName = toCamelCase(goMsg.Name)

		for _, msgField := range msg.Fields {
			regField, err := parseField(msgField.RegField, types)
			if err != nil {
				return nil, err
			}
			if regField == nil {
				continue
			}
			goMsg.Fields = append(goMsg.Fields, regField)
			if len(msgField.DynFields) == 0 {
				continue
			}
			for _, df := range msgField.DynFields {
				dynField, err := parseField(df, types)
				if err != nil {
					return nil, err
				}
				if dynField != nil {
					regField.DynFields = append(regField.DynFields, dynField)
				}
			}
		}
		goMsgs = append(goMsgs, &goMsg)
	}

	return goMsgs, nil
}

func parseField(mff []string, types map[string]*GoType) (*GoField, error) {
	var (
		f     GoField
		ft    string
		found bool
	)

	if mff[mexample] == "" {
		return nil, nil
	}

	f.DefNum = trimFloat(mff[mfdefn])
	f.Name = mff[mfname]
	f.CamelCaseName = toCamelCase(f.Name)

	typeName := mff[mftype]
	if rewrite, tfound := typeQuirks[typeName]; tfound {
		typeName = rewrite
	}
	typeDef, found := types[toCamelCase(typeName)]
	if found {
		f.BaseType = "fit" + toCamelCase(typeDef.BaseType)
	} else {
		_, found = timestampTypes[typeName]
		if found {
			f.BaseType = "fitUint32"
		} else if typeName == "bool" {
			f.BaseType = "fitEnum"
		} else {
			f.BaseType = "fit" + toCamelCase(typeName)
		}
	}

	switch {
	case mff[mscale] != "" && mff[mcomps] == "":
		scaleSplit := strings.Split(mff[mscale], ",")
		f.Scale = scaleSplit[0]
		if mff[moffset] != "" {
			f.Offset = trimFloat(mff[moffset])
		} else {
			f.Offset = "0"
		}
		f.Type = "float64"
		f.GoType = "float"
		f.GoInvalid = "0xFFFFFFFFFFFFFFFF"
	case strings.HasSuffix(f.Name, "_lat"):
		f.Type = "Latitude"
		f.Scale = "0"
		f.Offset = "0"
		f.GoType = "lat"
		f.GoInvalid = "0x7FFFFFFF"
	case strings.HasSuffix(f.Name, "_long"):
		f.Type = "Longitude"
		f.Scale = "0"
		f.Offset = "0"
		f.GoType = "lng"
		f.GoInvalid = "0x7FFFFFFF"
	default:
		ft = mff[mftype]
		_, found = timestampTypes[ft]
		if found {
			f.Type = "time.Time"
			f.Scale = "0"
			f.Offset = "0"
			f.GoInvalid = "timeBase"
			if strings.HasPrefix(ft, "local") {
				f.GoType = "timelocal"
			} else {
				f.GoType = "timeutc"
			}
		} else {
			if renamed, shouldRename := typeQuirks[ft]; shouldRename {
				f.Type = toCamelCase(renamed)
			} else {
				if btype, isBaseType := baseTypeToGoType[ft]; isBaseType {
					f.Type = btype
				} else {
					f.Type = toCamelCase(ft)
				}
			}

			f.Scale = "0"
			f.Offset = "0"
			f.GoType = "fit"

			typeDef, tfound := types[f.Type]
			if !tfound {
				// special case for now
				if f.Type == "Bool" {
					f.GoInvalid = "0xFF"
				} else {
					// Assume base type.
					invalidValue, bfound := goBaseTypeToInvalidValue[f.Type]
					if !bfound {
						return nil, fmt.Errorf(
							"TransformMsgs: base type for type %q not found",
							f.Type,
						)
					}
					f.GoInvalid = invalidValue
				}
			} else {
				f.GoInvalid = typeDef.InvalidValue
			}
		}
	}

	arrayStr := strings.TrimFunc(
		mff[marray], func(r rune) bool {
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

	if mff[mcomps] != "" {
		f.Components = strings.Split(mff[mcomps], ",")
		f.Bits = strings.Split(mff[mbits], ",")
	}

	f.Units = mff[munits]

	f.RefFieldName = strings.Split(mff[mrfname], ",")
	if len(f.RefFieldName) > 0 {
		for i, rfn := range f.RefFieldName {
			tmp := strings.TrimSpace(rfn)
			f.RefFieldName[i] = toCamelCase(tmp)
		}
	}

	f.RefFieldValue = strings.Split(mff[mrfval], ",")
	if len(f.RefFieldValue) > 0 {
		for i, rfv := range f.RefFieldValue {
			tmp := strings.TrimSpace(rfv)
			f.RefFieldValue[i] = toCamelCase(tmp)
		}
	}
	f.Comment = mff[mcomment]
	f.Example = trimFloat(mff[mexample])

	return &f, nil
}
