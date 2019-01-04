package fit

import (
	"bytes"
	"encoding/binary"
	"reflect"
	"testing"
	"time"

	"github.com/tormoder/fit/internal/types"
)

func TestEncodeWriteField(t *testing.T) {
	type writeFieldTest struct {
		field field
		value interface{}
		le    []byte
		be    []byte
	}

	tests := []writeFieldTest{
		{
			field: field{
				t:      types.MakeNative(types.BaseEnum, false),
				length: 1,
			},
			value: byte(0x42),
			le:    []byte{0x42},
			be:    []byte{0x42},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseSint8, false),
				length: 1,
			},
			value: int8(-0x80),
			le:    []byte{0x80},
			be:    []byte{0x80},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseUint8, false),
				length: 1,
			},
			value: uint8(0x80),
			le:    []byte{0x80},
			be:    []byte{0x80},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseSint16, false),
				length: 1,
			},
			value: int16(-0x1234),
			le:    []byte{0xCC, 0xED},
			be:    []byte{0xED, 0xCC},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseUint16, false),
				length: 1,
			},
			value: uint16(0x1234),
			le:    []byte{0x34, 0x12},
			be:    []byte{0x12, 0x34},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseSint32, false),
				length: 1,
			},
			value: int32(-0x12345678),
			le:    []byte{0x88, 0xA9, 0xCB, 0xED},
			be:    []byte{0xED, 0xCB, 0xA9, 0x88},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseUint32, false),
				length: 1,
			},
			value: uint32(0x12345678),
			le:    []byte{0x78, 0x56, 0x34, 0x12},
			be:    []byte{0x12, 0x34, 0x56, 0x78},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseString, false),
				length: 6,
			},
			value: string("Hello"),
			le:    []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x00},
			be:    []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x00},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseString, false),
				length: 7,
			},
			value: string("Hello"),
			le:    []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x00, 0x00},
			be:    []byte{0x48, 0x65, 0x6C, 0x6C, 0x6F, 0x00, 0x00},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseString, false),
				length: 5,
			},
			value: string("Hello"),
			le:    []byte{0x48, 0x65, 0x6C, 0x6C, 0x00},
			be:    []byte{0x48, 0x65, 0x6C, 0x6C, 0x00},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseFloat32, false),
				length: 1,
			},
			value: float32(3.142),
			le:    []byte{0x87, 0x16, 0x49, 0x40},
			be:    []byte{0x40, 0x49, 0x16, 0x87},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseFloat64, false),
				length: 1,
			},
			value: float64(3.142),
			le:    []byte{0x89, 0x41, 0x60, 0xE5, 0xD0, 0x22, 0x09, 0x40},
			be:    []byte{0x40, 0x09, 0x22, 0xD0, 0xE5, 0x60, 0x41, 0x89},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseUint8z, false),
				length: 1,
			},
			value: uint8(0x80),
			le:    []byte{0x80},
			be:    []byte{0x80},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseUint16z, false),
				length: 1,
			},
			value: uint16(0x1234),
			le:    []byte{0x34, 0x12},
			be:    []byte{0x12, 0x34},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseUint32z, false),
				length: 1,
			},
			value: uint32(0x12345678),
			le:    []byte{0x78, 0x56, 0x34, 0x12},
			be:    []byte{0x12, 0x34, 0x56, 0x78},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseByte, false),
				length: 1,
			},
			value: byte(0x42),
			le:    []byte{0x42},
			be:    []byte{0x42},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseSint64, false),
				length: 1,
			},
			value: int64(-0x12345678ABCDEF00),
			le:    []byte{0x00, 0x11, 0x32, 0x54, 0x87, 0xA9, 0xCB, 0xED},
			be:    []byte{0xED, 0xCB, 0xA9, 0x87, 0x54, 0x32, 0x11, 0x00},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseUint64, false),
				length: 1,
			},
			value: uint64(0x12345678ABCDEF00),
			le:    []byte{0x00, 0xEF, 0xCD, 0xAB, 0x78, 0x56, 0x34, 0x12},
			be:    []byte{0x12, 0x34, 0x56, 0x78, 0xAB, 0xCD, 0xEF, 0x00},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseUint64z, false),
				length: 1,
			},
			value: uint64(0x12345678ABCDEF00),
			le:    []byte{0x00, 0xEF, 0xCD, 0xAB, 0x78, 0x56, 0x34, 0x12},
			be:    []byte{0x12, 0x34, 0x56, 0x78, 0xAB, 0xCD, 0xEF, 0x00},
		},
		{
			field: field{
				t:      types.Make(types.TimeUTC, false),
				length: 1,
			},
			value: timeBase,
			le:    []byte{0x00, 0x00, 0x00, 0x00},
			be:    []byte{0x00, 0x00, 0x00, 0x00},
		},
		{
			field: field{
				t:      types.Make(types.TimeUTC, false),
				length: 1,
			},
			// Less than 0x10000000, should be encoded as-is
			value: timeBase.Add(3600 * time.Second),
			le:    []byte{0x10, 0x0E, 0x00, 0x00},
			be:    []byte{0x00, 0x00, 0x0E, 0x10},
		},
		{
			field: field{
				t:      types.Make(types.TimeUTC, false),
				length: 1,
			},
			// 10 years - greater than 0x10000000
			value: timeBase.Add(315532800 * time.Second),
			le:    []byte{0x00, 0xA6, 0xCE, 0x12},
			be:    []byte{0x12, 0xCE, 0xA6, 0x00},
		},
		{
			field: field{
				t:      types.Make(types.TimeLocal, false),
				length: 1,
			},
			// 10 years - greater than 0x10000000
			value: timeBase.In(time.FixedZone("FITLOCAL", -3600)).Add(315532800 * time.Second),
			le:    []byte{0xF0, 0x97, 0xCE, 0x12},
			be:    []byte{0x12, 0xCE, 0x97, 0xF0},
		},
		{
			field: field{
				t:      types.Make(types.TimeLocal, false),
				length: 1,
			},
			// Less than 0x10000000, should be encoded as-is
			value: timeBase.Add(3600 * time.Second),
			le:    []byte{0x10, 0x0E, 0x00, 0x00},
			be:    []byte{0x00, 0x00, 0x0E, 0x10},
		},
		{
			field: field{
				t:      types.Make(types.Lat, false),
				length: 1,
			},
			value: NewLatitudeDegrees(52.2053),
			le:    []byte{0x51, 0xAF, 0x1F, 0x25},
			be:    []byte{0x25, 0x1F, 0xAF, 0x51},
		},
		{
			field: field{
				t:      types.Make(types.Lng, false),
				length: 1,
			},
			value: NewLongitudeDegrees(0.1218),
			le:    []byte{0x4A, 0x2C, 0x16, 0x00},
			be:    []byte{0x00, 0x16, 0x2C, 0x4A},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseByte, true),
				length: 4,
			},
			value: []byte{0x01, 0x02, 0x03, 0x04},
			le:    []byte{0x01, 0x02, 0x03, 0x04},
			be:    []byte{0x01, 0x02, 0x03, 0x04},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseByte, true),
				length: 3,
			},
			value: []byte{0x01, 0x02, 0x03, 0x04},
			le:    []byte{0x01, 0x02, 0x03},
			be:    []byte{0x01, 0x02, 0x03},
		},
		{
			field: field{
				t:      types.MakeNative(types.BaseByte, true),
				length: 5,
			},
			value: []byte{0x01, 0x02, 0x03, 0x04},
			le:    []byte{0x01, 0x02, 0x03, 0x04, 0xFF},
			be:    []byte{0x01, 0x02, 0x03, 0x04, 0xFF},
		},
	}

	buf := &bytes.Buffer{}

	e := &encoder{
		w: buf,
	}

	e.arch = binary.LittleEndian
	for i, test := range tests {
		buf.Reset()

		err := e.writeField(reflect.ValueOf(test.value), &test.field)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(buf.Bytes(), test.le) {
			t.Errorf("LE %d (%s): Expected '%v' got '%v'", i, test.field.t, test.le, buf.Bytes())
		}
	}

	e.arch = binary.BigEndian
	for i, test := range tests {
		buf.Reset()

		err := e.writeField(reflect.ValueOf(test.value), &test.field)
		if err != nil {
			t.Fatal(err)
		}

		if !bytes.Equal(buf.Bytes(), test.be) {
			t.Errorf("BE %d (%s): Expected '%v' got '%v'", i, test.field.t, test.be, buf.Bytes())
		}
	}
}
