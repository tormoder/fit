package fit

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math"
	"os"
	"reflect"
	"strconv"
	"time"

	"github.com/tormoder/fit/dyncrc16"
	"github.com/tormoder/fit/internal/base"
)

var debug, _ = strconv.ParseBool(os.Getenv("FIT_DEBUG"))

func init() {
	log.SetPrefix("fit: ")
	log.SetFlags(0)
}

var (
	le = binary.LittleEndian
	be = binary.BigEndian
)

type reader interface {
	io.Reader
	io.ByteReader
}

type decoder struct {
	r       reader
	n       uint32
	crc     dyncrc16.Hash16
	tmp     [maxFieldSize]byte
	defmsgs [maxLocalMesgs]*defmsg

	timestamp      uint32
	lastTimeOffset int32

	opts decodeOptions

	h   Header
	fit *Fit
}

// CheckIntegrity verifies the FIT header and file CRC. Only the header CRC is
// verified if headerOnly is true.
func CheckIntegrity(r io.Reader, headerOnly bool) error {
	var d decoder
	return d.decode(r, headerOnly, false, true)
}

// DecodeHeader returns the FIT file header without decoding the entire FIT
// file.
func DecodeHeader(r io.Reader) (Header, error) {
	var d decoder
	if err := d.decode(r, true, false, false); err != nil {
		return Header{}, err
	}
	return d.h, nil
}

// DecodeHeaderAndFileID returns the FIT file header and FileId message without
// decoding the entire FIT file. The FileId message must be present in all FIT
// files.
func DecodeHeaderAndFileID(r io.Reader) (Header, FileIdMsg, error) {
	var d decoder
	if err := d.decode(r, false, true, false); err != nil {
		return Header{}, FileIdMsg{}, err
	}
	return d.h, d.fit.FileId, nil
}

// Decode reads a FIT file from r and returns it as a *Fit.
func Decode(r io.Reader, opts ...DecodeOption) (*Fit, error) {
	var d decoder
	for _, opt := range opts {
		opt(&d.opts)
	}
	if err := d.decode(r, false, false, false); err != nil {
		return nil, err
	}
	return d.fit, nil
}

func (d *decoder) decode(r io.Reader, headerOnly, fileIDOnly, crcOnly bool) error {
	d.crc = dyncrc16.New()
	tr := io.TeeReader(r, d.crc)

	// Add buffering if r does not provide ReadByte.
	if rr, ok := tr.(reader); ok {
		d.r = rr
	} else {
		d.r = bufio.NewReader(tr)
	}

	err := d.decodeHeader()
	if err != nil {
		return fmt.Errorf("error decoding header: %v", err)
	}

	d.fit = new(Fit)
	d.fit.Header = d.h

	if debug {
		log.Println("header decoded:", d.h)
	}

	if headerOnly {
		return nil
	}

	if crcOnly {
		_, err = io.CopyN(ioutil.Discard, d.r, int64(d.h.DataSize))
		if err != nil {
			return fmt.Errorf("error parsing data: %v", err)
		}
		goto crc
	}

	d.fit.UnknownMessages = make(map[MesgNum]int)
	d.fit.UnknownFields = make(map[UnknownField]int)

	err = d.parseFileIdMsg()
	if err != nil {
		return fmt.Errorf("error parsing file id message: %v", err)
	}
	if fileIDOnly {
		return nil
	}

	err = d.initFileType()
	if err != nil {
		return err
	}

	for d.n < d.h.DataSize-2 {
		var (
			b   byte
			dm  *defmsg
			msg reflect.Value
		)

		b, err = d.readByte()
		if err != nil {
			return fmt.Errorf("error parsing record header: %v", err)
		}

		switch {

		case (b & compressedHeaderMask) == compressedHeaderMask:
			msg, err = d.parseDataMessage(b, true)
			if err != nil {
				return fmt.Errorf("compressed timestamp message: %v", err)
			}
			if msg.IsValid() {
				d.fit.add(msg)
			}
		case (b & headerTypeMask) == mesgDefinitionMask:
			dm, err = d.parseDefinitionMessage(b)
			if err != nil {
				return fmt.Errorf("parsing definition message: %v", err)
			}
			d.defmsgs[dm.localMsgType] = dm
		case (b & mesgHeaderMask) == mesgHeaderMask:
			msg, err = d.parseDataMessage(b, false)
			if err != nil {
				return fmt.Errorf("parsing data message: %v", err)
			}
			if msg.IsValid() {
				d.fit.add(msg)
			}
		default:
			return fmt.Errorf("unknown record header, got: %#x", b)
		}
	}

crc:
	if err = binary.Read(d.r, binary.LittleEndian, &d.fit.CRC); err != nil {
		err = noEOF(err)
		return fmt.Errorf("error parsing file CRC: %v", err)
	}

	if d.crc.Sum16() != 0x0000 {
		return IntegrityError("file checksum failed")
	}

	return nil
}

