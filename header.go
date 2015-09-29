package fit

import (
	"fmt"

	"github.com/tormoder/gofit/dyncrc16"
)

// Header represents a FIT file header.
type Header struct {
	Size            byte
	ProtocolVersion byte
	ProfileVersion  uint16
	DataSize        uint32
	DataType        [4]byte
	CRC             uint16
}

var (
	errNotFit     = FormatError("header data type was not '.FIT'")
	errHeaderSize = FormatError("illegal header size")
	errHdrCRC     = IntegrityError("header checksum failed")
)

func (d *decoder) decodeHeader() error {
	size, err := d.r.ReadByte()
	if err != nil {
		return err
	}
	if size != headerSizeCRC && size != headerSizeNoCRC {
		return errHeaderSize
	}
	d.h.Size = size

	if err := d.readFull(d.tmp[0 : size-1]); err != nil {
		return err
	}

	if err = checkProtocolVersion(d.tmp[0]); err != nil {
		return err
	}
	d.h.ProtocolVersion = d.tmp[0]
	d.h.ProfileVersion = le.Uint16(d.tmp[1:3])
	d.h.DataSize = le.Uint32(d.tmp[3:7])

	if string(d.tmp[7:11]) != fitDataTypeString {
		return errNotFit
	}
	copy(d.h.DataType[:], d.tmp[7:11])

	if size == headerSizeNoCRC {
		return nil
	}

	d.h.CRC = le.Uint16(d.tmp[11:13])
	if d.h.CRC == 0x0000 {
		return nil
	}

	checksum := dyncrc16.New()
	checksum.Write([]byte{size})
	checksum.Write(d.tmp[0 : size-1])

	if checksum.Sum16() != 0x0000 {
		return errHdrCRC
	}

	return nil
}

func checkProtocolVersion(b byte) error {
	if (b & protocolVersionMajorMask) > (protocolMajorVersion << protocolVersionMajorShift) {
		err := fmt.Sprintf(
			"protocol version %d.x not supported by sdk protocol version %d.%d",
			(b&protocolVersionMajorMask)>>protocolVersionMajorShift,
			protocolMajorVersion,
			protocolMinorVersion)
		return NotSupportedError(err)
	}
	return nil
}

// CheckIntegrity verifies the FIT header CRC.
func (h *Header) CheckIntegrity() error {
	if err := checkProtocolVersion(h.ProtocolVersion); err != nil {
		return err
	}
	if string(h.DataType[:len(h.DataType)]) != fitDataTypeString {
		return errNotFit
	}
	if h.Size == headerSizeNoCRC {
		return nil
	}
	if h.CRC == 0 {
		return nil
	}

	crc := dyncrc16.New()
	bh := make([]byte, h.Size)
	bh[0] = h.Size
	bh[1] = h.ProtocolVersion
	le.PutUint16(bh[2:4], h.ProfileVersion)
	le.PutUint32(bh[4:8], h.DataSize)
	copy(bh[8:12], h.DataType[:])
	le.PutUint16(bh[12:14], h.CRC)

	if crc.Sum16() != 0x0000 {
		return errHdrCRC
	}

	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (h *Header) MarshalJSON() ([]byte, error) {
	e := newJSONEncodeState()
	e.open()
	e.writeFieldName("Size")
	e.writeUint(uint64(h.Size))
	e.c()
	e.writeFieldName("ProtocolVersion")
	e.writeUint(uint64(h.ProtocolVersion))
	e.c()
	e.writeFieldName("ProfileVersion")
	e.writeUint(uint64(h.ProfileVersion))
	e.c()
	e.writeFieldName("DataSize")
	e.writeUint(uint64(h.DataSize))
	e.c()
	e.writeFieldName("DataType")
	e.writeStringBytes(h.DataType[:4])
	e.c()
	e.writeFieldName("CRC")
	e.writeUint(uint64(h.CRC))
	e.close()
	return e.Bytes(), nil
}

func (h Header) String() string {
	return fmt.Sprintf(
		"size: %d | protover: %d | profver: %d | dsize: %d | dtype: %s | crc: %d",
		h.Size, h.ProtocolVersion, h.ProfileVersion, h.DataSize, h.DataType, h.CRC,
	)
}
