// This file is auto-generated using the
// program found in 'cmd/fitgen/main.go'
// DO NOT EDIT.
// SDK Version: 16.20
// Generation time: Fri Oct 16 15:58:20 UTC 2015

package fit

// ActivityClass represents the activity_class FIT type.
type ActivityClass byte

const (
	ActivityClassLevel    ActivityClass = 0x7F // 0 to 100
	ActivityClassLevelMax ActivityClass = 100
	ActivityClassAthlete  ActivityClass = 0x80
	ActivityClassInvalid  ActivityClass = 0xFF
)

// ActivityLevel represents the activity_level FIT type.
type ActivityLevel byte

const (
	ActivityLevelLow     ActivityLevel = 0
	ActivityLevelMedium  ActivityLevel = 1
	ActivityLevelHigh    ActivityLevel = 2
	ActivityLevelInvalid ActivityLevel = 0xFF
)

// ActivityMode represents the activity FIT type.
type ActivityMode byte

const (
	ActivityModeManual         ActivityMode = 0
	ActivityModeAutoMultiSport ActivityMode = 1
	ActivityModeInvalid        ActivityMode = 0xFF
)

// ActivitySubtype represents the activity_subtype FIT type.
type ActivitySubtype byte

const (
	ActivitySubtypeGeneric       ActivitySubtype = 0
	ActivitySubtypeTreadmill     ActivitySubtype = 1  // Run
	ActivitySubtypeStreet        ActivitySubtype = 2  // Run
	ActivitySubtypeTrail         ActivitySubtype = 3  // Run
	ActivitySubtypeTrack         ActivitySubtype = 4  // Run
	ActivitySubtypeSpin          ActivitySubtype = 5  // Cycling
	ActivitySubtypeIndoorCycling ActivitySubtype = 6  // Cycling
	ActivitySubtypeRoad          ActivitySubtype = 7  // Cycling
	ActivitySubtypeMountain      ActivitySubtype = 8  // Cycling
	ActivitySubtypeDownhill      ActivitySubtype = 9  // Cycling
	ActivitySubtypeRecumbent     ActivitySubtype = 10 // Cycling
	ActivitySubtypeCyclocross    ActivitySubtype = 11 // Cycling
	ActivitySubtypeHandCycling   ActivitySubtype = 12 // Cycling
	ActivitySubtypeTrackCycling  ActivitySubtype = 13 // Cycling
	ActivitySubtypeIndoorRowing  ActivitySubtype = 14 // Fitness Equipment
	ActivitySubtypeElliptical    ActivitySubtype = 15 // Fitness Equipment
	ActivitySubtypeStairClimbing ActivitySubtype = 16 // Fitness Equipment
	ActivitySubtypeLapSwimming   ActivitySubtype = 17 // Swimming
	ActivitySubtypeOpenWater     ActivitySubtype = 18 // Swimming
	ActivitySubtypeAll           ActivitySubtype = 254
	ActivitySubtypeInvalid       ActivitySubtype = 0xFF
)

// ActivityType represents the activity_type FIT type.
type ActivityType byte

const (
	ActivityTypeGeneric          ActivityType = 0
	ActivityTypeRunning          ActivityType = 1
	ActivityTypeCycling          ActivityType = 2
	ActivityTypeTransition       ActivityType = 3 // Mulitsport transition
	ActivityTypeFitnessEquipment ActivityType = 4
	ActivityTypeSwimming         ActivityType = 5
	ActivityTypeWalking          ActivityType = 6
	ActivityTypeAll              ActivityType = 254 // All is for goals only to include all sports.
	ActivityTypeInvalid          ActivityType = 0xFF
)

// AntNetwork represents the ant_network FIT type.
type AntNetwork byte

const (
	AntNetworkPublic  AntNetwork = 0
	AntNetworkAntplus AntNetwork = 1
	AntNetworkAntfs   AntNetwork = 2
	AntNetworkPrivate AntNetwork = 3
	AntNetworkInvalid AntNetwork = 0xFF
)

// AntplusDeviceType represents the antplus_device_type FIT type.
type AntplusDeviceType uint8

const (
	AntplusDeviceTypeAntfs                   AntplusDeviceType = 1
	AntplusDeviceTypeBikePower               AntplusDeviceType = 11
	AntplusDeviceTypeEnvironmentSensorLegacy AntplusDeviceType = 12
	AntplusDeviceTypeMultiSportSpeedDistance AntplusDeviceType = 15
	AntplusDeviceTypeControl                 AntplusDeviceType = 16
	AntplusDeviceTypeFitnessEquipment        AntplusDeviceType = 17
	AntplusDeviceTypeBloodPressure           AntplusDeviceType = 18
	AntplusDeviceTypeGeocacheNode            AntplusDeviceType = 19
	AntplusDeviceTypeLightElectricVehicle    AntplusDeviceType = 20
	AntplusDeviceTypeEnvSensor               AntplusDeviceType = 25
	AntplusDeviceTypeRacquet                 AntplusDeviceType = 26
	AntplusDeviceTypeWeightScale             AntplusDeviceType = 119
	AntplusDeviceTypeHeartRate               AntplusDeviceType = 120
	AntplusDeviceTypeBikeSpeedCadence        AntplusDeviceType = 121
	AntplusDeviceTypeBikeCadence             AntplusDeviceType = 122
	AntplusDeviceTypeBikeSpeed               AntplusDeviceType = 123
	AntplusDeviceTypeStrideSpeedDistance     AntplusDeviceType = 124
	AntplusDeviceTypeInvalid                 AntplusDeviceType = 0xFF
)

// AttitudeStage represents the attitude_stage FIT type.
type AttitudeStage byte

const (
	AttitudeStageFailed   AttitudeStage = 0
	AttitudeStageAligning AttitudeStage = 1
	AttitudeStageDegraded AttitudeStage = 2
	AttitudeStageValid    AttitudeStage = 3
	AttitudeStageInvalid  AttitudeStage = 0xFF
)

// AttitudeValidity represents the attitude_validity FIT type.
type AttitudeValidity uint16

const (
	AttitudeValidityTrackAngleHeadingValid AttitudeValidity = 0x0001
	AttitudeValidityPitchValid             AttitudeValidity = 0x0002
	AttitudeValidityRollValid              AttitudeValidity = 0x0004
	AttitudeValidityLateralBodyAccelValid  AttitudeValidity = 0x0008
	AttitudeValidityNormalBodyAccelValid   AttitudeValidity = 0x0010
	AttitudeValidityTurnRateValid          AttitudeValidity = 0x0020
	AttitudeValidityHwFail                 AttitudeValidity = 0x0040
	AttitudeValidityMagInvalid             AttitudeValidity = 0x0080
	AttitudeValidityNoGps                  AttitudeValidity = 0x0100
	AttitudeValidityGpsInvalid             AttitudeValidity = 0x0200
	AttitudeValiditySolutionCoasting       AttitudeValidity = 0x0400
	AttitudeValidityTrueTrackAngle         AttitudeValidity = 0x0800
	AttitudeValidityMagneticHeading        AttitudeValidity = 0x1000
	AttitudeValidityInvalid                AttitudeValidity = 0xFFFF
)

// AutolapTrigger represents the autolap_trigger FIT type.
type AutolapTrigger byte

const (
	AutolapTriggerTime             AutolapTrigger = 0
	AutolapTriggerDistance         AutolapTrigger = 1
	AutolapTriggerPositionStart    AutolapTrigger = 2
	AutolapTriggerPositionLap      AutolapTrigger = 3
	AutolapTriggerPositionWaypoint AutolapTrigger = 4
	AutolapTriggerPositionMarked   AutolapTrigger = 5
	AutolapTriggerOff              AutolapTrigger = 6
	AutolapTriggerInvalid          AutolapTrigger = 0xFF
)

// BatteryStatus represents the battery_status FIT type.
type BatteryStatus uint8

const (
	BatteryStatusNew      BatteryStatus = 1
	BatteryStatusGood     BatteryStatus = 2
	BatteryStatusOk       BatteryStatus = 3
	BatteryStatusLow      BatteryStatus = 4
	BatteryStatusCritical BatteryStatus = 5
	BatteryStatusUnknown  BatteryStatus = 7
	BatteryStatusInvalid  BatteryStatus = 0xFF
)

// BodyLocation represents the body_location FIT type.
type BodyLocation byte

const (
	BodyLocationLeftLeg               BodyLocation = 0
	BodyLocationLeftCalf              BodyLocation = 1
	BodyLocationLeftShin              BodyLocation = 2
	BodyLocationLeftHamstring         BodyLocation = 3
	BodyLocationLeftQuad              BodyLocation = 4
	BodyLocationLeftGlute             BodyLocation = 5
	BodyLocationRightLeg              BodyLocation = 6
	BodyLocationRightCalf             BodyLocation = 7
	BodyLocationRightShin             BodyLocation = 8
	BodyLocationRightHamstring        BodyLocation = 9
	BodyLocationRightQuad             BodyLocation = 10
	BodyLocationRightGlute            BodyLocation = 11
	BodyLocationTorsoBack             BodyLocation = 12
	BodyLocationLeftLowerBack         BodyLocation = 13
	BodyLocationLeftUpperBack         BodyLocation = 14
	BodyLocationRightLowerBack        BodyLocation = 15
	BodyLocationRightUpperBack        BodyLocation = 16
	BodyLocationTorsoFront            BodyLocation = 17
	BodyLocationLeftAbdomen           BodyLocation = 18
	BodyLocationLeftChest             BodyLocation = 19
	BodyLocationRightAbdomen          BodyLocation = 20
	BodyLocationRightChest            BodyLocation = 21
	BodyLocationLeftArm               BodyLocation = 22
	BodyLocationLeftShoulder          BodyLocation = 23
	BodyLocationLeftBicep             BodyLocation = 24
	BodyLocationLeftTricep            BodyLocation = 25
	BodyLocationLeftBrachioradialis   BodyLocation = 26 // Left anterior forearm
	BodyLocationLeftForearmExtensors  BodyLocation = 27 // Left posterior forearm
	BodyLocationRightArm              BodyLocation = 28
	BodyLocationRightShoulder         BodyLocation = 29
	BodyLocationRightBicep            BodyLocation = 30
	BodyLocationRightTricep           BodyLocation = 31
	BodyLocationRightBrachioradialis  BodyLocation = 32 // Right anterior forearm
	BodyLocationRightForearmExtensors BodyLocation = 33 // Right posterior forearm
	BodyLocationNeck                  BodyLocation = 34
	BodyLocationThroat                BodyLocation = 35
	BodyLocationInvalid               BodyLocation = 0xFF
)