func (d *decoder) readByte() (c byte, err error) {
	c, err = d.r.ReadByte()
	if err == nil {
		d.n++
		return c, nil
	}
	err = noEOF(err)
	return c, err
}

func (d *decoder) skipByte() error {
	_, err := d.readByte()
	return err
}

func (d *decoder) readFull(buf []byte) error {
	n, err := io.ReadFull(d.r, buf)
	if err == nil {
		d.n += uint32(n)
		return nil
	}
	err = noEOF(err)
	return err
}

type defmsg struct {
	localMsgType uint8
	arch         binary.ByteOrder
	globalMsgNum MesgNum
	fields       byte
	fieldDefs    []fieldDef
}

func (dm defmsg) String() string {
	return fmt.Sprintf(
		"local: %d | global: %v | arch: %v | fields: %d",
		dm.localMsgType, dm.globalMsgNum, dm.arch, dm.fields,
	)
}

type fieldDef struct {
	num   byte
	size  byte
	btype base.Type
}

func (fd fieldDef) String() string {
	return fmt.Sprintf("num: %d | size: %d | btype: %v", fd.num, fd.size, fd.btype)
}

func (d *decoder) parseFileIdMsg() error {
	b, err := d.readByte()
	if err != nil {
		return fmt.Errorf("error parsing record header: %v", err)
	}

	if !((b & mesgDefinitionMask) == mesgDefinitionMask) {
		return fmt.Errorf("expected record header byte for definition message, got %#x - %8b", b, b)
	}

	dm, err := d.parseDefinitionMessage(b)
	if err != nil {
		return fmt.Errorf("error parsing definition message: %v", err)
	}
	if dm.globalMsgNum != MesgNumFileId {
		return errors.New("parsed definiton message was not for file_id")
	}
	d.defmsgs[dm.localMsgType] = dm

	b, err = d.readByte()
	if err != nil {
		return fmt.Errorf("error parsing record header: %v", err)
	}

	if !((b & mesgHeaderMask) == mesgHeaderMask) {
		return fmt.Errorf("expected record header byte for data message, got %#x - %8b", b, b)
	}
	msg, err := d.parseDataMessage(b, false)
	if err != nil {
		return fmt.Errorf("error reading data message:  %v", err)
	}

	_, ok := msg.Interface().(FileIdMsg)
	if !ok {
		return errors.New("parsed message was not of type file_id")
	}

	d.fit.add(msg)

	return nil
}

func (d *decoder) initFileType() error {
	t := d.fit.FileId.Type
	switch t {
	case FileActivity:
		d.fit.activity = new(ActivityFile)
		d.fit.msgAdder = d.fit.activity
	case FileDevice:
		d.fit.device = new(DeviceFile)
		d.fit.msgAdder = d.fit.device
	case FileSettings:
		d.fit.settings = new(SettingsFile)
		d.fit.msgAdder = d.fit.settings
	case FileSport:
		d.fit.sport = new(SportFile)
		d.fit.msgAdder = d.fit.sport
	case FileWorkout:
		d.fit.workout = new(WorkoutFile)
		d.fit.msgAdder = d.fit.workout
	case FileCourse:
		d.fit.course = new(CourseFile)
		d.fit.msgAdder = d.fit.course
	case FileSchedules:
		d.fit.schedules = new(SchedulesFile)
		d.fit.msgAdder = d.fit.schedules
	case FileWeight:
		d.fit.weight = new(WeightFile)
		d.fit.msgAdder = d.fit.weight
	case FileTotals:
		d.fit.totals = new(TotalsFile)
		d.fit.msgAdder = d.fit.totals
	case FileGoals:
		d.fit.goals = new(GoalsFile)
		d.fit.msgAdder = d.fit.goals
	case FileBloodPressure:
		d.fit.bloodPressure = new(BloodPressureFile)
		d.fit.msgAdder = d.fit.bloodPressure
	case FileMonitoringA:
		d.fit.monitoringA = new(MonitoringAFile)
		d.fit.msgAdder = d.fit.monitoringA
	case FileActivitySummary:
		d.fit.activitySummary = new(ActivitySummaryFile)
		d.fit.msgAdder = d.fit.activitySummary
	case FileMonitoringDaily:
		d.fit.monitoringDaily = new(MonitoringDailyFile)
		d.fit.msgAdder = d.fit.monitoringDaily
	case FileMonitoringB:
		d.fit.monitoringB = new(MonitoringBFile)
		d.fit.msgAdder = d.fit.monitoringB
	case FileSegment:
		d.fit.segment = new(SegmentFile)
		d.fit.msgAdder = d.fit.segment
	case FileSegmentList:
		d.fit.segmentList = new(SegmentListFile)
		d.fit.msgAdder = d.fit.segmentList
	case FileInvalid:
		return FormatError("file type was set invalid")
	default:
		switch {
		case t > FileMonitoringB && t < FileMfgRangeMin:
			return FormatError(
				fmt.Sprintf("unknown file type: %v", t),
			)
		case t >= FileMfgRangeMin && t <= FileMfgRangeMax:
			return NotSupportedError("manufacturer specific file types")
		default:
			return FormatError(
				fmt.Sprintf("unknown file type: %v", t),
			)
		}
	}
	return nil
}

