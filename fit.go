package fit

import (
	"fmt"
	"reflect"
)

// Fit represents a decoded FIT file.
type Fit struct {
	// Header is the FIT file header.
	Header Header

	// CRC is the FIT file CRC.
	CRC uint16

	// FileId is a message required for all FIT files.
	FileId FileIdMsg

	// Common messages for all FIT file types.
	FileCreator          *FileCreatorMsg
	TimestampCorrelation *TimestampCorrelationMsg
	DeviceInfo           *DeviceInfoMsg

	// UnknownMessages is a map that maps an unknown message number to how
	// many times the message was encountered during encoding.
	UnknownMessages map[MesgNum]int

	// UnknownFields is a map that maps an unknown field to how many times
	// the field was encountered during encoding.
	UnknownFields map[UnknownField]int

	msgAdder msgAdder

	activity        *ActivityFile
	device          *DeviceFile
	settings        *SettingsFile
	sport           *SportFile
	workout         *WorkoutFile
	course          *CourseFile
	schedules       *SchedulesFile
	weight          *WeightFile
	totals          *TotalsFile
	goals           *GoalsFile
	bloodPressure   *BloodPressureFile
	monitoringA     *MonitoringAFile
	activitySummary *ActivitySummaryFile
	monitoringDaily *MonitoringDailyFile
	monitoringB     *MonitoringBFile
	segment         *SegmentFile
	segmentList     *SegmentListFile
}

// UnknownField represents an unknown FIT message field not found in the
// official profile. It contains the global message and field number.
type UnknownField struct {
	MesgNum  MesgNum
	FieldNum byte
}

type msgAdder interface {
	add(reflect.Value)
}

func (f *Fit) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case FileIdMsg:
		f.FileId = x.(FileIdMsg)
	case FileCreatorMsg:
		tmp := x.(FileCreatorMsg)
		f.FileCreator = &tmp
	case TimestampCorrelationMsg:
		tmp := x.(TimestampCorrelationMsg)
		f.TimestampCorrelation = &tmp
	case DeviceInfoMsg:
		tmp := x.(DeviceInfoMsg)
		f.DeviceInfo = &tmp
	default:
		f.msgAdder.add(msg)
	}
}

// FileType returns the FIT file type.
func (f *Fit) FileType() File {
	return f.FileId.Type
}

type wrongFileTypeError struct {
	actual, requested File
}

func (e wrongFileTypeError) Error() string {
	return fmt.Sprintf("fit file is type is %v, not %v", e.actual, e.requested)
}

// Activity returns f's Activity file. An error is returned if the FIT file is
// not of type activity.
func (f *Fit) Activity() (*ActivityFile, error) {
	if !(f.FileId.Type == FileActivity) {
		return nil, wrongFileTypeError{f.FileId.Type, FileActivity}
	}
	return f.activity, nil
}

// Device returns f's Device file. An error is returned if the FIT file is
// not of type device.
func (f *Fit) Device() (*DeviceFile, error) {
	if !(f.FileId.Type == FileDevice) {
		return nil, wrongFileTypeError{f.FileId.Type, FileDevice}
	}
	return f.device, nil
}

// Settings returns f's Settings file. An error is returned if the FIT file is
// not of type settings.
func (f *Fit) Settings() (*SettingsFile, error) {
	if !(f.FileId.Type == FileSettings) {
		return nil, wrongFileTypeError{f.FileId.Type, FileSettings}
	}
	return f.settings, nil
}

// Sport returns f's Sport file. An error is returned if the FIT file is
// not of type sport.
func (f *Fit) Sport() (*SportFile, error) {
	if !(f.FileId.Type == FileSport) {
		return nil, wrongFileTypeError{f.FileId.Type, FileSport}
	}
	return f.sport, nil
}

// Workout returns f's Workout file. An error is returned if the FIT file is
// not of type workout.
func (f *Fit) Workout() (*WorkoutFile, error) {
	if !(f.FileId.Type == FileWorkout) {
		return nil, wrongFileTypeError{f.FileId.Type, FileWorkout}
	}
	return f.workout, nil
}

