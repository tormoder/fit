// Code generated using the program found in 'cmd/fitgen/main.go'. DO NOT EDIT.

// SDK Version: 20.43

package fit

import (
	"math"
	"time"
)

// FileIdMsg represents the file_id FIT message type.
type FileIdMsg struct {
	Type         FileType
	Manufacturer Manufacturer
	Product      uint16
	SerialNumber uint32
	TimeCreated  time.Time // Only set for files that are can be created/erased.
	Number       uint16    // Only set for files that are not created/erased.
	ProductName  string    // Optional free form string to indicate the devices name or model
}

// GetProduct returns the appropriate Product
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *FileIdMsg) GetProduct() interface{} {
	switch x.Manufacturer {
	case ManufacturerGarmin, ManufacturerDynastream, ManufacturerDynastreamOem:
		return GarminProduct(x.Product)
	default:
		return x.Product
	}
}

// FileCreatorMsg represents the file_creator FIT message type.
type FileCreatorMsg struct {
	SoftwareVersion uint16
	HardwareVersion uint8
}

// TimestampCorrelationMsg represents the timestamp_correlation FIT message type.
type TimestampCorrelationMsg struct {
}

// SoftwareMsg represents the software FIT message type.
type SoftwareMsg struct {
	MessageIndex MessageIndex
	Version      uint16
	PartNumber   string
}

// GetVersionScaled returns Version
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
func (x *SoftwareMsg) GetVersionScaled() float64 {
	if x.Version == 0xFFFF {
		return math.NaN()
	}
	return float64(x.Version) / 100
}

// SlaveDeviceMsg represents the slave_device FIT message type.
type SlaveDeviceMsg struct {
	Manufacturer Manufacturer
	Product      uint16
}

// GetProduct returns the appropriate Product
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *SlaveDeviceMsg) GetProduct() interface{} {
	switch x.Manufacturer {
	case ManufacturerGarmin, ManufacturerDynastream, ManufacturerDynastreamOem:
		return GarminProduct(x.Product)
	default:
		return x.Product
	}
}

// CapabilitiesMsg represents the capabilities FIT message type.
type CapabilitiesMsg struct {
	Languages             []uint8      // Use language_bits_x types where x is index of array.
	Sports                []SportBits0 // Use sport_bits_x types where x is index of array.
	WorkoutsSupported     WorkoutCapabilities
	ConnectivitySupported ConnectivityCapabilities
}

// FileCapabilitiesMsg represents the file_capabilities FIT message type.
type FileCapabilitiesMsg struct {
	MessageIndex MessageIndex
	Type         FileType
	Flags        FileFlags
	Directory    string
	MaxCount     uint16
	MaxSize      uint32
}

// MesgCapabilitiesMsg represents the mesg_capabilities FIT message type.
type MesgCapabilitiesMsg struct {
	MessageIndex MessageIndex
	File         FileType
	MesgNum      MesgNum
	CountType    MesgCount
	Count        uint16
}

// GetCount returns the appropriate Count
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *MesgCapabilitiesMsg) GetCount() interface{} {
	switch x.CountType {
	case MesgCountNumPerFile:
		return uint16(x.Count)
	case MesgCountMaxPerFile:
		return uint16(x.Count)
	case MesgCountMaxPerFileType:
		return uint16(x.Count)
	default:
		return x.Count
	}
}

// FieldCapabilitiesMsg represents the field_capabilities FIT message type.
type FieldCapabilitiesMsg struct {
	MessageIndex MessageIndex
	File         FileType
	MesgNum      MesgNum
	FieldNum     uint8
	Count        uint16
}

// DeviceSettingsMsg represents the device_settings FIT message type.
type DeviceSettingsMsg struct {
	ActiveTimeZone         uint8         // Index into time zone arrays.
	UtcOffset              uint32        // Offset from system time. Required to convert timestamp from system time to UTC.
	TimeOffset             []uint32      // Offset from system time.
	TimeMode               []TimeMode    // Display mode for the time
	TimeZoneOffset         []int8        // timezone offset in 1/4 hour increments
	BacklightMode          BacklightMode // Mode for backlight
	ActivityTrackerEnabled Bool          // Enabled state of the activity tracker functionality
	ClockTime              time.Time     // UTC timestamp used to set the devices clock and date
	PagesEnabled           []uint16      // Bitfield  to configure enabled screens for each supported loop
	MoveAlertEnabled       Bool          // Enabled state of the move alert
	DateMode               DateMode      // Display mode for the date
	DisplayOrientation     DisplayOrientation
	MountingSide           Side
	DefaultPage            []uint16 // Bitfield to indicate one page as default for each supported loop
	AutosyncMinSteps       uint16   // Minimum steps before an autosync can occur
	AutosyncMinTime        uint16   // Minimum minutes before an autosync can occur
}

// GetTimeZoneOffsetScaled returns TimeZoneOffset
// as a slice with scale and any offset applied to every element.
// Units: hr
func (x *DeviceSettingsMsg) GetTimeZoneOffsetScaled() []float64 {
	if len(x.TimeZoneOffset) == 0 {
		return nil
	}
	s := make([]float64, len(x.TimeZoneOffset))
	for i, v := range x.TimeZoneOffset {
		s[i] = float64(v) / 4
	}
	return s
}

// UserProfileMsg represents the user_profile FIT message type.
type UserProfileMsg struct {
	MessageIndex               MessageIndex
	FriendlyName               string
	Gender                     Gender
	Age                        uint8
	Height                     uint8
	Weight                     uint16
	Language                   Language
	ElevSetting                DisplayMeasure
	WeightSetting              DisplayMeasure
	RestingHeartRate           uint8
	DefaultMaxRunningHeartRate uint8
	DefaultMaxBikingHeartRate  uint8
	DefaultMaxHeartRate        uint8
	HrSetting                  DisplayHeart
	SpeedSetting               DisplayMeasure
	DistSetting                DisplayMeasure
	PowerSetting               DisplayPower
	ActivityClass              ActivityClass
	PositionSetting            DisplayPosition
	TemperatureSetting         DisplayMeasure
	LocalId                    UserLocalId
	GlobalId                   []byte
	HeightSetting              DisplayMeasure
	UserRunningStepLength      uint16 // User defined running step length set to 0 for auto length
	UserWalkingStepLength      uint16 // User defined walking step length set to 0 for auto length
}

// GetHeightScaled returns Height
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *UserProfileMsg) GetHeightScaled() float64 {
	if x.Height == 0xFF {
		return math.NaN()
	}
	return float64(x.Height) / 100
}

// GetWeightScaled returns Weight
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: kg
func (x *UserProfileMsg) GetWeightScaled() float64 {
	if x.Weight == 0xFFFF {
		return math.NaN()
	}
	return float64(x.Weight) / 10
}

// GetUserRunningStepLengthScaled returns UserRunningStepLength
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *UserProfileMsg) GetUserRunningStepLengthScaled() float64 {
	if x.UserRunningStepLength == 0xFFFF {
		return math.NaN()
	}
	return float64(x.UserRunningStepLength) / 1000
}

// GetUserWalkingStepLengthScaled returns UserWalkingStepLength
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *UserProfileMsg) GetUserWalkingStepLengthScaled() float64 {
	if x.UserWalkingStepLength == 0xFFFF {
		return math.NaN()
	}
	return float64(x.UserWalkingStepLength) / 1000
}

// HrmProfileMsg represents the hrm_profile FIT message type.
type HrmProfileMsg struct {
	MessageIndex      MessageIndex
	Enabled           Bool
	HrmAntId          uint16
	LogHrv            Bool
	HrmAntIdTransType uint8
}

// SdmProfileMsg represents the sdm_profile FIT message type.
type SdmProfileMsg struct {
	MessageIndex      MessageIndex
	Enabled           Bool
	SdmAntId          uint16
	SdmCalFactor      uint16
	Odometer          uint32
	SpeedSource       Bool // Use footpod for speed source instead of GPS
	SdmAntIdTransType uint8
	OdometerRollover  uint8 // Rollover counter that can be used to extend the odometer
}

// GetSdmCalFactorScaled returns SdmCalFactor
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *SdmProfileMsg) GetSdmCalFactorScaled() float64 {
	if x.SdmCalFactor == 0xFFFF {
		return math.NaN()
	}
	return float64(x.SdmCalFactor) / 10
}

// GetOdometerScaled returns Odometer
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SdmProfileMsg) GetOdometerScaled() float64 {
	if x.Odometer == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.Odometer) / 100
}

// BikeProfileMsg represents the bike_profile FIT message type.
type BikeProfileMsg struct {
	MessageIndex             MessageIndex
	Name                     string
	Sport                    Sport
	SubSport                 SubSport
	Odometer                 uint32
	BikeSpdAntId             uint16
	BikeCadAntId             uint16
	BikeSpdcadAntId          uint16
	BikePowerAntId           uint16
	CustomWheelsize          uint16
	AutoWheelsize            uint16
	BikeWeight               uint16
	PowerCalFactor           uint16
	AutoWheelCal             Bool
	AutoPowerZero            Bool
	Id                       uint8
	SpdEnabled               Bool
	CadEnabled               Bool
	SpdcadEnabled            Bool
	PowerEnabled             Bool
	CrankLength              uint8
	Enabled                  Bool
	BikeSpdAntIdTransType    uint8
	BikeCadAntIdTransType    uint8
	BikeSpdcadAntIdTransType uint8
	BikePowerAntIdTransType  uint8
	OdometerRollover         uint8   // Rollover counter that can be used to extend the odometer
	FrontGearNum             uint8   // Number of front gears
	FrontGear                []uint8 // Number of teeth on each gear 0 is innermost
	RearGearNum              uint8   // Number of rear gears
	RearGear                 []uint8 // Number of teeth on each gear 0 is innermost
	ShimanoDi2Enabled        Bool
}

// GetOdometerScaled returns Odometer
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *BikeProfileMsg) GetOdometerScaled() float64 {
	if x.Odometer == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.Odometer) / 100
}

// GetCustomWheelsizeScaled returns CustomWheelsize
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *BikeProfileMsg) GetCustomWheelsizeScaled() float64 {
	if x.CustomWheelsize == 0xFFFF {
		return math.NaN()
	}
	return float64(x.CustomWheelsize) / 1000
}

// GetAutoWheelsizeScaled returns AutoWheelsize
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *BikeProfileMsg) GetAutoWheelsizeScaled() float64 {
	if x.AutoWheelsize == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AutoWheelsize) / 1000
}

// GetBikeWeightScaled returns BikeWeight
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: kg
func (x *BikeProfileMsg) GetBikeWeightScaled() float64 {
	if x.BikeWeight == 0xFFFF {
		return math.NaN()
	}
	return float64(x.BikeWeight) / 10
}

// GetPowerCalFactorScaled returns PowerCalFactor
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *BikeProfileMsg) GetPowerCalFactorScaled() float64 {
	if x.PowerCalFactor == 0xFFFF {
		return math.NaN()
	}
	return float64(x.PowerCalFactor) / 10
}

// GetCrankLengthScaled returns CrankLength
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: mm
func (x *BikeProfileMsg) GetCrankLengthScaled() float64 {
	if x.CrankLength == 0xFF {
		return math.NaN()
	}
	return float64(x.CrankLength)/2 - -110
}

// ConnectivityMsg represents the connectivity FIT message type.
type ConnectivityMsg struct {
	BluetoothEnabled            Bool // Use Bluetooth for connectivity features
	BluetoothLeEnabled          Bool // Use Bluetooth Low Energy for connectivity features
	AntEnabled                  Bool // Use ANT for connectivity features
	Name                        string
	LiveTrackingEnabled         Bool
	WeatherConditionsEnabled    Bool
	WeatherAlertsEnabled        Bool
	AutoActivityUploadEnabled   Bool
	CourseDownloadEnabled       Bool
	WorkoutDownloadEnabled      Bool
	GpsEphemerisDownloadEnabled Bool
	IncidentDetectionEnabled    Bool
	GrouptrackEnabled           Bool
}

// WatchfaceSettingsMsg represents the watchface_settings FIT message type.
type WatchfaceSettingsMsg struct {
}

// OhrSettingsMsg represents the ohr_settings FIT message type.
type OhrSettingsMsg struct {
}

// ZonesTargetMsg represents the zones_target FIT message type.
type ZonesTargetMsg struct {
	MaxHeartRate             uint8
	ThresholdHeartRate       uint8
	FunctionalThresholdPower uint16
	HrCalcType               HrZoneCalc
	PwrCalcType              PwrZoneCalc
}

// SportMsg represents the sport FIT message type.
type SportMsg struct {
	Sport    Sport
	SubSport SubSport
	Name     string
}

// HrZoneMsg represents the hr_zone FIT message type.
type HrZoneMsg struct {
	MessageIndex MessageIndex
	HighBpm      uint8
	Name         string
}

// SpeedZoneMsg represents the speed_zone FIT message type.
type SpeedZoneMsg struct {
	MessageIndex MessageIndex
	HighValue    uint16
	Name         string
}

// GetHighValueScaled returns HighValue
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SpeedZoneMsg) GetHighValueScaled() float64 {
	if x.HighValue == 0xFFFF {
		return math.NaN()
	}
	return float64(x.HighValue) / 1000
}

// CadenceZoneMsg represents the cadence_zone FIT message type.
type CadenceZoneMsg struct {
	MessageIndex MessageIndex
	HighValue    uint8
	Name         string
}

// PowerZoneMsg represents the power_zone FIT message type.
type PowerZoneMsg struct {
	MessageIndex MessageIndex
	HighValue    uint16
	Name         string
}

// MetZoneMsg represents the met_zone FIT message type.
type MetZoneMsg struct {
	MessageIndex MessageIndex
	HighBpm      uint8
	Calories     uint16
	FatCalories  uint8
}