// BpStatus represents the bp_status FIT type.
type BpStatus byte

const (
	BpStatusNoError                 BpStatus = 0
	BpStatusErrorIncompleteData     BpStatus = 1
	BpStatusErrorNoMeasurement      BpStatus = 2
	BpStatusErrorDataOutOfRange     BpStatus = 3
	BpStatusErrorIrregularHeartRate BpStatus = 4
	BpStatusInvalid                 BpStatus = 0xFF
)

// CameraEventType represents the camera_event_type FIT type.
type CameraEventType byte

const (
	CameraEventTypeVideoStart                  CameraEventType = 0 // Start of video recording
	CameraEventTypeVideoSplit                  CameraEventType = 1 // Mark of video file split (end of one file, beginning of the other)
	CameraEventTypeVideoEnd                    CameraEventType = 2 // End of video recording
	CameraEventTypePhotoTaken                  CameraEventType = 3 // Still photo taken
	CameraEventTypeVideoSecondStreamStart      CameraEventType = 4
	CameraEventTypeVideoSecondStreamSplit      CameraEventType = 5
	CameraEventTypeVideoSecondStreamEnd        CameraEventType = 6
	CameraEventTypeVideoSplitStart             CameraEventType = 7 // Mark of video file split start
	CameraEventTypeVideoSecondStreamSplitStart CameraEventType = 8
	CameraEventTypeInvalid                     CameraEventType = 0xFF
)

// CameraOrientationType represents the camera_orientation_type FIT type.
type CameraOrientationType byte

const (
	CameraOrientationTypeCameraOrientation0   CameraOrientationType = 0
	CameraOrientationTypeCameraOrientation90  CameraOrientationType = 1
	CameraOrientationTypeCameraOrientation180 CameraOrientationType = 2
	CameraOrientationTypeCameraOrientation270 CameraOrientationType = 3
	CameraOrientationTypeInvalid              CameraOrientationType = 0xFF
)

// Checksum represents the checksum FIT type.
type Checksum uint8

const (
	ChecksumClear   Checksum = 0 // Allows clear of checksum for flash memory where can only write 1 to 0 without erasing sector.
	ChecksumOk      Checksum = 1 // Set to mark checksum as valid if computes to invalid values 0 or 0xFF.  Checksum can also be set to ok to save encoding computation time.
	ChecksumInvalid Checksum = 0xFF
)

// CommTimeoutType represents the comm_timeout_type FIT type.
type CommTimeoutType uint16

const (
	CommTimeoutTypeWildcardPairingTimeout CommTimeoutType = 0 // Timeout pairing to any device
	CommTimeoutTypePairingTimeout         CommTimeoutType = 1 // Timeout pairing to previously paired device
	CommTimeoutTypeConnectionLost         CommTimeoutType = 2 // Temporary loss of communications
	CommTimeoutTypeConnectionTimeout      CommTimeoutType = 3 // Connection closed due to extended bad communications
	CommTimeoutTypeInvalid                CommTimeoutType = 0xFFFF
)

// ConnectivityCapabilities represents the connectivity_capabilities FIT type.
type ConnectivityCapabilities uint32

const (
	ConnectivityCapabilitiesBluetooth                       ConnectivityCapabilities = 0x00000001
	ConnectivityCapabilitiesBluetoothLe                     ConnectivityCapabilities = 0x00000002
	ConnectivityCapabilitiesAnt                             ConnectivityCapabilities = 0x00000004
	ConnectivityCapabilitiesActivityUpload                  ConnectivityCapabilities = 0x00000008
	ConnectivityCapabilitiesCourseDownload                  ConnectivityCapabilities = 0x00000010
	ConnectivityCapabilitiesWorkoutDownload                 ConnectivityCapabilities = 0x00000020
	ConnectivityCapabilitiesLiveTrack                       ConnectivityCapabilities = 0x00000040
	ConnectivityCapabilitiesWeatherConditions               ConnectivityCapabilities = 0x00000080
	ConnectivityCapabilitiesWeatherAlerts                   ConnectivityCapabilities = 0x00000100
	ConnectivityCapabilitiesGpsEphemerisDownload            ConnectivityCapabilities = 0x00000200
	ConnectivityCapabilitiesExplicitArchive                 ConnectivityCapabilities = 0x00000400
	ConnectivityCapabilitiesSetupIncomplete                 ConnectivityCapabilities = 0x00000800
	ConnectivityCapabilitiesContinueSyncAfterSoftwareUpdate ConnectivityCapabilities = 0x00001000
	ConnectivityCapabilitiesConnectIqAppDownload            ConnectivityCapabilities = 0x00002000
	ConnectivityCapabilitiesInvalid                         ConnectivityCapabilities = 0x00000000
)

// CourseCapabilities represents the course_capabilities FIT type.
type CourseCapabilities uint32

const (
	CourseCapabilitiesProcessed  CourseCapabilities = 0x00000001
	CourseCapabilitiesValid      CourseCapabilities = 0x00000002
	CourseCapabilitiesTime       CourseCapabilities = 0x00000004
	CourseCapabilitiesDistance   CourseCapabilities = 0x00000008
	CourseCapabilitiesPosition   CourseCapabilities = 0x00000010
	CourseCapabilitiesHeartRate  CourseCapabilities = 0x00000020
	CourseCapabilitiesPower      CourseCapabilities = 0x00000040
	CourseCapabilitiesCadence    CourseCapabilities = 0x00000080
	CourseCapabilitiesTraining   CourseCapabilities = 0x00000100
	CourseCapabilitiesNavigation CourseCapabilities = 0x00000200
	CourseCapabilitiesBikeway    CourseCapabilities = 0x00000400
	CourseCapabilitiesInvalid    CourseCapabilities = 0x00000000
)

// CoursePoint represents the course_point FIT type.
type CoursePoint byte

const (
	CoursePointGeneric        CoursePoint = 0
	CoursePointSummit         CoursePoint = 1
	CoursePointValley         CoursePoint = 2
	CoursePointWater          CoursePoint = 3
	CoursePointFood           CoursePoint = 4
	CoursePointDanger         CoursePoint = 5
	CoursePointLeft           CoursePoint = 6
	CoursePointRight          CoursePoint = 7
	CoursePointStraight       CoursePoint = 8
	CoursePointFirstAid       CoursePoint = 9
	CoursePointFourthCategory CoursePoint = 10
	CoursePointThirdCategory  CoursePoint = 11
	CoursePointSecondCategory CoursePoint = 12
	CoursePointFirstCategory  CoursePoint = 13
	CoursePointHorsCategory   CoursePoint = 14
	CoursePointSprint         CoursePoint = 15
	CoursePointLeftFork       CoursePoint = 16
	CoursePointRightFork      CoursePoint = 17
	CoursePointMiddleFork     CoursePoint = 18
	CoursePointSlightLeft     CoursePoint = 19
	CoursePointSharpLeft      CoursePoint = 20
	CoursePointSlightRight    CoursePoint = 21
	CoursePointSharpRight     CoursePoint = 22
	CoursePointUTurn          CoursePoint = 23
	CoursePointInvalid        CoursePoint = 0xFF
)

// DeviceIndex represents the device_index FIT type.
type DeviceIndex uint8

const (
	DeviceIndexCreator DeviceIndex = 0 // Creator of the file is always device index 0.
	DeviceIndexInvalid DeviceIndex = 0xFF
)

// DisplayHeart represents the display_heart FIT type.
type DisplayHeart byte

const (
	DisplayHeartBpm     DisplayHeart = 0
	DisplayHeartMax     DisplayHeart = 1
	DisplayHeartReserve DisplayHeart = 2
	DisplayHeartInvalid DisplayHeart = 0xFF
)

// DisplayMeasure represents the display_measure FIT type.
type DisplayMeasure byte

const (
	DisplayMeasureMetric  DisplayMeasure = 0
	DisplayMeasureStatute DisplayMeasure = 1
	DisplayMeasureInvalid DisplayMeasure = 0xFF
)

// DisplayPosition represents the display_position FIT type.
type DisplayPosition byte

