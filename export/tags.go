package export

// WorkoutMetadataKey Metadata keys
type WorkoutMetadataKey string

// WorkoutMetadataKeys slice of multiple WorkoutMetadataKey items
type WorkoutMetadataKeys []WorkoutMetadataKey

const (
	// AverageMETs example <MetadataEntry key="HKAverageMETs" value="5.11259 kcal/hr·kg"/>
	AverageMETs WorkoutMetadataKey = "HKAverageMETs"
	// ElevationAscended <MetadataEntry key="HKElevationAscended" value="67254 cm"/>
	ElevationAscended WorkoutMetadataKey = "HKElevationAscended"
	// IndoorWorkout <MetadataEntry key="HKIndoorWorkout" value="1"/> or "0"
	IndoorWorkout WorkoutMetadataKey = "HKIndoorWorkout"
	//  TimeZone <MetadataEntry key="HKTimeZone" value="America/Los_Angeles"/>
	TimeZone WorkoutMetadataKey = "HKTimeZone"
	// Humidity  <MetadataEntry key="HKWeatherHumidity" value="10000 %"/>
	Humidity WorkoutMetadataKey = "HKWeatherHumidity"
	// Temperature <MetadataEntry key="HKWeatherTemperature" value="54 degF"/>
	Temperature WorkoutMetadataKey = "HKWeatherTemperature"
)

// WorkoutMetadataKeysKnown list of data seen in keys until now
var WorkoutMetadataKeysKnown = WorkoutMetadataKeys{AverageMETs, ElevationAscended, IndoorWorkout,
	TimeZone, Humidity, Temperature,
}

// WorkOutActivityType all activity types seen
type WorkOutActivityType string

// WorkOutActivityTypes
type WorkOutActivityTypes []WorkOutActivityType

const (
	// Elliptical
	Elliptical = "HKWorkoutActivityTypeElliptical"
	// Hiking
	Hiking = "HKWorkoutActivityTypeHiking"
	// Other
	Other = "HKWorkoutActivityTypeOther"
	// Walking
	Walking = "HKWorkoutActivityTypeWalking"
)

// WorkOutActivityTypesKnown analyzed so far
var WorkOutActivityTypesKnown = WorkOutActivityTypes{Elliptical, Hiking, Other, Walking}

// RecordMetadataKey
type RecordMetadataKey string
type RecordMetadataKeys []RecordMetadataKey

const (
	// Asleep example <MetadataEntry key="Asleep" value="20580"/>
	Asleep RecordMetadataKey = "Asleep"
	// AverageHR  record example <MetadataEntry key="Average HR" value="50.51"/>
	AverageHR RecordMetadataKey = "Average HR"
	// DaytimeHR record example <MetadataEntry key="Daytime HR" value="60.15"/>
	DaytimeHR RecordMetadataKey = "Daytime HR"
	// DeepSleep record example <MetadataEntry key="Deep Sleep" value="8053"/>
	DeepSleep RecordMetadataKey = "Deep Sleep"
	// DevicePlacementSide, example <Record type="HKQuantityTypeIdentifierWalkingAsymmetryPercentage" ..>
	//    <MetadataEntry key="HKMetadataKeyDevicePlacementSide" value="2"/>
	DevicePlacementSide RecordMetadataKey = "HKMetadataKeyDevicePlacementSide"
	// EditSlots record example  <MetadataEntry key="Edit Slots" value="0415,0630,0500,0445,0530,0615,0430,0515,0545"/>
	EditSlots RecordMetadataKey = "Edit Slots"
	// EnergyThreshold record example <MetadataEntry key="Energy Threshold" value="1200"/>
	EnergyThreshold RecordMetadataKey = "Energy Threshold"
	// Calibrated example <MetadataEntry key="HKMetadataKeyAppleDeviceCalibrated" value="1"/>
	Calibrated RecordMetadataKey = "HKMetadataKeyAppleDeviceCalibrated"
	// MotionContext example <MetadataEntry key="HKMetadataKeyHeartRateMotionContext" value="1"/>
	MotionContext RecordMetadataKey = "HKMetadataKeyHeartRateMotionContext"
	// SyncID example <MetadataEntry key="HKMetadataKeySyncIdentifier" value="3:32FAC695-B2D3-4358-9154-CA13801043FA:613624195.01513:613624791.01513:119"/>
	SyncID RecordMetadataKey = "HKMetadataKeySyncIdentifier"
	// SyncVersion example <MetadataEntry key="HKMetadataKeySyncVersion" value="1"/>
	SyncVersion RecordMetadataKey = "HKMetadataKeySyncVersion"
	// RecordTimeZone example  <MetadataEntry key="HKTimeZone" value="America/Los_Angeles"/>
	RecordTimeZone RecordMetadataKey = "HKTimeZone"
	// HKVO2MaxTestType  <MetadataEntry key="HKVO2MaxTestType" value="2"/>
	HKVO2MaxTestType RecordMetadataKey = "HKVO2MaxTestType"
	// Lights example <MetadataEntry key="Lights" value="NO"/>
	Lights RecordMetadataKey = "Lights"
	// Meal <MetadataEntry key="Meal" value="Lunch"/> "Dinner" "Snacks" or "Breakfast"
	// meal is also seen in some records. Ignore case. <MetadataEntry key="meal" value="Lunch"/>
	Meal          RecordMetadataKey = "Meal"
	MealLowercase RecordMetadataKey = "meal"
	// Nap <MetadataEntry key="Nap" value="YES"/>
	Nap RecordMetadataKey = "Nap"
	// Rating examples <MetadataEntry key="Rating" value="7.62"/> or <MetadataEntry key="Rating" value="83.25"/>
	Rating RecordMetadataKey = "Rating"
	// Recharge example <MetadataEntry key="Recharge" value="9"/> or <MetadataEntry key="Recharge" value="100"/>
	Recharge RecordMetadataKey = "Recharge"
	// Tags example <MetadataEntry key="Tags" value="1"/>
	Tags RecordMetadataKey = "Tags"
)

