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

type TestMesg struct {
	Type         byte
	Timestamp    time.Time
	PositionLat  Latitude
	PositionLong Longitude
}

func testMesgDef() *encodeMesgDef {
	return &encodeMesgDef{
		localMesgNum: 3,
		fields: []*field{
			{
				sindex: 0,
				num:    0,
				t:      types.MakeNative(types.BaseEnum, false),
				length: byte(types.BaseEnum.Size()),
			},
			{
				sindex: 1,
				num:    253,
				t:      types.Make(types.TimeUTC, false),
				length: byte(types.BaseUint32.Size()),
			},
			// PositionLat intentionally omitted
			{
				sindex: 3,
				num:    3,
				t:      types.Make(types.Lng, false),
				length: byte(types.BaseSint32.Size()),
			},
		},
	}
}

func TestEncodeWriteMesg(t *testing.T) {
	mesg := TestMesg{
		Type:         0x10,
		Timestamp:    timeBase.Add(32 * time.Second),
		PositionLat:  NewLatitudeDegrees(50.2053),
		PositionLong: NewLongitudeDegrees(0.1218),
	}

	def := testMesgDef()

	expect := []byte{
		0x03,
		0x10,
		0x20, 0x00, 0x00, 0x00,
		0x4A, 0x2C, 0x16, 0x00,
	}

	buf := &bytes.Buffer{}

	e := &encoder{
		w:    buf,
		arch: binary.LittleEndian,
	}

	err := e.writeMesg(reflect.ValueOf(mesg), def)
	if err != nil {
		t.Fatal(err)
	}

	if !bytes.Equal(buf.Bytes(), expect) {
		t.Errorf("Expected '%v', got '%v'", expect, buf.Bytes())
	}
}

func (a *encodeMesgDef) Equals(b *encodeMesgDef) bool {
	if a.globalMesgNum != b.globalMesgNum {
		return false
	}

	if a.localMesgNum != b.localMesgNum {
		return false
	}

	if len(a.fields) != len(b.fields) {
		return false
	}

	for i, f := range a.fields {
		if *f != *b.fields[i] {
			return false
		}
	}

	return true
}

func testFileIdMsg() FileIdMsg {
	return FileIdMsg{
		Type:         FileTypeActivity,
		Manufacturer: ManufacturerDynastream,
		Product:      uint16(GarminProductEdge25),
		SerialNumber: 0x00,     // Invalid field, should be skipped
		TimeCreated:  timeBase, // Invalid field, should be skipped
		Number:       0xffff,   // Invalid field, should be skipped
		ProductName:  "",       // Invalid field, should be skipped
	}
}

func TestGetEncodeMesgDef(t *testing.T) {
	mesg := testFileIdMsg()

	def := &encodeMesgDef{
		globalMesgNum: MesgNumFileId,
		localMesgNum:  2,
		fields: []*field{
			{
				sindex: 0,
				num:    0,
				t:      types.Fit(0),
				length: 1,
			},
			{
				sindex: 1,
				num:    1,
				t:      types.Fit(4),
				length: 1,
			},
			{
				sindex: 2,
				num:    2,
				t:      types.Fit(4),
				length: 1,
			},
		},
	}

	got := getEncodeMesgDef(reflect.ValueOf(mesg), 2)

	if !got.Equals(def) {
		t.Errorf("Expected '%+v', got '%+v'", def, got)
	}
}

func TestWriteDefMesg(t *testing.T) {
	mesg := testFileIdMsg()
	def := getEncodeMesgDef(reflect.ValueOf(mesg), 2)

	buf := &bytes.Buffer{}

	e := &encoder{
		w:    buf,
		arch: binary.LittleEndian,
	}

	err := e.writeDefMesg(def)
	if err != nil {
		t.Fatal(err)
	}

	expect := []byte{
		(1 << 6) | 2,
		0,
		0,
		byte(MesgNumFileId & 0xFF), byte(MesgNumFileId >> 8),
		3,
		0, 1, 0,
		1, 2, 4,
		2, 2, 4,
	}

	if !bytes.Equal(buf.Bytes(), expect) {
		t.Errorf("Expected '%v', got '%v'", expect, buf.Bytes())
	}
}

func TestWriteDefMesgArray(t *testing.T) {
	mesg := CapabilitiesMsg{
		Languages: []uint8{0x1},
		Sports:    []SportBits0{0x1, 0x2},
	}
	def := getEncodeMesgDef(reflect.ValueOf(mesg), 2)

	buf := &bytes.Buffer{}

	e := &encoder{
		w:    buf,
		arch: binary.LittleEndian,
	}

	err := e.writeDefMesg(def)
	if err != nil {
		t.Fatal(err)
	}

	expect := []byte{
		(1 << 6) | 2,
		0,
		0,
		byte(MesgNumCapabilities & 0xFF), byte(MesgNumCapabilities >> 8),
		2,
		0, 4, 10,
		1, 1, 10,
	}

	if !bytes.Equal(buf.Bytes(), expect) {
		t.Errorf("Expected '%v', got '%v'", expect, buf.Bytes())
	}
}
