package fit

import (
	"fmt"
	"reflect"
)

// File represents a decoded FIT file.
type File struct {
	// Header is the FIT file header.
	Header Header

	// CRC is the FIT file CRC.
	CRC uint16

	// FileId is a message required for all FIT files.
	FileId FileIdMsg

	// Common messages for all FIT file types.
	FileCreator          *FileCreatorMsg
	TimestampCorrelation *TimestampCorrelationMsg

	// UnknownMessages is a slice of unknown messages encountered during
	// decoding. It is sorted by message number.
	UnknownMessages []UnknownMessage

	// UnknownFields is a slice of unknown fields for known messages
	// encountered during decoding. It is sorted by message number.
	UnknownFields []UnknownField

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

type msgAdder interface {
	add(reflect.Value)
}

// NewFile creates a new File of the given type.
func NewFile(t FileType, h Header) (*File, error) {
	f := new(File)
	f.FileId.Type = t
	f.Header = h
	if err := f.checkAndSetMsgAdder(); err != nil {
		return nil, fmt.Errorf("error creating file: %w", err)
	}
	return f, nil
}

func (f *File) add(msg reflect.Value) {
	x := msg.Interface()
	_ = f.checkAndSetMsgAdder()
	switch tmp := x.(type) {
	case FileIdMsg:
		f.FileId = tmp
	case FileCreatorMsg:
		f.FileCreator = &tmp
	case TimestampCorrelationMsg:
		f.TimestampCorrelation = &tmp
	default:
		f.msgAdder.add(msg)
	}
}

// checkAndSetMsgAdder check message type and set msgAdder
func (f *File) checkAndSetMsgAdder() error {
	t := f.Type()
	switch t {
	case FileTypeActivity:
		if f.activity == nil {
			f.activity = new(ActivityFile)
		}
		f.msgAdder = f.activity
	case FileTypeDevice:
		if f.device == nil {
			f.device = new(DeviceFile)
		}
		f.msgAdder = f.device
	case FileTypeSettings:
		if f.settings == nil {
			f.settings = new(SettingsFile)
		}
		f.msgAdder = f.settings
	case FileTypeSport:
		if f.sport == nil {
			f.sport = new(SportFile)
		}
		f.msgAdder = f.sport
	case FileTypeWorkout:
		if f.workout == nil {
			f.workout = new(WorkoutFile)
		}
		f.msgAdder = f.workout
	case FileTypeCourse:
		if f.course == nil {
			f.course = new(CourseFile)
		}
		f.msgAdder = f.course
	case FileTypeSchedules:
		if f.schedules == nil {
			f.schedules = new(SchedulesFile)
		}
		f.msgAdder = f.schedules
	case FileTypeWeight:
		if f.weight == nil {
			f.weight = new(WeightFile)
		}
		f.msgAdder = f.weight
	case FileTypeTotals:
		if f.totals == nil {
			f.totals = new(TotalsFile)
		}
		f.msgAdder = f.totals
	case FileTypeGoals:
		if f.goals == nil {
			f.goals = new(GoalsFile)
		}
		f.msgAdder = f.goals
	case FileTypeBloodPressure:
		if f.bloodPressure == nil {
			f.bloodPressure = new(BloodPressureFile)
		}
		f.msgAdder = f.bloodPressure
	case FileTypeMonitoringA:
		if f.monitoringA == nil {
			f.monitoringA = new(MonitoringAFile)
		}

		f.msgAdder = f.monitoringA
	case FileTypeActivitySummary:
		if f.activitySummary == nil {
			f.activitySummary = new(ActivitySummaryFile)
		}
		f.msgAdder = f.activitySummary
	case FileTypeMonitoringDaily:
		if f.monitoringDaily == nil {
			f.monitoringDaily = new(MonitoringDailyFile)
		}
		f.msgAdder = f.monitoringDaily
	case FileTypeMonitoringB:
		if f.monitoringB == nil {
			f.monitoringB = new(MonitoringBFile)
		}
		f.msgAdder = f.monitoringB
	case FileTypeSegment:
		if f.segment == nil {
			f.segment = new(SegmentFile)
		}
		f.msgAdder = f.segment
	case FileTypeSegmentList:
		if f.segmentList == nil {
			f.segmentList = new(SegmentListFile)
		}
		f.msgAdder = f.segmentList
	case FileTypeInvalid:
		return FormatError("file type was set invalid")
	default:
		switch {
		case t > FileTypeMonitoringB && t < FileTypeMfgRangeMin:
			return FormatError(
				fmt.Sprintf("unknown file type: %v", t),
			)
		case t >= FileTypeMfgRangeMin && t <= FileTypeMfgRangeMax:
			return NotSupportedError("manufacturer specific file types")
		default:
			return FormatError(
				fmt.Sprintf("unknown file type: %v", t),
			)
		}
	}
	return nil
}

// Type returns the FIT file type.
func (f *File) Type() FileType {
	return f.FileId.Type
}

type wrongFileTypeError struct {
	actual, requested FileType
}

func (e wrongFileTypeError) Error() string {
	return fmt.Sprintf("fit file type is %v, not %v", e.actual, e.requested)
}

// Activity returns f's Activity file. An error is returned if the FIT file is
// not of type activity.
func (f *File) Activity() (*ActivityFile, error) {
	if !(f.FileId.Type == FileTypeActivity) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeActivity}
	}
	return f.activity, nil
}