// RecordMetadataKeysKnown so far in data set
var RecordMetadataKeysKnown = RecordMetadataKeys{Asleep, AverageHR, DaytimeHR, DeepSleep, DevicePlacementSide,
	EditSlots, EnergyThreshold, Calibrated, MotionContext, SyncID, SyncVersion, RecordTimeZone, HKVO2MaxTestType, Lights,
	Meal, MealLowercase, Nap, Rating, Recharge, Tags,
}

// RecordType all Record types seen so far
type RecordType string
type RecordTypes []RecordType // Array of RecordTypes
const (
	// AppleStandHour eg: <Record type="HKCategoryTypeIdentifierAppleStandHour" sourceName="🕛🕐🕑🕒" sourceVersion="5.0.1" device="&lt;&lt;HKDevice: 0x2834ef0c0&gt;, name:Apple Watch, manufacturer:Apple, model:Watch, hardware:Watch4,2, software:5.0.1&gt;" creationDate="2018-10-23 08:10:57 -0800" startDate="2018-10-23 08:00:00 -0800" endDate="2018-10-23 09:00:00 -0800" value="HKCategoryValueAppleStandHourStood"/>
	AppleStandHour RecordType = "HKCategoryTypeIdentifierAppleStandHour"
	// HandwashingEvent eg: <Record type="HKCategoryTypeIdentifierHandwashingEvent" sourceName="🕛🕐🕑🕒" sourceVersion="5.3" creationDate="2020-09-20 07:24:43 -0800" startDate="2020-09-20 07:24:08 -0800" endDate="2020-09-20 07:24:43 -0800"/>
	HandwashingEvent RecordType = "HKCategoryTypeIdentifierHandwashingEvent"
	// MindfulSession eg: <Record type="HKCategoryTypeIdentifierMindfulSession" sourceName="Headspace" sourceVersion="10878" creationDate="2018-05-21 19:32:46 -0800" startDate="2018-05-21 19:21:15 -0800" endDate="2018-05-21 19:32:46 -0800"/>
	MindfulSession RecordType = "HKCategoryTypeIdentifierMindfulSession"
	// SleepAnalysis eg: <Record type="HKCategoryTypeIdentifierSleepAnalysis" sourceName="Clock" sourceVersion="50" device="&lt;&lt;HKDevice: 0x283704280&gt;, name:iPhone, manufacturer:Apple, model:iPhone, hardware:iPhone9,4, software:11.0&gt;" creationDate="2017-09-23 06:01:10 -0800" startDate="2017-09-22 23:49:16 -0800" endDate="2017-09-23 06:01:10 -0800" value="HKCategoryValueSleepAnalysisInBed">
	SleepAnalysis RecordType = "HKCategoryTypeIdentifierSleepAnalysis"
	// ActiveEnergyBurned eg: <Record type="HKQuantityTypeIdentifierActiveEnergyBurned" sourceName="🕛🕐🕑🕒" sourceVersion="7.0.1" device="&lt;&lt;HKDevice: 0x2834eed50&gt;, name:Apple Watch, manufacturer:Apple Inc., model:Watch, hardware:Watch4,2, software:7.0.1&gt;" unit="Cal" creationDate="2020-10-05 12:30:30 -0800" startDate="2020-10-05 12:29:55 -0800" endDate="2020-10-05 12:30:05 -0800" value="0.041"/>
	ActiveEnergyBurned RecordType = "HKQuantityTypeIdentifierActiveEnergyBurned"
	// AppleExerciseTime eg: <Record type="HKQuantityTypeIdentifierAppleExerciseTime" sourceName="🕛🕐🕑🕒" sourceVersion="5.0.1" device="&lt;&lt;HKDevice: 0x2837100a0&gt;, name:Apple Watch, manufacturer:Apple, model:Watch, hardware:Watch4,2, software:5.0.1&gt;" unit="min" creationDate="2018-10-23 10:55:55 -0800" startDate="2018-10-23 10:54:40 -0800" endDate="2018-10-23 10:55:40 -0800" value="1"/>
	AppleExerciseTime RecordType = "HKQuantityTypeIdentifierAppleExerciseTime"
	// AppleStandTime eg: <Record type="HKQuantityTypeIdentifierAppleStandTime" sourceName="🕛🕐🕑🕒" unit="min" creationDate="2019-09-19 17:10:11 -0800" startDate="2019-09-19 17:05:00 -0800" endDate="2019-09-19 17:10:00 -0800" value="2"/>
	AppleStandTime RecordType = "HKQuantityTypeIdentifierAppleStandTime"
	// BasalEnergyBurned eg: <Record type="HKQuantityTypeIdentifierBasalEnergyBurned" sourceName="🕛🕐🕑🕒" sourceVersion="5.0.1" device="&lt;&lt;HKDevice: 0x2837100a0&gt;, name:Apple Watch, manufacturer:Apple, model:Watch, hardware:Watch4,2, software:5.0.1&gt;" unit="Cal" creationDate="2018-10-22 17:40:20 -0800" startDate="2018-10-22 17:24:09 -0800" endDate="2018-10-22 17:38:00 -0800" value="20.157"/>
	BasalEnergyBurned RecordType = "HKQuantityTypeIdentifierBasalEnergyBurned"
	// BodyFatPercentage eg:  <Record type="HKQuantityTypeIdentifierBodyFatPercentage" sourceName="Renpho" sourceVersion="4" unit="%" creationDate="2020-04-19 06:53:44 -0800" startDate="2020-04-15 05:58:55 -0800" endDate="2020-04-15 05:58:55 -0800" value="0.161"/>
	BodyFatPercentage RecordType = "HKQuantityTypeIdentifierBodyFatPercentage"
	// BodyMass eg: <Record type="HKQuantityTypeIdentifierBodyMass" sourceName="🔜" sourceVersion="14.2" unit="lb" creationDate="2020-10-05 12:02:41 -0800" startDate="2020-10-05 12:02:41 -0800" endDate="2020-10-05 12:02:41 -0800" value="140"/>
	BodyMass RecordType = "HKQuantityTypeIdentifierBodyMass"
	// BodyMassIndex eg: <Record type="HKQuantityTypeIdentifierBodyMassIndex" sourceName="Renpho" sourceVersion="4" unit="count" creationDate="2020-04-19 06:53:45 -0800" startDate="2020-04-14 14:48:12 -0800" endDate="2020-04-14 14:48:12 -0800" value="22.2"/>
	BodyMassIndex RecordType = "HKQuantityTypeIdentifierBodyMassIndex"
	// DietaryCalcium eg: <Record type="HKQuantityTypeIdentifierDietaryCalcium" sourceName="MyFitnessPal" sourceVersion="22314" unit="mg" creationDate="2018-12-28 22:28:18 -0800" startDate="2018-12-28 22:27:00 -0800" endDate="2018-12-28 22:27:00 -0800" value="75.6628">
	DietaryCalcium RecordType = "HKQuantityTypeIdentifierDietaryCalcium"
	// DietaryCarbohydrates eg:  <Record type="HKQuantityTypeIdentifierDietaryCarbohydrates" sourceName="MyFitnessPal" sourceVersion="22314" unit="g" creationDate="2018-12-28 20:53:07 -0800" startDate="2018-12-28 17:57:00 -0800" endDate="2018-12-28 17:57:00 -0800" value="39.94">
	DietaryCarbohydrates RecordType = "HKQuantityTypeIdentifierDietaryCarbohydrates"
	// DietaryCholesterol eg:  <Record type="HKQuantityTypeIdentifierDietaryCholesterol" sourceName="MyFitnessPal" sourceVersion="22314" unit="mg" creationDate="2018-12-28 22:28:18 -0800" startDate="2018-12-28 22:27:00 -0800" endDate="2018-12-28 22:27:00 -0800" value="262.157">
	DietaryCholesterol RecordType = "HKQuantityTypeIdentifierDietaryCholesterol"
	// DietaryEnergyConsumed eg: <Record type="HKQuantityTypeIdentifierDietaryEnergyConsumed" sourceName="MyFitnessPal" sourceVersion="22314" unit="Cal" creationDate="2018-12-28 15:44:53 -0800" startDate="2018-12-28 15:44:00 -0800" endDate="2018-12-28 15:44:00 -0800" value="741">
	DietaryEnergyConsumed RecordType = "HKQuantityTypeIdentifierDietaryEnergyConsumed"
	// DietaryFatMonounsaturated eg:  <Record type="HKQuantityTypeIdentifierDietaryFatMonounsaturated" sourceName="MyFitnessPal" sourceVersion="22314" unit="g" creationDate="2018-12-28 22:28:18 -0800" startDate="2018-12-28 22:27:00 -0800" endDate="2018-12-28 22:27:00 -0800" value="7.979">
	DietaryFatMonounsaturated RecordType = "HKQuantityTypeIdentifierDietaryFatMonounsaturated"
	// DietaryFatPolyunsaturated eg: <Record type="HKQuantityTypeIdentifierDietaryFatPolyunsaturated" sourceName="MyFitnessPal" sourceVersion="22314" unit="g" creationDate="2018-12-28 22:28:18 -0800" startDate="2018-12-28 22:27:00 -0800" endDate="2018-12-28 22:27:00 -0800" value="2.0911">
	DietaryFatPolyunsaturated RecordType = "HKQuantityTypeIdentifierDietaryFatPolyunsaturated"
	// DietaryFatSaturated eg: <Record type="HKQuantityTypeIdentifierDietaryFatSaturated" sourceName="MyFitnessPal" sourceVersion="22314" unit="g" creationDate="2018-12-28 22:28:18 -0800" startDate="2018-12-28 22:27:00 -0800" endDate="2018-12-28 22:27:00 -0800" value="11.0707">
	DietaryFatSaturated RecordType = "HKQuantityTypeIdentifierDietaryFatSaturated"
	// DietaryFatTotal eg: <Record type="HKQuantityTypeIdentifierDietaryFatTotal" sourceName="MyFitnessPal" sourceVersion="22314" unit="g" creationDate="2018-12-28 22:28:18 -0800" startDate="2018-12-28 22:27:00 -0800" endDate="2018-12-28 22:27:00 -0800" value="39.0836">
	DietaryFatTotal RecordType = "HKQuantityTypeIdentifierDietaryFatTotal"
	// DietaryFiber eg: <Record type="HKQuantityTypeIdentifierDietaryFiber" sourceName="MyFitnessPal" sourceVersion="22314" unit="g" creationDate="2018-12-28 15:44:54 -0800" startDate="2018-12-28 15:44:00 -0800" endDate="2018-12-28 15:44:00 -0800" value="0">
	DietaryFiber RecordType = "HKQuantityTypeIdentifierDietaryFiber"
	// DietaryIron eg:  <Record type="HKQuantityTypeIdentifierDietaryIron" sourceName="MyFitnessPal" sourceVersion="22314" unit="mg" creationDate="2018-12-28 22:28:18 -0800" startDate="2018-12-28 22:27:00 -0800" endDate="2018-12-28 22:27:00 -0800" value="4.1756">
	DietaryIron RecordType = "HKQuantityTypeIdentifierDietaryIron"
	// DietaryPotassium eg: <Record type="HKQuantityTypeIdentifierDietaryPotassium" sourceName="MyFitnessPal" sourceVersion="22314" unit="mg" creationDate="2018-12-28 22:28:18 -0800" startDate="2018-12-28 22:27:00 -0800" endDate="2018-12-28 22:27:00 -0800" value="915.111">
	DietaryPotassium RecordType = "HKQuantityTypeIdentifierDietaryPotassium"
	// DietaryProtein eg: <Record type="HKQuantityTypeIdentifierDietaryProtein" sourceName="MyFitnessPal" sourceVersion="22314" unit="g" creationDate="2018-12-28 22:28:18 -0800" startDate="2018-12-28 22:27:00 -0800" endDate="2018-12-28 22:27:00 -0800" value="81.2453">
	DietaryProtein RecordType = "HKQuantityTypeIdentifierDietaryProtein"
	// DietarySodium eg: <Record type="HKQuantityTypeIdentifierDietarySodium" sourceName="MyFitnessPal" sourceVersion="22314" unit="mg" creationDate="2018-12-28 22:28:18 -0800" startDate="2018-12-28 22:27:00 -0800" endDate="2018-12-28 22:27:00 -0800" value="941.965">
	DietarySodium RecordType = "HKQuantityTypeIdentifierDietarySodium"
	// DietarySugar eg: <Record type="HKQuantityTypeIdentifierDietarySugar" sourceName="MyFitnessPal" sourceVersion="22314" unit="g" creationDate="2018-12-28 20:53:08 -0800" startDate="2018-12-28 17:57:00 -0800" endDate="2018-12-28 17:57:00 -0800" value="25.54">
	DietarySugar RecordType = "HKQuantityTypeIdentifierDietarySugar"
	// DietaryVitaminC eg: <Record type="HKQuantityTypeIdentifierDietaryVitaminC" sourceName="MyFitnessPal" sourceVersion="22314" unit="mg" creationDate="2018-12-28 22:28:18 -0800" startDate="2018-12-28 22:27:00 -0800" endDate="2018-12-28 22:27:00 -0800" value="22.7521">
	DietaryVitaminC RecordType = "HKQuantityTypeIdentifierDietaryVitaminC"
	// DistanceWalkingRunning eg: <Record type="HKQuantityTypeIdentifierDistanceWalkingRunning" sourceName="🕛🕐🕑🕒" sourceVersion="7.1" device="&lt;&lt;HKDevice: 0x283704460&gt;, name:Apple Watch, manufacturer:Apple Inc., model:Watch, hardware:Watch4,2, software:7.1&gt;" unit="mi" creationDate="2020-12-01 15:42:18 -0800" startDate="2020-12-01 15:42:12 -0800" endDate="2020-12-01 15:42:14 -0800" value="0.00153413"/>
	DistanceWalkingRunning RecordType = "HKQuantityTypeIdentifierDistanceWalkingRunning"
	// FlightsClimbed eg:   <Record type="HKQuantityTypeIdentifierFlightsClimbed" sourceName="🔜" sourceVersion="10.1.1" device="&lt;&lt;HKDevice: 0x2834f21c0&gt;, name:iPhone, manufacturer:Apple, model:iPhone, hardware:iPhone9,4, software:10.1.1&gt;" unit="count" creationDate="2016-12-10 15:57:14 -0800" startDate="2016-12-10 15:08:33 -0800" endDate="2016-12-10 15:08:33 -0800" value="1"/>
	FlightsClimbed RecordType = "HKQuantityTypeIdentifierFlightsClimbed"
	// EnvironmentalAudioExposure eg:  <Record type="HKQuantityTypeIdentifierEnvironmentalAudioExposure" sourceName="🕛🕐🕑🕒" sourceVersion="6.1" device="&lt;&lt;HKDevice: 0x2837080a0&gt;, name:Apple Watch, manufacturer:Apple Inc., model:Watch, hardware:Watch4,2, software:6.1&gt;" unit="dBASPL" creationDate="2019-10-31 06:09:24 -0800" startDate="2019-10-31 06:01:21 -0800" endDate="2019-10-31 06:01:41 -0800" value="62.1341"/>
	EnvironmentalAudioExposure RecordType = "HKQuantityTypeIdentifierEnvironmentalAudioExposure"
	// HeadphoneAudioExposure eg: <Record type="HKQuantityTypeIdentifierHeadphoneAudioExposure" sourceName="🔜" sourceVersion="13.0" device="&lt;&lt;HKDevice: 0x283737ed0&gt;, name:AirPods, manufacturer:Apple Inc., model:0x2002, localIdentifier:60:F4:45:EC:48:62-tacl&gt;" unit="dBASPL" creationDate="2019-09-24 11:13:28 -0800" startDate="2019-09-24 11:09:58 -0800" endDate="2019-09-24 11:10:07 -0800" value="71.7391"/>
	HeadphoneAudioExposure RecordType = "HKQuantityTypeIdentifierHeadphoneAudioExposure"
	// HeartRate eg:  <Record type="HKQuantityTypeIdentifierHeartRate" sourceName="🕛🕐🕑🕒" sourceVersion="5.0.1" device="&lt;&lt;HKDevice: 0x283735e00&gt;, name:Apple Watch, manufacturer:Apple, model:Watch, hardware:Watch4,2, software:5.0.1&gt;" unit="count/min" creationDate="2018-10-22 17:20:16 -0800" startDate="2018-10-22 17:16:02 -0800" endDate="2018-10-22 17:16:02 -0800" value="60">
	HeartRate RecordType = "HKQuantityTypeIdentifierHeartRate"
	// HeartRateVariabilitySDNN eg: <Record type="HKQuantityTypeIdentifierHeartRateVariabilitySDNN" sourceName="🕛🕐🕑🕒" sourceVersion="5.0.1" device="&lt;&lt;HKDevice: 0x28347bac0&gt;, name:Apple Watch, manufacturer:Apple, model:Watch, hardware:Watch4,2, software:5.0.1&gt;" unit="ms" creationDate="2018-10-24 10:16:25 -0800" startDate="2018-10-24 10:15:20 -0800" endDate="2018-10-24 10:16:25 -0800" value="67.1854">
	HeartRateVariabilitySDNN RecordType = "HKQuantityTypeIdentifierHeartRateVariabilitySDNN"
	// Height eg: <Record type="HKQuantityTypeIdentifierHeight" sourceName="🔜" sourceVersion="10.1.1" unit="ft" creationDate="2016-12-08 20:06:35 -0800" startDate="2016-12-08 20:06:35 -0800" endDate="2016-12-08 20:06:35 -0800" value="5.5"/>
	Height RecordType = "HKQuantityTypeIdentifierHeight"
	// LeanBodyMass eg: <Record type="HKQuantityTypeIdentifierLeanBodyMass" sourceName="Renpho" sourceVersion="4" unit="lb" creationDate="2020-04-19 06:53:44 -0800" startDate="2020-04-15 05:58:55 -0800" endDate="2020-04-15 05:58:55 -0800" value="115.8"/>
	LeanBodyMass RecordType = "HKQuantityTypeIdentifierLeanBodyMass"
	// RestingHeartRate eg:
	//   <Record type="HKQuantityTypeIdentifierRestingHeartRate" sourceName="🕛🕐🕑🕒" sourceVersion="5.0.1" unit="count/min" creationDate="2018-10-27 20:40:57 -0800" startDate="2018-10-27 18:28:47 -0800" endDate="2018-10-27 20:37:20 -0800" value="59"/>
	RestingHeartRate RecordType = "HKQuantityTypeIdentifierRestingHeartRate"
	// SixMinuteWalkTestDistance eg:
	//   <Record type="HKQuantityTypeIdentifierSixMinuteWalkTestDistance" sourceName="🔜" sourceVersion="2420.8.11" unit="m" creationDate="2020-11-12 18:31:58 -0800" startDate="2020-11-12 18:31:49 -0800" endDate="2020-11-12 18:31:49 -0800" value="500">
	SixMinuteWalkTestDistance RecordType = "HKQuantityTypeIdentifierSixMinuteWalkTestDistance"
	// StepCount eg: <Record type="HKQuantityTypeIdentifierStepCount" sourceName="🔜" sourceVersion="10.1.1" device="&lt;&lt;HKDevice: 0x283714aa0&gt;, name:iPhone, manufacturer:Apple, model:iPhone, hardware:iPhone9,4, software:10.1.1&gt;" unit="count" creationDate="2016-12-09 09:52:21 -0800" startDate="2016-12-09 09:37:52 -0800" endDate="2016-12-09 09:43:12 -0800" value="155"/>
	StepCount RecordType = "HKQuantityTypeIdentifierStepCount"
	// VO2Max eg:
	//   <Record type="HKQuantityTypeIdentifierVO2Max" sourceName="🕛🕐🕑🕒" unit="mL/min·kg" creationDate="2018-12-19 17:49:50 -0800" startDate="2018-12-19 17:49:50 -0800" endDate="2018-12-19 17:49:50 -0800" value="32.7204">
	VO2Max RecordType = "HKQuantityTypeIdentifierVO2Max"
	// WalkingAsymmetryPercentage A quantity sample measuring the percentage of steps in which one foot moves at a
	// different speed than the other when walking steadily over flat ground. eg:
	//     <Record type="HKQuantityTypeIdentifierWalkingAsymmetryPercentage" sourceName="Dharti Shah’s iPhone" sourceVersion="14.0" device="&lt;&lt;HKDevice: 0x281326260&gt;, name:iPhone, manufacturer:Apple Inc., model:iPhone, hardware:iPhone12,3, software:14.0&gt;" unit="%" creationDate="2020-09-16 17:56:41 -0800" startDate="2020-09-16 17:52:45 -0800" endDate="2020-09-16 17:53:09 -0800" value="0.07">
	//        <MetadataEntry key="HKMetadataKeyDevicePlacementSide" value="2"/>
	//    </Record>
	WalkingAsymmetryPercentage RecordType = "HKQuantityTypeIdentifierWalkingAsymmetryPercentage"
	// WalkingDoubleSupportPercentage A quantity sample that measures the percentage of time when both of the user’s feet
	// are touching the ground while walking steadily over flat ground. eg:
	//      <Record type="HKQuantityTypeIdentifierWalkingDoubleSupportPercentage" sourceName="Dharti Shah’s iPhone" sourceVersion="14.0" device="&lt;&lt;HKDevice: 0x28134ec60&gt;, name:iPhone, manufacturer:Apple Inc., model:iPhone, hardware:iPhone12,3, software:14.0&gt;" unit="%" creationDate="2020-09-16 18:35:24 -0800" startDate="2020-09-16 18:33:42 -0800" endDate="2020-09-16 18:34:15 -0800" value="0.294"/>
	WalkingDoubleSupportPercentage RecordType = "HKQuantityTypeIdentifierWalkingDoubleSupportPercentage"
	// WalkingHeartRateAverage eg:
	//   <Record type="HKQuantityTypeIdentifierWalkingHeartRateAverage" sourceName="🕛🕐🕑🕒" sourceVersion="5.0.1" unit="count/min" creationDate="2018-10-28 15:27:01 -0800" startDate="2018-10-27 23:03:10 -0800" endDate="2018-10-28 15:26:56 -0800" value="97.5"/>
	WalkingHeartRateAverage RecordType = "HKQuantityTypeIdentifierWalkingHeartRateAverage"
	// WalkingSpeed eg:  2.52774 mi/hr
	//   <Record type="HKQuantityTypeIdentifierWalkingSpeed" sourceName="Dharti Shah’s iPhone" sourceVersion="14.2" device="&lt;&lt;HKDevice: 0x28135db30&gt;, name:iPhone, manufacturer:Apple Inc., model:iPhone, hardware:iPhone12,3, software:14.2&gt;" unit="mi/hr" creationDate="2021-01-02 11:07:49 -0800" startDate="2021-01-02 11:03:45 -0800" endDate="2021-01-02 11:03:58 -0800" value="2.52774"/>
	WalkingSpeed RecordType = "HKQuantityTypeIdentifierWalkingSpeed"
	// WalkingStepLength eg: 17.3228 in step
	//   <Record type="HKQuantityTypeIdentifierWalkingStepLength" sourceName="Dharti Shah’s iPhone" sourceVersion="14.0" device="&lt;&lt;HKDevice: 0x281307c50&gt;, name:iPhone, manufacturer:Apple Inc., model:iPhone, hardware:iPhone12,3, software:14.0&gt;" unit="in" creationDate="2020-09-18 08:16:42 -0800" startDate="2020-09-18 07:49:05 -0800" endDate="2020-09-18 07:49:10 -0800" value="17.3228"/>
	WalkingStepLength RecordType = "HKQuantityTypeIdentifierWalkingStepLength"
)