func (d *decoder) parseDefinitionMessage(recordHeader byte) (*defmsg, error) {
	dm := defmsg{}
	dm.localMsgType = recordHeader & localMesgNumMask
	if dm.localMsgType > localMesgNumMask {
		if debug {
			log.Printf("illegal local message number: %d\n", dm.localMsgType)
		}
		return nil, FormatError("illegal local message number")
	}

	// next byte reserved
	if err := d.skipByte(); err != nil {
		return nil, err
	}

	arch, err := d.readByte()
	if err != nil {
		return nil, err
	}

	switch arch {
	case littleEndian:
		dm.arch = le
	case bigEndian:
		dm.arch = be
	default:
		return nil, fmt.Errorf("unknow arch: %#x", arch)
	}

	if err = d.readFull(d.tmp[:2]); err != nil {
		return nil, fmt.Errorf("error parsing global message number: %v", err)
	}
	dm.globalMsgNum = MesgNum(dm.arch.Uint16(d.tmp[:2]))
	if dm.globalMsgNum == MesgNumInvalid {
		return nil, FormatError("global message number was set invalid")
	}

	dm.fields, err = d.readByte()
	if err != nil {
		return nil, err
	}
	if dm.fields == 0 {
		return &dm, nil
	}

	if err = d.readFull(d.tmp[0 : 3*dm.fields]); err != nil {
		return nil, fmt.Errorf("error parsing fields: %v", err)
	}

	dm.fieldDefs = make([]fieldDef, dm.fields)
	for i, fd := range dm.fieldDefs {
		fd.num = d.tmp[i*3]
		fd.size = d.tmp[(i*3)+1]
		fd.btype = base.Type(d.tmp[(i*3)+2])
		if err = d.validateFieldDef(dm.globalMsgNum, fd); err != nil {
			return nil, fmt.Errorf(
				"validating %v failed: %v",
				dm.globalMsgNum, err)
		}
		dm.fieldDefs[i] = fd
	}

	if debug {
		log.Println("definition messages parsed:", dm)
	}

	return &dm, nil
}