// GetCaloriesScaled returns Calories
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: kcal / min
func (x *MetZoneMsg) GetCaloriesScaled() float64 {
	if x.Calories == 0xFFFF {
		return math.NaN()
	}
	return float64(x.Calories) / 10
}

// GetFatCaloriesScaled returns FatCalories
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: kcal / min
func (x *MetZoneMsg) GetFatCaloriesScaled() float64 {
	if x.FatCalories == 0xFF {
		return math.NaN()
	}
	return float64(x.FatCalories) / 10
}

// GoalMsg represents the goal FIT message type.
type GoalMsg struct {
	MessageIndex    MessageIndex
	Sport           Sport
	SubSport        SubSport
	StartDate       time.Time
	EndDate         time.Time
	Type            Goal
	Value           uint32
	Repeat          Bool
	TargetValue     uint32
	Recurrence      GoalRecurrence
	RecurrenceValue uint16
	Enabled         Bool
	Source          GoalSource
}

// ActivityMsg represents the activity FIT message type.
type ActivityMsg struct {
	Timestamp      time.Time
	TotalTimerTime uint32 // Exclude pauses
	NumSessions    uint16
	Type           ActivityMode
	Event          Event
	EventType      EventType
	LocalTimestamp time.Time // timestamp epoch expressed in local time, used to convert activity timestamps to local time
	EventGroup     uint8
}

// GetTotalTimerTimeScaled returns TotalTimerTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *ActivityMsg) GetTotalTimerTimeScaled() float64 {
	if x.TotalTimerTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalTimerTime) / 1000
}

// SessionMsg represents the session FIT message type.
type SessionMsg struct {
	MessageIndex                 MessageIndex // Selected bit is set for the current session.
	Timestamp                    time.Time    // Sesson end time.
	Event                        Event        // session
	EventType                    EventType    // stop
	StartTime                    time.Time
	StartPositionLat             Latitude
	StartPositionLong            Longitude
	Sport                        Sport
	SubSport                     SubSport
	TotalElapsedTime             uint32 // Time (includes pauses)
	TotalTimerTime               uint32 // Timer Time (excludes pauses)
	TotalDistance                uint32
	TotalCycles                  uint32
	TotalCalories                uint16
	TotalFatCalories             uint16
	AvgSpeed                     uint16 // total_distance / total_timer_time
	MaxSpeed                     uint16
	AvgHeartRate                 uint8 // average heart rate (excludes pause time)
	MaxHeartRate                 uint8
	AvgCadence                   uint8 // total_cycles / total_timer_time if non_zero_avg_cadence otherwise total_cycles / total_elapsed_time
	MaxCadence                   uint8
	AvgPower                     uint16 // total_power / total_timer_time if non_zero_avg_power otherwise total_power / total_elapsed_time
	MaxPower                     uint16
	TotalAscent                  uint16
	TotalDescent                 uint16
	TotalTrainingEffect          uint8
	FirstLapIndex                uint16
	NumLaps                      uint16
	EventGroup                   uint8
	Trigger                      SessionTrigger
	NecLat                       Latitude
	NecLong                      Longitude
	SwcLat                       Latitude
	SwcLong                      Longitude
	NormalizedPower              uint16
	TrainingStressScore          uint16
	IntensityFactor              uint16
	LeftRightBalance             LeftRightBalance100
	AvgStrokeCount               uint32
	AvgStrokeDistance            uint16
	SwimStroke                   SwimStroke
	PoolLength                   uint16
	ThresholdPower               uint16
	PoolLengthUnit               DisplayMeasure
	NumActiveLengths             uint16 // # of active lengths of swim pool
	TotalWork                    uint32
	AvgAltitude                  uint16
	MaxAltitude                  uint16
	GpsAccuracy                  uint8
	AvgGrade                     int16
	AvgPosGrade                  int16
	AvgNegGrade                  int16
	MaxPosGrade                  int16
	MaxNegGrade                  int16
	AvgTemperature               int8
	MaxTemperature               int8
	TotalMovingTime              uint32
	AvgPosVerticalSpeed          int16
	AvgNegVerticalSpeed          int16
	MaxPosVerticalSpeed          int16
	MaxNegVerticalSpeed          int16
	MinHeartRate                 uint8
	TimeInHrZone                 []uint32
	TimeInSpeedZone              []uint32
	TimeInCadenceZone            []uint32
	TimeInPowerZone              []uint32
	AvgLapTime                   uint32
	BestLapIndex                 uint16
	MinAltitude                  uint16
	PlayerScore                  uint16
	OpponentScore                uint16
	OpponentName                 string
	StrokeCount                  []uint16 // stroke_type enum used as the index
	ZoneCount                    []uint16 // zone number used as the index
	MaxBallSpeed                 uint16
	AvgBallSpeed                 uint16
	AvgVerticalOscillation       uint16
	AvgStanceTimePercent         uint16
	AvgStanceTime                uint16
	AvgFractionalCadence         uint8 // fractional part of the avg_cadence
	MaxFractionalCadence         uint8 // fractional part of the max_cadence
	TotalFractionalCycles        uint8 // fractional part of the total_cycles
	SportIndex                   uint8
	EnhancedAvgSpeed             uint32 // total_distance / total_timer_time
	EnhancedMaxSpeed             uint32
	EnhancedAvgAltitude          uint32
	EnhancedMinAltitude          uint32
	EnhancedMaxAltitude          uint32
	TotalAnaerobicTrainingEffect uint8
	AvgVam                       uint16
}

// GetTotalElapsedTimeScaled returns TotalElapsedTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *SessionMsg) GetTotalElapsedTimeScaled() float64 {
	if x.TotalElapsedTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalElapsedTime) / 1000
}

// GetTotalTimerTimeScaled returns TotalTimerTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *SessionMsg) GetTotalTimerTimeScaled() float64 {
	if x.TotalTimerTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalTimerTime) / 1000
}

// GetTotalDistanceScaled returns TotalDistance
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SessionMsg) GetTotalDistanceScaled() float64 {
	if x.TotalDistance == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalDistance) / 100
}

// GetAvgSpeedScaled returns AvgSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SessionMsg) GetAvgSpeedScaled() float64 {
	if x.AvgSpeed == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgSpeed) / 1000
}

// GetMaxSpeedScaled returns MaxSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SessionMsg) GetMaxSpeedScaled() float64 {
	if x.MaxSpeed == 0xFFFF {
		return math.NaN()
	}
	return float64(x.MaxSpeed) / 1000
}

// GetTotalTrainingEffectScaled returns TotalTrainingEffect
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
func (x *SessionMsg) GetTotalTrainingEffectScaled() float64 {
	if x.TotalTrainingEffect == 0xFF {
		return math.NaN()
	}
	return float64(x.TotalTrainingEffect) / 10
}

// GetTrainingStressScoreScaled returns TrainingStressScore
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: tss
func (x *SessionMsg) GetTrainingStressScoreScaled() float64 {
	if x.TrainingStressScore == 0xFFFF {
		return math.NaN()
	}
	return float64(x.TrainingStressScore) / 10
}

// GetIntensityFactorScaled returns IntensityFactor
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: if
func (x *SessionMsg) GetIntensityFactorScaled() float64 {
	if x.IntensityFactor == 0xFFFF {
		return math.NaN()
	}
	return float64(x.IntensityFactor) / 1000
}

// GetAvgStrokeCountScaled returns AvgStrokeCount
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: strokes/lap
func (x *SessionMsg) GetAvgStrokeCountScaled() float64 {
	if x.AvgStrokeCount == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.AvgStrokeCount) / 10
}

// GetAvgStrokeDistanceScaled returns AvgStrokeDistance
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SessionMsg) GetAvgStrokeDistanceScaled() float64 {
	if x.AvgStrokeDistance == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgStrokeDistance) / 100
}

// GetPoolLengthScaled returns PoolLength
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SessionMsg) GetPoolLengthScaled() float64 {
	if x.PoolLength == 0xFFFF {
		return math.NaN()
	}
	return float64(x.PoolLength) / 100
}

// GetAvgAltitudeScaled returns AvgAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SessionMsg) GetAvgAltitudeScaled() float64 {
	if x.AvgAltitude == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgAltitude)/5 - 500
}

// GetMaxAltitudeScaled returns MaxAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SessionMsg) GetMaxAltitudeScaled() float64 {
	if x.MaxAltitude == 0xFFFF {
		return math.NaN()
	}
	return float64(x.MaxAltitude)/5 - 500
}

// GetAvgGradeScaled returns AvgGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *SessionMsg) GetAvgGradeScaled() float64 {
	if x.AvgGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgGrade) / 100
}

// GetAvgPosGradeScaled returns AvgPosGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *SessionMsg) GetAvgPosGradeScaled() float64 {
	if x.AvgPosGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgPosGrade) / 100
}

// GetAvgNegGradeScaled returns AvgNegGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *SessionMsg) GetAvgNegGradeScaled() float64 {
	if x.AvgNegGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgNegGrade) / 100
}

// GetMaxPosGradeScaled returns MaxPosGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *SessionMsg) GetMaxPosGradeScaled() float64 {
	if x.MaxPosGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.MaxPosGrade) / 100
}

// GetMaxNegGradeScaled returns MaxNegGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *SessionMsg) GetMaxNegGradeScaled() float64 {
	if x.MaxNegGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.MaxNegGrade) / 100
}

// GetTotalMovingTimeScaled returns TotalMovingTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *SessionMsg) GetTotalMovingTimeScaled() float64 {
	if x.TotalMovingTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalMovingTime) / 1000
}

// GetAvgPosVerticalSpeedScaled returns AvgPosVerticalSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SessionMsg) GetAvgPosVerticalSpeedScaled() float64 {
	if x.AvgPosVerticalSpeed == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgPosVerticalSpeed) / 1000
}

// GetAvgNegVerticalSpeedScaled returns AvgNegVerticalSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SessionMsg) GetAvgNegVerticalSpeedScaled() float64 {
	if x.AvgNegVerticalSpeed == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgNegVerticalSpeed) / 1000
}

// GetMaxPosVerticalSpeedScaled returns MaxPosVerticalSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SessionMsg) GetMaxPosVerticalSpeedScaled() float64 {
	if x.MaxPosVerticalSpeed == 0x7FFF {
		return math.NaN()
	}
	return float64(x.MaxPosVerticalSpeed) / 1000
}

// GetMaxNegVerticalSpeedScaled returns MaxNegVerticalSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SessionMsg) GetMaxNegVerticalSpeedScaled() float64 {
	if x.MaxNegVerticalSpeed == 0x7FFF {
		return math.NaN()
	}
	return float64(x.MaxNegVerticalSpeed) / 1000
}

// GetTimeInHrZoneScaled returns TimeInHrZone
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *SessionMsg) GetTimeInHrZoneScaled() []float64 {
	if len(x.TimeInHrZone) == 0 {
		return nil
	}
	s := make([]float64, len(x.TimeInHrZone))
	for i, v := range x.TimeInHrZone {
		s[i] = float64(v) / 1000
	}
	return s
}

// GetTimeInSpeedZoneScaled returns TimeInSpeedZone
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *SessionMsg) GetTimeInSpeedZoneScaled() []float64 {
	if len(x.TimeInSpeedZone) == 0 {
		return nil
	}
	s := make([]float64, len(x.TimeInSpeedZone))
	for i, v := range x.TimeInSpeedZone {
		s[i] = float64(v) / 1000
	}
	return s
}

// GetTimeInCadenceZoneScaled returns TimeInCadenceZone
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *SessionMsg) GetTimeInCadenceZoneScaled() []float64 {
	if len(x.TimeInCadenceZone) == 0 {
		return nil
	}
	s := make([]float64, len(x.TimeInCadenceZone))
	for i, v := range x.TimeInCadenceZone {
		s[i] = float64(v) / 1000
	}
	return s
}

// GetTimeInPowerZoneScaled returns TimeInPowerZone
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *SessionMsg) GetTimeInPowerZoneScaled() []float64 {
	if len(x.TimeInPowerZone) == 0 {
		return nil
	}
	s := make([]float64, len(x.TimeInPowerZone))
	for i, v := range x.TimeInPowerZone {
		s[i] = float64(v) / 1000
	}
	return s
}

// GetAvgLapTimeScaled returns AvgLapTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *SessionMsg) GetAvgLapTimeScaled() float64 {
	if x.AvgLapTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.AvgLapTime) / 1000
}

// GetMinAltitudeScaled returns MinAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SessionMsg) GetMinAltitudeScaled() float64 {
	if x.MinAltitude == 0xFFFF {
		return math.NaN()
	}
	return float64(x.MinAltitude)/5 - 500
}

// GetMaxBallSpeedScaled returns MaxBallSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SessionMsg) GetMaxBallSpeedScaled() float64 {
	if x.MaxBallSpeed == 0xFFFF {
		return math.NaN()
	}
	return float64(x.MaxBallSpeed) / 100
}

// GetAvgBallSpeedScaled returns AvgBallSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SessionMsg) GetAvgBallSpeedScaled() float64 {
	if x.AvgBallSpeed == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgBallSpeed) / 100
}

// GetAvgVerticalOscillationScaled returns AvgVerticalOscillation
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: mm
func (x *SessionMsg) GetAvgVerticalOscillationScaled() float64 {
	if x.AvgVerticalOscillation == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgVerticalOscillation) / 10
}

// GetAvgStanceTimePercentScaled returns AvgStanceTimePercent
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: percent
func (x *SessionMsg) GetAvgStanceTimePercentScaled() float64 {
	if x.AvgStanceTimePercent == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgStanceTimePercent) / 100
}

// GetAvgStanceTimeScaled returns AvgStanceTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: ms
func (x *SessionMsg) GetAvgStanceTimeScaled() float64 {
	if x.AvgStanceTime == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgStanceTime) / 10
}