const (
	DisplayPositionDegree               DisplayPosition = 0  // dd.dddddd
	DisplayPositionDegreeMinute         DisplayPosition = 1  // dddmm.mmm
	DisplayPositionDegreeMinuteSecond   DisplayPosition = 2  // dddmmss
	DisplayPositionAustrianGrid         DisplayPosition = 3  // Austrian Grid (BMN)
	DisplayPositionBritishGrid          DisplayPosition = 4  // British National Grid
	DisplayPositionDutchGrid            DisplayPosition = 5  // Dutch grid system
	DisplayPositionHungarianGrid        DisplayPosition = 6  // Hungarian grid system
	DisplayPositionFinnishGrid          DisplayPosition = 7  // Finnish grid system Zone3 KKJ27
	DisplayPositionGermanGrid           DisplayPosition = 8  // Gausss Krueger (German)
	DisplayPositionIcelandicGrid        DisplayPosition = 9  // Icelandic Grid
	DisplayPositionIndonesianEquatorial DisplayPosition = 10 // Indonesian Equatorial LCO
	DisplayPositionIndonesianIrian      DisplayPosition = 11 // Indonesian Irian LCO
	DisplayPositionIndonesianSouthern   DisplayPosition = 12 // Indonesian Southern LCO
	DisplayPositionIndiaZone0           DisplayPosition = 13 // India zone 0
	DisplayPositionIndiaZoneIA          DisplayPosition = 14 // India zone IA
	DisplayPositionIndiaZoneIB          DisplayPosition = 15 // India zone IB
	DisplayPositionIndiaZoneIIA         DisplayPosition = 16 // India zone IIA
	DisplayPositionIndiaZoneIIB         DisplayPosition = 17 // India zone IIB
	DisplayPositionIndiaZoneIIIA        DisplayPosition = 18 // India zone IIIA
	DisplayPositionIndiaZoneIIIB        DisplayPosition = 19 // India zone IIIB
	DisplayPositionIndiaZoneIVA         DisplayPosition = 20 // India zone IVA
	DisplayPositionIndiaZoneIVB         DisplayPosition = 21 // India zone IVB
	DisplayPositionIrishTransverse      DisplayPosition = 22 // Irish Transverse Mercator
	DisplayPositionIrishGrid            DisplayPosition = 23 // Irish Grid
	DisplayPositionLoran                DisplayPosition = 24 // Loran TD
	DisplayPositionMaidenheadGrid       DisplayPosition = 25 // Maidenhead grid system
	DisplayPositionMgrsGrid             DisplayPosition = 26 // MGRS grid system
	DisplayPositionNewZealandGrid       DisplayPosition = 27 // New Zealand grid system
	DisplayPositionNewZealandTransverse DisplayPosition = 28 // New Zealand Transverse Mercator
	DisplayPositionQatarGrid            DisplayPosition = 29 // Qatar National Grid
	DisplayPositionModifiedSwedishGrid  DisplayPosition = 30 // Modified RT-90 (Sweden)
	DisplayPositionSwedishGrid          DisplayPosition = 31 // RT-90 (Sweden)
	DisplayPositionSouthAfricanGrid     DisplayPosition = 32 // South African Grid
	DisplayPositionSwissGrid            DisplayPosition = 33 // Swiss CH-1903 grid
	DisplayPositionTaiwanGrid           DisplayPosition = 34 // Taiwan Grid
	DisplayPositionUnitedStatesGrid     DisplayPosition = 35 // United States National Grid
	DisplayPositionUtmUpsGrid           DisplayPosition = 36 // UTM/UPS grid system
	DisplayPositionWestMalayan          DisplayPosition = 37 // West Malayan RSO
	DisplayPositionBorneoRso            DisplayPosition = 38 // Borneo RSO
	DisplayPositionEstonianGrid         DisplayPosition = 39 // Estonian grid system
	DisplayPositionLatvianGrid          DisplayPosition = 40 // Latvian Transverse Mercator
	DisplayPositionSwedishRef99Grid     DisplayPosition = 41 // Reference Grid 99 TM (Swedish)
	DisplayPositionInvalid              DisplayPosition = 0xFF
)

// DisplayPower represents the display_power FIT type.
type DisplayPower byte

const (
	DisplayPowerWatts      DisplayPower = 0
	DisplayPowerPercentFtp DisplayPower = 1
	DisplayPowerInvalid    DisplayPower = 0xFF
)

// Event represents the event FIT type.
type Event byte

const (
	EventTimer                 Event = 0  // Group 0.  Start / stop_all
	EventWorkout               Event = 3  // start / stop
	EventWorkoutStep           Event = 4  // Start at beginning of workout.  Stop at end of each step.
	EventPowerDown             Event = 5  // stop_all group 0
	EventPowerUp               Event = 6  // stop_all group 0
	EventOffCourse             Event = 7  // start / stop group 0
	EventSession               Event = 8  // Stop at end of each session.
	EventLap                   Event = 9  // Stop at end of each lap.
	EventCoursePoint           Event = 10 // marker
	EventBattery               Event = 11 // marker
	EventVirtualPartnerPace    Event = 12 // Group 1. Start at beginning of activity if VP enabled, when VP pace is changed during activity or VP enabled mid activity.  stop_disable when VP disabled.
	EventHrHighAlert           Event = 13 // Group 0.  Start / stop when in alert condition.
	EventHrLowAlert            Event = 14 // Group 0.  Start / stop when in alert condition.
	EventSpeedHighAlert        Event = 15 // Group 0.  Start / stop when in alert condition.
	EventSpeedLowAlert         Event = 16 // Group 0.  Start / stop when in alert condition.
	EventCadHighAlert          Event = 17 // Group 0.  Start / stop when in alert condition.
	EventCadLowAlert           Event = 18 // Group 0.  Start / stop when in alert condition.
	EventPowerHighAlert        Event = 19 // Group 0.  Start / stop when in alert condition.
	EventPowerLowAlert         Event = 20 // Group 0.  Start / stop when in alert condition.
	EventRecoveryHr            Event = 21 // marker
	EventBatteryLow            Event = 22 // marker
	EventTimeDurationAlert     Event = 23 // Group 1.  Start if enabled mid activity (not required at start of activity). Stop when duration is reached.  stop_disable if disabled.
	EventDistanceDurationAlert Event = 24 // Group 1.  Start if enabled mid activity (not required at start of activity). Stop when duration is reached.  stop_disable if disabled.
	EventCalorieDurationAlert  Event = 25 // Group 1.  Start if enabled mid activity (not required at start of activity). Stop when duration is reached.  stop_disable if disabled.
	EventActivity              Event = 26 // Group 1..  Stop at end of activity.
	EventFitnessEquipment      Event = 27 // marker
	EventLength                Event = 28 // Stop at end of each length.
	EventUserMarker            Event = 32 // marker
	EventSportPoint            Event = 33 // marker
	EventCalibration           Event = 36 // start/stop/marker
	EventFrontGearChange       Event = 42 // marker
	EventRearGearChange        Event = 43 // marker
	EventRiderPositionChange   Event = 44 // marker
	EventElevHighAlert         Event = 45 // Group 0.  Start / stop when in alert condition.
	EventElevLowAlert          Event = 46 // Group 0.  Start / stop when in alert condition.
	EventCommTimeout           Event = 47 // marker
	EventInvalid               Event = 0xFF
)

// EventType represents the event_type FIT type.
type EventType byte

const (
	EventTypeStart                  EventType = 0
	EventTypeStop                   EventType = 1
	EventTypeConsecutiveDepreciated EventType = 2
	EventTypeMarker                 EventType = 3
	EventTypeStopAll                EventType = 4
	EventTypeBeginDepreciated       EventType = 5
	EventTypeEndDepreciated         EventType = 6
	EventTypeEndAllDepreciated      EventType = 7
	EventTypeStopDisable            EventType = 8
	EventTypeStopDisableAll         EventType = 9
	EventTypeInvalid                EventType = 0xFF
)

// File represents the file FIT type.
type File byte

const (
	FileDevice          File = 1  // Read only, single file. Must be in root directory.
	FileSettings        File = 2  // Read/write, single file. Directory=Settings
	FileSport           File = 3  // Read/write, multiple files, file number = sport type. Directory=Sports
	FileActivity        File = 4  // Read/erase, multiple files. Directory=Activities
	FileWorkout         File = 5  // Read/write/erase, multiple files. Directory=Workouts
	FileCourse          File = 6  // Read/write/erase, multiple files. Directory=Courses
	FileSchedules       File = 7  // Read/write, single file. Directory=Schedules
	FileWeight          File = 9  // Read only, single file. Circular buffer. All message definitions at start of file. Directory=Weight
	FileTotals          File = 10 // Read only, single file. Directory=Totals
	FileGoals           File = 11 // Read/write, single file. Directory=Goals
	FileBloodPressure   File = 14 // Read only. Directory=Blood Pressure
	FileMonitoringA     File = 15 // Read only. Directory=Monitoring. File number=sub type.
	FileActivitySummary File = 20 // Read/erase, multiple files. Directory=Activities
	FileMonitoringDaily File = 28
	FileMonitoringB     File = 32   // Read only. Directory=Monitoring. File number=identifier
	FileSegment         File = 34   // Read/write/erase. Multiple Files.  Directory=Segments
	FileSegmentList     File = 35   // Read/write/erase. Single File.  Directory=Segments
	FileMfgRangeMin     File = 0xF7 // 0xF7 - 0xFE reserved for manufacturer specific file types
	FileMfgRangeMax     File = 0xFE // 0xF7 - 0xFE reserved for manufacturer specific file types
	FileInvalid         File = 0xFF
)

// FileFlags represents the file_flags FIT type.
type FileFlags uint8

const (
	FileFlagsRead    FileFlags = 0x02
	FileFlagsWrite   FileFlags = 0x04
	FileFlagsErase   FileFlags = 0x08
	FileFlagsInvalid FileFlags = 0x00
)

// FitnessEquipmentState represents the fitness_equipment_state FIT type.
type FitnessEquipmentState byte

const (
	FitnessEquipmentStateReady   FitnessEquipmentState = 0
	FitnessEquipmentStateInUse   FitnessEquipmentState = 1
	FitnessEquipmentStatePaused  FitnessEquipmentState = 2
	FitnessEquipmentStateUnknown FitnessEquipmentState = 3 // lost connection to fitness equipment
	FitnessEquipmentStateInvalid FitnessEquipmentState = 0xFF
)

// GarminProduct represents the garmin_product FIT type.
type GarminProduct uint16