// RecordTypesKnown so far in data. As you see more types add them here.
var RecordTypesKnown = RecordTypes{AppleStandHour, HandwashingEvent, MindfulSession, SleepAnalysis,
	ActiveEnergyBurned, AppleExerciseTime, AppleStandTime, BasalEnergyBurned, BodyFatPercentage, BodyMass,
	BodyMassIndex, DietaryCalcium, DietaryCarbohydrates, DietaryCholesterol, DietaryEnergyConsumed,
	DietaryFatMonounsaturated, DietaryFatPolyunsaturated, DietaryFatSaturated, DietaryFatTotal, DietaryFiber,
	DietaryIron, DietaryPotassium, DietaryProtein, DietarySodium, DietarySugar, DietaryVitaminC, DistanceWalkingRunning,
	EnvironmentalAudioExposure, FlightsClimbed, HeadphoneAudioExposure, HeartRate, HeartRateVariabilitySDNN,
	Height, LeanBodyMass, RestingHeartRate, SixMinuteWalkTestDistance, StepCount, VO2Max,
	WalkingAsymmetryPercentage, WalkingDoubleSupportPercentage, WalkingHeartRateAverage, WalkingSpeed, WalkingStepLength,
}

// RecordSource defines enum space of known "sourceName" in a Record
type RecordSource string