// GetAvgFractionalCadenceScaled returns AvgFractionalCadence
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: rpm
func (x *SessionMsg) GetAvgFractionalCadenceScaled() float64 {
	if x.AvgFractionalCadence == 0xFF {
		return math.NaN()
	}
	return float64(x.AvgFractionalCadence) / 128
}

// GetMaxFractionalCadenceScaled returns MaxFractionalCadence
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: rpm
func (x *SessionMsg) GetMaxFractionalCadenceScaled() float64 {
	if x.MaxFractionalCadence == 0xFF {
		return math.NaN()
	}
	return float64(x.MaxFractionalCadence) / 128
}

// GetTotalFractionalCyclesScaled returns TotalFractionalCycles
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: cycles
func (x *SessionMsg) GetTotalFractionalCyclesScaled() float64 {
	if x.TotalFractionalCycles == 0xFF {
		return math.NaN()
	}
	return float64(x.TotalFractionalCycles) / 128
}

// GetEnhancedAvgSpeedScaled returns EnhancedAvgSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SessionMsg) GetEnhancedAvgSpeedScaled() float64 {
	if x.EnhancedAvgSpeed == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.EnhancedAvgSpeed) / 1000
}

// GetEnhancedMaxSpeedScaled returns EnhancedMaxSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SessionMsg) GetEnhancedMaxSpeedScaled() float64 {
	if x.EnhancedMaxSpeed == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.EnhancedMaxSpeed) / 1000
}

// GetEnhancedAvgAltitudeScaled returns EnhancedAvgAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SessionMsg) GetEnhancedAvgAltitudeScaled() float64 {
	if x.EnhancedAvgAltitude == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.EnhancedAvgAltitude)/5 - 500
}

// GetEnhancedMinAltitudeScaled returns EnhancedMinAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SessionMsg) GetEnhancedMinAltitudeScaled() float64 {
	if x.EnhancedMinAltitude == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.EnhancedMinAltitude)/5 - 500
}

// GetEnhancedMaxAltitudeScaled returns EnhancedMaxAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SessionMsg) GetEnhancedMaxAltitudeScaled() float64 {
	if x.EnhancedMaxAltitude == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.EnhancedMaxAltitude)/5 - 500
}

// GetTotalAnaerobicTrainingEffectScaled returns TotalAnaerobicTrainingEffect
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
func (x *SessionMsg) GetTotalAnaerobicTrainingEffectScaled() float64 {
	if x.TotalAnaerobicTrainingEffect == 0xFF {
		return math.NaN()
	}
	return float64(x.TotalAnaerobicTrainingEffect) / 10
}

// GetAvgVamScaled returns AvgVam
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SessionMsg) GetAvgVamScaled() float64 {
	if x.AvgVam == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgVam) / 1000
}

// GetTotalCycles returns the appropriate TotalCycles
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *SessionMsg) GetTotalCycles() interface{} {
	switch x.Sport {
	case SportRunning, SportWalking:
		return uint32(x.TotalCycles)
	default:
		return x.TotalCycles
	}
}

// GetAvgCadence returns the appropriate AvgCadence
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *SessionMsg) GetAvgCadence() interface{} {
	switch x.Sport {
	case SportRunning:
		return uint8(x.AvgCadence)
	default:
		return x.AvgCadence
	}
}

// GetMaxCadence returns the appropriate MaxCadence
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *SessionMsg) GetMaxCadence() interface{} {
	switch x.Sport {
	case SportRunning:
		return uint8(x.MaxCadence)
	default:
		return x.MaxCadence
	}
}

func (x *SessionMsg) expandComponents() {
	if x.AvgSpeed != 0xFFFF {
		x.EnhancedAvgSpeed = uint32(
			(x.AvgSpeed >> 0) & ((1 << 16) - 1),
		)
	}
	if x.MaxSpeed != 0xFFFF {
		x.EnhancedMaxSpeed = uint32(
			(x.MaxSpeed >> 0) & ((1 << 16) - 1),
		)
	}
	if x.AvgAltitude != 0xFFFF {
		x.EnhancedAvgAltitude = uint32(
			(x.AvgAltitude >> 0) & ((1 << 16) - 1),
		)
	}
	if x.MaxAltitude != 0xFFFF {
		x.EnhancedMaxAltitude = uint32(
			(x.MaxAltitude >> 0) & ((1 << 16) - 1),
		)
	}
	if x.MinAltitude != 0xFFFF {
		x.EnhancedMinAltitude = uint32(
			(x.MinAltitude >> 0) & ((1 << 16) - 1),
		)
	}
}

// LapMsg represents the lap FIT message type.
type LapMsg struct {
	MessageIndex                  MessageIndex
	Timestamp                     time.Time // Lap end time.
	Event                         Event
	EventType                     EventType
	StartTime                     time.Time
	StartPositionLat              Latitude
	StartPositionLong             Longitude
	EndPositionLat                Latitude
	EndPositionLong               Longitude
	TotalElapsedTime              uint32 // Time (includes pauses)
	TotalTimerTime                uint32 // Timer Time (excludes pauses)
	TotalDistance                 uint32
	TotalCycles                   uint32
	TotalCalories                 uint16
	TotalFatCalories              uint16 // If New Leaf
	AvgSpeed                      uint16
	MaxSpeed                      uint16
	AvgHeartRate                  uint8
	MaxHeartRate                  uint8
	AvgCadence                    uint8 // total_cycles / total_timer_time if non_zero_avg_cadence otherwise total_cycles / total_elapsed_time
	MaxCadence                    uint8
	AvgPower                      uint16 // total_power / total_timer_time if non_zero_avg_power otherwise total_power / total_elapsed_time
	MaxPower                      uint16
	TotalAscent                   uint16
	TotalDescent                  uint16
	Intensity                     Intensity
	LapTrigger                    LapTrigger
	Sport                         Sport
	EventGroup                    uint8
	NumLengths                    uint16 // # of lengths of swim pool
	NormalizedPower               uint16
	LeftRightBalance              LeftRightBalance100
	FirstLengthIndex              uint16
	AvgStrokeDistance             uint16
	SwimStroke                    SwimStroke
	SubSport                      SubSport
	NumActiveLengths              uint16 // # of active lengths of swim pool
	TotalWork                     uint32
	AvgAltitude                   uint16
	MaxAltitude                   uint16
	GpsAccuracy                   uint8
	AvgGrade                      int16
	AvgPosGrade                   int16
	AvgNegGrade                   int16
	MaxPosGrade                   int16
	MaxNegGrade                   int16
	AvgTemperature                int8
	MaxTemperature                int8
	TotalMovingTime               uint32
	AvgPosVerticalSpeed           int16
	AvgNegVerticalSpeed           int16
	MaxPosVerticalSpeed           int16
	MaxNegVerticalSpeed           int16
	TimeInHrZone                  []uint32
	TimeInSpeedZone               []uint32
	TimeInCadenceZone             []uint32
	TimeInPowerZone               []uint32
	RepetitionNum                 uint16
	MinAltitude                   uint16
	MinHeartRate                  uint8
	WktStepIndex                  MessageIndex
	OpponentScore                 uint16
	StrokeCount                   []uint16 // stroke_type enum used as the index
	ZoneCount                     []uint16 // zone number used as the index
	AvgVerticalOscillation        uint16
	AvgStanceTimePercent          uint16
	AvgStanceTime                 uint16
	AvgFractionalCadence          uint8 // fractional part of the avg_cadence
	MaxFractionalCadence          uint8 // fractional part of the max_cadence
	TotalFractionalCycles         uint8 // fractional part of the total_cycles
	PlayerScore                   uint16
	AvgTotalHemoglobinConc        []uint16 // Avg saturated and unsaturated hemoglobin
	MinTotalHemoglobinConc        []uint16 // Min saturated and unsaturated hemoglobin
	MaxTotalHemoglobinConc        []uint16 // Max saturated and unsaturated hemoglobin
	AvgSaturatedHemoglobinPercent []uint16 // Avg percentage of hemoglobin saturated with oxygen
	MinSaturatedHemoglobinPercent []uint16 // Min percentage of hemoglobin saturated with oxygen
	MaxSaturatedHemoglobinPercent []uint16 // Max percentage of hemoglobin saturated with oxygen
	EnhancedAvgSpeed              uint32
	EnhancedMaxSpeed              uint32
	EnhancedAvgAltitude           uint32
	EnhancedMinAltitude           uint32
	EnhancedMaxAltitude           uint32
	AvgVam                        uint16
}

// GetTotalElapsedTimeScaled returns TotalElapsedTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *LapMsg) GetTotalElapsedTimeScaled() float64 {
	if x.TotalElapsedTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalElapsedTime) / 1000
}

// GetTotalTimerTimeScaled returns TotalTimerTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *LapMsg) GetTotalTimerTimeScaled() float64 {
	if x.TotalTimerTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalTimerTime) / 1000
}

// GetTotalDistanceScaled returns TotalDistance
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *LapMsg) GetTotalDistanceScaled() float64 {
	if x.TotalDistance == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalDistance) / 100
}

// GetAvgSpeedScaled returns AvgSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *LapMsg) GetAvgSpeedScaled() float64 {
	if x.AvgSpeed == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgSpeed) / 1000
}

// GetMaxSpeedScaled returns MaxSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *LapMsg) GetMaxSpeedScaled() float64 {
	if x.MaxSpeed == 0xFFFF {
		return math.NaN()
	}
	return float64(x.MaxSpeed) / 1000
}

// GetAvgStrokeDistanceScaled returns AvgStrokeDistance
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *LapMsg) GetAvgStrokeDistanceScaled() float64 {
	if x.AvgStrokeDistance == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgStrokeDistance) / 100
}

// GetAvgAltitudeScaled returns AvgAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *LapMsg) GetAvgAltitudeScaled() float64 {
	if x.AvgAltitude == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgAltitude)/5 - 500
}

// GetMaxAltitudeScaled returns MaxAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *LapMsg) GetMaxAltitudeScaled() float64 {
	if x.MaxAltitude == 0xFFFF {
		return math.NaN()
	}
	return float64(x.MaxAltitude)/5 - 500
}

// GetAvgGradeScaled returns AvgGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *LapMsg) GetAvgGradeScaled() float64 {
	if x.AvgGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgGrade) / 100
}

// GetAvgPosGradeScaled returns AvgPosGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *LapMsg) GetAvgPosGradeScaled() float64 {
	if x.AvgPosGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgPosGrade) / 100
}

// GetAvgNegGradeScaled returns AvgNegGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *LapMsg) GetAvgNegGradeScaled() float64 {
	if x.AvgNegGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgNegGrade) / 100
}

// GetMaxPosGradeScaled returns MaxPosGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *LapMsg) GetMaxPosGradeScaled() float64 {
	if x.MaxPosGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.MaxPosGrade) / 100
}

// GetMaxNegGradeScaled returns MaxNegGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *LapMsg) GetMaxNegGradeScaled() float64 {
	if x.MaxNegGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.MaxNegGrade) / 100
}

// GetTotalMovingTimeScaled returns TotalMovingTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *LapMsg) GetTotalMovingTimeScaled() float64 {
	if x.TotalMovingTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalMovingTime) / 1000
}

// GetAvgPosVerticalSpeedScaled returns AvgPosVerticalSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *LapMsg) GetAvgPosVerticalSpeedScaled() float64 {
	if x.AvgPosVerticalSpeed == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgPosVerticalSpeed) / 1000
}

// GetAvgNegVerticalSpeedScaled returns AvgNegVerticalSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *LapMsg) GetAvgNegVerticalSpeedScaled() float64 {
	if x.AvgNegVerticalSpeed == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgNegVerticalSpeed) / 1000
}

// GetMaxPosVerticalSpeedScaled returns MaxPosVerticalSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *LapMsg) GetMaxPosVerticalSpeedScaled() float64 {
	if x.MaxPosVerticalSpeed == 0x7FFF {
		return math.NaN()
	}
	return float64(x.MaxPosVerticalSpeed) / 1000
}

// GetMaxNegVerticalSpeedScaled returns MaxNegVerticalSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *LapMsg) GetMaxNegVerticalSpeedScaled() float64 {
	if x.MaxNegVerticalSpeed == 0x7FFF {
		return math.NaN()
	}
	return float64(x.MaxNegVerticalSpeed) / 1000
}

// GetTimeInHrZoneScaled returns TimeInHrZone
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *LapMsg) GetTimeInHrZoneScaled() []float64 {
	if len(x.TimeInHrZone) == 0 {
		return nil
	}
	s := make([]float64, len(x.TimeInHrZone))
	for i, v := range x.TimeInHrZone {
		s[i] = float64(v) / 1000
	}
	return s
}

// GetTimeInSpeedZoneScaled returns TimeInSpeedZone
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *LapMsg) GetTimeInSpeedZoneScaled() []float64 {
	if len(x.TimeInSpeedZone) == 0 {
		return nil
	}
	s := make([]float64, len(x.TimeInSpeedZone))
	for i, v := range x.TimeInSpeedZone {
		s[i] = float64(v) / 1000
	}
	return s
}

// GetTimeInCadenceZoneScaled returns TimeInCadenceZone
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *LapMsg) GetTimeInCadenceZoneScaled() []float64 {
	if len(x.TimeInCadenceZone) == 0 {
		return nil
	}
	s := make([]float64, len(x.TimeInCadenceZone))
	for i, v := range x.TimeInCadenceZone {
		s[i] = float64(v) / 1000
	}
	return s
}

// GetTimeInPowerZoneScaled returns TimeInPowerZone
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *LapMsg) GetTimeInPowerZoneScaled() []float64 {
	if len(x.TimeInPowerZone) == 0 {
		return nil
	}
	s := make([]float64, len(x.TimeInPowerZone))
	for i, v := range x.TimeInPowerZone {
		s[i] = float64(v) / 1000
	}
	return s
}

// GetMinAltitudeScaled returns MinAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *LapMsg) GetMinAltitudeScaled() float64 {
	if x.MinAltitude == 0xFFFF {
		return math.NaN()
	}
	return float64(x.MinAltitude)/5 - 500
}