const (
	GarminProductHrm1                      GarminProduct = 1
	GarminProductAxh01                     GarminProduct = 2 // AXH01 HRM chipset
	GarminProductAxb01                     GarminProduct = 3
	GarminProductAxb02                     GarminProduct = 4
	GarminProductHrm2ss                    GarminProduct = 5
	GarminProductDsiAlf02                  GarminProduct = 6
	GarminProductHrm3ss                    GarminProduct = 7
	GarminProductHrmRunSingleByteProductId GarminProduct = 8  // hrm_run model for HRM ANT+ messaging
	GarminProductBsm                       GarminProduct = 9  // BSM model for ANT+ messaging
	GarminProductBcm                       GarminProduct = 10 // BCM model for ANT+ messaging
	GarminProductAxs01                     GarminProduct = 11 // AXS01 HRM Bike Chipset model for ANT+ messaging
	GarminProductHrmTriSingleByteProductId GarminProduct = 12 // hrm_tri model for HRM ANT+ messaging
	GarminProductFr225SingleByteProductId  GarminProduct = 14 // fr225 model for HRM ANT+ messaging
	GarminProductFr301China                GarminProduct = 473
	GarminProductFr301Japan                GarminProduct = 474
	GarminProductFr301Korea                GarminProduct = 475
	GarminProductFr301Taiwan               GarminProduct = 494
	GarminProductFr405                     GarminProduct = 717 // Forerunner 405
	GarminProductFr50                      GarminProduct = 782 // Forerunner 50
	GarminProductFr405Japan                GarminProduct = 987
	GarminProductFr60                      GarminProduct = 988 // Forerunner 60
	GarminProductDsiAlf01                  GarminProduct = 1011
	GarminProductFr310xt                   GarminProduct = 1018 // Forerunner 310
	GarminProductEdge500                   GarminProduct = 1036
	GarminProductFr110                     GarminProduct = 1124 // Forerunner 110
	GarminProductEdge800                   GarminProduct = 1169
	GarminProductEdge500Taiwan             GarminProduct = 1199
	GarminProductEdge500Japan              GarminProduct = 1213
	GarminProductChirp                     GarminProduct = 1253
	GarminProductFr110Japan                GarminProduct = 1274
	GarminProductEdge200                   GarminProduct = 1325
	GarminProductFr910xt                   GarminProduct = 1328
	GarminProductEdge800Taiwan             GarminProduct = 1333
	GarminProductEdge800Japan              GarminProduct = 1334
	GarminProductAlf04                     GarminProduct = 1341
	GarminProductFr610                     GarminProduct = 1345
	GarminProductFr210Japan                GarminProduct = 1360
	GarminProductVectorSs                  GarminProduct = 1380
	GarminProductVectorCp                  GarminProduct = 1381
	GarminProductEdge800China              GarminProduct = 1386
	GarminProductEdge500China              GarminProduct = 1387
	GarminProductFr610Japan                GarminProduct = 1410
	GarminProductEdge500Korea              GarminProduct = 1422
	GarminProductFr70                      GarminProduct = 1436
	GarminProductFr310xt4t                 GarminProduct = 1446
	GarminProductAmx                       GarminProduct = 1461
	GarminProductFr10                      GarminProduct = 1482
	GarminProductEdge800Korea              GarminProduct = 1497
	GarminProductSwim                      GarminProduct = 1499
	GarminProductFr910xtChina              GarminProduct = 1537
	GarminProductFenix                     GarminProduct = 1551
	GarminProductEdge200Taiwan             GarminProduct = 1555
	GarminProductEdge510                   GarminProduct = 1561
	GarminProductEdge810                   GarminProduct = 1567
	GarminProductTempe                     GarminProduct = 1570
	GarminProductFr910xtJapan              GarminProduct = 1600
	GarminProductFr620                     GarminProduct = 1623
	GarminProductFr220                     GarminProduct = 1632
	GarminProductFr910xtKorea              GarminProduct = 1664
	GarminProductFr10Japan                 GarminProduct = 1688
	GarminProductEdge810Japan              GarminProduct = 1721
	GarminProductVirbElite                 GarminProduct = 1735
	GarminProductEdgeTouring               GarminProduct = 1736 // Also Edge Touring Plus
	GarminProductEdge510Japan              GarminProduct = 1742
	GarminProductHrmTri                    GarminProduct = 1743
	GarminProductHrmRun                    GarminProduct = 1752
	GarminProductFr920xt                   GarminProduct = 1765
	GarminProductEdge510Asia               GarminProduct = 1821
	GarminProductEdge810China              GarminProduct = 1822
	GarminProductEdge810Taiwan             GarminProduct = 1823
	GarminProductEdge1000                  GarminProduct = 1836
	GarminProductVivoFit                   GarminProduct = 1837
	GarminProductVirbRemote                GarminProduct = 1853
	GarminProductVivoKi                    GarminProduct = 1885
	GarminProductFr15                      GarminProduct = 1903
	GarminProductVivoActive                GarminProduct = 1907
	GarminProductEdge510Korea              GarminProduct = 1918
	GarminProductFr620Japan                GarminProduct = 1928
	GarminProductFr620China                GarminProduct = 1929
	GarminProductFr220Japan                GarminProduct = 1930
	GarminProductFr220China                GarminProduct = 1931
	GarminProductApproachS6                GarminProduct = 1936
	GarminProductVivoSmart                 GarminProduct = 1956
	GarminProductFenix2                    GarminProduct = 1967
	GarminProductEpix                      GarminProduct = 1988
	GarminProductFenix3                    GarminProduct = 2050
	GarminProductEdge1000Taiwan            GarminProduct = 2052
	GarminProductEdge1000Japan             GarminProduct = 2053
	GarminProductFr15Japan                 GarminProduct = 2061
	GarminProductEdge520                   GarminProduct = 2067
	GarminProductEdge1000China             GarminProduct = 2070
	GarminProductFr620Russia               GarminProduct = 2072
	GarminProductFr220Russia               GarminProduct = 2073
	GarminProductVectorS                   GarminProduct = 2079
	GarminProductEdge1000Korea             GarminProduct = 2100
	GarminProductFr920xtTaiwan             GarminProduct = 2130
	GarminProductFr920xtChina              GarminProduct = 2131
	GarminProductFr920xtJapan              GarminProduct = 2132
	GarminProductVirbx                     GarminProduct = 2134
	GarminProductVivoSmartApac             GarminProduct = 2135
	GarminProductEtrexTouch                GarminProduct = 2140
	GarminProductEdge25                    GarminProduct = 2147
	GarminProductVivoFit2                  GarminProduct = 2150
	GarminProductFr225                     GarminProduct = 2153
	GarminProductVivoActiveApac            GarminProduct = 2160
	GarminProductVector2                   GarminProduct = 2161
	GarminProductVector2s                  GarminProduct = 2162
	GarminProductVirbxe                    GarminProduct = 2172
	GarminProductFr620Taiwan               GarminProduct = 2173
	GarminProductFr220Taiwan               GarminProduct = 2174
	GarminProductFenix3China               GarminProduct = 2188
	GarminProductFenix3Twn                 GarminProduct = 2189
	GarminProductVariaHeadlight            GarminProduct = 2192
	GarminProductVariaTaillightOld         GarminProduct = 2193
	GarminProductFr225Asia                 GarminProduct = 2219
	GarminProductVariaRadarTaillight       GarminProduct = 2225
	GarminProductVariaRadarDisplay         GarminProduct = 2226
	GarminProductEdge20                    GarminProduct = 2238
	GarminProductD2Bravo                   GarminProduct = 2262
	GarminProductVariaRemote               GarminProduct = 2276
	GarminProductSdm4                      GarminProduct = 10007 // SDM4 footpod
	GarminProductEdgeRemote                GarminProduct = 10014
	GarminProductTrainingCenter            GarminProduct = 20119
	GarminProductAndroidAntplusPlugin      GarminProduct = 65532
	GarminProductConnect                   GarminProduct = 65534 // Garmin Connect website
	GarminProductInvalid                   GarminProduct = 0xFFFF
)

// Gender represents the gender FIT type.
type Gender byte

const (
	GenderFemale  Gender = 0
	GenderMale    Gender = 1
	GenderInvalid Gender = 0xFF
)

// Goal represents the goal FIT type.
type Goal byte

const (
	GoalTime      Goal = 0
	GoalDistance  Goal = 1
	GoalCalories  Goal = 2
	GoalFrequency Goal = 3
	GoalSteps     Goal = 4
	GoalInvalid   Goal = 0xFF
)

// GoalRecurrence represents the goal_recurrence FIT type.
type GoalRecurrence byte

const (
	GoalRecurrenceOff     GoalRecurrence = 0
	GoalRecurrenceDaily   GoalRecurrence = 1
	GoalRecurrenceWeekly  GoalRecurrence = 2
	GoalRecurrenceMonthly GoalRecurrence = 3
	GoalRecurrenceYearly  GoalRecurrence = 4
	GoalRecurrenceCustom  GoalRecurrence = 5
	GoalRecurrenceInvalid GoalRecurrence = 0xFF
)

// HrType represents the hr_type FIT type.
type HrType byte

const (
	HrTypeNormal    HrType = 0
	HrTypeIrregular HrType = 1
	HrTypeInvalid   HrType = 0xFF
)

// HrZoneCalc represents the hr_zone_calc FIT type.
type HrZoneCalc byte

const (
	HrZoneCalcCustom       HrZoneCalc = 0
	HrZoneCalcPercentMaxHr HrZoneCalc = 1
	HrZoneCalcPercentHrr   HrZoneCalc = 2
	HrZoneCalcInvalid      HrZoneCalc = 0xFF
)

// Intensity represents the intensity FIT type.
type Intensity byte

const (
	IntensityActive   Intensity = 0
	IntensityRest     Intensity = 1
	IntensityWarmup   Intensity = 2
	IntensityCooldown Intensity = 3
	IntensityInvalid  Intensity = 0xFF
)

// Language represents the language FIT type.
type Language byte

