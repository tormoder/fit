// This file is auto-generated using the
// program found in 'cmd/fitgen/main.go'
// DO NOT EDIT.
// SDK Version: 16.10
// Generation time: Tue Oct  6 17:37:03 UTC 2015

package fit

import "time"

// FileId represents the file_id FIT message type.
type FileIdMsg struct {
	Type         File
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

// FileCreator represents the file_creator FIT message type.
type FileCreatorMsg struct {
	SoftwareVersion uint16
	HardwareVersion uint8
}

// TimestampCorrelation represents the timestamp_correlation FIT message type.
type TimestampCorrelationMsg struct {
}

// Software represents the software FIT message type.
type SoftwareMsg struct {
	MessageIndex MessageIndex
	Version      float64
	PartNumber   string
}

// SlaveDevice represents the slave_device FIT message type.
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

// Capabilities represents the capabilities FIT message type.
type CapabilitiesMsg struct {
	Languages             []uint8      // Use language_bits_x types where x is index of array.
	Sports                []SportBits0 // Use sport_bits_x types where x is index of array.
	WorkoutsSupported     WorkoutCapabilities
	ConnectivitySupported ConnectivityCapabilities
}

// FileCapabilities represents the file_capabilities FIT message type.
type FileCapabilitiesMsg struct {
	MessageIndex MessageIndex
	Type         File
	Flags        FileFlags
	Directory    string
	MaxCount     uint16
	MaxSize      uint32
}

// MesgCapabilities represents the mesg_capabilities FIT message type.
type MesgCapabilitiesMsg struct {
	MessageIndex MessageIndex
	File         File
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

// FieldCapabilities represents the field_capabilities FIT message type.
type FieldCapabilitiesMsg struct {
	MessageIndex MessageIndex
	File         File
	MesgNum      MesgNum
	FieldNum     uint8
	Count        uint16
}

// DeviceSettings represents the device_settings FIT message type.
type DeviceSettingsMsg struct {
	ActiveTimeZone uint8     // Index into time zone arrays.
	UtcOffset      uint32    // Offset from system time. Required to convert timestamp from system time to UTC.
	TimeZoneOffset []float64 // timezone offset in 1/4 hour increments
}

// UserProfile represents the user_profile FIT message type.
type UserProfileMsg struct {
	MessageIndex               MessageIndex
	FriendlyName               string
	Gender                     Gender
	Age                        uint8
	Height                     float64
	Weight                     float64
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
}

// HrmProfile represents the hrm_profile FIT message type.
type HrmProfileMsg struct {
	MessageIndex      MessageIndex
	Enabled           Bool
	HrmAntId          uint16
	LogHrv            Bool
	HrmAntIdTransType uint8
}

// SdmProfile represents the sdm_profile FIT message type.
type SdmProfileMsg struct {
	MessageIndex      MessageIndex
	Enabled           Bool
	SdmAntId          uint16
	SdmCalFactor      float64
	Odometer          float64
	SpeedSource       Bool // Use footpod for speed source instead of GPS
	SdmAntIdTransType uint8
	OdometerRollover  uint8 // Rollover counter that can be used to extend the odometer
}

// BikeProfile represents the bike_profile FIT message type.
type BikeProfileMsg struct {
	MessageIndex             MessageIndex
	Name                     string
	Sport                    Sport
	SubSport                 SubSport
	Odometer                 float64
	BikeSpdAntId             uint16
	BikeCadAntId             uint16
	BikeSpdcadAntId          uint16
	BikePowerAntId           uint16
	CustomWheelsize          float64
	AutoWheelsize            float64
	BikeWeight               float64
	PowerCalFactor           float64
	AutoWheelCal             Bool
	AutoPowerZero            Bool
	Id                       uint8
	SpdEnabled               Bool
	CadEnabled               Bool
	SpdcadEnabled            Bool
	PowerEnabled             Bool
	CrankLength              float64
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

// ZonesTarget represents the zones_target FIT message type.
type ZonesTargetMsg struct {
	MaxHeartRate             uint8
	ThresholdHeartRate       uint8
	FunctionalThresholdPower uint16
	HrCalcType               HrZoneCalc
	PwrCalcType              PwrZoneCalc
}

// Sport represents the sport FIT message type.
type SportMsg struct {
	Sport    Sport
	SubSport SubSport
	Name     string
}

// HrZone represents the hr_zone FIT message type.
type HrZoneMsg struct {
	MessageIndex MessageIndex
	HighBpm      uint8
	Name         string
}

// SpeedZone represents the speed_zone FIT message type.
type SpeedZoneMsg struct {
	MessageIndex MessageIndex
	HighValue    float64
	Name         string
}

// CadenceZone represents the cadence_zone FIT message type.
type CadenceZoneMsg struct {
	MessageIndex MessageIndex
	HighValue    uint8
	Name         string
}

// PowerZone represents the power_zone FIT message type.
type PowerZoneMsg struct {
	MessageIndex MessageIndex
	HighValue    uint16
	Name         string
}

// MetZone represents the met_zone FIT message type.
type MetZoneMsg struct {
	MessageIndex MessageIndex
	HighBpm      uint8
	Calories     float64
	FatCalories  float64
}

// Goal represents the goal FIT message type.
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
}

// Activity represents the activity FIT message type.
type ActivityMsg struct {
	Timestamp      time.Time
	TotalTimerTime float64 // Exclude pauses
	NumSessions    uint16
	Type           ActivityMode
	Event          Event
	EventType      EventType
	LocalTimestamp time.Time // timestamp epoch expressed in local time, used to convert activity timestamps to local time
	EventGroup     uint8
}

// Session represents the session FIT message type.
type SessionMsg struct {
	MessageIndex           MessageIndex // Selected bit is set for the current session.
	Timestamp              time.Time    // Sesson end time.
	Event                  Event        // session
	EventType              EventType    // stop
	StartTime              time.Time
	StartPositionLat       Latitude
	StartPositionLong      Longitude
	Sport                  Sport
	SubSport               SubSport
	TotalElapsedTime       float64 // Time (includes pauses)
	TotalTimerTime         float64 // Timer Time (excludes pauses)
	TotalDistance          float64
	TotalCycles            uint32
	TotalCalories          uint16
	TotalFatCalories       uint16
	AvgSpeed               uint16 // total_distance / total_timer_time
	MaxSpeed               uint16
	AvgHeartRate           uint8 // average heart rate (excludes pause time)
	MaxHeartRate           uint8
	AvgCadence             uint8 // total_cycles / total_timer_time if non_zero_avg_cadence otherwise total_cycles / total_elapsed_time
	MaxCadence             uint8
	AvgPower               uint16 // total_power / total_timer_time if non_zero_avg_power otherwise total_power / total_elapsed_time
	MaxPower               uint16
	TotalAscent            uint16
	TotalDescent           uint16
	TotalTrainingEffect    float64
	FirstLapIndex          uint16
	NumLaps                uint16
	EventGroup             uint8
	Trigger                SessionTrigger
	NecLat                 Latitude
	NecLong                Longitude
	SwcLat                 Latitude
	SwcLong                Longitude
	NormalizedPower        uint16
	TrainingStressScore    float64
	IntensityFactor        float64
	LeftRightBalance       LeftRightBalance100
	AvgStrokeCount         float64
	AvgStrokeDistance      float64
	SwimStroke             SwimStroke
	PoolLength             float64
	ThresholdPower         uint16
	PoolLengthUnit         DisplayMeasure
	NumActiveLengths       uint16 // # of active lengths of swim pool
	TotalWork              uint32
	AvgAltitude            uint16
	MaxAltitude            uint16
	GpsAccuracy            uint8
	AvgGrade               float64
	AvgPosGrade            float64
	AvgNegGrade            float64
	MaxPosGrade            float64
	MaxNegGrade            float64
	AvgTemperature         int8
	MaxTemperature         int8
	TotalMovingTime        float64
	AvgPosVerticalSpeed    float64
	AvgNegVerticalSpeed    float64
	MaxPosVerticalSpeed    float64
	MaxNegVerticalSpeed    float64
	MinHeartRate           uint8
	TimeInHrZone           []float64
	TimeInSpeedZone        []float64
	TimeInCadenceZone      []float64
	TimeInPowerZone        []float64
	AvgLapTime             float64
	BestLapIndex           uint16
	MinAltitude            uint16
	PlayerScore            uint16
	OpponentScore          uint16
	OpponentName           string
	StrokeCount            []uint16 // stroke_type enum used as the index
	ZoneCount              []uint16 // zone number used as the index
	MaxBallSpeed           float64
	AvgBallSpeed           float64
	AvgVerticalOscillation float64
	AvgStanceTimePercent   float64
	AvgStanceTime          float64
	AvgFractionalCadence   float64 // fractional part of the avg_cadence
	MaxFractionalCadence   float64 // fractional part of the max_cadence
	TotalFractionalCycles  float64 // fractional part of the total_cycles
	SportIndex             uint8
	EnhancedAvgSpeed       float64 // total_distance / total_timer_time
	EnhancedMaxSpeed       float64
	EnhancedAvgAltitude    float64
	EnhancedMinAltitude    float64
	EnhancedMaxAltitude    float64
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

// Lap represents the lap FIT message type.
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
	TotalElapsedTime              float64 // Time (includes pauses)
	TotalTimerTime                float64 // Timer Time (excludes pauses)
	TotalDistance                 float64
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
	AvgStrokeDistance             float64
	SwimStroke                    SwimStroke
	SubSport                      SubSport
	NumActiveLengths              uint16 // # of active lengths of swim pool
	TotalWork                     uint32
	AvgAltitude                   uint16
	MaxAltitude                   uint16
	GpsAccuracy                   uint8
	AvgGrade                      float64
	AvgPosGrade                   float64
	AvgNegGrade                   float64
	MaxPosGrade                   float64
	MaxNegGrade                   float64
	AvgTemperature                int8
	MaxTemperature                int8
	TotalMovingTime               float64
	AvgPosVerticalSpeed           float64
	AvgNegVerticalSpeed           float64
	MaxPosVerticalSpeed           float64
	MaxNegVerticalSpeed           float64
	TimeInHrZone                  []float64
	TimeInSpeedZone               []float64
	TimeInCadenceZone             []float64
	TimeInPowerZone               []float64
	RepetitionNum                 uint16
	MinAltitude                   uint16
	MinHeartRate                  uint8
	WktStepIndex                  MessageIndex
	OpponentScore                 uint16
	StrokeCount                   []uint16 // stroke_type enum used as the index
	ZoneCount                     []uint16 // zone number used as the index
	AvgVerticalOscillation        float64
	AvgStanceTimePercent          float64
	AvgStanceTime                 float64
	AvgFractionalCadence          float64 // fractional part of the avg_cadence
	MaxFractionalCadence          float64 // fractional part of the max_cadence
	TotalFractionalCycles         float64 // fractional part of the total_cycles
	PlayerScore                   uint16
	AvgTotalHemoglobinConc        []float64 // Avg saturated and unsaturated hemoglobin
	MinTotalHemoglobinConc        []float64 // Min saturated and unsaturated hemoglobin
	MaxTotalHemoglobinConc        []float64 // Max saturated and unsaturated hemoglobin
	AvgSaturatedHemoglobinPercent []float64 // Avg percentage of hemoglobin saturated with oxygen
	MinSaturatedHemoglobinPercent []float64 // Min percentage of hemoglobin saturated with oxygen
	MaxSaturatedHemoglobinPercent []float64 // Max percentage of hemoglobin saturated with oxygen
	EnhancedAvgSpeed              float64
	EnhancedMaxSpeed              float64
	EnhancedAvgAltitude           float64
	EnhancedMinAltitude           float64
	EnhancedMaxAltitude           float64
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

// Length represents the length FIT message type.
type LengthMsg struct {
	MessageIndex       MessageIndex
	Timestamp          time.Time
	Event              Event
	EventType          EventType
	StartTime          time.Time
	TotalElapsedTime   float64
	TotalTimerTime     float64
	TotalStrokes       uint16
	AvgSpeed           float64
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

// Record represents the record FIT message type.
type RecordMsg struct {
	Timestamp                     time.Time
	PositionLat                   Latitude
	PositionLong                  Longitude
	Altitude                      uint16
	HeartRate                     uint8
	Cadence                       uint8
	Distance                      float64
	Speed                         uint16
	Power                         uint16
	CompressedSpeedDistance       []byte
	Grade                         float64
	Resistance                    uint8 // Relative. 0 is none  254 is Max.
	TimeFromCourse                float64
	CycleLength                   float64
	Temperature                   int8
	Speed1s                       []float64 // Speed at 1s intervals.  Timestamp field indicates time of last array element.
	Cycles                        uint8
	TotalCycles                   uint32
	CompressedAccumulatedPower    uint16
	AccumulatedPower              uint32
	LeftRightBalance              LeftRightBalance
	GpsAccuracy                   uint8
	VerticalSpeed                 float64
	Calories                      uint16
	VerticalOscillation           float64
	StanceTimePercent             float64
	StanceTime                    float64
	ActivityType                  ActivityType
	LeftTorqueEffectiveness       float64
	RightTorqueEffectiveness      float64
	LeftPedalSmoothness           float64
	RightPedalSmoothness          float64
	CombinedPedalSmoothness       float64
	Time128                       float64
	StrokeType                    StrokeType
	Zone                          uint8
	BallSpeed                     float64
	Cadence256                    float64 // Log cadence and fractional cadence for backwards compatability
	FractionalCadence             float64
	TotalHemoglobinConc           float64 // Total saturated and unsaturated hemoglobin
	TotalHemoglobinConcMin        float64 // Min saturated and unsaturated hemoglobin
	TotalHemoglobinConcMax        float64 // Max saturated and unsaturated hemoglobin
	SaturatedHemoglobinPercent    float64 // Percentage of hemoglobin saturated with oxygen
	SaturatedHemoglobinPercentMin float64 // Min percentage of hemoglobin saturated with oxygen
	SaturatedHemoglobinPercentMax float64 // Max percentage of hemoglobin saturated with oxygen
	DeviceIndex                   DeviceIndex
	EnhancedSpeed                 float64
	EnhancedAltitude              float64
}

// Event represents the event FIT message type.
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
		return float64(float32(x.Data)/1000.0 - float32(0))
	case EventVirtualPartnerPace:
		return float64(float32(x.Data)/1000.0 - float32(0))
	case EventHrHighAlert:
		return uint8(x.Data)
	case EventHrLowAlert:
		return uint8(x.Data)
	case EventSpeedHighAlert:
		return float64(float32(x.Data)/1000.0 - float32(0))
	case EventSpeedLowAlert:
		return float64(float32(x.Data)/1000.0 - float32(0))
	case EventCadHighAlert:
		return uint16(x.Data)
	case EventCadLowAlert:
		return uint16(x.Data)
	case EventPowerHighAlert:
		return uint16(x.Data)
	case EventPowerLowAlert:
		return uint16(x.Data)
	case EventTimeDurationAlert:
		return float64(float32(x.Data)/1000.0 - float32(0))
	case EventDistanceDurationAlert:
		return float64(float32(x.Data)/100.0 - float32(0))
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

// DeviceInfo represents the device_info FIT message type.
type DeviceInfoMsg struct {
	Timestamp           time.Time
	DeviceIndex         DeviceIndex
	DeviceType          uint8
	Manufacturer        Manufacturer
	SerialNumber        uint32
	Product             uint16
	SoftwareVersion     float64
	HardwareVersion     uint8
	CumOperatingTime    uint32 // Reset by new battery or charge.
	BatteryVoltage      float64
	BatteryStatus       BatteryStatus
	SensorPosition      BodyLocation // Indicates the location of the sensor
	Descriptor          string       // Used to describe the sensor or location
	AntTransmissionType uint8
	AntDeviceNumber     uint16
	AntNetwork          AntNetwork
	SourceType          SourceType
	ProductName         string // Optional free form string to indicate the devices name or model
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

// TrainingFile represents the training_file FIT message type.
type TrainingFileMsg struct {
	Timestamp    time.Time
	Type         File
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

// Hrv represents the hrv FIT message type.
type HrvMsg struct {
	Time []float64 // Time between beats
}

// CameraEvent represents the camera_event FIT message type.
type CameraEventMsg struct {
}

// GyroscopeData represents the gyroscope_data FIT message type.
type GyroscopeDataMsg struct {
}

// AccelerometerData represents the accelerometer_data FIT message type.
type AccelerometerDataMsg struct {
}

// ThreeDSensorCalibration represents the three_d_sensor_calibration FIT message type.
type ThreeDSensorCalibrationMsg struct {
}

// VideoFrame represents the video_frame FIT message type.
type VideoFrameMsg struct {
}

// ObdiiData represents the obdii_data FIT message type.
type ObdiiDataMsg struct {
}

// NmeaSentence represents the nmea_sentence FIT message type.
type NmeaSentenceMsg struct {
	Timestamp   time.Time // Timestamp message was output
	TimestampMs uint16    // Fractional part of timestamp, added to timestamp
	Sentence    string    // NMEA sentence
}

// AviationAttitude represents the aviation_attitude FIT message type.
type AviationAttitudeMsg struct {
	Timestamp             time.Time // Timestamp message was output
	TimestampMs           uint16    // Fractional part of timestamp, added to timestamp
	SystemTime            []uint32  // System time associated with sample expressed in ms.
	Pitch                 []float64 // Range -PI/2 to +PI/2
	Roll                  []float64 // Range -PI to +PI
	AccelLateral          []float64 // Range -78.4 to +78.4 (-8 Gs to 8 Gs)
	AccelNormal           []float64 // Range -78.4 to +78.4 (-8 Gs to 8 Gs)
	TurnRate              []float64 // Range -8.727 to +8.727 (-500 degs/sec to +500 degs/sec)
	Stage                 []AttitudeStage
	AttitudeStageComplete []uint8   // The percent complete of the current attitude stage.  Set to 0 for attitude stages 0, 1 and 2 and to 100 for attitude stage 3 by AHRS modules that do not support it.  Range - 100
	Track                 []float64 // Track Angle/Heading Range 0 - 2pi
	Validity              []AttitudeValidity
}

// Video represents the video FIT message type.
type VideoMsg struct {
}

// VideoTitle represents the video_title FIT message type.
type VideoTitleMsg struct {
	MessageIndex MessageIndex // Long titles will be split into multiple parts
	MessageCount uint16       // Total number of title parts
	Text         string
}

// VideoDescription represents the video_description FIT message type.
type VideoDescriptionMsg struct {
	MessageIndex MessageIndex // Long descriptions will be split into multiple parts
	MessageCount uint16       // Total number of description parts
	Text         string
}

// VideoClip represents the video_clip FIT message type.
type VideoClipMsg struct {
}

// Course represents the course FIT message type.
type CourseMsg struct {
	Sport        Sport
	Name         string
	Capabilities CourseCapabilities
}

// CoursePoint represents the course_point FIT message type.
type CoursePointMsg struct {
	MessageIndex MessageIndex
	Timestamp    time.Time
	PositionLat  Latitude
	PositionLong Longitude
	Distance     float64
	Type         CoursePoint
	Name         string
	Favorite     Bool
}

// SegmentId represents the segment_id FIT message type.
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

// SegmentLeaderboardEntry represents the segment_leaderboard_entry FIT message type.
type SegmentLeaderboardEntryMsg struct {
	MessageIndex    MessageIndex
	Name            string                 // Friendly name assigned to leader
	Type            SegmentLeaderboardType // Leader classification
	GroupPrimaryKey uint32                 // Primary user ID of this leader
	ActivityId      uint32                 // ID of the activity associated with this leader time
	SegmentTime     float64                // Segment Time (includes pauses)
}

// SegmentPoint represents the segment_point FIT message type.
type SegmentPointMsg struct {
	MessageIndex MessageIndex
	PositionLat  Latitude
	PositionLong Longitude
	Distance     float64   // Accumulated distance along the segment at the described point
	Altitude     float64   // Accumulated altitude along the segment at the described point
	LeaderTime   []float64 // Accumualted time each leader board member required to reach the described point. This value is zero for all leader board members at the starting point of the segment.
}

// SegmentLap represents the segment_lap FIT message type.
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
	TotalElapsedTime            float64 // Time (includes pauses)
	TotalTimerTime              float64 // Timer Time (excludes pauses)
	TotalDistance               float64
	TotalCycles                 uint32
	TotalCalories               uint16
	TotalFatCalories            uint16 // If New Leaf
	AvgSpeed                    float64
	MaxSpeed                    float64
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
	AvgAltitude                 float64
	MaxAltitude                 float64
	GpsAccuracy                 uint8
	AvgGrade                    float64
	AvgPosGrade                 float64
	AvgNegGrade                 float64
	MaxPosGrade                 float64
	MaxNegGrade                 float64
	AvgTemperature              int8
	MaxTemperature              int8
	TotalMovingTime             float64
	AvgPosVerticalSpeed         float64
	AvgNegVerticalSpeed         float64
	MaxPosVerticalSpeed         float64
	MaxNegVerticalSpeed         float64
	TimeInHrZone                []float64
	TimeInSpeedZone             []float64
	TimeInCadenceZone           []float64
	TimeInPowerZone             []float64
	RepetitionNum               uint16
	MinAltitude                 float64
	MinHeartRate                uint8
	ActiveTime                  float64
	WktStepIndex                MessageIndex
	SportEvent                  SportEvent
	AvgLeftTorqueEffectiveness  float64
	AvgRightTorqueEffectiveness float64
	AvgLeftPedalSmoothness      float64
	AvgRightPedalSmoothness     float64
	AvgCombinedPedalSmoothness  float64
	Status                      SegmentLapStatus
	Uuid                        string
	AvgFractionalCadence        float64 // fractional part of the avg_cadence
	MaxFractionalCadence        float64 // fractional part of the max_cadence
	TotalFractionalCycles       float64 // fractional part of the total_cycles
	FrontGearShiftCount         uint16
	RearGearShiftCount          uint16
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

// SegmentFile represents the segment_file FIT message type.
type SegmentFileMsg struct {
	MessageIndex          MessageIndex
	FileUuid              string                   // UUID of the segment file
	Enabled               Bool                     // Enabled state of the segment file
	UserProfilePrimaryKey uint32                   // Primary key of the user that created the segment file
	LeaderType            []SegmentLeaderboardType // Leader type of each leader in the segment file
	LeaderGroupPrimaryKey []uint32                 // Group primary key of each leader in the segment file
	LeaderActivityId      []uint32                 // Activity ID of each leader in the segment file
}

// Workout represents the workout FIT message type.
type WorkoutMsg struct {
	Sport         Sport
	Capabilities  WorkoutCapabilities
	NumValidSteps uint16 // number of valid steps
	WktName       string
}

// WorkoutStep represents the workout_step FIT message type.
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
}

// GetDurationValue returns the appropriate DurationValue
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *WorkoutStepMsg) GetDurationValue() interface{} {
	switch x.DurationType {
	case WktStepDurationTime, WktStepDurationRepetitionTime:
		return float64(float32(x.DurationValue)/1000.0 - float32(0))
	case WktStepDurationDistance:
		return float64(float32(x.DurationValue)/100.0 - float32(0))
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
	case x.TargetType == WktStepTargetHeartRate:
		return uint32(x.TargetValue)
	case x.TargetType == WktStepTargetPower:
		return uint32(x.TargetValue)
	case x.DurationType == WktStepDurationRepeatUntilStepsCmplt:
		return uint32(x.TargetValue)
	case x.DurationType == WktStepDurationRepeatUntilTime:
		return float64(float32(x.TargetValue)/1000.0 - float32(0))
	case x.DurationType == WktStepDurationRepeatUntilDistance:
		return float64(float32(x.TargetValue)/100.0 - float32(0))
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
		return float64(float32(x.CustomTargetValueLow)/1000.0 - float32(0))
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
		return float64(float32(x.CustomTargetValueHigh)/1000.0 - float32(0))
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

// Schedule represents the schedule FIT message type.
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

// Totals represents the totals FIT message type.
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

// WeightScale represents the weight_scale FIT message type.
type WeightScaleMsg struct {
	Timestamp         time.Time
	Weight            float64
	PercentFat        float64
	PercentHydration  float64
	VisceralFatMass   float64
	BoneMass          float64
	MuscleMass        float64
	BasalMet          float64
	PhysiqueRating    uint8
	ActiveMet         float64 // ~4kJ per kcal, 0.25 allows max 16384 kcal
	MetabolicAge      uint8
	VisceralFatRating uint8
	UserProfileIndex  MessageIndex // Associates this weight scale message to a user.  This corresponds to the index of the user profile message in the weight scale file.
}

// BloodPressure represents the blood_pressure FIT message type.
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

// MonitoringInfo represents the monitoring_info FIT message type.
type MonitoringInfoMsg struct {
	Timestamp      time.Time
	LocalTimestamp time.Time // Use to convert activity timestamps to local time if device does not support time zone and daylight savings time correction.
}

// Monitoring represents the monitoring FIT message type.
type MonitoringMsg struct {
	Timestamp       time.Time   // Must align to logging interval, for example, time must be 00:00:00 for daily log.
	DeviceIndex     DeviceIndex // Associates this data to device_info message.  Not required for file with single device (sensor).
	Calories        uint16      // Accumulated total calories.  Maintained by MonitoringReader for each activity_type.  See SDK documentation
	Distance        float64     // Accumulated distance.  Maintained by MonitoringReader for each activity_type.  See SDK documentation.
	Cycles          float64     // Accumulated cycles.  Maintained by MonitoringReader for each activity_type.  See SDK documentation.
	ActiveTime      float64
	ActivityType    ActivityType
	ActivitySubtype ActivitySubtype
	Distance16      uint16
	Cycles16        uint16
	ActiveTime16    uint16
	LocalTimestamp  time.Time // Must align to logging interval, for example, time must be 00:00:00 for daily log.
}

// GetCycles returns the appropriate Cycles
// subfield if a matching reference field/value combination is found.
// If none of the reference field/value combinations are true
// then the main field is returned.
func (x *MonitoringMsg) GetCycles() interface{} {
	switch x.ActivityType {
	case ActivityTypeCycling, ActivityTypeSwimming:
		return float64(float32(x.Cycles)/2.0 - float32(0))
	default:
		return x.Cycles
	}
}

// MemoGlob represents the memo_glob FIT message type.
type MemoGlobMsg struct {
}