// GetAvgVerticalOscillationScaled returns AvgVerticalOscillation
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: mm
func (x *LapMsg) GetAvgVerticalOscillationScaled() float64 {
	if x.AvgVerticalOscillation == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgVerticalOscillation) / 10
}

// GetAvgStanceTimePercentScaled returns AvgStanceTimePercent
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: percent
func (x *LapMsg) GetAvgStanceTimePercentScaled() float64 {
	if x.AvgStanceTimePercent == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgStanceTimePercent) / 100
}

// GetAvgStanceTimeScaled returns AvgStanceTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: ms
func (x *LapMsg) GetAvgStanceTimeScaled() float64 {
	if x.AvgStanceTime == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgStanceTime) / 10
}

// GetAvgFractionalCadenceScaled returns AvgFractionalCadence
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: rpm
func (x *LapMsg) GetAvgFractionalCadenceScaled() float64 {
	if x.AvgFractionalCadence == 0xFF {
		return math.NaN()
	}
	return float64(x.AvgFractionalCadence) / 128
}

// GetMaxFractionalCadenceScaled returns MaxFractionalCadence
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: rpm
func (x *LapMsg) GetMaxFractionalCadenceScaled() float64 {
	if x.MaxFractionalCadence == 0xFF {
		return math.NaN()
	}
	return float64(x.MaxFractionalCadence) / 128
}

// GetTotalFractionalCyclesScaled returns TotalFractionalCycles
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: cycles
func (x *LapMsg) GetTotalFractionalCyclesScaled() float64 {
	if x.TotalFractionalCycles == 0xFF {
		return math.NaN()
	}
	return float64(x.TotalFractionalCycles) / 128
}

// GetAvgTotalHemoglobinConcScaled returns AvgTotalHemoglobinConc
// as a slice with scale and any offset applied to every element.
// Units: g/dL
func (x *LapMsg) GetAvgTotalHemoglobinConcScaled() []float64 {
	if len(x.AvgTotalHemoglobinConc) == 0 {
		return nil
	}
	s := make([]float64, len(x.AvgTotalHemoglobinConc))
	for i, v := range x.AvgTotalHemoglobinConc {
		s[i] = float64(v) / 100
	}
	return s
}

// GetMinTotalHemoglobinConcScaled returns MinTotalHemoglobinConc
// as a slice with scale and any offset applied to every element.
// Units: g/dL
func (x *LapMsg) GetMinTotalHemoglobinConcScaled() []float64 {
	if len(x.MinTotalHemoglobinConc) == 0 {
		return nil
	}
	s := make([]float64, len(x.MinTotalHemoglobinConc))
	for i, v := range x.MinTotalHemoglobinConc {
		s[i] = float64(v) / 100
	}
	return s
}

// GetMaxTotalHemoglobinConcScaled returns MaxTotalHemoglobinConc
// as a slice with scale and any offset applied to every element.
// Units: g/dL
func (x *LapMsg) GetMaxTotalHemoglobinConcScaled() []float64 {
	if len(x.MaxTotalHemoglobinConc) == 0 {
		return nil
	}
	s := make([]float64, len(x.MaxTotalHemoglobinConc))
	for i, v := range x.MaxTotalHemoglobinConc {
		s[i] = float64(v) / 100
	}
	return s
}

// GetAvgSaturatedHemoglobinPercentScaled returns AvgSaturatedHemoglobinPercent
// as a slice with scale and any offset applied to every element.
// Units: %
func (x *LapMsg) GetAvgSaturatedHemoglobinPercentScaled() []float64 {
	if len(x.AvgSaturatedHemoglobinPercent) == 0 {
		return nil
	}
	s := make([]float64, len(x.AvgSaturatedHemoglobinPercent))
	for i, v := range x.AvgSaturatedHemoglobinPercent {
		s[i] = float64(v) / 10
	}
	return s
}

// GetMinSaturatedHemoglobinPercentScaled returns MinSaturatedHemoglobinPercent
// as a slice with scale and any offset applied to every element.
// Units: %
func (x *LapMsg) GetMinSaturatedHemoglobinPercentScaled() []float64 {
	if len(x.MinSaturatedHemoglobinPercent) == 0 {
		return nil
	}
	s := make([]float64, len(x.MinSaturatedHemoglobinPercent))
	for i, v := range x.MinSaturatedHemoglobinPercent {
		s[i] = float64(v) / 10
	}
	return s
}

// GetMaxSaturatedHemoglobinPercentScaled returns MaxSaturatedHemoglobinPercent
// as a slice with scale and any offset applied to every element.
// Units: %
func (x *LapMsg) GetMaxSaturatedHemoglobinPercentScaled() []float64 {
	if len(x.MaxSaturatedHemoglobinPercent) == 0 {
		return nil
	}
	s := make([]float64, len(x.MaxSaturatedHemoglobinPercent))
	for i, v := range x.MaxSaturatedHemoglobinPercent {
		s[i] = float64(v) / 10
	}
	return s
}

// GetEnhancedAvgSpeedScaled returns EnhancedAvgSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *LapMsg) GetEnhancedAvgSpeedScaled() float64 {
	if x.EnhancedAvgSpeed == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.EnhancedAvgSpeed) / 1000
}

// GetEnhancedMaxSpeedScaled returns EnhancedMaxSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *LapMsg) GetEnhancedMaxSpeedScaled() float64 {
	if x.EnhancedMaxSpeed == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.EnhancedMaxSpeed) / 1000
}

// GetEnhancedAvgAltitudeScaled returns EnhancedAvgAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *LapMsg) GetEnhancedAvgAltitudeScaled() float64 {
	if x.EnhancedAvgAltitude == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.EnhancedAvgAltitude)/5 - 500
}

// GetEnhancedMinAltitudeScaled returns EnhancedMinAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *LapMsg) GetEnhancedMinAltitudeScaled() float64 {
	if x.EnhancedMinAltitude == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.EnhancedMinAltitude)/5 - 500
}

// GetEnhancedMaxAltitudeScaled returns EnhancedMaxAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *LapMsg) GetEnhancedMaxAltitudeScaled() float64 {
	if x.EnhancedMaxAltitude == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.EnhancedMaxAltitude)/5 - 500
}

// GetAvgVamScaled returns AvgVam
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *LapMsg) GetAvgVamScaled() float64 {
	if x.AvgVam == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgVam) / 1000
}

// GetTotalCycles returns the appropriate TotalCycles
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *LapMsg) GetTotalCycles() interface{} {
	switch x.Sport {
	case SportRunning, SportWalking:
		return uint32(x.TotalCycles)
	default:
		return x.TotalCycles
	}
}

// GetAvgCadence returns the appropriate AvgCadence
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *LapMsg) GetAvgCadence() interface{} {
	switch x.Sport {
	case SportRunning:
		return uint8(x.AvgCadence)
	default:
		return x.AvgCadence
	}
}

// GetMaxCadence returns the appropriate MaxCadence
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *LapMsg) GetMaxCadence() interface{} {
	switch x.Sport {
	case SportRunning:
		return uint8(x.MaxCadence)
	default:
		return x.MaxCadence
	}
}

func (x *LapMsg) expandComponents() {
	if x.AvgSpeed != 0xFFFF {
		x.EnhancedAvgSpeed = uint32(
			(x.AvgSpeed >> 0) & ((1 << 16) - 1),
		)
	}
	if x.MaxSpeed != 0xFFFF {
		x.EnhancedMaxSpeed = uint32(
			(x.MaxSpeed >> 0) & ((1 << 16) - 1),
		)
	}
	if x.AvgAltitude != 0xFFFF {
		x.EnhancedAvgAltitude = uint32(
			(x.AvgAltitude >> 0) & ((1 << 16) - 1),
		)
	}
	if x.MaxAltitude != 0xFFFF {
		x.EnhancedMaxAltitude = uint32(
			(x.MaxAltitude >> 0) & ((1 << 16) - 1),
		)
	}
	if x.MinAltitude != 0xFFFF {
		x.EnhancedMinAltitude = uint32(
			(x.MinAltitude >> 0) & ((1 << 16) - 1),
		)
	}
}

// LengthMsg represents the length FIT message type.
type LengthMsg struct {
	MessageIndex       MessageIndex
	Timestamp          time.Time
	Event              Event
	EventType          EventType
	StartTime          time.Time
	TotalElapsedTime   uint32
	TotalTimerTime     uint32
	TotalStrokes       uint16
	AvgSpeed           uint16
	SwimStroke         SwimStroke
	AvgSwimmingCadence uint8
	EventGroup         uint8
	TotalCalories      uint16
	LengthType         LengthType
	PlayerScore        uint16
	OpponentScore      uint16
	StrokeCount        []uint16 // stroke_type enum used as the index
	ZoneCount          []uint16 // zone number used as the index
}

// GetTotalElapsedTimeScaled returns TotalElapsedTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *LengthMsg) GetTotalElapsedTimeScaled() float64 {
	if x.TotalElapsedTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalElapsedTime) / 1000
}

// GetTotalTimerTimeScaled returns TotalTimerTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *LengthMsg) GetTotalTimerTimeScaled() float64 {
	if x.TotalTimerTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalTimerTime) / 1000
}

// GetAvgSpeedScaled returns AvgSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *LengthMsg) GetAvgSpeedScaled() float64 {
	if x.AvgSpeed == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgSpeed) / 1000
}

// RecordMsg represents the record FIT message type.
type RecordMsg struct {
	Timestamp                     time.Time
	PositionLat                   Latitude
	PositionLong                  Longitude
	Altitude                      uint16
	HeartRate                     uint8
	Cadence                       uint8
	Distance                      uint32
	Speed                         uint16
	Power                         uint16
	CompressedSpeedDistance       []byte
	Grade                         int16
	Resistance                    uint8 // Relative. 0 is none  254 is Max.
	TimeFromCourse                int32
	CycleLength                   uint8
	Temperature                   int8
	Speed1s                       []uint8 // Speed at 1s intervals.  Timestamp field indicates time of last array element.
	Cycles                        uint8
	TotalCycles                   uint32
	CompressedAccumulatedPower    uint16
	AccumulatedPower              uint32
	LeftRightBalance              LeftRightBalance
	GpsAccuracy                   uint8
	VerticalSpeed                 int16
	Calories                      uint16
	VerticalOscillation           uint16
	StanceTimePercent             uint16
	StanceTime                    uint16
	ActivityType                  ActivityType
	LeftTorqueEffectiveness       uint8
	RightTorqueEffectiveness      uint8
	LeftPedalSmoothness           uint8
	RightPedalSmoothness          uint8
	CombinedPedalSmoothness       uint8
	Time128                       uint8
	StrokeType                    StrokeType
	Zone                          uint8
	BallSpeed                     uint16
	Cadence256                    uint16 // Log cadence and fractional cadence for backwards compatability
	FractionalCadence             uint8
	TotalHemoglobinConc           uint16 // Total saturated and unsaturated hemoglobin
	TotalHemoglobinConcMin        uint16 // Min saturated and unsaturated hemoglobin
	TotalHemoglobinConcMax        uint16 // Max saturated and unsaturated hemoglobin
	SaturatedHemoglobinPercent    uint16 // Percentage of hemoglobin saturated with oxygen
	SaturatedHemoglobinPercentMin uint16 // Min percentage of hemoglobin saturated with oxygen
	SaturatedHemoglobinPercentMax uint16 // Max percentage of hemoglobin saturated with oxygen
	DeviceIndex                   DeviceIndex
	EnhancedSpeed                 uint32
	EnhancedAltitude              uint32
}

// GetAltitudeScaled returns Altitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *RecordMsg) GetAltitudeScaled() float64 {
	if x.Altitude == 0xFFFF {
		return math.NaN()
	}
	return float64(x.Altitude)/5 - 500
}

// GetDistanceScaled returns Distance
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *RecordMsg) GetDistanceScaled() float64 {
	if x.Distance == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.Distance) / 100
}

// GetSpeedScaled returns Speed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *RecordMsg) GetSpeedScaled() float64 {
	if x.Speed == 0xFFFF {
		return math.NaN()
	}
	return float64(x.Speed) / 1000
}

// GetGradeScaled returns Grade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *RecordMsg) GetGradeScaled() float64 {
	if x.Grade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.Grade) / 100
}

// GetTimeFromCourseScaled returns TimeFromCourse
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *RecordMsg) GetTimeFromCourseScaled() float64 {
	if x.TimeFromCourse == 0x7FFFFFFF {
		return math.NaN()
	}
	return float64(x.TimeFromCourse) / 1000
}

// GetCycleLengthScaled returns CycleLength
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *RecordMsg) GetCycleLengthScaled() float64 {
	if x.CycleLength == 0xFF {
		return math.NaN()
	}
	return float64(x.CycleLength) / 100
}

// GetSpeed1sScaled returns Speed1s
// as a slice with scale and any offset applied to every element.
// Units: m/s
func (x *RecordMsg) GetSpeed1sScaled() []float64 {
	if len(x.Speed1s) == 0 {
		return nil
	}
	s := make([]float64, len(x.Speed1s))
	for i, v := range x.Speed1s {
		s[i] = float64(v) / 16
	}
	return s
}

// GetVerticalSpeedScaled returns VerticalSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *RecordMsg) GetVerticalSpeedScaled() float64 {
	if x.VerticalSpeed == 0x7FFF {
		return math.NaN()
	}
	return float64(x.VerticalSpeed) / 1000
}

// GetVerticalOscillationScaled returns VerticalOscillation
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: mm
func (x *RecordMsg) GetVerticalOscillationScaled() float64 {
	if x.VerticalOscillation == 0xFFFF {
		return math.NaN()
	}
	return float64(x.VerticalOscillation) / 10
}