const (
	LanguageEnglish    Language = 0
	LanguageFrench     Language = 1
	LanguageItalian    Language = 2
	LanguageGerman     Language = 3
	LanguageSpanish    Language = 4
	LanguageCroatian   Language = 5
	LanguageCzech      Language = 6
	LanguageDanish     Language = 7
	LanguageDutch      Language = 8
	LanguageFinnish    Language = 9
	LanguageGreek      Language = 10
	LanguageHungarian  Language = 11
	LanguageNorwegian  Language = 12
	LanguagePolish     Language = 13
	LanguagePortuguese Language = 14
	LanguageSlovakian  Language = 15
	LanguageSlovenian  Language = 16
	LanguageSwedish    Language = 17
	LanguageRussian    Language = 18
	LanguageTurkish    Language = 19
	LanguageLatvian    Language = 20
	LanguageUkrainian  Language = 21
	LanguageArabic     Language = 22
	LanguageFarsi      Language = 23
	LanguageBulgarian  Language = 24
	LanguageRomanian   Language = 25
	LanguageCustom     Language = 254
	LanguageInvalid    Language = 0xFF
)

// LapTrigger represents the lap_trigger FIT type.
type LapTrigger byte

const (
	LapTriggerManual           LapTrigger = 0
	LapTriggerTime             LapTrigger = 1
	LapTriggerDistance         LapTrigger = 2
	LapTriggerPositionStart    LapTrigger = 3
	LapTriggerPositionLap      LapTrigger = 4
	LapTriggerPositionWaypoint LapTrigger = 5
	LapTriggerPositionMarked   LapTrigger = 6
	LapTriggerSessionEnd       LapTrigger = 7
	LapTriggerFitnessEquipment LapTrigger = 8
	LapTriggerInvalid          LapTrigger = 0xFF
)

// LeftRightBalance represents the left_right_balance FIT type.
type LeftRightBalance uint8

const (
	LeftRightBalanceMask    LeftRightBalance = 0x7F // % contribution
	LeftRightBalanceRight   LeftRightBalance = 0x80 // data corresponds to right if set, otherwise unknown
	LeftRightBalanceInvalid LeftRightBalance = 0xFF
)

// LeftRightBalance100 represents the left_right_balance_100 FIT type.
type LeftRightBalance100 uint16

const (
	LeftRightBalance100Mask    LeftRightBalance100 = 0x3FFF // % contribution scaled by 100
	LeftRightBalance100Right   LeftRightBalance100 = 0x8000 // data corresponds to right if set, otherwise unknown
	LeftRightBalance100Invalid LeftRightBalance100 = 0xFFFF
)

// LengthType represents the length_type FIT type.
type LengthType byte

const (
	LengthTypeIdle    LengthType = 0 // Rest period. Length with no strokes
	LengthTypeActive  LengthType = 1 // Length with strokes.
	LengthTypeInvalid LengthType = 0xFF
)

// Manufacturer represents the manufacturer FIT type.
type Manufacturer uint16

const (
	ManufacturerGarmin                 Manufacturer = 1
	ManufacturerGarminFr405Antfs       Manufacturer = 2 // Do not use.  Used by FR405 for ANTFS man id.
	ManufacturerZephyr                 Manufacturer = 3
	ManufacturerDayton                 Manufacturer = 4
	ManufacturerIdt                    Manufacturer = 5
	ManufacturerSrm                    Manufacturer = 6
	ManufacturerQuarq                  Manufacturer = 7
	ManufacturerIbike                  Manufacturer = 8
	ManufacturerSaris                  Manufacturer = 9
	ManufacturerSparkHk                Manufacturer = 10
	ManufacturerTanita                 Manufacturer = 11
	ManufacturerEchowell               Manufacturer = 12
	ManufacturerDynastreamOem          Manufacturer = 13
	ManufacturerNautilus               Manufacturer = 14
	ManufacturerDynastream             Manufacturer = 15
	ManufacturerTimex                  Manufacturer = 16
	ManufacturerMetrigear              Manufacturer = 17
	ManufacturerXelic                  Manufacturer = 18
	ManufacturerBeurer                 Manufacturer = 19
	ManufacturerCardiosport            Manufacturer = 20
	ManufacturerAAndD                  Manufacturer = 21
	ManufacturerHmm                    Manufacturer = 22
	ManufacturerSuunto                 Manufacturer = 23
	ManufacturerThitaElektronik        Manufacturer = 24
	ManufacturerGpulse                 Manufacturer = 25
	ManufacturerCleanMobile            Manufacturer = 26
	ManufacturerPedalBrain             Manufacturer = 27
	ManufacturerPeaksware              Manufacturer = 28
	ManufacturerSaxonar                Manufacturer = 29
	ManufacturerLemondFitness          Manufacturer = 30
	ManufacturerDexcom                 Manufacturer = 31
	ManufacturerWahooFitness           Manufacturer = 32
	ManufacturerOctaneFitness          Manufacturer = 33
	ManufacturerArchinoetics           Manufacturer = 34
	ManufacturerTheHurtBox             Manufacturer = 35
	ManufacturerCitizenSystems         Manufacturer = 36
	ManufacturerMagellan               Manufacturer = 37
	ManufacturerOsynce                 Manufacturer = 38
	ManufacturerHolux                  Manufacturer = 39
	ManufacturerConcept2               Manufacturer = 40
	ManufacturerOneGiantLeap           Manufacturer = 42
	ManufacturerAceSensor              Manufacturer = 43
	ManufacturerBrimBrothers           Manufacturer = 44
	ManufacturerXplova                 Manufacturer = 45
	ManufacturerPerceptionDigital      Manufacturer = 46
	ManufacturerBf1systems             Manufacturer = 47
	ManufacturerPioneer                Manufacturer = 48
	ManufacturerSpantec                Manufacturer = 49
	ManufacturerMetalogics             Manufacturer = 50
	Manufacturer4iiiis                 Manufacturer = 51
	ManufacturerSeikoEpson             Manufacturer = 52
	ManufacturerSeikoEpsonOem          Manufacturer = 53
	ManufacturerIforPowell             Manufacturer = 54
	ManufacturerMaxwellGuider          Manufacturer = 55
	ManufacturerStarTrac               Manufacturer = 56
	ManufacturerBreakaway              Manufacturer = 57
	ManufacturerAlatechTechnologyLtd   Manufacturer = 58
	ManufacturerMioTechnologyEurope    Manufacturer = 59
	ManufacturerRotor                  Manufacturer = 60
	ManufacturerGeonaute               Manufacturer = 61
	ManufacturerIdBike                 Manufacturer = 62
	ManufacturerSpecialized            Manufacturer = 63
	ManufacturerWtek                   Manufacturer = 64
	ManufacturerPhysicalEnterprises    Manufacturer = 65
	ManufacturerNorthPoleEngineering   Manufacturer = 66
	ManufacturerBkool                  Manufacturer = 67
	ManufacturerCateye                 Manufacturer = 68
	ManufacturerStagesCycling          Manufacturer = 69
	ManufacturerSigmasport             Manufacturer = 70
	ManufacturerTomtom                 Manufacturer = 71
	ManufacturerPeripedal              Manufacturer = 72
	ManufacturerWattbike               Manufacturer = 73
	ManufacturerMoxy                   Manufacturer = 76
	ManufacturerCiclosport             Manufacturer = 77
	ManufacturerPowerbahn              Manufacturer = 78
	ManufacturerAcornProjectsAps       Manufacturer = 79
	ManufacturerLifebeam               Manufacturer = 80
	ManufacturerBontrager              Manufacturer = 81
	ManufacturerWellgo                 Manufacturer = 82
	ManufacturerScosche                Manufacturer = 83
	ManufacturerMagura                 Manufacturer = 84
	ManufacturerWoodway                Manufacturer = 85
	ManufacturerElite                  Manufacturer = 86
	ManufacturerNielsenKellerman       Manufacturer = 87
	ManufacturerDkCity                 Manufacturer = 88
	ManufacturerTacx                   Manufacturer = 89
	ManufacturerDirectionTechnology    Manufacturer = 90
	ManufacturerMagtonic               Manufacturer = 91
	Manufacturer1partcarbon            Manufacturer = 92
	ManufacturerInsideRideTechnologies Manufacturer = 93
	ManufacturerSoundOfMotion          Manufacturer = 94
	ManufacturerStryd                  Manufacturer = 95
	ManufacturerIcg                    Manufacturer = 96 // Indoorcycling Group
	ManufacturerMiPulse                Manufacturer = 97
	ManufacturerBsxAthletics           Manufacturer = 98
	ManufacturerLook                   Manufacturer = 99
	ManufacturerDevelopment            Manufacturer = 255
	ManufacturerHealthandlife          Manufacturer = 257
	ManufacturerLezyne                 Manufacturer = 258
	ManufacturerScribeLabs             Manufacturer = 259
	ManufacturerZwift                  Manufacturer = 260
	ManufacturerWatteam                Manufacturer = 261
	ManufacturerRecon                  Manufacturer = 262
	ManufacturerFaveroElectronics      Manufacturer = 263
	ManufacturerDynovelo               Manufacturer = 264
	ManufacturerStrava                 Manufacturer = 265
	ManufacturerActigraphcorp          Manufacturer = 5759
	ManufacturerInvalid                Manufacturer = 0xFFFF
)

// MesgCount represents the mesg_count FIT type.
type MesgCount byte

const (
	MesgCountNumPerFile     MesgCount = 0
	MesgCountMaxPerFile     MesgCount = 1
	MesgCountMaxPerFileType MesgCount = 2
	MesgCountInvalid        MesgCount = 0xFF
)

// MesgNum represents the mesg_num FIT type.
type MesgNum uint16