func (d *decoder) validateFieldDef(gmsgnum MesgNum, dfield fieldDef) error {
	if !dfield.btype.Known() {
		return fmt.Errorf("field %d: unknown base type 0X%X", dfield.num, dfield.btype)
	}

	var pfield *field
	pfound := false
	if knownMsgNums[gmsgnum] {
		pfield, pfound = getField(gmsgnum, dfield.num)
	}

	if dfield.btype == base.String {
		if !pfound {
			return nil
		}
		if pfield.btype == dfield.btype {
			return nil
		}
		return fmt.Errorf(
			"field %d: field base type is string, but profile lists it as %v, not compatible",
			dfield.num, pfield.btype)
	}

	// Verify that field definition size is not less than field definition
	// base type size.
	if int(dfield.size) < dfield.btype.Size() {
		return fmt.Errorf(
			"field %d: size (%d) is less than base type size (%d)",
			dfield.num, dfield.size, dfield.btype.Size())
	}

	if !pfound {
		return nil
	}

	// Profile field.
	if pfield.array == 0 {

		// Profile field not an array. Verify that the field size is
		// not greater than the profile base type size. A smaller size
		// is allowed due to dynamic fields.
		switch {

		case int(dfield.size) > pfield.btype.Size():
			return fmt.Errorf(
				"field %d: size (%d) is greater than size of profile base type %v (%d)",
				dfield.num, dfield.size, dfield.btype, dfield.btype.Size())

		case int(dfield.size) <= pfield.btype.Size() && dfield.btype != pfield.btype:
			// Size is less or equal, but we can only allow
			// "compatible" types that will not panic when setting
			// fields using reflection.
			switch {
			case pfield.btype.Signed() != dfield.btype.Signed():
				fallthrough
			case dfield.btype.Float() && !pfield.btype.Float():
				fallthrough
			case pfield.btype == base.String && dfield.btype != base.String:
				return fmt.Errorf(
					"field %d: type %v is not compatible with profile type %v",
					dfield.num, dfield.btype, pfield.btype)
			}
		}

		return nil
	}

	// Profile field is an array.
	switch {
	case (int(dfield.size) % dfield.btype.Size()) != 0:
		return fmt.Errorf(
			"field %d: array, but size (%d) is not a multiple of base type %v size (%d)",
			dfield.num, dfield.size, dfield.btype, dfield.btype.Size())
	case dfield.btype != pfield.btype:
		// Require correct base type if an array. I have not seen a
		// dynamic field that is an array and have a smaller base type
		// for array elements. Maybe allow equal sized compatible types
		// later if needed (like for non-array fields).
		return fmt.Errorf(
			"field %d: array, but definition (%v) and profile (%v) base types differ",
			dfield.num, dfield.btype, dfield.btype.Size())
	default:
		return nil
	}
}

func (d *decoder) parseDataMessage(recordHeader byte, compressed bool) (reflect.Value, error) {
	var localMsgNum byte
	if compressed {
		localMsgNum = (recordHeader & compressedLocalMesgNumMask) >> 5
	} else {
		localMsgNum = recordHeader & localMesgNumMask
	}

	dm := d.defmsgs[localMsgNum]
	if dm == nil {
		return reflect.Value{}, fmt.Errorf(
			"missing data definition message for local message number %d",
			localMsgNum)
	}

	var msgv reflect.Value
	knownMsg := knownMsgNums[dm.globalMsgNum]
	if knownMsg {
		msgv = getMesgAllInvalid(dm.globalMsgNum)
	} else {
		d.fit.UnknownMessages[dm.globalMsgNum]++
	}

	if !compressed {
		return d.parseDataFields(dm, knownMsg, msgv)
	}

	// Data message has compressed timestamp header.
	if d.timestamp == 0 {
		if debug {
			log.Println(
				"warning: parsing compressed timestamp",
				"header, but have no previous reference",
				"time, skipping setting timestamp for message",
			)
		}
		return d.parseDataFields(dm, knownMsg, msgv)
	}

	timeOffset := int32(recordHeader & compressedTimeMask)
	d.timestamp += uint32((timeOffset - d.lastTimeOffset) & compressedTimeMask)
	d.lastTimeOffset = timeOffset

	fieldTimestamp, found := getField(dm.globalMsgNum, fieldNumTimeStamp)
	if found {
		fieldval := msgv.Field(fieldTimestamp.sindex)
		t := decodeDateTime(d.timestamp)
		fieldval.Set(reflect.ValueOf(t))
		return d.parseDataFields(dm, knownMsg, msgv)
	}

	if debug {
		log.Println(
			"warning: parsing message with compressed timestamp header,",
			"but did not find timestamp field in message of type", dm.globalMsgNum)

	}

	return d.parseDataFields(dm, knownMsg, msgv)
}