// GetStanceTimePercentScaled returns StanceTimePercent
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: percent
func (x *RecordMsg) GetStanceTimePercentScaled() float64 {
	if x.StanceTimePercent == 0xFFFF {
		return math.NaN()
	}
	return float64(x.StanceTimePercent) / 100
}

// GetStanceTimeScaled returns StanceTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: ms
func (x *RecordMsg) GetStanceTimeScaled() float64 {
	if x.StanceTime == 0xFFFF {
		return math.NaN()
	}
	return float64(x.StanceTime) / 10
}

// GetLeftTorqueEffectivenessScaled returns LeftTorqueEffectiveness
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: percent
func (x *RecordMsg) GetLeftTorqueEffectivenessScaled() float64 {
	if x.LeftTorqueEffectiveness == 0xFF {
		return math.NaN()
	}
	return float64(x.LeftTorqueEffectiveness) / 2
}

// GetRightTorqueEffectivenessScaled returns RightTorqueEffectiveness
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: percent
func (x *RecordMsg) GetRightTorqueEffectivenessScaled() float64 {
	if x.RightTorqueEffectiveness == 0xFF {
		return math.NaN()
	}
	return float64(x.RightTorqueEffectiveness) / 2
}

// GetLeftPedalSmoothnessScaled returns LeftPedalSmoothness
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: percent
func (x *RecordMsg) GetLeftPedalSmoothnessScaled() float64 {
	if x.LeftPedalSmoothness == 0xFF {
		return math.NaN()
	}
	return float64(x.LeftPedalSmoothness) / 2
}

// GetRightPedalSmoothnessScaled returns RightPedalSmoothness
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: percent
func (x *RecordMsg) GetRightPedalSmoothnessScaled() float64 {
	if x.RightPedalSmoothness == 0xFF {
		return math.NaN()
	}
	return float64(x.RightPedalSmoothness) / 2
}

// GetCombinedPedalSmoothnessScaled returns CombinedPedalSmoothness
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: percent
func (x *RecordMsg) GetCombinedPedalSmoothnessScaled() float64 {
	if x.CombinedPedalSmoothness == 0xFF {
		return math.NaN()
	}
	return float64(x.CombinedPedalSmoothness) / 2
}

// GetTime128Scaled returns Time128
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *RecordMsg) GetTime128Scaled() float64 {
	if x.Time128 == 0xFF {
		return math.NaN()
	}
	return float64(x.Time128) / 128
}

// GetBallSpeedScaled returns BallSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *RecordMsg) GetBallSpeedScaled() float64 {
	if x.BallSpeed == 0xFFFF {
		return math.NaN()
	}
	return float64(x.BallSpeed) / 100
}

// GetCadence256Scaled returns Cadence256
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: rpm
func (x *RecordMsg) GetCadence256Scaled() float64 {
	if x.Cadence256 == 0xFFFF {
		return math.NaN()
	}
	return float64(x.Cadence256) / 256
}

// GetFractionalCadenceScaled returns FractionalCadence
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: rpm
func (x *RecordMsg) GetFractionalCadenceScaled() float64 {
	if x.FractionalCadence == 0xFF {
		return math.NaN()
	}
	return float64(x.FractionalCadence) / 128
}

// GetTotalHemoglobinConcScaled returns TotalHemoglobinConc
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: g/dL
func (x *RecordMsg) GetTotalHemoglobinConcScaled() float64 {
	if x.TotalHemoglobinConc == 0xFFFF {
		return math.NaN()
	}
	return float64(x.TotalHemoglobinConc) / 100
}

// GetTotalHemoglobinConcMinScaled returns TotalHemoglobinConcMin
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: g/dL
func (x *RecordMsg) GetTotalHemoglobinConcMinScaled() float64 {
	if x.TotalHemoglobinConcMin == 0xFFFF {
		return math.NaN()
	}
	return float64(x.TotalHemoglobinConcMin) / 100
}

// GetTotalHemoglobinConcMaxScaled returns TotalHemoglobinConcMax
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: g/dL
func (x *RecordMsg) GetTotalHemoglobinConcMaxScaled() float64 {
	if x.TotalHemoglobinConcMax == 0xFFFF {
		return math.NaN()
	}
	return float64(x.TotalHemoglobinConcMax) / 100
}

// GetSaturatedHemoglobinPercentScaled returns SaturatedHemoglobinPercent
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *RecordMsg) GetSaturatedHemoglobinPercentScaled() float64 {
	if x.SaturatedHemoglobinPercent == 0xFFFF {
		return math.NaN()
	}
	return float64(x.SaturatedHemoglobinPercent) / 10
}

// GetSaturatedHemoglobinPercentMinScaled returns SaturatedHemoglobinPercentMin
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *RecordMsg) GetSaturatedHemoglobinPercentMinScaled() float64 {
	if x.SaturatedHemoglobinPercentMin == 0xFFFF {
		return math.NaN()
	}
	return float64(x.SaturatedHemoglobinPercentMin) / 10
}

// GetSaturatedHemoglobinPercentMaxScaled returns SaturatedHemoglobinPercentMax
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *RecordMsg) GetSaturatedHemoglobinPercentMaxScaled() float64 {
	if x.SaturatedHemoglobinPercentMax == 0xFFFF {
		return math.NaN()
	}
	return float64(x.SaturatedHemoglobinPercentMax) / 10
}

// GetEnhancedSpeedScaled returns EnhancedSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *RecordMsg) GetEnhancedSpeedScaled() float64 {
	if x.EnhancedSpeed == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.EnhancedSpeed) / 1000
}

// GetEnhancedAltitudeScaled returns EnhancedAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *RecordMsg) GetEnhancedAltitudeScaled() float64 {
	if x.EnhancedAltitude == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.EnhancedAltitude)/5 - 500
}

// GetSpeedFromCompressedSpeedDistance returns
// Speed with the scale and offset defined by the "Speed"
// component in the CompressedSpeedDistance field. NaN is
// if the field has an invalid value (i.e. has not been set).
func (x *RecordMsg) GetSpeedFromCompressedSpeedDistance() float64 {
	if x.Speed == 0xFFFF {
		return math.NaN()
	}
	return float64(x.Speed) / 100
}

// GetDistanceFromCompressedSpeedDistance returns
// Distance with the scale and offset defined by the "Distance"
// component in the CompressedSpeedDistance field. NaN is
// if the field has an invalid value (i.e. has not been set).
func (x *RecordMsg) GetDistanceFromCompressedSpeedDistance() float64 {
	if x.Distance == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.Distance) / 16
}

func (x *RecordMsg) expandComponents() {
	if x.Altitude != 0xFFFF {
		x.EnhancedAltitude = uint32(
			(x.Altitude >> 0) & ((1 << 16) - 1),
		)
	}
	if x.Speed != 0xFFFF {
		x.EnhancedSpeed = uint32(
			(x.Speed >> 0) & ((1 << 16) - 1),
		)
	}
	expand := false
	if len(x.CompressedSpeedDistance) == 3 {
		for _, v := range x.CompressedSpeedDistance {
			if v != 0xFF {
				expand = true
				break
			}
		}
	}
	if expand {
		x.Speed = uint16(x.CompressedSpeedDistance[0]) | uint16(x.CompressedSpeedDistance[1]&0x0F)<<8
		if accumuDistance == nil {
			accumuDistance = uint32NewAccumulator(12)
		}
		x.Distance = accumuDistance.accumulate(
			uint32(x.CompressedSpeedDistance[1]>>4) | uint32(x.CompressedSpeedDistance[2]<<4),
		)
	}
	if x.Cycles != 0xFF {
		if accumuTotalCycles == nil {
			accumuTotalCycles = new(uint32Accumulator)
		}
		x.TotalCycles = accumuTotalCycles.accumulate(
			uint32(
				(x.Cycles >> 0) & ((1 << 8) - 1),
			),
		)
	}
	if x.CompressedAccumulatedPower != 0xFFFF {
		if accumuAccumulatedPower == nil {
			accumuAccumulatedPower = new(uint32Accumulator)
		}
		x.AccumulatedPower = accumuAccumulatedPower.accumulate(
			uint32(
				(x.CompressedAccumulatedPower >> 0) & ((1 << 16) - 1),
			),
		)
	}
}

// EventMsg represents the event FIT message type.
type EventMsg struct {
	Timestamp     time.Time
	Event         Event
	EventType     EventType
	Data16        uint16
	Data          uint32
	EventGroup    uint8
	Score         uint16 // Do not populate directly.  Autogenerated by decoder for sport_point subfield components
	OpponentScore uint16 // Do not populate directly.  Autogenerated by decoder for sport_point subfield components
	FrontGearNum  uint8  // Do not populate directly.  Autogenerated by decoder for gear_change subfield components.  Front gear number. 1 is innermost.
	FrontGear     uint8  // Do not populate directly.  Autogenerated by decoder for gear_change subfield components.  Number of front teeth.
	RearGearNum   uint8  // Do not populate directly.  Autogenerated by decoder for gear_change subfield components.  Rear gear number. 1 is innermost.
	RearGear      uint8  // Do not populate directly.  Autogenerated by decoder for gear_change subfield components.  Number of rear teeth.
}

// GetData returns the appropriate Data
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *EventMsg) GetData() interface{} {
	switch x.Event {
	case EventTimer:
		return TimerTrigger(x.Data)
	case EventCoursePoint:
		return MessageIndex(x.Data)
	case EventBattery:
		return float64(x.Data) / 1000
	case EventVirtualPartnerPace:
		return float64(x.Data) / 1000
	case EventHrHighAlert:
		return uint8(x.Data)
	case EventHrLowAlert:
		return uint8(x.Data)
	case EventSpeedHighAlert:
		return float64(x.Data) / 1000
	case EventSpeedLowAlert:
		return float64(x.Data) / 1000
	case EventCadHighAlert:
		return uint16(x.Data)
	case EventCadLowAlert:
		return uint16(x.Data)
	case EventPowerHighAlert:
		return uint16(x.Data)
	case EventPowerLowAlert:
		return uint16(x.Data)
	case EventTimeDurationAlert:
		return float64(x.Data) / 1000
	case EventDistanceDurationAlert:
		return float64(x.Data) / 100
	case EventCalorieDurationAlert:
		return uint32(x.Data)
	case EventFitnessEquipment:
		return FitnessEquipmentState(x.Data)
	case EventSportPoint:
		return uint32(x.Data)
	case EventFrontGearChange, EventRearGearChange:
		return uint32(x.Data)
	default:
		return x.Data
	}
}

func (x *EventMsg) expandComponents() {
	if x.Data16 != 0xFFFF {
		x.Data = uint32(
			(x.Data16 >> 0) & ((1 << 16) - 1),
		)
	}
	if x.Data != 0xFFFFFFFF {
		switch x.Event {
		case EventSportPoint:
			x.Score = uint16(
				(x.Data >> 0) & ((1 << 16) - 1),
			)
			x.OpponentScore = uint16(
				(x.Data >> 16) & ((1 << 16) - 1),
			)
		case EventFrontGearChange, EventRearGearChange:
			x.RearGearNum = uint8(
				(x.Data >> 0) & ((1 << 8) - 1),
			)
			x.RearGear = uint8(
				(x.Data >> 8) & ((1 << 8) - 1),
			)
			x.FrontGearNum = uint8(
				(x.Data >> 16) & ((1 << 8) - 1),
			)
			x.FrontGear = uint8(
				(x.Data >> 24) & ((1 << 8) - 1),
			)
		}
	}
}

// DeviceInfoMsg represents the device_info FIT message type.
type DeviceInfoMsg struct {
	Timestamp           time.Time
	DeviceIndex         DeviceIndex
	DeviceType          uint8
	Manufacturer        Manufacturer
	SerialNumber        uint32
	Product             uint16
	SoftwareVersion     uint16
	HardwareVersion     uint8
	CumOperatingTime    uint32 // Reset by new battery or charge.
	BatteryVoltage      uint16
	BatteryStatus       BatteryStatus
	SensorPosition      BodyLocation // Indicates the location of the sensor
	Descriptor          string       // Used to describe the sensor or location
	AntTransmissionType uint8
	AntDeviceNumber     uint16
	AntNetwork          AntNetwork
	SourceType          SourceType
	ProductName         string // Optional free form string to indicate the devices name or model
}

// GetSoftwareVersionScaled returns SoftwareVersion
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
func (x *DeviceInfoMsg) GetSoftwareVersionScaled() float64 {
	if x.SoftwareVersion == 0xFFFF {
		return math.NaN()
	}
	return float64(x.SoftwareVersion) / 100
}

// GetBatteryVoltageScaled returns BatteryVoltage
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: V
func (x *DeviceInfoMsg) GetBatteryVoltageScaled() float64 {
	if x.BatteryVoltage == 0xFFFF {
		return math.NaN()
	}
	return float64(x.BatteryVoltage) / 256
}

// GetDeviceType returns the appropriate DeviceType
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *DeviceInfoMsg) GetDeviceType() interface{} {
	switch x.SourceType {
	case SourceTypeAntplus:
		return AntplusDeviceType(x.DeviceType)
	case SourceTypeAnt:
		return uint8(x.DeviceType)
	default:
		return x.DeviceType
	}
}

// GetProduct returns the appropriate Product
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *DeviceInfoMsg) GetProduct() interface{} {
	switch x.Manufacturer {
	case ManufacturerGarmin, ManufacturerDynastream, ManufacturerDynastreamOem:
		return GarminProduct(x.Product)
	default:
		return x.Product
	}
}

// TrainingFileMsg represents the training_file FIT message type.
type TrainingFileMsg struct {
	Timestamp    time.Time
	Type         FileType
	Manufacturer Manufacturer
	Product      uint16
	SerialNumber uint32
	TimeCreated  time.Time
}

// GetProduct returns the appropriate Product
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *TrainingFileMsg) GetProduct() interface{} {
	switch x.Manufacturer {
	case ManufacturerGarmin, ManufacturerDynastream, ManufacturerDynastreamOem:
		return GarminProduct(x.Product)
	default:
		return x.Product
	}
}