// Device returns f's Device file. An error is returned if the FIT file is
// not of type device.
func (f *File) Device() (*DeviceFile, error) {
	if !(f.FileId.Type == FileTypeDevice) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeDevice}
	}
	return f.device, nil
}

// Settings returns f's Settings file. An error is returned if the FIT file is
// not of type settings.
func (f *File) Settings() (*SettingsFile, error) {
	if !(f.FileId.Type == FileTypeSettings) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeSettings}
	}
	return f.settings, nil
}

// Sport returns f's Sport file. An error is returned if the FIT file is
// not of type sport.
func (f *File) Sport() (*SportFile, error) {
	if !(f.FileId.Type == FileTypeSport) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeSport}
	}
	return f.sport, nil
}

// Workout returns f's Workout file. An error is returned if the FIT file is
// not of type workout.
func (f *File) Workout() (*WorkoutFile, error) {
	if !(f.FileId.Type == FileTypeWorkout) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeWorkout}
	}
	return f.workout, nil
}

// Course returns f's Course file. An error is returned if the FIT file is
// not of type course.
func (f *File) Course() (*CourseFile, error) {
	if !(f.FileId.Type == FileTypeCourse) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeCourse}
	}
	return f.course, nil
}

// Schedules returns f's Schedules file. An error is returned if the FIT file is
// not of type schedules.
func (f *File) Schedules() (*SchedulesFile, error) {
	if !(f.FileId.Type == FileTypeSchedules) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeSchedules}
	}
	return f.schedules, nil
}

// Weight returns f's Weight file. An error is returned if the FIT file is
// not of type weight.
func (f *File) Weight() (*WeightFile, error) {
	if !(f.FileId.Type == FileTypeWeight) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeWeight}
	}
	return f.weight, nil
}

// Totals returns f's Totals file. An error is returned if the FIT file is
// not of type totals.
func (f *File) Totals() (*TotalsFile, error) {
	if !(f.FileId.Type == FileTypeTotals) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeTotals}
	}
	return f.totals, nil
}

// Goals returns f's Goals file. An error is returned if the FIT file is
// not of type goals.
func (f *File) Goals() (*GoalsFile, error) {
	if !(f.FileId.Type == FileTypeGoals) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeGoals}
	}
	return f.goals, nil
}

// BloodPressure returns f's BloodPressure file. An error is returned if the FIT file is
// not of type blood pressure.
func (f *File) BloodPressure() (*BloodPressureFile, error) {
	if !(f.FileId.Type == FileTypeBloodPressure) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeBloodPressure}
	}
	return f.bloodPressure, nil
}

// MonitoringA returns f's MonitoringA file. An error is returned if the FIT file is
// not of type monitoring A.
func (f *File) MonitoringA() (*MonitoringAFile, error) {
	if !(f.FileId.Type == FileTypeMonitoringA) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeMonitoringA}
	}
	return f.monitoringA, nil
}

// ActivitySummary returns f's ActivitySummary file. An error is returned if the FIT file is
// not of type activity summary.
func (f *File) ActivitySummary() (*ActivitySummaryFile, error) {
	if !(f.FileId.Type == FileTypeActivitySummary) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeActivitySummary}
	}
	return f.activitySummary, nil
}

// MonitoringDaily returns f's MonitoringDaily file. An error is returned if the FIT file is
// not of type monitoring daily.
func (f *File) MonitoringDaily() (*MonitoringDailyFile, error) {
	if !(f.FileId.Type == FileTypeMonitoringDaily) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeMonitoringDaily}
	}
	return f.monitoringDaily, nil
}

// MonitoringB returns f's MonitoringB file. An error is returned if the FIT file is
// not of type monitoring B.
func (f *File) MonitoringB() (*MonitoringBFile, error) {
	if !(f.FileId.Type == FileTypeMonitoringB) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeMonitoringB}
	}
	return f.monitoringB, nil
}

// Segment returns f's Segment file. An error is returned if the FIT file is
// not of type segment.
func (f *File) Segment() (*SegmentFile, error) {
	if !(f.FileId.Type == FileTypeSegment) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeSegment}
	}
	return f.segment, nil
}

// SegmentList returns f's SegmentList file. An error is returned if the FIT file is
// not of type segment list.
func (f *File) SegmentList() (*SegmentListFile, error) {
	if !(f.FileId.Type == FileTypeSegmentList) {
		return nil, wrongFileTypeError{f.FileId.Type, FileTypeSegmentList}
	}
	return f.segmentList, nil
}