func (d *decoder) parseDataFields(dm *defmsg, knownMsg bool, msgv reflect.Value) (reflect.Value, error) {
	for i, dfield := range dm.fieldDefs {

		dsize := int(dfield.size)
		padding := 0

		pfield, pfound := getField(dm.globalMsgNum, dfield.num)
		if pfound {
			if pfield.btype != base.String && pfield.array == 0 {
				padding = pfield.btype.Size() - dsize
			}
		} else {
			d.fit.UnknownFields[UnknownField{dm.globalMsgNum, dfield.num}]++
		}

		err := d.readFull(d.tmp[0:dsize])
		if err != nil {
			return reflect.Value{}, fmt.Errorf(
				"error parsing data message: %v (field %d [%v] for [%v])",
				err, i, dfield, dm)
		}

		if padding != 0 {
			if dm.arch == le {
				for j := dsize; j < pfield.btype.Size(); j++ {
					d.tmp[j] = 0x00
				}
			} else {
				for j := 0; j < pfield.btype.Size(); j++ {
					d.tmp[j], d.tmp[j+padding] = 0x00, d.tmp[j]
				}
			}
		}

		if !knownMsg || !pfound {
			continue
		}

		fieldv := msgv.Field(pfield.sindex)

		switch pfield.t {
		case fit:
			if pfield.array == 0 {
				err = d.parseFitField(dm, dfield, fieldv)
			} else {
				err = d.parseFitFieldArray(dm, dfield, fieldv)
			}
			if err == nil {
				continue
			}
			return reflect.Value{}, fmt.Errorf("error parsing data message: %v", err)
		case timeutc, timelocal:
			d.parseTimeStamp(dm, fieldv, pfield)
		case lat:
			i32 := dm.arch.Uint32(d.tmp[:base.Sint32.Size()])
			lat := NewLatitude(int32(i32))
			fieldv.Set(reflect.ValueOf(lat))
		case lng:
			i32 := dm.arch.Uint32(d.tmp[:base.Sint32.Size()])
			lng := NewLongitude(int32(i32))
			fieldv.Set(reflect.ValueOf(lng))
		default:
			panic("parseDataFields: unreachable")
		}
	}

	return msgv, nil
}

func (d *decoder) parseFitField(dm *defmsg, dfield fieldDef, fieldv reflect.Value) error {
	dsize := int(dfield.size)
	switch dfield.btype {
	case base.Byte, base.Enum, base.Uint8, base.Uint8z:
		fieldv.SetUint(uint64(d.tmp[0]))
	case base.Sint8:
		fieldv.SetInt(int64(d.tmp[0]))
	case base.Sint16:
		i16 := int64(dm.arch.Uint16(d.tmp[:dsize]))
		fieldv.SetInt(i16)
	case base.Uint16, base.Uint16z:
		u16 := uint64(dm.arch.Uint16(d.tmp[:dsize]))
		fieldv.SetUint(u16)
	case base.Sint32:
		i32 := int64(dm.arch.Uint32(d.tmp[:dsize]))
		fieldv.SetInt(i32)
	case base.Uint32, base.Uint32z:
		u32 := uint64(dm.arch.Uint32(d.tmp[:dsize]))
		fieldv.SetUint(u32)
	case base.Float32:
		bits := dm.arch.Uint32(d.tmp[:dsize])
		f32 := float64(math.Float32frombits(bits))
		fieldv.SetFloat(f32)
	case base.Float64:
		bits := dm.arch.Uint64(d.tmp[:dsize])
		f64 := math.Float64frombits(bits)
		fieldv.SetFloat(f64)
	case base.String:
		for j := 0; j < dsize; j++ {
			if d.tmp[j] == 0x00 {
				if j > 0 {
					fieldv.SetString(string(d.tmp[:j]))
				}
				break
			}
			if j == dsize-1 {
				fieldv.SetString(string(d.tmp[:j]))
			}
		}
	default:
		return fmt.Errorf(
			"unknown base type %d for field %v in definition message %v",
			dfield.btype, dfield, dm)
	}

	return nil
}