// HrvMsg represents the hrv FIT message type.
type HrvMsg struct {
	Time []uint16 // Time between beats
}

// GetTimeScaled returns Time
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *HrvMsg) GetTimeScaled() []float64 {
	if len(x.Time) == 0 {
		return nil
	}
	s := make([]float64, len(x.Time))
	for i, v := range x.Time {
		s[i] = float64(v) / 1000
	}
	return s
}

// WeatherConditionsMsg represents the weather_conditions FIT message type.
type WeatherConditionsMsg struct {
	Timestamp                time.Time     // time of update for current conditions, else forecast time
	WeatherReport            WeatherReport // Current or forecast
	Temperature              int8
	Condition                WeatherStatus // Corresponds to GSC Response weatherIcon field
	WindDirection            uint16
	WindSpeed                uint16
	PrecipitationProbability uint8 // range 0-100
	TemperatureFeelsLike     int8  // Heat Index if  GCS heatIdx above or equal to 90F or wind chill if GCS windChill below or equal to 32F
	RelativeHumidity         uint8
	Location                 string // string corresponding to GCS response location string
	ObservedAtTime           time.Time
	ObservedLocationLat      Latitude
	ObservedLocationLong     Longitude
	DayOfWeek                DayOfWeek
	HighTemperature          int8
	LowTemperature           int8
}

// GetWindSpeedScaled returns WindSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *WeatherConditionsMsg) GetWindSpeedScaled() float64 {
	if x.WindSpeed == 0xFFFF {
		return math.NaN()
	}
	return float64(x.WindSpeed) / 1000
}

// WeatherAlertMsg represents the weather_alert FIT message type.
type WeatherAlertMsg struct {
	Timestamp  time.Time
	ReportId   string            // Unique identifier from GCS report ID string, length is 12
	IssueTime  time.Time         // Time alert was issued
	ExpireTime time.Time         // Time alert expires
	Severity   WeatherSeverity   // Warning, Watch, Advisory, Statement
	Type       WeatherSevereType // Tornado, Severe Thunderstorm, etc.
}

// GpsMetadataMsg represents the gps_metadata FIT message type.
type GpsMetadataMsg struct {
}

// CameraEventMsg represents the camera_event FIT message type.
type CameraEventMsg struct {
}

// GyroscopeDataMsg represents the gyroscope_data FIT message type.
type GyroscopeDataMsg struct {
}

// AccelerometerDataMsg represents the accelerometer_data FIT message type.
type AccelerometerDataMsg struct {
}

// MagnetometerDataMsg represents the magnetometer_data FIT message type.
type MagnetometerDataMsg struct {
}

// ThreeDSensorCalibrationMsg represents the three_d_sensor_calibration FIT message type.
type ThreeDSensorCalibrationMsg struct {
}

// VideoFrameMsg represents the video_frame FIT message type.
type VideoFrameMsg struct {
}

// ObdiiDataMsg represents the obdii_data FIT message type.
type ObdiiDataMsg struct {
}

// NmeaSentenceMsg represents the nmea_sentence FIT message type.
type NmeaSentenceMsg struct {
	Timestamp   time.Time // Timestamp message was output
	TimestampMs uint16    // Fractional part of timestamp, added to timestamp
	Sentence    string    // NMEA sentence
}

// AviationAttitudeMsg represents the aviation_attitude FIT message type.
type AviationAttitudeMsg struct {
	Timestamp             time.Time // Timestamp message was output
	TimestampMs           uint16    // Fractional part of timestamp, added to timestamp
	SystemTime            []uint32  // System time associated with sample expressed in ms.
	Pitch                 []int16   // Range -PI/2 to +PI/2
	Roll                  []int16   // Range -PI to +PI
	AccelLateral          []int16   // Range -78.4 to +78.4 (-8 Gs to 8 Gs)
	AccelNormal           []int16   // Range -78.4 to +78.4 (-8 Gs to 8 Gs)
	TurnRate              []int16   // Range -8.727 to +8.727 (-500 degs/sec to +500 degs/sec)
	Stage                 []AttitudeStage
	AttitudeStageComplete []uint8  // The percent complete of the current attitude stage.  Set to 0 for attitude stages 0, 1 and 2 and to 100 for attitude stage 3 by AHRS modules that do not support it.  Range - 100
	Track                 []uint16 // Track Angle/Heading Range 0 - 2pi
	Validity              []AttitudeValidity
}

// GetPitchScaled returns Pitch
// as a slice with scale and any offset applied to every element.
// Units: radians
func (x *AviationAttitudeMsg) GetPitchScaled() []float64 {
	if len(x.Pitch) == 0 {
		return nil
	}
	s := make([]float64, len(x.Pitch))
	for i, v := range x.Pitch {
		s[i] = float64(v) / 10430.38
	}
	return s
}

// GetRollScaled returns Roll
// as a slice with scale and any offset applied to every element.
// Units: radians
func (x *AviationAttitudeMsg) GetRollScaled() []float64 {
	if len(x.Roll) == 0 {
		return nil
	}
	s := make([]float64, len(x.Roll))
	for i, v := range x.Roll {
		s[i] = float64(v) / 10430.38
	}
	return s
}

// GetAccelLateralScaled returns AccelLateral
// as a slice with scale and any offset applied to every element.
// Units: m/s^2
func (x *AviationAttitudeMsg) GetAccelLateralScaled() []float64 {
	if len(x.AccelLateral) == 0 {
		return nil
	}
	s := make([]float64, len(x.AccelLateral))
	for i, v := range x.AccelLateral {
		s[i] = float64(v) / 100
	}
	return s
}

// GetAccelNormalScaled returns AccelNormal
// as a slice with scale and any offset applied to every element.
// Units: m/s^2
func (x *AviationAttitudeMsg) GetAccelNormalScaled() []float64 {
	if len(x.AccelNormal) == 0 {
		return nil
	}
	s := make([]float64, len(x.AccelNormal))
	for i, v := range x.AccelNormal {
		s[i] = float64(v) / 100
	}
	return s
}

// GetTurnRateScaled returns TurnRate
// as a slice with scale and any offset applied to every element.
// Units: radians/second
func (x *AviationAttitudeMsg) GetTurnRateScaled() []float64 {
	if len(x.TurnRate) == 0 {
		return nil
	}
	s := make([]float64, len(x.TurnRate))
	for i, v := range x.TurnRate {
		s[i] = float64(v) / 1024
	}
	return s
}

// GetTrackScaled returns Track
// as a slice with scale and any offset applied to every element.
// Units: radians
func (x *AviationAttitudeMsg) GetTrackScaled() []float64 {
	if len(x.Track) == 0 {
		return nil
	}
	s := make([]float64, len(x.Track))
	for i, v := range x.Track {
		s[i] = float64(v) / 10430.38
	}
	return s
}

// VideoMsg represents the video FIT message type.
type VideoMsg struct {
}

// VideoTitleMsg represents the video_title FIT message type.
type VideoTitleMsg struct {
	MessageIndex MessageIndex // Long titles will be split into multiple parts
	MessageCount uint16       // Total number of title parts
	Text         string
}

// VideoDescriptionMsg represents the video_description FIT message type.
type VideoDescriptionMsg struct {
	MessageIndex MessageIndex // Long descriptions will be split into multiple parts
	MessageCount uint16       // Total number of description parts
	Text         string
}

// VideoClipMsg represents the video_clip FIT message type.
type VideoClipMsg struct {
}

// CourseMsg represents the course FIT message type.
type CourseMsg struct {
	Sport        Sport
	Name         string
	Capabilities CourseCapabilities
	SubSport     SubSport
}

// CoursePointMsg represents the course_point FIT message type.
type CoursePointMsg struct {
	MessageIndex MessageIndex
	Timestamp    time.Time
	PositionLat  Latitude
	PositionLong Longitude
	Distance     uint32
	Type         CoursePoint
	Name         string
	Favorite     Bool
}

// GetDistanceScaled returns Distance
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *CoursePointMsg) GetDistanceScaled() float64 {
	if x.Distance == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.Distance) / 100
}

// SegmentIdMsg represents the segment_id FIT message type.
type SegmentIdMsg struct {
	Name                  string               // Friendly name assigned to segment
	Uuid                  string               // UUID of the segment
	Sport                 Sport                // Sport associated with the segment
	Enabled               Bool                 // Segment enabled for evaluation
	UserProfilePrimaryKey uint32               // Primary key of the user that created the segment
	DeviceId              uint32               // ID of the device that created the segment
	DefaultRaceLeader     uint8                // Index for the Leader Board entry selected as the default race participant
	DeleteStatus          SegmentDeleteStatus  // Indicates if any segments should be deleted
	SelectionType         SegmentSelectionType // Indicates how the segment was selected to be sent to the device
}

// SegmentLeaderboardEntryMsg represents the segment_leaderboard_entry FIT message type.
type SegmentLeaderboardEntryMsg struct {
	MessageIndex    MessageIndex
	Name            string                 // Friendly name assigned to leader
	Type            SegmentLeaderboardType // Leader classification
	GroupPrimaryKey uint32                 // Primary user ID of this leader
	ActivityId      uint32                 // ID of the activity associated with this leader time
	SegmentTime     uint32                 // Segment Time (includes pauses)
}

// GetSegmentTimeScaled returns SegmentTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *SegmentLeaderboardEntryMsg) GetSegmentTimeScaled() float64 {
	if x.SegmentTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.SegmentTime) / 1000
}

// SegmentPointMsg represents the segment_point FIT message type.
type SegmentPointMsg struct {
	MessageIndex MessageIndex
	PositionLat  Latitude
	PositionLong Longitude
	Distance     uint32   // Accumulated distance along the segment at the described point
	Altitude     uint16   // Accumulated altitude along the segment at the described point
	LeaderTime   []uint32 // Accumualted time each leader board member required to reach the described point. This value is zero for all leader board members at the starting point of the segment.
}

// GetDistanceScaled returns Distance
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SegmentPointMsg) GetDistanceScaled() float64 {
	if x.Distance == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.Distance) / 100
}

// GetAltitudeScaled returns Altitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SegmentPointMsg) GetAltitudeScaled() float64 {
	if x.Altitude == 0xFFFF {
		return math.NaN()
	}
	return float64(x.Altitude)/5 - 500
}

// GetLeaderTimeScaled returns LeaderTime
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *SegmentPointMsg) GetLeaderTimeScaled() []float64 {
	if len(x.LeaderTime) == 0 {
		return nil
	}
	s := make([]float64, len(x.LeaderTime))
	for i, v := range x.LeaderTime {
		s[i] = float64(v) / 1000
	}
	return s
}

// SegmentLapMsg represents the segment_lap FIT message type.
type SegmentLapMsg struct {
	MessageIndex                MessageIndex
	Timestamp                   time.Time // Lap end time.
	Event                       Event
	EventType                   EventType
	StartTime                   time.Time
	StartPositionLat            Latitude
	StartPositionLong           Longitude
	EndPositionLat              Latitude
	EndPositionLong             Longitude
	TotalElapsedTime            uint32 // Time (includes pauses)
	TotalTimerTime              uint32 // Timer Time (excludes pauses)
	TotalDistance               uint32
	TotalCycles                 uint32
	TotalCalories               uint16
	TotalFatCalories            uint16 // If New Leaf
	AvgSpeed                    uint16
	MaxSpeed                    uint16
	AvgHeartRate                uint8
	MaxHeartRate                uint8
	AvgCadence                  uint8 // total_cycles / total_timer_time if non_zero_avg_cadence otherwise total_cycles / total_elapsed_time
	MaxCadence                  uint8
	AvgPower                    uint16 // total_power / total_timer_time if non_zero_avg_power otherwise total_power / total_elapsed_time
	MaxPower                    uint16
	TotalAscent                 uint16
	TotalDescent                uint16
	Sport                       Sport
	EventGroup                  uint8
	NecLat                      Latitude  // North east corner latitude.
	NecLong                     Longitude // North east corner longitude.
	SwcLat                      Latitude  // South west corner latitude.
	SwcLong                     Longitude // South west corner latitude.
	Name                        string
	NormalizedPower             uint16
	LeftRightBalance            LeftRightBalance100
	SubSport                    SubSport
	TotalWork                   uint32
	AvgAltitude                 uint16
	MaxAltitude                 uint16
	GpsAccuracy                 uint8
	AvgGrade                    int16
	AvgPosGrade                 int16
	AvgNegGrade                 int16
	MaxPosGrade                 int16
	MaxNegGrade                 int16
	AvgTemperature              int8
	MaxTemperature              int8
	TotalMovingTime             uint32
	AvgPosVerticalSpeed         int16
	AvgNegVerticalSpeed         int16
	MaxPosVerticalSpeed         int16
	MaxNegVerticalSpeed         int16
	TimeInHrZone                []uint32
	TimeInSpeedZone             []uint32
	TimeInCadenceZone           []uint32
	TimeInPowerZone             []uint32
	RepetitionNum               uint16
	MinAltitude                 uint16
	MinHeartRate                uint8
	ActiveTime                  uint32
	WktStepIndex                MessageIndex
	SportEvent                  SportEvent
	AvgLeftTorqueEffectiveness  uint8
	AvgRightTorqueEffectiveness uint8
	AvgLeftPedalSmoothness      uint8
	AvgRightPedalSmoothness     uint8
	AvgCombinedPedalSmoothness  uint8
	Status                      SegmentLapStatus
	Uuid                        string
	AvgFractionalCadence        uint8 // fractional part of the avg_cadence
	MaxFractionalCadence        uint8 // fractional part of the max_cadence
	TotalFractionalCycles       uint8 // fractional part of the total_cycles
	FrontGearShiftCount         uint16
	RearGearShiftCount          uint16
}

