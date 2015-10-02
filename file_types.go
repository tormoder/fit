package fit

import "reflect"

// ActivityFile represents the Activity FIT file type.
// Records sensor data and events from active sessions.
type ActivityFile struct {
	Activity *ActivityMsg
	Sessions []*SessionMsg
	Laps     []*LapMsg
	Lengths  []*LengthMsg
	Records  []*RecordMsg
	Events   []*EventMsg
	Hrvs     []*HrvMsg
}

// DeviceFile represents the Device FIT file type.
// Describes a device's file structure and capabilities.
type DeviceFile struct {
	Softwares         []*SoftwareMsg
	Capabilities      []*CapabilitiesMsg
	FileCapabilities  []*FileCapabilitiesMsg
	MesgCapabilities  []*MesgCapabilitiesMsg
	FieldCapabilities []*FieldCapabilitiesMsg
}

// SettingsFile represents the Settings FIT file type.
// Describes a user’s parameters such as Age & Weight as well as device
// settings.
type SettingsFile struct {
	UserProfiles   []*UserProfileMsg
	HrmProfiles    []*HrmProfileMsg
	SdmProfiles    []*SdmProfileMsg
	BikeProfiles   []*BikeProfileMsg
	DeviceSettings []*DeviceSettingsMsg
}

// SportFile represents the Sport Settings FIT file type.
// Describes a user’s desired sport/zone settings.
type SportFile struct {
	ZonesTarget  *ZonesTargetMsg
	Sport        *SportMsg
	HrZones      []*HrZoneMsg
	PowerZones   []*PowerZoneMsg
	MetZones     []*MetZoneMsg
	SpeedZones   []*SpeedZoneMsg
	CadenceZones []*CadenceZoneMsg
}

// WorkoutFile represents the Workout FIT file type.
// Describes a structured activity that can be designed on a computer and
// transferred to a display device to guide a user through the activity.
type WorkoutFile struct {
	Workout      *WorkoutMsg
	WorkoutSteps []*WorkoutStepMsg
}

// CourseFile represents the Course FIT file type.
// Uses data from an activity to recreate a course.
type CourseFile struct {
	Course       *CourseMsg
	Laps         []*LapMsg
	CoursePoints []*CoursePointMsg
	Records      []*RecordMsg
}

// SchedulesFile represents the Schedules FIT file type.
// Provides scheduling of workouts and courses.
type SchedulesFile struct {
	Schedules []*ScheduleMsg
}

// WeightFile represents the Weight FIT file type.
// Records weight scale data.
type WeightFile struct {
	UserProfile  *UserProfileMsg
	WeightScales []*WeightScaleMsg
}

// TotalsFile represents the Totals FIT file type.
// Summarizes a user’s total activity, characterized by sport.
type TotalsFile struct {
	Totals []*TotalsMsg
}

// GoalsFile represents the Goals FIT file type.
// Describes a user’s exercise/health goals.
type GoalsFile struct {
	Goals []*GoalMsg
}

// BloodPressureFile represents the Bload Pressure FIT file type.
// Records blood pressure data.
type BloodPressureFile struct {
	UserProfile    *UserProfileMsg
	BloodPressures []*BloodPressureMsg
}

// MonitoringAFile represents the MonitoringA FIT file type.
// Records detailed monitoring data (i.e. logging interval < 24 Hr).
type MonitoringAFile struct {
	MonitoringInfo *MonitoringInfoMsg
	Monitorings    []*MonitoringMsg
}

// ActivitySummaryFile represents the Activity Summary FIT file type.
// Similar to Activity file, contains summary information only.
type ActivitySummaryFile struct {
	Activity *ActivityMsg
	Sessions []*SessionMsg
	Laps     []*LapMsg
}

// MonitoringDailyFile represents the Daily Monitoring FIT file type.
// Records daily summary monitoring data (i.e. logging interval = 24 hour).
type MonitoringDailyFile struct {
	MonitoringInfo *MonitoringInfoMsg
	Monitorings    []*MonitoringMsg
}

// MonitoringBFile represents the MonitoringB FIT file type.
// Records detailed monitoring data (i.e. logging interval < 24 Hr).
type MonitoringBFile struct {
	MonitoringInfo *MonitoringInfoMsg
	Monitorings    []*MonitoringMsg
}