// RecordSources slice of "sourceName" values known so far
type RecordSources []RecordSource

const (
	// AutoSleep eg: <Record type="HKCategoryTypeIdentifierSleepAnalysis" sourceName="AutoSleep" sourceVersion="6.4.0" creationDate="2020-08-17 18:29:28 -0800" startDate="2020-08-16 22:00:00 -0800" endDate="2020-08-17 04:22:00 -0800" value="HKCategoryValueSleepAnalysisAsleep"/>
	AutoSleep RecordSource = "AutoSleep"
	// Clock eg:  <Record type="HKCategoryTypeIdentifierSleepAnalysis" sourceName="Clock" sourceVersion="50" device="&lt;&lt;HKDevice: 0x283704280&gt;, name:iPhone, manufacturer:Apple, model:iPhone, hardware:iPhone9,4, software:11.0&gt;" creationDate="2017-09-24 06:01:00 -0800" startDate="2017-09-23 23:24:48 -0800" endDate="2017-09-24 06:00:52 -0800" value="HKCategoryValueSleepAnalysisInBed">
	Clock RecordSource = "Clock"
	// Headspace eg:  <Record type="HKCategoryTypeIdentifierMindfulSession" sourceName="Headspace" sourceVersion="10878" creationDate="2018-05-21 19:32:46 -0800" startDate="2018-05-21 19:21:15 -0800" endDate="2018-05-21 19:32:46 -0800"/>
	Headspace RecordSource = "Headspace"
	// MyFitnessPal eg: <Record type="HKQuantityTypeIdentifierBodyMass" sourceName="MyFitnessPal" sourceVersion="22314" unit="lb" creationDate="2019-01-21 20:50:22 -0800" startDate="2019-01-04 20:50:00 -0800" endDate="2019-01-04 20:50:00 -0800" value="185.078"/>
	MyFitnessPal RecordSource = "MyFitnessPal"
	// Pacer eg:  <Record type="HKQuantityTypeIdentifierActiveEnergyBurned" sourceName="Pacer" sourceVersion="3.9.1.307" unit="Cal" creationDate="2019-02-07 13:08:13 -0800" startDate="2019-02-07 12:15:00 -0800" endDate="2019-02-07 12:30:00 -0800" value="12.0754"/>
	Pacer RecordSource = "Pacer"
	// Renpho eg: <Record type="HKQuantityTypeIdentifierActiveEnergyBurned" sourceName="Pacer" sourceVersion="3.9.1.307" unit="Cal" creationDate="2019-02-07 13:08:13 -0800" startDate="2019-02-07 12:15:00 -0800" endDate="2019-02-07 12:30:00 -0800" value="12.0754"/>
	Renpho RecordSource = "Renpho"
	// SleepPP eg:  <Record type="HKCategoryTypeIdentifierSleepAnalysis" sourceName="Sleep++" sourceVersion="58" creationDate="2019-05-27 13:41:31 -0800" startDate="2019-05-18 22:03:49 -0800" endDate="2019-05-19 02:47:07 -0800" value="HKCategoryValueSleepAnalysisAsleep">
	SleepPP RecordSource = "Sleep++"
	// SleepWatch eg: <Record type="HKCategoryTypeIdentifierSleepAnalysis" sourceName="SleepWatch" sourceVersion="5.6.0.0" creationDate="2019-03-19 21:09:51 -0800" startDate="2019-03-12 21:55:00 -0800" endDate="2019-03-13 02:10:00 -0800" value="HKCategoryValueSleepAnalysisInBed"/>
	SleepWatch RecordSource = "SleepWatch"
	// IPhone eg: <Record type="HKQuantityTypeIdentifierStepCount" sourceName="🔜" sourceVersion="10.1.1" device="&lt;&lt;HKDevice: 0x283714aa0&gt;, name:iPhone, manufacturer:Apple, model:iPhone, hardware:iPhone9,4, software:10.1.1&gt;" unit="count" creationDate="2016-12-08 19:46:37 -0800" startDate="2016-12-08 19:20:57 -0800" endDate="2016-12-08 19:21:58 -0800" value="9"/>
	IPhone RecordSource = "🔜"
	// AppleWatch eg: <Record type="HKQuantityTypeIdentifierHeartRate" sourceName="🕛🕐🕑🕒" sourceVersion="5.0.1" device="&lt;&lt;HKDevice: 0x283735e00&gt;, name:Apple Watch, manufacturer:Apple, model:Watch, hardware:Watch4,2, software:5.0.1&gt;" unit="count/min" creationDate="2018-10-22 17:30:36 -0800" startDate="2018-10-22 17:26:28 -0800" endDate="2018-10-22 17:26:28 -0800" value="64">
	AppleWatch RecordSource = "🕛🕐🕑🕒"
)