// GetTotalElapsedTimeScaled returns TotalElapsedTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *SegmentLapMsg) GetTotalElapsedTimeScaled() float64 {
	if x.TotalElapsedTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalElapsedTime) / 1000
}

// GetTotalTimerTimeScaled returns TotalTimerTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *SegmentLapMsg) GetTotalTimerTimeScaled() float64 {
	if x.TotalTimerTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalTimerTime) / 1000
}

// GetTotalDistanceScaled returns TotalDistance
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SegmentLapMsg) GetTotalDistanceScaled() float64 {
	if x.TotalDistance == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalDistance) / 100
}

// GetAvgSpeedScaled returns AvgSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SegmentLapMsg) GetAvgSpeedScaled() float64 {
	if x.AvgSpeed == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgSpeed) / 1000
}

// GetMaxSpeedScaled returns MaxSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SegmentLapMsg) GetMaxSpeedScaled() float64 {
	if x.MaxSpeed == 0xFFFF {
		return math.NaN()
	}
	return float64(x.MaxSpeed) / 1000
}

// GetAvgAltitudeScaled returns AvgAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SegmentLapMsg) GetAvgAltitudeScaled() float64 {
	if x.AvgAltitude == 0xFFFF {
		return math.NaN()
	}
	return float64(x.AvgAltitude)/5 - 500
}

// GetMaxAltitudeScaled returns MaxAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SegmentLapMsg) GetMaxAltitudeScaled() float64 {
	if x.MaxAltitude == 0xFFFF {
		return math.NaN()
	}
	return float64(x.MaxAltitude)/5 - 500
}

// GetAvgGradeScaled returns AvgGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *SegmentLapMsg) GetAvgGradeScaled() float64 {
	if x.AvgGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgGrade) / 100
}

// GetAvgPosGradeScaled returns AvgPosGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *SegmentLapMsg) GetAvgPosGradeScaled() float64 {
	if x.AvgPosGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgPosGrade) / 100
}

// GetAvgNegGradeScaled returns AvgNegGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *SegmentLapMsg) GetAvgNegGradeScaled() float64 {
	if x.AvgNegGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgNegGrade) / 100
}

// GetMaxPosGradeScaled returns MaxPosGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *SegmentLapMsg) GetMaxPosGradeScaled() float64 {
	if x.MaxPosGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.MaxPosGrade) / 100
}

// GetMaxNegGradeScaled returns MaxNegGrade
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *SegmentLapMsg) GetMaxNegGradeScaled() float64 {
	if x.MaxNegGrade == 0x7FFF {
		return math.NaN()
	}
	return float64(x.MaxNegGrade) / 100
}

// GetTotalMovingTimeScaled returns TotalMovingTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *SegmentLapMsg) GetTotalMovingTimeScaled() float64 {
	if x.TotalMovingTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.TotalMovingTime) / 1000
}

// GetAvgPosVerticalSpeedScaled returns AvgPosVerticalSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SegmentLapMsg) GetAvgPosVerticalSpeedScaled() float64 {
	if x.AvgPosVerticalSpeed == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgPosVerticalSpeed) / 1000
}

// GetAvgNegVerticalSpeedScaled returns AvgNegVerticalSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SegmentLapMsg) GetAvgNegVerticalSpeedScaled() float64 {
	if x.AvgNegVerticalSpeed == 0x7FFF {
		return math.NaN()
	}
	return float64(x.AvgNegVerticalSpeed) / 1000
}

// GetMaxPosVerticalSpeedScaled returns MaxPosVerticalSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SegmentLapMsg) GetMaxPosVerticalSpeedScaled() float64 {
	if x.MaxPosVerticalSpeed == 0x7FFF {
		return math.NaN()
	}
	return float64(x.MaxPosVerticalSpeed) / 1000
}

// GetMaxNegVerticalSpeedScaled returns MaxNegVerticalSpeed
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m/s
func (x *SegmentLapMsg) GetMaxNegVerticalSpeedScaled() float64 {
	if x.MaxNegVerticalSpeed == 0x7FFF {
		return math.NaN()
	}
	return float64(x.MaxNegVerticalSpeed) / 1000
}

// GetTimeInHrZoneScaled returns TimeInHrZone
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *SegmentLapMsg) GetTimeInHrZoneScaled() []float64 {
	if len(x.TimeInHrZone) == 0 {
		return nil
	}
	s := make([]float64, len(x.TimeInHrZone))
	for i, v := range x.TimeInHrZone {
		s[i] = float64(v) / 1000
	}
	return s
}

// GetTimeInSpeedZoneScaled returns TimeInSpeedZone
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *SegmentLapMsg) GetTimeInSpeedZoneScaled() []float64 {
	if len(x.TimeInSpeedZone) == 0 {
		return nil
	}
	s := make([]float64, len(x.TimeInSpeedZone))
	for i, v := range x.TimeInSpeedZone {
		s[i] = float64(v) / 1000
	}
	return s
}

// GetTimeInCadenceZoneScaled returns TimeInCadenceZone
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *SegmentLapMsg) GetTimeInCadenceZoneScaled() []float64 {
	if len(x.TimeInCadenceZone) == 0 {
		return nil
	}
	s := make([]float64, len(x.TimeInCadenceZone))
	for i, v := range x.TimeInCadenceZone {
		s[i] = float64(v) / 1000
	}
	return s
}

// GetTimeInPowerZoneScaled returns TimeInPowerZone
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *SegmentLapMsg) GetTimeInPowerZoneScaled() []float64 {
	if len(x.TimeInPowerZone) == 0 {
		return nil
	}
	s := make([]float64, len(x.TimeInPowerZone))
	for i, v := range x.TimeInPowerZone {
		s[i] = float64(v) / 1000
	}
	return s
}

// GetMinAltitudeScaled returns MinAltitude
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *SegmentLapMsg) GetMinAltitudeScaled() float64 {
	if x.MinAltitude == 0xFFFF {
		return math.NaN()
	}
	return float64(x.MinAltitude)/5 - 500
}

// GetActiveTimeScaled returns ActiveTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *SegmentLapMsg) GetActiveTimeScaled() float64 {
	if x.ActiveTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.ActiveTime) / 1000
}

// GetAvgLeftTorqueEffectivenessScaled returns AvgLeftTorqueEffectiveness
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: percent
func (x *SegmentLapMsg) GetAvgLeftTorqueEffectivenessScaled() float64 {
	if x.AvgLeftTorqueEffectiveness == 0xFF {
		return math.NaN()
	}
	return float64(x.AvgLeftTorqueEffectiveness) / 2
}

// GetAvgRightTorqueEffectivenessScaled returns AvgRightTorqueEffectiveness
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: percent
func (x *SegmentLapMsg) GetAvgRightTorqueEffectivenessScaled() float64 {
	if x.AvgRightTorqueEffectiveness == 0xFF {
		return math.NaN()
	}
	return float64(x.AvgRightTorqueEffectiveness) / 2
}

// GetAvgLeftPedalSmoothnessScaled returns AvgLeftPedalSmoothness
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: percent
func (x *SegmentLapMsg) GetAvgLeftPedalSmoothnessScaled() float64 {
	if x.AvgLeftPedalSmoothness == 0xFF {
		return math.NaN()
	}
	return float64(x.AvgLeftPedalSmoothness) / 2
}

// GetAvgRightPedalSmoothnessScaled returns AvgRightPedalSmoothness
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: percent
func (x *SegmentLapMsg) GetAvgRightPedalSmoothnessScaled() float64 {
	if x.AvgRightPedalSmoothness == 0xFF {
		return math.NaN()
	}
	return float64(x.AvgRightPedalSmoothness) / 2
}

// GetAvgCombinedPedalSmoothnessScaled returns AvgCombinedPedalSmoothness
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: percent
func (x *SegmentLapMsg) GetAvgCombinedPedalSmoothnessScaled() float64 {
	if x.AvgCombinedPedalSmoothness == 0xFF {
		return math.NaN()
	}
	return float64(x.AvgCombinedPedalSmoothness) / 2
}

// GetAvgFractionalCadenceScaled returns AvgFractionalCadence
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: rpm
func (x *SegmentLapMsg) GetAvgFractionalCadenceScaled() float64 {
	if x.AvgFractionalCadence == 0xFF {
		return math.NaN()
	}
	return float64(x.AvgFractionalCadence) / 128
}

// GetMaxFractionalCadenceScaled returns MaxFractionalCadence
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: rpm
func (x *SegmentLapMsg) GetMaxFractionalCadenceScaled() float64 {
	if x.MaxFractionalCadence == 0xFF {
		return math.NaN()
	}
	return float64(x.MaxFractionalCadence) / 128
}

// GetTotalFractionalCyclesScaled returns TotalFractionalCycles
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: cycles
func (x *SegmentLapMsg) GetTotalFractionalCyclesScaled() float64 {
	if x.TotalFractionalCycles == 0xFF {
		return math.NaN()
	}
	return float64(x.TotalFractionalCycles) / 128
}

// GetTotalCycles returns the appropriate TotalCycles
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *SegmentLapMsg) GetTotalCycles() interface{} {
	switch x.Sport {
	case SportCycling:
		return uint32(x.TotalCycles)
	default:
		return x.TotalCycles
	}
}

// SegmentFileMsg represents the segment_file FIT message type.
type SegmentFileMsg struct {
	MessageIndex          MessageIndex
	FileUuid              string                   // UUID of the segment file
	Enabled               Bool                     // Enabled state of the segment file
	UserProfilePrimaryKey uint32                   // Primary key of the user that created the segment file
	LeaderType            []SegmentLeaderboardType // Leader type of each leader in the segment file
	LeaderGroupPrimaryKey []uint32                 // Group primary key of each leader in the segment file
	LeaderActivityId      []uint32                 // Activity ID of each leader in the segment file
}

// WorkoutMsg represents the workout FIT message type.
type WorkoutMsg struct {
	Sport          Sport
	Capabilities   WorkoutCapabilities
	NumValidSteps  uint16 // number of valid steps
	WktName        string
	SubSport       SubSport
	PoolLength     uint16
	PoolLengthUnit DisplayMeasure
}

// GetPoolLengthScaled returns PoolLength
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *WorkoutMsg) GetPoolLengthScaled() float64 {
	if x.PoolLength == 0xFFFF {
		return math.NaN()
	}
	return float64(x.PoolLength) / 100
}

// WorkoutSessionMsg represents the workout_session FIT message type.
type WorkoutSessionMsg struct {
	MessageIndex   MessageIndex
	Sport          Sport
	SubSport       SubSport
	NumValidSteps  uint16
	FirstStepIndex uint16
	PoolLength     uint16
	PoolLengthUnit DisplayMeasure
}

// GetPoolLengthScaled returns PoolLength
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *WorkoutSessionMsg) GetPoolLengthScaled() float64 {
	if x.PoolLength == 0xFFFF {
		return math.NaN()
	}
	return float64(x.PoolLength) / 100
}

// WorkoutStepMsg represents the workout_step FIT message type.
type WorkoutStepMsg struct {
	MessageIndex          MessageIndex
	WktStepName           string
	DurationType          WktStepDuration
	DurationValue         uint32
	TargetType            WktStepTarget
	TargetValue           uint32
	CustomTargetValueLow  uint32
	CustomTargetValueHigh uint32
	Intensity             Intensity
	Notes                 string
	Equipment             WorkoutEquipment
}

// GetDurationValue returns the appropriate DurationValue
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *WorkoutStepMsg) GetDurationValue() interface{} {
	switch x.DurationType {
	case WktStepDurationTime, WktStepDurationRepetitionTime:
		return float64(x.DurationValue) / 1000
	case WktStepDurationDistance:
		return float64(x.DurationValue) / 100
	case WktStepDurationHrLessThan, WktStepDurationHrGreaterThan:
		return WorkoutHr(x.DurationValue)
	case WktStepDurationCalories:
		return uint32(x.DurationValue)
	case WktStepDurationRepeatUntilStepsCmplt, WktStepDurationRepeatUntilTime, WktStepDurationRepeatUntilDistance, WktStepDurationRepeatUntilCalories, WktStepDurationRepeatUntilHrLessThan, WktStepDurationRepeatUntilHrGreaterThan, WktStepDurationRepeatUntilPowerLessThan, WktStepDurationRepeatUntilPowerGreaterThan:
		return uint32(x.DurationValue)
	case WktStepDurationPowerLessThan, WktStepDurationPowerGreaterThan:
		return WorkoutPower(x.DurationValue)
	default:
		return x.DurationValue
	}
}

// GetTargetValue returns the appropriate TargetValue
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *WorkoutStepMsg) GetTargetValue() interface{} {
	switch {
	case x.TargetType == WktStepTargetSpeed:
		return uint32(x.TargetValue)
	case x.TargetType == WktStepTargetHeartRate:
		return uint32(x.TargetValue)
	case x.TargetType == WktStepTargetCadence:
		return uint32(x.TargetValue)
	case x.TargetType == WktStepTargetPower:
		return uint32(x.TargetValue)
	case x.DurationType == WktStepDurationRepeatUntilStepsCmplt:
		return uint32(x.TargetValue)
	case x.DurationType == WktStepDurationRepeatUntilTime:
		return float64(x.TargetValue) / 1000
	case x.DurationType == WktStepDurationRepeatUntilDistance:
		return float64(x.TargetValue) / 100
	case x.DurationType == WktStepDurationRepeatUntilCalories:
		return uint32(x.TargetValue)
	case x.DurationType == WktStepDurationRepeatUntilHrLessThan:
		return WorkoutHr(x.TargetValue)
	case x.DurationType == WktStepDurationRepeatUntilHrGreaterThan:
		return WorkoutHr(x.TargetValue)
	case x.DurationType == WktStepDurationRepeatUntilPowerLessThan:
		return WorkoutPower(x.TargetValue)
	case x.DurationType == WktStepDurationRepeatUntilPowerGreaterThan:
		return WorkoutPower(x.TargetValue)
	case x.TargetType == WktStepTargetSwimStroke:
		return SwimStroke(x.TargetValue)
	default:
		return x.TargetValue
	}
}