const (
	MesgNumFileId                  MesgNum = 0
	MesgNumCapabilities            MesgNum = 1
	MesgNumDeviceSettings          MesgNum = 2
	MesgNumUserProfile             MesgNum = 3
	MesgNumHrmProfile              MesgNum = 4
	MesgNumSdmProfile              MesgNum = 5
	MesgNumBikeProfile             MesgNum = 6
	MesgNumZonesTarget             MesgNum = 7
	MesgNumHrZone                  MesgNum = 8
	MesgNumPowerZone               MesgNum = 9
	MesgNumMetZone                 MesgNum = 10
	MesgNumSport                   MesgNum = 12
	MesgNumGoal                    MesgNum = 15
	MesgNumSession                 MesgNum = 18
	MesgNumLap                     MesgNum = 19
	MesgNumRecord                  MesgNum = 20
	MesgNumEvent                   MesgNum = 21
	MesgNumDeviceInfo              MesgNum = 23
	MesgNumWorkout                 MesgNum = 26
	MesgNumWorkoutStep             MesgNum = 27
	MesgNumSchedule                MesgNum = 28
	MesgNumWeightScale             MesgNum = 30
	MesgNumCourse                  MesgNum = 31
	MesgNumCoursePoint             MesgNum = 32
	MesgNumTotals                  MesgNum = 33
	MesgNumActivity                MesgNum = 34
	MesgNumSoftware                MesgNum = 35
	MesgNumFileCapabilities        MesgNum = 37
	MesgNumMesgCapabilities        MesgNum = 38
	MesgNumFieldCapabilities       MesgNum = 39
	MesgNumFileCreator             MesgNum = 49
	MesgNumBloodPressure           MesgNum = 51
	MesgNumSpeedZone               MesgNum = 53
	MesgNumMonitoring              MesgNum = 55
	MesgNumTrainingFile            MesgNum = 72
	MesgNumHrv                     MesgNum = 78
	MesgNumLength                  MesgNum = 101
	MesgNumMonitoringInfo          MesgNum = 103
	MesgNumPad                     MesgNum = 105
	MesgNumSlaveDevice             MesgNum = 106
	MesgNumCadenceZone             MesgNum = 131
	MesgNumSegmentLap              MesgNum = 142
	MesgNumMemoGlob                MesgNum = 145
	MesgNumSegmentId               MesgNum = 148
	MesgNumSegmentLeaderboardEntry MesgNum = 149
	MesgNumSegmentPoint            MesgNum = 150
	MesgNumSegmentFile             MesgNum = 151
	MesgNumGpsMetadata             MesgNum = 160
	MesgNumCameraEvent             MesgNum = 161
	MesgNumTimestampCorrelation    MesgNum = 162
	MesgNumGyroscopeData           MesgNum = 164
	MesgNumAccelerometerData       MesgNum = 165
	MesgNumThreeDSensorCalibration MesgNum = 167
	MesgNumVideoFrame              MesgNum = 169
	MesgNumObdiiData               MesgNum = 174
	MesgNumNmeaSentence            MesgNum = 177
	MesgNumAviationAttitude        MesgNum = 178
	MesgNumVideo                   MesgNum = 184
	MesgNumVideoTitle              MesgNum = 185
	MesgNumVideoDescription        MesgNum = 186
	MesgNumVideoClip               MesgNum = 187
	MesgNumMfgRangeMin             MesgNum = 0xFF00 // 0xFF00 - 0xFFFE reserved for manufacturer specific messages
	MesgNumMfgRangeMax             MesgNum = 0xFFFE // 0xFF00 - 0xFFFE reserved for manufacturer specific messages
	MesgNumInvalid                 MesgNum = 0xFFFF
)

// MessageIndex represents the message_index FIT type.
type MessageIndex uint16

const (
	MessageIndexSelected MessageIndex = 0x8000 // message is selected if set
	MessageIndexReserved MessageIndex = 0x7000 // reserved (default 0)
	MessageIndexMask     MessageIndex = 0x0FFF // index
	MessageIndexInvalid  MessageIndex = 0xFFFF
)

// PowerPhaseType represents the power_phase_type FIT type.
type PowerPhaseType byte

const (
	PowerPhaseTypePowerPhaseStartAngle PowerPhaseType = 0
	PowerPhaseTypePowerPhaseEndAngle   PowerPhaseType = 1
	PowerPhaseTypePowerPhaseArcLength  PowerPhaseType = 2
	PowerPhaseTypePowerPhaseCenter     PowerPhaseType = 3
	PowerPhaseTypeInvalid              PowerPhaseType = 0xFF
)

// PwrZoneCalc represents the pwr_zone_calc FIT type.
type PwrZoneCalc byte

const (
	PwrZoneCalcCustom     PwrZoneCalc = 0
	PwrZoneCalcPercentFtp PwrZoneCalc = 1
	PwrZoneCalcInvalid    PwrZoneCalc = 0xFF
)

// RiderPositionType represents the rider_position_type FIT type.
type RiderPositionType byte

const (
	RiderPositionTypeSeated   RiderPositionType = 0
	RiderPositionTypeStanding RiderPositionType = 1
	RiderPositionTypeInvalid  RiderPositionType = 0xFF
)

// Schedule represents the schedule FIT type.
type Schedule byte

const (
	ScheduleWorkout Schedule = 0
	ScheduleCourse  Schedule = 1
	ScheduleInvalid Schedule = 0xFF
)

// SegmentDeleteStatus represents the segment_delete_status FIT type.
type SegmentDeleteStatus byte

const (
	SegmentDeleteStatusDoNotDelete SegmentDeleteStatus = 0
	SegmentDeleteStatusDeleteOne   SegmentDeleteStatus = 1
	SegmentDeleteStatusDeleteAll   SegmentDeleteStatus = 2
	SegmentDeleteStatusInvalid     SegmentDeleteStatus = 0xFF
)

// SegmentLapStatus represents the segment_lap_status FIT type.
type SegmentLapStatus byte

const (
	SegmentLapStatusEnd     SegmentLapStatus = 0
	SegmentLapStatusFail    SegmentLapStatus = 1
	SegmentLapStatusInvalid SegmentLapStatus = 0xFF
)

// SegmentLeaderboardType represents the segment_leaderboard_type FIT type.
type SegmentLeaderboardType byte

const (
	SegmentLeaderboardTypeOverall      SegmentLeaderboardType = 0
	SegmentLeaderboardTypePersonalBest SegmentLeaderboardType = 1
	SegmentLeaderboardTypeConnections  SegmentLeaderboardType = 2
	SegmentLeaderboardTypeGroup        SegmentLeaderboardType = 3
	SegmentLeaderboardTypeChallenger   SegmentLeaderboardType = 4
	SegmentLeaderboardTypeKom          SegmentLeaderboardType = 5
	SegmentLeaderboardTypeQom          SegmentLeaderboardType = 6
	SegmentLeaderboardTypePr           SegmentLeaderboardType = 7
	SegmentLeaderboardTypeGoal         SegmentLeaderboardType = 8
	SegmentLeaderboardTypeRival        SegmentLeaderboardType = 9
	SegmentLeaderboardTypeClubLeader   SegmentLeaderboardType = 10
	SegmentLeaderboardTypeInvalid      SegmentLeaderboardType = 0xFF
)

// SegmentSelectionType represents the segment_selection_type FIT type.
type SegmentSelectionType byte

const (
	SegmentSelectionTypeStarred   SegmentSelectionType = 0
	SegmentSelectionTypeSuggested SegmentSelectionType = 1
	SegmentSelectionTypeInvalid   SegmentSelectionType = 0xFF
)

// SensorType represents the sensor_type FIT type.
type SensorType byte

const (
	SensorTypeAccelerometer SensorType = 0
	SensorTypeGyroscope     SensorType = 1
	SensorTypeCompass       SensorType = 2
	SensorTypeInvalid       SensorType = 0xFF
)

// SessionTrigger represents the session_trigger FIT type.
type SessionTrigger byte

const (
	SessionTriggerActivityEnd      SessionTrigger = 0
	SessionTriggerManual           SessionTrigger = 1 // User changed sport.
	SessionTriggerAutoMultiSport   SessionTrigger = 2 // Auto multi-sport feature is enabled and user pressed lap button to advance session.
	SessionTriggerFitnessEquipment SessionTrigger = 3 // Auto sport change caused by user linking to fitness equipment.
	SessionTriggerInvalid          SessionTrigger = 0xFF
)

// SourceType represents the source_type FIT type.
type SourceType byte

const (
	SourceTypeAnt                SourceType = 0 // External device connected with ANT
	SourceTypeAntplus            SourceType = 1 // External device connected with ANT+
	SourceTypeBluetooth          SourceType = 2 // External device connected with BT
	SourceTypeBluetoothLowEnergy SourceType = 3 // External device connected with BLE
	SourceTypeWifi               SourceType = 4 // External device connected with Wifi
	SourceTypeLocal              SourceType = 5 // Onboard device
	SourceTypeInvalid            SourceType = 0xFF
)

// Sport represents the sport FIT type.
type Sport byte

const (
	SportGeneric               Sport = 0
	SportRunning               Sport = 1
	SportCycling               Sport = 2
	SportTransition            Sport = 3 // Mulitsport transition
	SportFitnessEquipment      Sport = 4
	SportSwimming              Sport = 5
	SportBasketball            Sport = 6
	SportSoccer                Sport = 7
	SportTennis                Sport = 8
	SportAmericanFootball      Sport = 9
	SportTraining              Sport = 10
	SportWalking               Sport = 11
	SportCrossCountrySkiing    Sport = 12
	SportAlpineSkiing          Sport = 13
	SportSnowboarding          Sport = 14
	SportRowing                Sport = 15
	SportMountaineering        Sport = 16
	SportHiking                Sport = 17
	SportMultisport            Sport = 18
	SportPaddling              Sport = 19
	SportFlying                Sport = 20
	SportEBiking               Sport = 21
	SportMotorcycling          Sport = 22
	SportBoating               Sport = 23
	SportDriving               Sport = 24
	SportGolf                  Sport = 25
	SportHangGliding           Sport = 26
	SportHorsebackRiding       Sport = 27
	SportHunting               Sport = 28
	SportFishing               Sport = 29
	SportInlineSkating         Sport = 30
	SportRockClimbing          Sport = 31
	SportSailing               Sport = 32
	SportIceSkating            Sport = 33
	SportSkyDiving             Sport = 34
	SportSnowshoeing           Sport = 35
	SportSnowmobiling          Sport = 36
	SportStandUpPaddleboarding Sport = 37
	SportSurfing               Sport = 38
	SportWakeboarding          Sport = 39
	SportWaterSkiing           Sport = 40
	SportKayaking              Sport = 41
	SportRafting               Sport = 42
	SportWindsurfing           Sport = 43
	SportKitesurfing           Sport = 44
	SportAll                   Sport = 254 // All is for goals only to include all sports.
	SportInvalid               Sport = 0xFF
)