func (d *decoder) parseFitFieldArray(dm *defmsg, dfield fieldDef, fieldv reflect.Value) error {
	dbt := dfield.btype
	dsize := int(dfield.size)

	if dbt == base.Byte {
		byteArray := make([]byte, dsize, dsize)
		copy(byteArray, d.tmp[:dsize])
		fieldv.SetBytes(byteArray)
		return nil
	}

	slicev := reflect.MakeSlice(
		fieldv.Type(),
		dsize/dbt.Size(),
		dsize/dbt.Size(),
	)

	switch dbt {
	case base.Uint8, base.Uint8z, base.Enum:
		for j := 0; j < dsize; j++ {
			slicev.Index(j).SetUint(uint64(d.tmp[j]))
		}
	case base.Sint8:
		for j := 0; j < dsize; j++ {
			slicev.Index(j).SetInt(int64(d.tmp[j]))
		}
	case base.Sint16:
		for j, k := 0, 0; j < dsize; j, k = j+dbt.Size(), k+1 {
			i16 := int64(dm.arch.Uint16(d.tmp[j : j+dbt.Size()]))
			slicev.Index(k).SetInt(i16)
		}
	case base.Uint16, base.Uint16z:
		for j, k := 0, 0; j < dsize; j, k = j+dbt.Size(), k+1 {
			ui16 := uint64(dm.arch.Uint16(d.tmp[j : j+dbt.Size()]))
			slicev.Index(k).SetUint(ui16)
		}
	case base.Sint32:
		for j, k := 0, 0; j < dsize; j, k = j+dbt.Size(), k+1 {
			i32 := int64(dm.arch.Uint32(d.tmp[j : j+dbt.Size()]))
			slicev.Index(k).SetInt(i32)
		}
	case base.Uint32, base.Uint32z:
		for j, k := 0, 0; j < dsize; j, k = j+dbt.Size(), k+1 {
			ui32 := uint64(dm.arch.Uint32(d.tmp[j : j+dbt.Size()]))
			slicev.Index(k).SetUint(ui32)
		}
	case base.Float32:
		for j, k := 0, 0; j < dsize; j, k = j+dbt.Size(), k+1 {
			bits := dm.arch.Uint32(d.tmp[j : j+dbt.Size()])
			f32 := float64(math.Float32frombits(bits))
			slicev.Index(k).SetFloat(f32)
		}
	case base.Float64:
		for j, k := 0, 0; j < dsize; j, k = j+dbt.Size(), k+1 {
			bits := dm.arch.Uint64(d.tmp[j : j+dbt.Size()])
			f64 := math.Float64frombits(bits)
			slicev.Index(k).SetFloat(f64)
		}
	case base.String:
		if dfield.size == 0 {
			return nil
		}
		var strings []string
		j, k := 0, 0
		for {
			if d.tmp[j+k] == 0x00 {
				if k == 0 {
					break
				}
				strings = append(strings, string(d.tmp[j:j+k]))
				j = j + k + 1
				if j >= dsize {
					break
				}
				k = 0
			} else {
				k++
				if j+k >= dsize {
					// We have not seen a 0x00 terminator,
					// but there's no room for one.
					// Take the string we have and exit loop.
					strings = append(strings, string(d.tmp[j:dsize]))
					break
				}
			}
		}
		fieldv.Set(reflect.ValueOf(strings))
		return nil // We don't want the Set after the switch.
	default:
		return fmt.Errorf(
			"unknown base type %d for field %v in definition message %v",
			dbt, dfield, dm)
	}

	fieldv.Set(slicev)
	return nil
}

func (d *decoder) parseTimeStamp(dm *defmsg, fieldv reflect.Value, pfield *field) {
	u32 := dm.arch.Uint32(d.tmp[:base.Uint32.Size()])
	if u32 == 0xFFFFFFFF {
		return
	}
	if u32 < systemTimeMarker {
		if debug {
			log.Println("parsing time: seconds from device power on")
		}
	}

	if pfield.t == timeutc {
		if pfield.num == fieldNumTimeStamp {
			d.timestamp = u32
			d.lastTimeOffset = int32(d.timestamp & compressedTimeMask)
		}
		t := decodeDateTime(u32)
		fieldv.Set(reflect.ValueOf(t))
		return
	}

	// Local timestamp.
	//
	// Use a custom timezone with the calculated offset to indicate that it
	// is not UTC.
	//
	// Also see the SetLocalTimeZone function in the timeutil subpackage.
	// For now not used due to an external dependency.
	var local time.Time
	switch {
	case d.timestamp == 0, d.timestamp < systemTimeMarker:
		// No time reference.
		// Set local with zero offset.
		d.timestamp = u32
		tzone := time.FixedZone(localZoneName, 0)
		local = decodeDateTime(u32)
		local = local.In(tzone)
	default:
		local = decodeDateTime(u32)
		utc := decodeDateTime(d.timestamp)
		offsetDur := local.Sub(utc)
		tzone := time.FixedZone(localZoneName, int(offsetDur.Seconds()))
		local = utc.In(tzone)
	}
	fieldv.Set(reflect.ValueOf(local))
}

func noEOF(err error) error {
	if err == io.EOF {
		return io.ErrUnexpectedEOF
	}
	return err
}
