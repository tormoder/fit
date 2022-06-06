// Package types provides the set of base types defined by the FIT protocol.
// Users of this package must validate base types parsed from a raw byte using
// the Known method before calling any other methods (except String).
package types

import (
	"fmt"
	"math"
)

const (
	pkg         = "types"
	typeNumMask = 0x1F
)

type Base byte

type baseType struct {
	size    int
	name    string
	signed  bool
	integer bool
	invalid interface{}
	// for codegen
	gotype    string
	goinvalid string
}

// Base types for fit data.
// The base 5 bits increase by 1 for each definition with all multi-byte number types having the MSB set.
const (
	BaseEnum    Base = 0x00
	BaseSint8   Base = 0x01 // 2's complement format
	BaseUint8   Base = 0x02
	BaseSint16  Base = 0x83 // 2's complement format
	BaseUint16  Base = 0x84
	BaseSint32  Base = 0x85 // 2's complement format
	BaseUint32  Base = 0x86
	BaseString  Base = 0x07 // Null terminated string encoded in UTF-8
	BaseFloat32 Base = 0x88
	BaseFloat64 Base = 0x89
	BaseUint8z  Base = 0x0A
	BaseUint16z Base = 0x8B
	BaseUint32z Base = 0x8C
	BaseByte    Base = 0x0D // Array of bytes. Field is invalid if all bytes are invalid
	BaseSint64  Base = 0x8E // 2's complement format
	BaseUint64  Base = 0x8F
	BaseUint64z Base = 0x90
)

// Internal compressed representation of certain base types.
// With this, we can fit all base types in 5 bits as opposed to 8 bits.
// Base types should be decompressed before use.
const (
	compressedSint16  byte = 0x03
	compressedUint16  byte = 0x04
	compressedSint32  byte = 0x05
	compressedUint32  byte = 0x06
	compressedFloat32 byte = 0x08
	compressedFloat64 byte = 0x09
	compressedUint16z byte = 0x0B
	compressedUint32z byte = 0x0C
	compressedSint64  byte = 0x0E
	compressedUint64  byte = 0x0F
	compressedUint64z byte = 0x10
)

func decompress(b byte) Base {
	b = b & typeNumMask
	switch b {
	case compressedSint16:
		return BaseSint16
	case compressedUint16:
		return BaseUint16
	case compressedSint32:
		return BaseSint32
	case compressedUint32:
		return BaseUint32
	case compressedFloat32:
		return BaseFloat32
	case compressedFloat64:
		return BaseFloat64
	case compressedUint16z:
		return BaseUint16z
	case compressedUint32z:
		return BaseUint32z
	case compressedSint64:
		return BaseSint64
	case compressedUint64:
		return BaseUint64
	case compressedUint64z:
		return BaseUint64z
	default:
		return Base(b)
	}
}

func (t Base) Float() bool {
	return !t.Integer() && t.Signed()
}

func (t Base) GoInvalidValue() string {
	return t.baseType().goinvalid
}

func (t Base) GoType() string {
	return t.baseType().gotype
}

func (t Base) Integer() bool {
	return t.baseType().integer
}

func (t Base) Known() bool {
	return t.baseType() != nil
}

func (t Base) PkgString() string {
	return pkg + "." + t.String()
}

func (t Base) Signed() bool {
	return t.baseType().signed
}

func (t Base) Size() int {
	return t.baseType().size
}

func (t Base) String() string {
	if b := t.baseType(); b != nil {
		return b.name
	}
	return fmt.Sprintf("unknown (0x%X)", byte(t))
}

func (t Base) Invalid() interface{} {
	if b := t.baseType(); b != nil {
		return b.invalid
	}
	return fmt.Sprintf("unknown (0x%X)", byte(t))
}

var baseTypeEnum = baseType{
	name:      "BaseEnum",
	size:      1,
	integer:   false,
	signed:    false,
	invalid:   byte(0xFF),
	gotype:    "byte",
	goinvalid: "0xFF",
}

var baseTypeSint8 = baseType{
	name:      "BaseSint8",
	size:      1,
	integer:   true,
	signed:    true,
	invalid:   int8(0x7F),
	gotype:    "int8",
	goinvalid: "0x7F",
}

var baseTypeUint8 = baseType{
	name:      "BaseUint8",
	size:      1,
	integer:   true,
	signed:    false,
	invalid:   uint8(0xFF),
	gotype:    "uint8",
	goinvalid: "0xFF",
}

var baseTypeSint16 = baseType{
	name:      "BaseSint16",
	size:      2,
	integer:   true,
	signed:    true,
	invalid:   int16(0x7FFF),
	gotype:    "int16",
	goinvalid: "0x7FFF",
}

var baseTypeUint16 = baseType{
	name:      "BaseUint16",
	size:      2,
	integer:   true,
	signed:    false,
	invalid:   uint16(0xFFFF),
	gotype:    "uint16",
	goinvalid: "0xFFFF",
}

var baseTypeSint32 = baseType{
	name:      "BaseSint32",
	size:      4,
	integer:   true,
	signed:    true,
	invalid:   int32(0x7FFFFFFF),
	gotype:    "int32",
	goinvalid: "0x7FFFFFFF",
}