// GetCustomTargetValueLow returns the appropriate CustomTargetValueLow
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *WorkoutStepMsg) GetCustomTargetValueLow() interface{} {
	switch x.TargetType {
	case WktStepTargetSpeed:
		return float64(x.CustomTargetValueLow) / 1000
	case WktStepTargetHeartRate:
		return WorkoutHr(x.CustomTargetValueLow)
	case WktStepTargetCadence:
		return uint32(x.CustomTargetValueLow)
	case WktStepTargetPower:
		return WorkoutPower(x.CustomTargetValueLow)
	default:
		return x.CustomTargetValueLow
	}
}

// GetCustomTargetValueHigh returns the appropriate CustomTargetValueHigh
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *WorkoutStepMsg) GetCustomTargetValueHigh() interface{} {
	switch x.TargetType {
	case WktStepTargetSpeed:
		return float64(x.CustomTargetValueHigh) / 1000
	case WktStepTargetHeartRate:
		return WorkoutHr(x.CustomTargetValueHigh)
	case WktStepTargetCadence:
		return uint32(x.CustomTargetValueHigh)
	case WktStepTargetPower:
		return WorkoutPower(x.CustomTargetValueHigh)
	default:
		return x.CustomTargetValueHigh
	}
}

// ScheduleMsg represents the schedule FIT message type.
type ScheduleMsg struct {
	Manufacturer  Manufacturer // Corresponds to file_id of scheduled workout / course.
	Product       uint16       // Corresponds to file_id of scheduled workout / course.
	SerialNumber  uint32       // Corresponds to file_id of scheduled workout / course.
	TimeCreated   time.Time    // Corresponds to file_id of scheduled workout / course.
	Completed     Bool         // TRUE if this activity has been started
	Type          Schedule
	ScheduledTime time.Time
}

// GetProduct returns the appropriate Product
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *ScheduleMsg) GetProduct() interface{} {
	switch x.Manufacturer {
	case ManufacturerGarmin, ManufacturerDynastream, ManufacturerDynastreamOem:
		return GarminProduct(x.Product)
	default:
		return x.Product
	}
}

// TotalsMsg represents the totals FIT message type.
type TotalsMsg struct {
	MessageIndex MessageIndex
	Timestamp    time.Time
	TimerTime    uint32 // Excludes pauses
	Distance     uint32
	Calories     uint32
	Sport        Sport
	ElapsedTime  uint32 // Includes pauses
	Sessions     uint16
	ActiveTime   uint32
}

// WeightScaleMsg represents the weight_scale FIT message type.
type WeightScaleMsg struct {
	Timestamp         time.Time
	Weight            Weight
	PercentFat        uint16
	PercentHydration  uint16
	VisceralFatMass   uint16
	BoneMass          uint16
	MuscleMass        uint16
	BasalMet          uint16
	PhysiqueRating    uint8
	ActiveMet         uint16 // ~4kJ per kcal, 0.25 allows max 16384 kcal
	MetabolicAge      uint8
	VisceralFatRating uint8
	UserProfileIndex  MessageIndex // Associates this weight scale message to a user.  This corresponds to the index of the user profile message in the weight scale file.
}

// GetWeightScaled returns Weight
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: kg
func (x *WeightScaleMsg) GetWeightScaled() float64 {
	if x.Weight == 0xFFFF {
		return math.NaN()
	}
	return float64(x.Weight) / 100
}

// GetPercentFatScaled returns PercentFat
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *WeightScaleMsg) GetPercentFatScaled() float64 {
	if x.PercentFat == 0xFFFF {
		return math.NaN()
	}
	return float64(x.PercentFat) / 100
}

// GetPercentHydrationScaled returns PercentHydration
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: %
func (x *WeightScaleMsg) GetPercentHydrationScaled() float64 {
	if x.PercentHydration == 0xFFFF {
		return math.NaN()
	}
	return float64(x.PercentHydration) / 100
}

// GetVisceralFatMassScaled returns VisceralFatMass
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: kg
func (x *WeightScaleMsg) GetVisceralFatMassScaled() float64 {
	if x.VisceralFatMass == 0xFFFF {
		return math.NaN()
	}
	return float64(x.VisceralFatMass) / 100
}

// GetBoneMassScaled returns BoneMass
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: kg
func (x *WeightScaleMsg) GetBoneMassScaled() float64 {
	if x.BoneMass == 0xFFFF {
		return math.NaN()
	}
	return float64(x.BoneMass) / 100
}

// GetMuscleMassScaled returns MuscleMass
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: kg
func (x *WeightScaleMsg) GetMuscleMassScaled() float64 {
	if x.MuscleMass == 0xFFFF {
		return math.NaN()
	}
	return float64(x.MuscleMass) / 100
}

// GetBasalMetScaled returns BasalMet
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: kcal/day
func (x *WeightScaleMsg) GetBasalMetScaled() float64 {
	if x.BasalMet == 0xFFFF {
		return math.NaN()
	}
	return float64(x.BasalMet) / 4
}

// GetActiveMetScaled returns ActiveMet
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: kcal/day
func (x *WeightScaleMsg) GetActiveMetScaled() float64 {
	if x.ActiveMet == 0xFFFF {
		return math.NaN()
	}
	return float64(x.ActiveMet) / 4
}

// BloodPressureMsg represents the blood_pressure FIT message type.
type BloodPressureMsg struct {
	Timestamp            time.Time
	SystolicPressure     uint16
	DiastolicPressure    uint16
	MeanArterialPressure uint16
	Map3SampleMean       uint16
	MapMorningValues     uint16
	MapEveningValues     uint16
	HeartRate            uint8
	HeartRateType        HrType
	Status               BpStatus
	UserProfileIndex     MessageIndex // Associates this blood pressure message to a user.  This corresponds to the index of the user profile message in the blood pressure file.
}

// MonitoringInfoMsg represents the monitoring_info FIT message type.
type MonitoringInfoMsg struct {
	Timestamp      time.Time
	LocalTimestamp time.Time // Use to convert activity timestamps to local time if device does not support time zone and daylight savings time correction.
}

// MonitoringMsg represents the monitoring FIT message type.
type MonitoringMsg struct {
	Timestamp       time.Time   // Must align to logging interval, for example, time must be 00:00:00 for daily log.
	DeviceIndex     DeviceIndex // Associates this data to device_info message.  Not required for file with single device (sensor).
	Calories        uint16      // Accumulated total calories.  Maintained by MonitoringReader for each activity_type.  See SDK documentation
	Distance        uint32      // Accumulated distance.  Maintained by MonitoringReader for each activity_type.  See SDK documentation.
	Cycles          uint32      // Accumulated cycles.  Maintained by MonitoringReader for each activity_type.  See SDK documentation.
	ActiveTime      uint32
	ActivityType    ActivityType
	ActivitySubtype ActivitySubtype
	Distance16      uint16
	Cycles16        uint16
	ActiveTime16    uint16
	LocalTimestamp  time.Time // Must align to logging interval, for example, time must be 00:00:00 for daily log.
}

// GetDistanceScaled returns Distance
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: m
func (x *MonitoringMsg) GetDistanceScaled() float64 {
	if x.Distance == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.Distance) / 100
}

// GetCyclesScaled returns Cycles
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: cycles
func (x *MonitoringMsg) GetCyclesScaled() float64 {
	if x.Cycles == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.Cycles) / 2
}

// GetActiveTimeScaled returns ActiveTime
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *MonitoringMsg) GetActiveTimeScaled() float64 {
	if x.ActiveTime == 0xFFFFFFFF {
		return math.NaN()
	}
	return float64(x.ActiveTime) / 1000
}

// GetCycles returns the appropriate Cycles
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *MonitoringMsg) GetCycles() interface{} {
	switch x.ActivityType {
	case ActivityTypeCycling, ActivityTypeSwimming:
		return float64(x.Cycles) / 2
	default:
		return x.Cycles
	}
}

// HrMsg represents the hr FIT message type.
type HrMsg struct {
	Timestamp           time.Time
	FractionalTimestamp uint16
	Time256             uint8
	FilteredBpm         []uint8
	EventTimestamp      []uint32
	EventTimestamp12    []byte
}

// GetFractionalTimestampScaled returns FractionalTimestamp
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *HrMsg) GetFractionalTimestampScaled() float64 {
	if x.FractionalTimestamp == 0xFFFF {
		return math.NaN()
	}
	return float64(x.FractionalTimestamp) / 32768
}

// GetTime256Scaled returns Time256
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *HrMsg) GetTime256Scaled() float64 {
	if x.Time256 == 0xFF {
		return math.NaN()
	}
	return float64(x.Time256) / 256
}

// GetEventTimestampScaled returns EventTimestamp
// as a slice with scale and any offset applied to every element.
// Units: s
func (x *HrMsg) GetEventTimestampScaled() []float64 {
	if len(x.EventTimestamp) == 0 {
		return nil
	}
	s := make([]float64, len(x.EventTimestamp))
	for i, v := range x.EventTimestamp {
		s[i] = float64(v) / 1024
	}
	return s
}

func (x *HrMsg) expandComponents() {
	if x.Time256 != 0xFF {
	}
	// TODO
}

// MemoGlobMsg represents the memo_glob FIT message type.
type MemoGlobMsg struct {
}

// AntChannelIdMsg represents the ant_channel_id FIT message type.
type AntChannelIdMsg struct {
}

// AntRxMsg represents the ant_rx FIT message type.
type AntRxMsg struct {
	Timestamp           time.Time
	FractionalTimestamp uint16
	MesgId              byte
	MesgData            []byte
	ChannelNumber       uint8
	Data                []byte
}

// GetFractionalTimestampScaled returns FractionalTimestamp
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *AntRxMsg) GetFractionalTimestampScaled() float64 {
	if x.FractionalTimestamp == 0xFFFF {
		return math.NaN()
	}
	return float64(x.FractionalTimestamp) / 32768
}

func (x *AntRxMsg) expandComponents() {
	if len(x.MesgData) != 0 {
		x.Data = make([]byte, len(x.MesgData)-1)
		for i, v := range x.MesgData {
			if v == 0xFF {
				break
			}
			if i == 0 {
				x.ChannelNumber = v
			} else {
				x.Data[i-1] = v
			}
		}
	}
}

// AntTxMsg represents the ant_tx FIT message type.
type AntTxMsg struct {
	Timestamp           time.Time
	FractionalTimestamp uint16
	MesgId              byte
	MesgData            []byte
	ChannelNumber       uint8
	Data                []byte
}

// GetFractionalTimestampScaled returns FractionalTimestamp
// with scale and any offset applied. NaN is returned if the
// field has an invalid value (i.e. has not been set).
// Units: s
func (x *AntTxMsg) GetFractionalTimestampScaled() float64 {
	if x.FractionalTimestamp == 0xFFFF {
		return math.NaN()
	}
	return float64(x.FractionalTimestamp) / 32768
}

func (x *AntTxMsg) expandComponents() {
	if len(x.MesgData) != 0 {
		x.Data = make([]byte, len(x.MesgData)-1)
		for i, v := range x.MesgData {
			if v == 0xFF {
				break
			}
			if i == 0 {
				x.ChannelNumber = v
			} else {
				x.Data[i-1] = v
			}
		}
	}
}

// ExdScreenConfigurationMsg represents the exd_screen_configuration FIT message type.
type ExdScreenConfigurationMsg struct {
	ScreenIndex   uint8
	FieldCount    uint8 // number of fields in screen
	Layout        ExdLayout
	ScreenEnabled Bool
}

// ExdDataFieldConfigurationMsg represents the exd_data_field_configuration FIT message type.
type ExdDataFieldConfigurationMsg struct {
	ScreenIndex  uint8
	ConceptField byte
	FieldId      uint8
	ConceptCount uint8
	DisplayType  ExdDisplayType
	Title        []string
}

func (x *ExdDataFieldConfigurationMsg) expandComponents() {
	if x.ConceptField != 0xFF {
		x.FieldId = uint8(
			(x.ConceptField >> 0) & ((1 << 4) - 1),
		)
		x.ConceptCount = uint8(
			(x.ConceptField >> 4) & ((1 << 4) - 1),
		)
	}
}

// ExdDataConceptConfigurationMsg represents the exd_data_concept_configuration FIT message type.
type ExdDataConceptConfigurationMsg struct {
	ScreenIndex  uint8
	ConceptField byte
	FieldId      uint8
	ConceptIndex uint8
	DataPage     uint8
	ConceptKey   uint8
	Scaling      uint8
	DataUnits    ExdDataUnits
	Qualifier    ExdQualifiers
	Descriptor   ExdDescriptors
	IsSigned     Bool
}

func (x *ExdDataConceptConfigurationMsg) expandComponents() {
	if x.ConceptField != 0xFF {
		x.FieldId = uint8(
			(x.ConceptField >> 0) & ((1 << 4) - 1),
		)
		x.ConceptIndex = uint8(
			(x.ConceptField >> 4) & ((1 << 4) - 1),
		)
	}
}

// FieldDescriptionMsg represents the field_description FIT message type.
type FieldDescriptionMsg struct {
	DeveloperDataIndex    uint8
	FieldDefinitionNumber uint8
	FitBaseTypeId         FitBaseType
	FieldName             []string
	Array                 uint8
	Components            string
	Scale                 uint8
	Offset                int8
	Units                 []string
	Bits                  string
	Accumulate            string
	FitBaseUnitId         FitBaseUnit
	NativeMesgNum         MesgNum
	NativeFieldNum        uint8
}

// DeveloperDataIdMsg represents the developer_data_id FIT message type.
type DeveloperDataIdMsg struct {
	DeveloperId        []byte
	ApplicationId      []byte
	ManufacturerId     Manufacturer
	DeveloperDataIndex uint8
	ApplicationVersion uint32
}
