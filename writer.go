package fit

import (
	"encoding/binary"
	"fmt"
	"io"
	"reflect"
	"time"
	"unicode/utf8"

	"github.com/tormoder/fit/internal/types"
)

type encoder struct {
	w    io.Writer
	arch binary.ByteOrder
}

func encodeString(str string, size byte) ([]byte, error) {
	length := len(str)
	if length > int(size)-1 {
		length = int(size) - 1
	}

	bstr := make([]byte, size)
	copy(bstr, str[:length])
	if !utf8.Valid(bstr) {
		return nil, fmt.Errorf("can't encode %+v as UTF-8 string", str)
	}
	return bstr, nil
}

func (e *encoder) encodeValue(value interface{}, f *field) error {
	switch f.t.Kind() {
	case types.TimeUTC:
		t := value.(time.Time)
		u32 := encodeTime(t)
		binary.Write(e.w, e.arch, u32)
	case types.TimeLocal:
		t := value.(time.Time)
		_, offs := t.Zone()
		u32 := uint32(int64(encodeTime(t)) + int64(offs))
		binary.Write(e.w, e.arch, u32)
	case types.Lat:
		lat := value.(Latitude)
		binary.Write(e.w, e.arch, lat.semicircles)
	case types.Lng:
		lng := value.(Longitude)
		binary.Write(e.w, e.arch, lng.semicircles)
	case types.NativeFit:
		if f.t.BaseType() == types.BaseString {
			str, ok := value.(string)
			if !ok {
				return fmt.Errorf("not a string: %+v", value)
			}

			var err error
			value, err = encodeString(str, f.length)
			if err != nil {
				return fmt.Errorf("can't encode %+v as UTF-8 string: %v", value, err)
			}
		}
		binary.Write(e.w, e.arch, value)
	default:
		return fmt.Errorf("unknown Fit type %+v", f.t)
	}

	return nil
}

func (e *encoder) writeField(value reflect.Value, f *field) error {
	if !f.t.Array() {
		return e.encodeValue(value.Interface(), f)
	}

	if f.t.BaseType() == types.BaseString {
		return fmt.Errorf("can't encode array of strings")
	}

	invalid := f.t.BaseType().Invalid()
	max := byte(value.Len())
	if max > f.length {
		max = f.length
	}
	for i := byte(0); i < max; i++ {
		elem := value.Index(int(i))
		err := e.encodeValue(elem.Interface(), f)
		if err != nil {
			return err
		}
	}
	for i := max; i < f.length; i++ {
		err := e.encodeValue(invalid, f)
		if err != nil {
			return err
		}
	}

	return nil
}

type encodeMesgDef struct {
	globalMesgNum MesgNum
	localMesgNum  byte
	fields        []*field
}

func (e *encoder) writeMesg(mesg reflect.Value, def *encodeMesgDef) error {
	hdr := def.localMesgNum & localMesgNumMask
	err := binary.Write(e.w, e.arch, hdr)
	if err != nil {
		return err
	}

	for _, f := range def.fields {
		value := mesg.Field(f.sindex)

		err := e.writeField(value, f)
		if err != nil {
			return err
		}
	}

	return nil
}

func profileFieldDef(m MesgNum) [256]*field {
	return _fields[m]
}

func getFieldBySindex(index int, fields [256]*field) *field {
	for _, f := range fields {
		if f != nil && index == f.sindex {
			return f
		}
	}

	return fields[255]
}

// getEncodeMesgDef generates an appropriate encodeMesgDef to will encode all
// of the valid fields in mesg. Any fields which are set to their respective
// invalid value will be skipped (not present in the returned encodeMesgDef)
func getEncodeMesgDef(mesg reflect.Value, localMesgNum byte) *encodeMesgDef {
	mesgNum := getGlobalMesgNum(mesg.Type())
	allInvalid := getMesgAllInvalid(mesgNum)
	profileFields := profileFieldDef(mesgNum)

	if mesg.NumField() != allInvalid.NumField() {
		panic(fmt.Sprintf("mismatched number of fields in type %+v", mesg.Type()))
	}

	def := &encodeMesgDef{
		globalMesgNum: mesgNum,
		localMesgNum:  localMesgNum,
		fields:        make([]*field, 0, mesg.NumField()),
	}

	for i := 0; i < mesg.NumField(); i++ {
		fval := mesg.Field(i)
		field := getFieldBySindex(i, profileFields)

		// Don't encode invalid fields
		if fval.Kind() == reflect.Slice {
			if fval.IsNil() {
				continue
			}

			skip := true
			invalid := field.t.BaseType().Invalid()
			for i := 0; i < fval.Len(); i++ {
				if fval.Interface() != invalid {
					skip = false
					break
				}
			}
			if skip {
				continue
			}
		} else if fval.Interface() == allInvalid.Field(i).Interface() {
			continue
		}

		// FIXME: No message can exceed 255 bytes
		def.fields = append(def.fields, field)
	}

	return def
}

func (e *encoder) writeDefMesg(def *encodeMesgDef) error {
	hdr := mesgDefinitionMask | (def.localMesgNum & localMesgNumMask)
	err := binary.Write(e.w, e.arch, hdr)
	if err != nil {
		return err
	}

	err = binary.Write(e.w, e.arch, byte(0))
	if err != nil {
		return err
	}

	switch e.arch {
	case binary.LittleEndian:
		err = binary.Write(e.w, e.arch, byte(0))
	case binary.BigEndian:
		err = binary.Write(e.w, e.arch, byte(1))
	}
	if err != nil {
		return err
	}

	err = binary.Write(e.w, e.arch, def.globalMesgNum)
	if err != nil {
		return err
	}

	err = binary.Write(e.w, e.arch, byte(len(def.fields)))
	if err != nil {
		return err
	}

	for _, f := range def.fields {
		fdef := fieldDef{
			num:   f.num,
			size:  byte(f.t.BaseType().Size()),
			btype: f.t.BaseType(),
		}
		if fdef.btype == types.BaseString {
			fdef.size = f.length
		} else if f.t.Array() {
			fdef.size = fdef.size * f.length
		}

		err := binary.Write(e.w, e.arch, fdef)
		if err != nil {
			return err
		}
	}

	return nil
}