// RecordSourcesKnown these "sourceName" in Records are seen in data
var RecordSourcesKnown = RecordSources{AutoSleep, Clock, Headspace, MyFitnessPal, MyFitnessPal, Pacer, Renpho,
	SleepPP, SleepWatch, IPhone, AppleWatch}

// RecordKnownSourcesTypes currently known "type=" emitted by a "sourceName" in a Record
var RecordKnownSourcesTypes = map[RecordSource]RecordTypes{
	Pacer:      {ActiveEnergyBurned},
	Clock:      {SleepAnalysis},
	SleepWatch: {SleepAnalysis},
	SleepPP:    {SleepAnalysis},
	AutoSleep:  {SleepAnalysis},
	Headspace:  {MindfulSession},
	Renpho:     {BasalEnergyBurned, BodyMassIndex, Height, BodyMass, BodyFatPercentage, LeanBodyMass},
	MyFitnessPal: {DietaryFatSaturated, DietaryFatMonounsaturated, DietaryFatPolyunsaturated, DietaryFatTotal,
		DietaryCholesterol, DietarySugar, BodyMass, DietaryEnergyConsumed,
		DietaryCarbohydrates, DietaryProtein, DietaryFiber, DietaryVitaminC,
		DietaryCalcium, DietarySodium, DietaryPotassium, DietaryIron},
	IPhone: {DistanceWalkingRunning, FlightsClimbed, HeadphoneAudioExposure, SixMinuteWalkTestDistance, Height,
		BodyMass, StepCount},
	AppleWatch: {HeartRate, StepCount, AppleStandHour, HeartRateVariabilitySDNN, DistanceWalkingRunning,
		ActiveEnergyBurned, VO2Max, EnvironmentalAudioExposure, BasalEnergyBurned, AppleExerciseTime,
		MindfulSession, HandwashingEvent, FlightsClimbed, RestingHeartRate, WalkingHeartRateAverage, AppleStandTime},
}