// SportBits0 represents the sport_bits_0 FIT type.
type SportBits0 uint8

const (
	SportBits0Generic          SportBits0 = 0x01
	SportBits0Running          SportBits0 = 0x02
	SportBits0Cycling          SportBits0 = 0x04
	SportBits0Transition       SportBits0 = 0x08 // Mulitsport transition
	SportBits0FitnessEquipment SportBits0 = 0x10
	SportBits0Swimming         SportBits0 = 0x20
	SportBits0Basketball       SportBits0 = 0x40
	SportBits0Soccer           SportBits0 = 0x80
	SportBits0Invalid          SportBits0 = 0x00
)

// SportBits1 represents the sport_bits_1 FIT type.
type SportBits1 uint8

const (
	SportBits1Tennis             SportBits1 = 0x01
	SportBits1AmericanFootball   SportBits1 = 0x02
	SportBits1Training           SportBits1 = 0x04
	SportBits1Walking            SportBits1 = 0x08
	SportBits1CrossCountrySkiing SportBits1 = 0x10
	SportBits1AlpineSkiing       SportBits1 = 0x20
	SportBits1Snowboarding       SportBits1 = 0x40
	SportBits1Rowing             SportBits1 = 0x80
	SportBits1Invalid            SportBits1 = 0x00
)

// SportBits2 represents the sport_bits_2 FIT type.
type SportBits2 uint8

const (
	SportBits2Mountaineering SportBits2 = 0x01
	SportBits2Hiking         SportBits2 = 0x02
	SportBits2Multisport     SportBits2 = 0x04
	SportBits2Paddling       SportBits2 = 0x08
	SportBits2Flying         SportBits2 = 0x10
	SportBits2EBiking        SportBits2 = 0x20
	SportBits2Motorcycling   SportBits2 = 0x40
	SportBits2Boating        SportBits2 = 0x80
	SportBits2Invalid        SportBits2 = 0x00
)

// SportBits3 represents the sport_bits_3 FIT type.
type SportBits3 uint8

const (
	SportBits3Driving         SportBits3 = 0x01
	SportBits3Golf            SportBits3 = 0x02
	SportBits3HangGliding     SportBits3 = 0x04
	SportBits3HorsebackRiding SportBits3 = 0x08
	SportBits3Hunting         SportBits3 = 0x10
	SportBits3Fishing         SportBits3 = 0x20
	SportBits3InlineSkating   SportBits3 = 0x40
	SportBits3RockClimbing    SportBits3 = 0x80
	SportBits3Invalid         SportBits3 = 0x00
)

// SportBits4 represents the sport_bits_4 FIT type.
type SportBits4 uint8

const (
	SportBits4Sailing               SportBits4 = 0x01
	SportBits4IceSkating            SportBits4 = 0x02
	SportBits4SkyDiving             SportBits4 = 0x04
	SportBits4Snowshoeing           SportBits4 = 0x08
	SportBits4Snowmobiling          SportBits4 = 0x10
	SportBits4StandUpPaddleboarding SportBits4 = 0x20
	SportBits4Surfing               SportBits4 = 0x40
	SportBits4Wakeboarding          SportBits4 = 0x80
	SportBits4Invalid               SportBits4 = 0x00
)

// SportBits5 represents the sport_bits_5 FIT type.
type SportBits5 uint8

const (
	SportBits5WaterSkiing SportBits5 = 0x01
	SportBits5Kayaking    SportBits5 = 0x02
	SportBits5Rafting     SportBits5 = 0x04
	SportBits5Windsurfing SportBits5 = 0x08
	SportBits5Kitesurfing SportBits5 = 0x10
	SportBits5Invalid     SportBits5 = 0x00
)

// SportEvent represents the sport_event FIT type.
type SportEvent byte

const (
	SportEventUncategorized  SportEvent = 0
	SportEventGeocaching     SportEvent = 1
	SportEventFitness        SportEvent = 2
	SportEventRecreation     SportEvent = 3
	SportEventRace           SportEvent = 4
	SportEventSpecialEvent   SportEvent = 5
	SportEventTraining       SportEvent = 6
	SportEventTransportation SportEvent = 7
	SportEventTouring        SportEvent = 8
	SportEventInvalid        SportEvent = 0xFF
)

// StrokeType represents the stroke_type FIT type.
type StrokeType byte

const (
	StrokeTypeNoEvent  StrokeType = 0
	StrokeTypeOther    StrokeType = 1 // stroke was detected but cannot be identified
	StrokeTypeServe    StrokeType = 2
	StrokeTypeForehand StrokeType = 3
	StrokeTypeBackhand StrokeType = 4
	StrokeTypeSmash    StrokeType = 5
	StrokeTypeInvalid  StrokeType = 0xFF
)

// SubSport represents the sub_sport FIT type.
type SubSport byte

const (
	SubSportGeneric              SubSport = 0
	SubSportTreadmill            SubSport = 1  // Run/Fitness Equipment
	SubSportStreet               SubSport = 2  // Run
	SubSportTrail                SubSport = 3  // Run
	SubSportTrack                SubSport = 4  // Run
	SubSportSpin                 SubSport = 5  // Cycling
	SubSportIndoorCycling        SubSport = 6  // Cycling/Fitness Equipment
	SubSportRoad                 SubSport = 7  // Cycling
	SubSportMountain             SubSport = 8  // Cycling
	SubSportDownhill             SubSport = 9  // Cycling
	SubSportRecumbent            SubSport = 10 // Cycling
	SubSportCyclocross           SubSport = 11 // Cycling
	SubSportHandCycling          SubSport = 12 // Cycling
	SubSportTrackCycling         SubSport = 13 // Cycling
	SubSportIndoorRowing         SubSport = 14 // Fitness Equipment
	SubSportElliptical           SubSport = 15 // Fitness Equipment
	SubSportStairClimbing        SubSport = 16 // Fitness Equipment
	SubSportLapSwimming          SubSport = 17 // Swimming
	SubSportOpenWater            SubSport = 18 // Swimming
	SubSportFlexibilityTraining  SubSport = 19 // Training
	SubSportStrengthTraining     SubSport = 20 // Training
	SubSportWarmUp               SubSport = 21 // Tennis
	SubSportMatch                SubSport = 22 // Tennis
	SubSportExercise             SubSport = 23 // Tennis
	SubSportChallenge            SubSport = 24 // Tennis
	SubSportIndoorSkiing         SubSport = 25 // Fitness Equipment
	SubSportCardioTraining       SubSport = 26 // Training
	SubSportIndoorWalking        SubSport = 27 // Walking/Fitness Equipment
	SubSportEBikeFitness         SubSport = 28 // E-Biking
	SubSportBmx                  SubSport = 29 // Cycling
	SubSportCasualWalking        SubSport = 30 // Walking
	SubSportSpeedWalking         SubSport = 31 // Walking
	SubSportBikeToRunTransition  SubSport = 32 // Transition
	SubSportRunToBikeTransition  SubSport = 33 // Transition
	SubSportSwimToBikeTransition SubSport = 34 // Transition
	SubSportAtv                  SubSport = 35 // Motorcycling
	SubSportMotocross            SubSport = 36 // Motorcycling
	SubSportBackcountry          SubSport = 37 // Alpine Skiing/Snowboarding
	SubSportResort               SubSport = 38 // Alpine Skiing/Snowboarding
	SubSportRcDrone              SubSport = 39 // Flying
	SubSportWingsuit             SubSport = 40 // Flying
	SubSportWhitewater           SubSport = 41 // Kayaking/Rafting
	SubSportAll                  SubSport = 254
	SubSportInvalid              SubSport = 0xFF
)

// SwimStroke represents the swim_stroke FIT type.
type SwimStroke byte

const (
	SwimStrokeFreestyle    SwimStroke = 0
	SwimStrokeBackstroke   SwimStroke = 1
	SwimStrokeBreaststroke SwimStroke = 2
	SwimStrokeButterfly    SwimStroke = 3
	SwimStrokeDrill        SwimStroke = 4
	SwimStrokeMixed        SwimStroke = 5
	SwimStrokeIm           SwimStroke = 6 // IM is a mixed interval containing the same number of lengths for each of: Butterfly, Backstroke, Breaststroke, Freestyle, swam in that order.
	SwimStrokeInvalid      SwimStroke = 0xFF
)

// TimeZone represents the time_zone FIT type.
type TimeZone byte