// SegmentFile represents the Segment FIT file type.
// Describes timing data for virtual races.
type SegmentFile struct {
	SegmentId               *SegmentIdMsg
	SegmentLeaderboardEntry *SegmentLeaderboardEntryMsg
	SegmentLap              *SegmentLapMsg
	SegmentPoints           []*SegmentPointMsg
}

// SegmentListFile represents the Segment List FIT file type.
// Describes available segments.
type SegmentListFile struct {
	SegmentFiles []*SegmentFileMsg
}

func (a *ActivityFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case ActivityMsg:
		tmp := x.(ActivityMsg)
		a.Activity = &tmp
	case SessionMsg:
		tmp := x.(SessionMsg)
		tmp.expandComponents()
		a.Sessions = append(a.Sessions, &tmp)
	case LapMsg:
		tmp := x.(LapMsg)
		tmp.expandComponents()
		a.Laps = append(a.Laps, &tmp)
	case LengthMsg:
		tmp := x.(LengthMsg)
		a.Lengths = append(a.Lengths, &tmp)
	case RecordMsg:
		tmp := x.(RecordMsg)
		tmp.expandComponents()
		a.Records = append(a.Records, &tmp)
	case EventMsg:
		tmp := x.(EventMsg)
		tmp.expandComponents()
		a.Events = append(a.Events, &tmp)
	case HrvMsg:
		tmp := x.(HrvMsg)
		a.Hrvs = append(a.Hrvs, &tmp)
	default:
	}
}

func (d *DeviceFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case SoftwareMsg:
		tmp := x.(SoftwareMsg)
		d.Softwares = append(d.Softwares, &tmp)
	case CapabilitiesMsg:
		tmp := x.(CapabilitiesMsg)
		d.Capabilities = append(d.Capabilities, &tmp)
	case FileCapabilitiesMsg:
		tmp := x.(FileCapabilitiesMsg)
		d.FileCapabilities = append(d.FileCapabilities, &tmp)
	case MesgCapabilitiesMsg:
		tmp := x.(MesgCapabilitiesMsg)
		d.MesgCapabilities = append(d.MesgCapabilities, &tmp)
	case FieldCapabilitiesMsg:
		tmp := x.(FieldCapabilitiesMsg)
		d.FieldCapabilities = append(d.FieldCapabilities, &tmp)
	default:
	}
}

func (s *SettingsFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case UserProfileMsg:
		tmp := x.(UserProfileMsg)
		s.UserProfiles = append(s.UserProfiles, &tmp)
	case HrmProfileMsg:
		tmp := x.(HrmProfileMsg)
		s.HrmProfiles = append(s.HrmProfiles, &tmp)
	case SdmProfileMsg:
		tmp := x.(SdmProfileMsg)
		s.SdmProfiles = append(s.SdmProfiles, &tmp)
	case BikeProfileMsg:
		tmp := x.(BikeProfileMsg)
		s.BikeProfiles = append(s.BikeProfiles, &tmp)
	case DeviceSettingsMsg:
		tmp := x.(DeviceSettingsMsg)
		s.DeviceSettings = append(s.DeviceSettings, &tmp)
	default:
	}
}

func (s *SportFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case ZonesTargetMsg:
		tmp := x.(ZonesTargetMsg)
		s.ZonesTarget = &tmp
	case SportMsg:
		tmp := x.(SportMsg)
		s.Sport = &tmp
	case HrZoneMsg:
		tmp := x.(HrZoneMsg)
		s.HrZones = append(s.HrZones, &tmp)
	case PowerZoneMsg:
		tmp := x.(PowerZoneMsg)
		s.PowerZones = append(s.PowerZones, &tmp)
	case MetZoneMsg:
		tmp := x.(MetZoneMsg)
		s.MetZones = append(s.MetZones, &tmp)
	case SpeedZoneMsg:
		tmp := x.(SpeedZoneMsg)
		s.SpeedZones = append(s.SpeedZones, &tmp)
	case CadenceZoneMsg:
		tmp := x.(CadenceZoneMsg)
		s.CadenceZones = append(s.CadenceZones, &tmp)
	default:
	}
}

