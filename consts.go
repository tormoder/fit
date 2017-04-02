package fit

import "fmt"

// ProfileVersion is the current supported profile version of the FIT SDK.
const ProfileVersion uint16 = ((ProfileMajorVersion * 100) + ProfileMinorVersion)

var currentProtocolVersion = V20

// CurrentProtocolVersion returns the current supported FIT protocol version.
func CurrentProtocolVersion() ProtocolVersion {
	return currentProtocolVersion
}

// ProtocolVersion represents the FIT protocol version.
type ProtocolVersion byte

// FIT protocol versions.
const (
	V10 ProtocolVersion = 0x10
	V20 ProtocolVersion = 0x20
)

// Version returns the full FIT protocol version encoded as a single byte.
func (p ProtocolVersion) Version() byte {
	return byte(p)
}

// Major returns the major FIT protocol version.
func (p ProtocolVersion) Major() byte {
	return byte(p&protocolVersionMajorMask) >> protocolVersionMajorShift
}

// Minor returns the minor FIT protocol version.
func (p ProtocolVersion) Minor() byte {
	return byte(p & protocolVersionMinorMask)
}

func (p ProtocolVersion) String() string {
	return fmt.Sprintf("%d.%d", p.Major(), p.Minor())
}

const (
	protocolVersionMajorShift = 4
	protocolVersionMajorMask  = 0x0F << protocolVersionMajorShift
	protocolVersionMinorMask  = 0x0F
)

const (
	headerTypeMask             byte = 0xF0
	compressedHeaderMask            = 0x80
	compressedTimeMask              = 0x1F
	compressedLocalMesgNumMask      = 0x60

	mesgDefinitionMask byte = 0x40
	mesgHeaderMask          = 0x00
	localMesgNumMask        = 0x0F

	maxLocalMesgs = localMesgNumMask + 1
	maxFieldSize  = 255

	littleEndian = 0x00
	bigEndian    = 0x01

	bytesForCRC          = 2
	headerSizeCRC   byte = 14
	headerSizeNoCRC      = (headerSizeCRC - bytesForCRC)

	fitDataTypeString = ".FIT"

	fieldNumTimeStamp = 253
)