// Course returns f's Course file. An error is returned if the FIT file is
// not of type course.
func (f *Fit) Course() (*CourseFile, error) {
	if !(f.FileId.Type == FileCourse) {
		return nil, wrongFileTypeError{f.FileId.Type, FileCourse}
	}
	return f.course, nil
}

// Schedules returns f's Schedules file. An error is returned if the FIT file is
// not of type schedules.
func (f *Fit) Schedules() (*SchedulesFile, error) {
	if !(f.FileId.Type == FileSchedules) {
		return nil, wrongFileTypeError{f.FileId.Type, FileSchedules}
	}
	return f.schedules, nil
}

// Weight returns f's Weight file. An error is returned if the FIT file is
// not of type weight.
func (f *Fit) Weight() (*WeightFile, error) {
	if !(f.FileId.Type == FileWeight) {
		return nil, wrongFileTypeError{f.FileId.Type, FileWeight}
	}
	return f.weight, nil
}

// Totals returns f's Totals file. An error is returned if the FIT file is
// not of type totals.
func (f *Fit) Totals() (*TotalsFile, error) {
	if !(f.FileId.Type == FileTotals) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTotals}
	}
	return f.totals, nil
}

// Goals returns f's Goals file. An error is returned if the FIT file is
// not of type goals.
func (f *Fit) Goals() (*GoalsFile, error) {
	if !(f.FileId.Type == FileGoals) {
		return nil, wrongFileTypeError{f.FileId.Type, FileGoals}
	}
	return f.goals, nil
}

// BloodPressure returns f's BloodPressure file. An error is returned if the FIT file is
// not of type blood pressure.
func (f *Fit) BloodPressure() (*BloodPressureFile, error) {
	if !(f.FileId.Type == FileBloodPressure) {
		return nil, wrongFileTypeError{f.FileId.Type, FileBloodPressure}
	}
	return f.bloodPressure, nil
}

// MonitoringA returns f's MonitoringA file. An error is returned if the FIT file is
// not of type monitoring A.
func (f *Fit) MonitoringA() (*MonitoringAFile, error) {
	if !(f.FileId.Type == FileMonitoringA) {
		return nil, wrongFileTypeError{f.FileId.Type, FileMonitoringA}
	}
	return f.monitoringA, nil
}

// ActivitySummary returns f's ActivitySummary file. An error is returned if the FIT file is
// not of type activity summary.
func (f *Fit) ActivitySummary() (*ActivitySummaryFile, error) {
	if !(f.FileId.Type == FileActivitySummary) {
		return nil, wrongFileTypeError{f.FileId.Type, FileActivitySummary}
	}
	return f.activitySummary, nil
}

// MonitoringDaily returns f's MonitoringDaily file. An error is returned if the FIT file is
// not of type monitoring daily.
func (f *Fit) MonitoringDaily() (*MonitoringDailyFile, error) {
	if !(f.FileId.Type == FileMonitoringDaily) {
		return nil, wrongFileTypeError{f.FileId.Type, FileMonitoringDaily}
	}
	return f.monitoringDaily, nil
}

// MonitoringB returns f's MonitoringB file. An error is returned if the FIT file is
// not of type monitoring B.
func (f *Fit) MonitoringB() (*MonitoringBFile, error) {
	if !(f.FileId.Type == FileMonitoringB) {
		return nil, wrongFileTypeError{f.FileId.Type, FileMonitoringB}
	}
	return f.monitoringB, nil
}

// Segment returns f's Segment file. An error is returned if the FIT file is
// not of type segment.
func (f *Fit) Segment() (*SegmentFile, error) {
	if !(f.FileId.Type == FileSegment) {
		return nil, wrongFileTypeError{f.FileId.Type, FileSegment}
	}
	return f.segment, nil
}

// SegmentList returns f's SegmentList file. An error is returned if the FIT file is
// not of type segment list.
func (f *Fit) SegmentList() (*SegmentListFile, error) {
	if !(f.FileId.Type == FileSegmentList) {
		return nil, wrongFileTypeError{f.FileId.Type, FileSegmentList}
	}
	return f.segmentList, nil
}