func (w *WorkoutFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case WorkoutMsg:
		tmp := x.(WorkoutMsg)
		w.Workout = &tmp
	case WorkoutStepMsg:
		tmp := x.(WorkoutStepMsg)
		w.WorkoutSteps = append(w.WorkoutSteps, &tmp)
	default:
	}
}

func (c *CourseFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case CourseMsg:
		tmp := x.(CourseMsg)
		c.Course = &tmp
	case LapMsg:
		tmp := x.(LapMsg)
		tmp.expandComponents()
		c.Laps = append(c.Laps, &tmp)
	case CoursePointMsg:
		tmp := x.(CoursePointMsg)
		c.CoursePoints = append(c.CoursePoints, &tmp)
	case RecordMsg:
		tmp := x.(RecordMsg)
		tmp.expandComponents()
		c.Records = append(c.Records, &tmp)
	default:
	}
}

func (s *SchedulesFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case ScheduleMsg:
		tmp := x.(ScheduleMsg)
		s.Schedules = append(s.Schedules, &tmp)
	default:
	}
}

func (w *WeightFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case UserProfileMsg:
		tmp := x.(UserProfileMsg)
		w.UserProfile = &tmp
	case WeightScaleMsg:
		tmp := x.(WeightScaleMsg)
		w.WeightScales = append(w.WeightScales, &tmp)
	default:
	}
}

func (t *TotalsFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case TotalsMsg:
		tmp := x.(TotalsMsg)
		t.Totals = append(t.Totals, &tmp)
	default:
	}
}

func (g *GoalsFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case GoalMsg:
		tmp := x.(GoalMsg)
		g.Goals = append(g.Goals, &tmp)
	default:
	}
}

func (b *BloodPressureFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case UserProfileMsg:
		tmp := x.(UserProfileMsg)
		b.UserProfile = &tmp
	case BloodPressureMsg:
		tmp := x.(BloodPressureMsg)
		b.BloodPressures = append(b.BloodPressures, &tmp)
	default:
	}
}

func (m *MonitoringAFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case MonitoringInfoMsg:
		tmp := x.(MonitoringInfoMsg)
		m.MonitoringInfo = &tmp
	case MonitoringMsg:
		tmp := x.(MonitoringMsg)
		m.Monitorings = append(m.Monitorings, &tmp)
	default:
	}
}

func (a *ActivitySummaryFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case ActivityMsg:
		tmp := x.(ActivityMsg)
		a.Activity = &tmp
	case SessionMsg:
		tmp := x.(SessionMsg)
		tmp.expandComponents()
		a.Sessions = append(a.Sessions, &tmp)
	case LapMsg:
		tmp := x.(LapMsg)
		tmp.expandComponents()
		a.Laps = append(a.Laps, &tmp)
	default:
	}
}

func (m *MonitoringDailyFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case MonitoringInfoMsg:
		tmp := x.(MonitoringInfoMsg)
		m.MonitoringInfo = &tmp
	case MonitoringMsg:
		tmp := x.(MonitoringMsg)
		m.Monitorings = append(m.Monitorings, &tmp)
	default:
	}
}

func (m *MonitoringBFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case MonitoringInfoMsg:
		tmp := x.(MonitoringInfoMsg)
		m.MonitoringInfo = &tmp
	case MonitoringMsg:
		tmp := x.(MonitoringMsg)
		m.Monitorings = append(m.Monitorings, &tmp)
	default:
	}
}

func (s *SegmentFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case SegmentIdMsg:
		tmp := x.(SegmentIdMsg)
		s.SegmentId = &tmp
	case SegmentLeaderboardEntryMsg:
		tmp := x.(SegmentLeaderboardEntryMsg)
		s.SegmentLeaderboardEntry = &tmp
	case SegmentLapMsg:
		tmp := x.(SegmentLapMsg)
		s.SegmentLap = &tmp
	case SegmentPointMsg:
		tmp := x.(SegmentPointMsg)
		s.SegmentPoints = append(s.SegmentPoints, &tmp)
	default:
	}
}

func (s *SegmentListFile) add(msg reflect.Value) {
	x := msg.Interface()
	switch x.(type) {
	case SegmentFileMsg:
		tmp := x.(SegmentFileMsg)
		s.SegmentFiles = append(s.SegmentFiles, &tmp)
	default:
	}
}
