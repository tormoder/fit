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
