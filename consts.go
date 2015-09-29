package fit

const (
	protocolVersion           byte = ((protocolMajorVersion << protocolVersionMajorShift) + protocolMinorVersion)
	protocolMajorVersion           = 1
	protocolMinorVersion           = 0
	protocolVersionMajorShift      = 4
	protocolVersionMajorMask       = (0x0F << protocolVersionMajorShift)

	profileVersion      uint16 = ((profileMajorVersion * 100) + profileMinorVersion)
	profileMajorVersion        = 14
	profileMinorVersion        = 0

	headerTypeMask             byte = 0xF0
	compressedHeaderMask            = 0x80
	compressedTimeMask              = 0x1F
	compressedLocalMesgNumMask      = 0x60

	mesgDefinitionMask byte = 0x40
	mesgHeaderMask          = 0x00
	localMesgNumMask        = 0x0F
	maxLocalMesgs           = localMesgNumMask + 1

	mesgDefinitionReserved byte = 0x00
	littleEndian                = 0x00
	bigEndian                   = 0x01

	maxMesgSize  byte = 255
	maxFieldSize      = 255

	bytesForCRC          = 2
	headerSizeCRC   byte = 14
	headerSizeNoCRC      = (headerSizeCRC - bytesForCRC)

	fitDataTypeString = ".FIT"

	fieldNumInvalid   byte = 255
	fieldNumTimeStamp      = 253

	subfieldIndexMainField      uint16 = subfieldIndexActiveSubfield + 1
	subfieldIndexActiveSubfield        = 0xFFFE
	subfieldNameMainField       string = ""
)
