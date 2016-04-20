package base

import "fmt"

const (
	pkg         = "base"
	typeNumMask = 0x1F
)

type Type byte

const (
	Enum    Type = 0x00
	Sint8   Type = 0x01 // 2's complement format
	Uint8   Type = 0x02
	Sint16  Type = 0x83 // 2's complement format
	Uint16  Type = 0x84
	Sint32  Type = 0x85 // 2's complement format
	Uint32  Type = 0x86
	String  Type = 0x07 // Null terminated string encoded in UTF-8
	Float32 Type = 0x88
	Float64 Type = 0x89
	Uint8z  Type = 0x0A
	Uint16z Type = 0x8B
	Uint32z Type = 0x8C
	Byte    Type = 0x0D // Array of bytes. Field is invalid if all bytes are invalid
)

func (t Type) nr() int {
	return int(t & typeNumMask)
}

func (t Type) Float() bool {
	return !t.Integer() && t.Signed()
}

func (t Type) GoInvalidValue() string {
	return tinvalid[t.nr()]
}

func (t Type) GoType() string {
	return tgotype[t.nr()]
}

func (t Type) Integer() bool {
	return tinteger[t.nr()]
}

func (t Type) Known() bool {
	return t.nr() < len(tname) && t.nr() >= 0
}

func (t Type) PkgString() string {
	return pkg + "." + t.String()
}

func (t Type) Signed() bool {
	return tsigned[t.nr()]
}

func (t Type) Size() int {
	return tsize[t.nr()]
}

func (t Type) String() string {
	return tname[t.nr()]
}

var tsize = [...]int{
	1,
	1,
	1,
	2,
	2,
	4,
	4,
	1,
	4,
	8,
	1,
	2,
	4,
	1,
}

var tname = [...]string{
	"Enum",
	"Sint8",
	"Uint8",
	"Sint16",
	"Uint16",
	"Sint32",
	"Uint32",
	"String",
	"Float32",
	"Float64",
	"Uint8z",
	"Uint16z",
	"Uint32z",
	"Byte",
}

var tinteger = [...]bool{
	false,
	true,
	true,
	true,
	true,
	true,
	true,
	false,
	false,
	false,
	true,
	true,
	true,
	false,
}

var tsigned = [...]bool{
	false,
	true,
	false,
	true,
	false,
	true,
	false,
	false,
	true,
	true,
	false,
	false,
	false,
	false,
}

var tgotype = [...]string{
	"byte",
	"int8",
	"uint8",
	"int16",
	"uint16",
	"int32",
	"uint32",
	"string",
	"float32",
	"float64",
	"uint8",
	"uint16",
	"uint32",
	"byte",
}

var tinvalid = [...]string{
	"0xFF",
	"0x7F",
	"0xFF",
	"0x7FFF",
	"0xFFFF",
	"0x7FFFFFFF",
	"0xFFFFFFFF",
	"\"\"",
	"0xFFFFFFFF",
	"0xFFFFFFFFFFFFFFFF",
	"0x00",
	"0x0000",
	"0x00000000",
	"0xFF",
}

func TypeFromString(s string) (Type, error) {
	t, found := baseStringToType[s]
	if !found {
		return 0xFF, fmt.Errorf("no base type found for string: %q", s)
	}
	return t, nil
}

var baseStringToType = map[string]Type{
	"enum":    Enum,
	"sint8":   Sint8,
	"uint8":   Uint8,
	"sint16":  Sint16,
	"uint16":  Uint16,
	"sint32":  Sint32,
	"uint32":  Uint32,
	"string":  String,
	"float32": Float32,
	"float64": Float64,
	"uint8z":  Uint8z,
	"uint16z": Uint16z,
	"uint32z": Uint32z,
	"byte":    Byte,
}
