package fit

// ProfileVersion is the current supported profile version of the FIT SDK.
const ProfileVersion uint16 = ((ProfileMajorVersion * 100) + ProfileMinorVersion)

// ProtocolVersion is the current supported FIT protocol version.
const ProtocolVersion byte = ((protocolMajorVersion << protocolVersionMajorShift) + protocolMinorVersion)

const (
	protocolMajorVersion      = 1
	protocolMinorVersion      = 0
	protocolVersionMajorShift = 4
	protocolVersionMajorMask  = (0x0F << protocolVersionMajorShift)

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

// Constants taken from offical SDK code but unused.
// TODO(tormoder): Check if they can be used somewhere or remove.
/*
const (
	mesgDefinitionReserved byte = 0x00
	maxMesgSize   byte = 255
	fieldNumInvalid byte = 255
	subfieldIndexMainField      uint16 = subfieldIndexActiveSubfield + 1
	subfieldIndexActiveSubfield        = 0xFFFE
	subfieldNameMainField       string = ""
)
*/