var baseTypeUint32 = baseType{
	name:      "BaseUint32",
	size:      4,
	integer:   true,
	signed:    false,
	invalid:   uint32(0xFFFFFFFF),
	gotype:    "uint32",
	goinvalid: "0xFFFFFFFF",
}

var baseTypeString = baseType{
	name:      "BaseString",
	size:      1, // size doesn't really apply here
	integer:   false,
	signed:    false,
	invalid:   string(""),
	gotype:    "string",
	goinvalid: "\"\"",
}

var baseTypeFloat32 = baseType{
	name:      "BaseFloat32",
	size:      4,
	integer:   false,
	signed:    true,
	invalid:   math.Float32frombits(0xFFFFFFFF),
	gotype:    "float32",
	goinvalid: "0xFFFFFFFF",
}

var baseTypeFloat64 = baseType{
	name:      "BaseFloat64",
	size:      8,
	integer:   false,
	signed:    true,
	invalid:   math.Float64frombits(0xFFFFFFFFFFFFFFFF),
	gotype:    "float64",
	goinvalid: "0xFFFFFFFFFFFFFFFF",
}

var baseTypeUint8z = baseType{
	name:      "BaseUint8z",
	size:      1,
	integer:   true,
	signed:    false,
	invalid:   uint8(0x00),
	gotype:    "uint8",
	goinvalid: "0x00",
}

var baseTypeUint16z = baseType{
	name:      "BaseUint16z",
	size:      2,
	integer:   true,
	signed:    false,
	invalid:   uint16(0x0000),
	gotype:    "uint16",
	goinvalid: "0x0000",
}

var baseTypeUint32z = baseType{
	name:      "BaseUint32z",
	size:      4,
	integer:   true,
	signed:    false,
	invalid:   uint32(0x00000000),
	gotype:    "uint32",
	goinvalid: "0x00000000",
}

var baseTypeByte = baseType{
	name:      "BaseByte",
	size:      1,
	integer:   false,
	signed:    false,
	invalid:   byte(0xFF),
	gotype:    "byte",
	goinvalid: "0xFF",
}

var baseTypeSint64 = baseType{
	name:      "BaseSint64",
	size:      8,
	integer:   true,
	signed:    true,
	invalid:   int64(0x7FFFFFFFFFFFFFFF),
	gotype:    "int64",
	goinvalid: "0x7FFFFFFFFFFFFFFF",
}

var baseTypeUint64 = baseType{
	name:      "BaseUint64",
	size:      8,
	integer:   true,
	signed:    false,
	invalid:   uint64(0xFFFFFFFFFFFFFFFF),
	gotype:    "uint64",
	goinvalid: "0xFFFFFFFFFFFFFFFF",
}

var baseTypeUint64z = baseType{
	name:      "BaseUint64z",
	size:      8,
	integer:   true,
	signed:    false,
	invalid:   uint64(0x0000000000000000),
	gotype:    "uint64",
	goinvalid: "0x0000000000000000",
}

func (t Base) baseType() *baseType {
	switch t {
	case BaseEnum:
		return &baseTypeEnum
	case BaseSint8:
		return &baseTypeSint8
	case BaseUint8:
		return &baseTypeUint8
	case BaseSint16:
		return &baseTypeSint16
	case BaseUint16:
		return &baseTypeUint16
	case BaseSint32:
		return &baseTypeSint32
	case BaseUint32:
		return &baseTypeUint32
	case BaseString:
		return &baseTypeString
	case BaseFloat32:
		return &baseTypeFloat32
	case BaseFloat64:
		return &baseTypeFloat64
	case BaseUint8z:
		return &baseTypeUint8z
	case BaseUint16z:
		return &baseTypeUint16z
	case BaseUint32z:
		return &baseTypeUint32z
	case BaseByte:
		return &baseTypeByte
	case BaseSint64:
		return &baseTypeSint64
	case BaseUint64:
		return &baseTypeUint64
	case BaseUint64z:
		return &baseTypeUint64z
	default:
		// XXX how careful should we be about nil pointers?
		return nil
	}
}

func BaseFromString(s string) (Base, error) {
	t, found := baseStringToType[s]
	if !found {
		return 0xFF, fmt.Errorf("no base type found for string: %q", s)
	}
	return t, nil
}

var baseStringToType = map[string]Base{
	"enum":    BaseEnum,
	"sint8":   BaseSint8,
	"uint8":   BaseUint8,
	"sint16":  BaseSint16,
	"uint16":  BaseUint16,
	"sint32":  BaseSint32,
	"uint32":  BaseUint32,
	"string":  BaseString,
	"float32": BaseFloat32,
	"float64": BaseFloat64,
	"uint8z":  BaseUint8z,
	"uint16z": BaseUint16z,
	"uint32z": BaseUint32z,
	"byte":    BaseByte,
	"sint64":  BaseSint64,
	"uint64":  BaseUint64,
	"uint64z": BaseUint64z,

	// Typo in SDK 20.14:
	"unit8": BaseUint8,
}
