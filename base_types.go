package fit

const baseTypeNumMask = 0x1F

type fitBaseType byte

const (
	fitEnum    fitBaseType = 0x00
	fitSint8   fitBaseType = 0x01 // 2's complement format
	fitUint8   fitBaseType = 0x02
	fitSint16  fitBaseType = 0x83 // 2's complement format
	fitUint16  fitBaseType = 0x84
	fitSint32  fitBaseType = 0x85 // 2's complement format
	fitUint32  fitBaseType = 0x86
	fitString  fitBaseType = 0x07 // Null terminated string encoded in UTF-8
	fitFloat32 fitBaseType = 0x88
	fitFloat64 fitBaseType = 0x89
	fitUint8z  fitBaseType = 0x0A
	fitUint16z fitBaseType = 0x8B
	fitUint32z fitBaseType = 0x8C
	fitByte    fitBaseType = 0x0D // Array of bytes. Field is invalid if all bytes are invalid
)

func (b fitBaseType) nr() int {
	return int(b & baseTypeNumMask)
}

func (b fitBaseType) endianAbility() bool {
	return bsize[b.nr()] > 0
}

func (b fitBaseType) integer() bool {
	return binteger[b.nr()]
}

func (b fitBaseType) signed() bool {
	return bsigned[b.nr()]
}

func (b fitBaseType) size() int {
	return bsize[b.nr()]
}

func (b fitBaseType) String() string {
	return bname[b.nr()]
}

var fitBaseTypes = [...]string{
	"enum",
	"sint8",
	"uint8",
	"sint16",
	"uint16",
	"sint32",
	"uint32",
	"string",
	"float32",
	"float64",
	"uint8z",
	"uint16z",
	"uint32z",
	"byte",
}

var bsize = [...]int{
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

var bname = [...]string{
	"enum",
	"sint8",
	"uint8",
	"sint16",
	"uint16",
	"sint32",
	"uint32",
	"string",
	"float32",
	"float64",
	"uint8z",
	"uint16z",
	"uint32z",
	"byte",
}

var binteger = [...]bool{
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

var bsigned = [...]bool{
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