const (
	TimeZoneAlmaty                   TimeZone = 0
	TimeZoneBangkok                  TimeZone = 1
	TimeZoneBombay                   TimeZone = 2
	TimeZoneBrasilia                 TimeZone = 3
	TimeZoneCairo                    TimeZone = 4
	TimeZoneCapeVerdeIs              TimeZone = 5
	TimeZoneDarwin                   TimeZone = 6
	TimeZoneEniwetok                 TimeZone = 7
	TimeZoneFiji                     TimeZone = 8
	TimeZoneHongKong                 TimeZone = 9
	TimeZoneIslamabad                TimeZone = 10
	TimeZoneKabul                    TimeZone = 11
	TimeZoneMagadan                  TimeZone = 12
	TimeZoneMidAtlantic              TimeZone = 13
	TimeZoneMoscow                   TimeZone = 14
	TimeZoneMuscat                   TimeZone = 15
	TimeZoneNewfoundland             TimeZone = 16
	TimeZoneSamoa                    TimeZone = 17
	TimeZoneSydney                   TimeZone = 18
	TimeZoneTehran                   TimeZone = 19
	TimeZoneTokyo                    TimeZone = 20
	TimeZoneUsAlaska                 TimeZone = 21
	TimeZoneUsAtlantic               TimeZone = 22
	TimeZoneUsCentral                TimeZone = 23
	TimeZoneUsEastern                TimeZone = 24
	TimeZoneUsHawaii                 TimeZone = 25
	TimeZoneUsMountain               TimeZone = 26
	TimeZoneUsPacific                TimeZone = 27
	TimeZoneOther                    TimeZone = 28
	TimeZoneAuckland                 TimeZone = 29
	TimeZoneKathmandu                TimeZone = 30
	TimeZoneEuropeWesternWet         TimeZone = 31
	TimeZoneEuropeCentralCet         TimeZone = 32
	TimeZoneEuropeEasternEet         TimeZone = 33
	TimeZoneJakarta                  TimeZone = 34
	TimeZonePerth                    TimeZone = 35
	TimeZoneAdelaide                 TimeZone = 36
	TimeZoneBrisbane                 TimeZone = 37
	TimeZoneTasmania                 TimeZone = 38
	TimeZoneIceland                  TimeZone = 39
	TimeZoneAmsterdam                TimeZone = 40
	TimeZoneAthens                   TimeZone = 41
	TimeZoneBarcelona                TimeZone = 42
	TimeZoneBerlin                   TimeZone = 43
	TimeZoneBrussels                 TimeZone = 44
	TimeZoneBudapest                 TimeZone = 45
	TimeZoneCopenhagen               TimeZone = 46
	TimeZoneDublin                   TimeZone = 47
	TimeZoneHelsinki                 TimeZone = 48
	TimeZoneLisbon                   TimeZone = 49
	TimeZoneLondon                   TimeZone = 50
	TimeZoneMadrid                   TimeZone = 51
	TimeZoneMunich                   TimeZone = 52
	TimeZoneOslo                     TimeZone = 53
	TimeZoneParis                    TimeZone = 54
	TimeZonePrague                   TimeZone = 55
	TimeZoneReykjavik                TimeZone = 56
	TimeZoneRome                     TimeZone = 57
	TimeZoneStockholm                TimeZone = 58
	TimeZoneVienna                   TimeZone = 59
	TimeZoneWarsaw                   TimeZone = 60
	TimeZoneZurich                   TimeZone = 61
	TimeZoneQuebec                   TimeZone = 62
	TimeZoneOntario                  TimeZone = 63
	TimeZoneManitoba                 TimeZone = 64
	TimeZoneSaskatchewan             TimeZone = 65
	TimeZoneAlberta                  TimeZone = 66
	TimeZoneBritishColumbia          TimeZone = 67
	TimeZoneBoise                    TimeZone = 68
	TimeZoneBoston                   TimeZone = 69
	TimeZoneChicago                  TimeZone = 70
	TimeZoneDallas                   TimeZone = 71
	TimeZoneDenver                   TimeZone = 72
	TimeZoneKansasCity               TimeZone = 73
	TimeZoneLasVegas                 TimeZone = 74
	TimeZoneLosAngeles               TimeZone = 75
	TimeZoneMiami                    TimeZone = 76
	TimeZoneMinneapolis              TimeZone = 77
	TimeZoneNewYork                  TimeZone = 78
	TimeZoneNewOrleans               TimeZone = 79
	TimeZonePhoenix                  TimeZone = 80
	TimeZoneSantaFe                  TimeZone = 81
	TimeZoneSeattle                  TimeZone = 82
	TimeZoneWashingtonDc             TimeZone = 83
	TimeZoneUsArizona                TimeZone = 84
	TimeZoneChita                    TimeZone = 85
	TimeZoneEkaterinburg             TimeZone = 86
	TimeZoneIrkutsk                  TimeZone = 87
	TimeZoneKaliningrad              TimeZone = 88
	TimeZoneKrasnoyarsk              TimeZone = 89
	TimeZoneNovosibirsk              TimeZone = 90
	TimeZonePetropavlovskKamchatskiy TimeZone = 91
	TimeZoneSamara                   TimeZone = 92
	TimeZoneVladivostok              TimeZone = 93
	TimeZoneMexicoCentral            TimeZone = 94
	TimeZoneMexicoMountain           TimeZone = 95
	TimeZoneMexicoPacific            TimeZone = 96
	TimeZoneCapeTown                 TimeZone = 97
	TimeZoneWinkhoek                 TimeZone = 98
	TimeZoneLagos                    TimeZone = 99
	TimeZoneRiyahd                   TimeZone = 100
	TimeZoneVenezuela                TimeZone = 101
	TimeZoneAustraliaLh              TimeZone = 102
	TimeZoneSantiago                 TimeZone = 103
	TimeZoneManual                   TimeZone = 253
	TimeZoneAutomatic                TimeZone = 254
	TimeZoneInvalid                  TimeZone = 0xFF
)

// TimerTrigger represents the timer_trigger FIT type.
type TimerTrigger byte

const (
	TimerTriggerManual           TimerTrigger = 0
	TimerTriggerAuto             TimerTrigger = 1
	TimerTriggerFitnessEquipment TimerTrigger = 2
	TimerTriggerInvalid          TimerTrigger = 0xFF
)

// UserLocalId represents the user_local_id FIT type.
type UserLocalId uint16

const (
	UserLocalIdLocalMin      UserLocalId = 0x0000
	UserLocalIdLocalMax      UserLocalId = 0x000F
	UserLocalIdStationaryMin UserLocalId = 0x0010
	UserLocalIdStationaryMax UserLocalId = 0x00FF
	UserLocalIdPortableMin   UserLocalId = 0x0100
	UserLocalIdPortableMax   UserLocalId = 0xFFFE
	UserLocalIdInvalid       UserLocalId = 0xFFFF
)

// Weight represents the weight FIT type.
type Weight uint16

const (
	WeightCalculating Weight = 0xFFFE
	WeightInvalid     Weight = 0xFFFF
)

// WktStepDuration represents the wkt_step_duration FIT type.
type WktStepDuration byte

const (
	WktStepDurationTime                        WktStepDuration = 0
	WktStepDurationDistance                    WktStepDuration = 1
	WktStepDurationHrLessThan                  WktStepDuration = 2
	WktStepDurationHrGreaterThan               WktStepDuration = 3
	WktStepDurationCalories                    WktStepDuration = 4
	WktStepDurationOpen                        WktStepDuration = 5
	WktStepDurationRepeatUntilStepsCmplt       WktStepDuration = 6
	WktStepDurationRepeatUntilTime             WktStepDuration = 7
	WktStepDurationRepeatUntilDistance         WktStepDuration = 8
	WktStepDurationRepeatUntilCalories         WktStepDuration = 9
	WktStepDurationRepeatUntilHrLessThan       WktStepDuration = 10
	WktStepDurationRepeatUntilHrGreaterThan    WktStepDuration = 11
	WktStepDurationRepeatUntilPowerLessThan    WktStepDuration = 12
	WktStepDurationRepeatUntilPowerGreaterThan WktStepDuration = 13
	WktStepDurationPowerLessThan               WktStepDuration = 14
	WktStepDurationPowerGreaterThan            WktStepDuration = 15
	WktStepDurationRepetitionTime              WktStepDuration = 28
	WktStepDurationInvalid                     WktStepDuration = 0xFF
)

// WktStepTarget represents the wkt_step_target FIT type.
type WktStepTarget byte

const (
	WktStepTargetSpeed      WktStepTarget = 0
	WktStepTargetHeartRate  WktStepTarget = 1
	WktStepTargetOpen       WktStepTarget = 2
	WktStepTargetCadence    WktStepTarget = 3
	WktStepTargetPower      WktStepTarget = 4
	WktStepTargetGrade      WktStepTarget = 5
	WktStepTargetResistance WktStepTarget = 6
	WktStepTargetInvalid    WktStepTarget = 0xFF
)

// WorkoutCapabilities represents the workout_capabilities FIT type.
type WorkoutCapabilities uint32

const (
	WorkoutCapabilitiesInterval         WorkoutCapabilities = 0x00000001
	WorkoutCapabilitiesCustom           WorkoutCapabilities = 0x00000002
	WorkoutCapabilitiesFitnessEquipment WorkoutCapabilities = 0x00000004
	WorkoutCapabilitiesFirstbeat        WorkoutCapabilities = 0x00000008
	WorkoutCapabilitiesNewLeaf          WorkoutCapabilities = 0x00000010
	WorkoutCapabilitiesTcx              WorkoutCapabilities = 0x00000020 // For backwards compatibility.  Watch should add missing id fields then clear flag.
	WorkoutCapabilitiesSpeed            WorkoutCapabilities = 0x00000080 // Speed source required for workout step.
	WorkoutCapabilitiesHeartRate        WorkoutCapabilities = 0x00000100 // Heart rate source required for workout step.
	WorkoutCapabilitiesDistance         WorkoutCapabilities = 0x00000200 // Distance source required for workout step.
	WorkoutCapabilitiesCadence          WorkoutCapabilities = 0x00000400 // Cadence source required for workout step.
	WorkoutCapabilitiesPower            WorkoutCapabilities = 0x00000800 // Power source required for workout step.
	WorkoutCapabilitiesGrade            WorkoutCapabilities = 0x00001000 // Grade source required for workout step.
	WorkoutCapabilitiesResistance       WorkoutCapabilities = 0x00002000 // Resistance source required for workout step.
	WorkoutCapabilitiesProtected        WorkoutCapabilities = 0x00004000
	WorkoutCapabilitiesInvalid          WorkoutCapabilities = 0x00000000
)

// WorkoutHr represents the workout_hr FIT type.
type WorkoutHr uint32

const (
	WorkoutHrBpmOffset WorkoutHr = 100
	WorkoutHrInvalid   WorkoutHr = 0xFFFFFFFF
)

// WorkoutPower represents the workout_power FIT type.
type WorkoutPower uint32

const (
	WorkoutPowerWattsOffset WorkoutPower = 1000
	WorkoutPowerInvalid     WorkoutPower = 0xFFFFFFFF
)
