package fit

import (
	"bytes"
	"testing"

	"github.com/tormoder/fit/dyncrc16"
)

var decodeHeaderTests = []struct {
	in  []byte
	err error
	h   Header
}{
	{[]byte{11}, errHeaderSize, Header{}},
	{[]byte{12}, errReadData, Header{}},
	{[]byte{13}, errHeaderSize, Header{}},
	{[]byte{14}, errReadData, Header{}},
	{[]byte{15}, errHeaderSize, Header{}},
	{[]byte{12, 0}, errReadData, Header{}},
	{[]byte{14, 0}, errReadData, Header{}},
	{
		[]byte{14, 0x30, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		NotSupportedError("protocol version 3.0 not supported by sdk protocol version 2.0"),
		Header{},
	},
	{[]byte{14, 0, 0, 0, 0, 0, 0, 0, '.', 'G', 'I', 'T', 0, 0}, errNotFit, Header{}},
	{
		[]byte{14, 0, 0, 0, 0, 0, 0, 0, '.', 'F', 'I', 'T', 0, 0},
		nil,
		Header{
			Size:     14,
			DataType: [4]byte{'.', 'F', 'I', 'T'},
		},
	},
	{
		[]byte{14, 0x10, 0, 0, 0, 0, 0, 0, '.', 'F', 'I', 'T', 0, 0},
		nil,
		Header{
			Size:            14,
			ProtocolVersion: 0x10,
			DataType:        [4]byte{'.', 'F', 'I', 'T'},
		},
	},
	{
		[]byte{14, 0x10, 0x57, 0x04, 0, 0, 0, 0, '.', 'F', 'I', 'T', 0, 0},
		nil,
		Header{
			Size:            14,
			ProtocolVersion: 0x10,
			ProfileVersion:  1111,
			DataType:        [4]byte{'.', 'F', 'I', 'T'},
		},
	},
	{
		[]byte{14, 0x10, 0x57, 0x04, 0x63, 0x6f, 0x01, 0x00, '.', 'F', 'I', 'T', 0, 0},
		nil,
		Header{
			Size:            14,
			ProtocolVersion: 0x10,
			ProfileVersion:  1111,
			DataSize:        94051,
			DataType:        [4]byte{'.', 'F', 'I', 'T'},
		},
	},
	{
		[]byte{14, 0x10, 0x57, 0x04, 0x63, 0x6f, 0x01, 0x00, '.', 'F', 'I', 'T', 0x3b, 0x34},
		nil,
		Header{
			Size:            14,
			ProtocolVersion: 0x10,
			ProfileVersion:  1111,
			DataSize:        94051,
			DataType:        [4]byte{'.', 'F', 'I', 'T'},
			CRC:             13371,
		},
	},
	{
		[]byte{14, 0x10, 0x57, 0x04, 0x63, 0x6f, 0x01, 0x00, '.', 'F', 'I', 'T', 0x3b, 0x35},
		errHdrCRC,
		Header{},
	},
}

func TestDecodeHeader(t *testing.T) {
	for i, dht := range decodeHeaderTests {
		var dec decoder
		dec.r = bytes.NewReader(dht.in)
		dec.crc = dyncrc16.New()
		err := dec.decodeHeader()
		if err != dht.err {
			t.Errorf("%d: got error: %v, want: %v", i, err, dht.err)
			continue
		}
		if dht.err != nil {
			continue
		}
		if dec.h != dht.h {
			t.Errorf("%d:\ngot header:\n%v\nwant header\n%v", i, dec.h, dht.h)
		}
	}
}
