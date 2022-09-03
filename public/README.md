# Sample HK Data

Here is sample health kit data points.

```text
Export date: 2020-12-06
Me: DOB:2020-12-06, Sex:Male, BloodType:NotSet, SkinType:NotSet, Medication:None
Records: 1856310, #Metadata: 325430, #HRV: 4145
Clinical: 392
Correlations: 0, #Metadata: 0, #Records: 0
Workouts: 1164, #Metadata: 6127, #Events: 9147, Routes: 78
ActivitiesSummary: 780
```

There are approximately 2 million (1,856,310) *Records* over a period of 4 years.

## Records
Various apps like MyFitnessPal or Clock save data in Apple HealthKit using "Record". These records have format
as shown in `Record` data-type. The main fields in `Record` are:
- **Type:** Shows type of record eg: Height or HeartRate
- **SourceName:** Source that entered this record
- **MetadataEntries:** a `[]` of `Key`/`Value` pairs
- **HRVMetadata:** a `[]` of `Key`/`Value` pairs related to HRV-BPM (Beats Per Minute)


### Records Summary

| Type                       | Source Name  | Metadata Entry Tag     |  Count |
|----------------------------|--------------|------------------------|-------:|
| ActiveEnergyBurned         | 🕛🕐🕑🕒     | -                      | 717147 |
| ActiveEnergyBurned         | 🕛🕐🕑🕒     | SyncVersion            |   1421 |
| ActiveEnergyBurned         | 🕛🕐🕑🕒     | SyncIdentifier         |   1421 |
| ActiveEnergyBurned         | Pacer        | -                      |  18559 |
| DistanceWalkingRunning     | 🕛🕐🕑🕒     | -                      | 240244 |
| DistanceWalkingRunning     | 🕛🕐🕑🕒     | SyncVersion            |   1416 |
| DistanceWalkingRunning     | 🕛🕐🕑🕒     | SyncIdentifier         |   1416 |
| DistanceWalkingRunning     | 🔜           | -                      |  37543 |
| BasalEnergyBurned          | 🕛🕐🕑🕒     | -                      | 275044 |
| BasalEnergyBurned          | 🕛🕐🕑🕒     | SyncVersion            |   1421 |
| BasalEnergyBurned          | 🕛🕐🕑🕒     | SyncIdentifier         |   1421 |
| BasalEnergyBurned          | Renpho       | -                      |    655 |
| HeartRate                  | 🕛🕐🕑🕒     | HeartRateMotionContext | 257488 |
| HeartRate                  | 🕛🕐🕑🕒     | SyncVersion            |    579 |
| HeartRate                  | 🕛🕐🕑🕒     | SyncIdentifier         |    579 |
| StepCount                  | 🕛🕐🕑🕒     | -                      | 104113 |
| StepCount                  | 🔜           | -                      |  47552 |
| AppleExerciseTime          | 🕛🕐🕑🕒     | -                      |  59054 |
| AppleStandTime             | 🕛🕐🕑🕒     | -                      |  24354 |
| AppleStandHour             | 🕛🕐🕑🕒     | -                      |  16378 |
| FlightsClimbed             | 🔜           | -                      |   5044 |
| FlightsClimbed             | 🕛🕐🕑🕒     | -                      |   3276 |
| HeartRateVariabilitySDNN   | 🕛🕐🕑🕒     | -                      |   4181 |
| DietarySugar               | MyFitnessPal | meal                   |   2070 |
| DietarySugar               | MyFitnessPal | Meal                   |   2070 |
| DietaryFatTotal            | MyFitnessPal | meal                   |   1989 |
| DietaryFatTotal            | MyFitnessPal | Meal                   |   1989 |
| DietaryEnergyConsumed      | MyFitnessPal | meal                   |   1889 |
| DietaryEnergyConsumed      | MyFitnessPal | Meal                   |   1889 |
| DietaryFatSaturated        | MyFitnessPal | meal                   |   1884 |
| DietaryFatSaturated        | MyFitnessPal | Meal                   |   1884 |
| DietaryFatPolyunsaturated  | MyFitnessPal | meal                   |   1884 |
| DietaryFatPolyunsaturated  | MyFitnessPal | Meal                   |   1884 |
| DietaryFatMonounsaturated  | MyFitnessPal | meal                   |   1874 |
| DietaryFatMonounsaturated  | MyFitnessPal | Meal                   |   1874 |
| DietaryCholesterol         | MyFitnessPal | Meal                   |   1864 |
| DietaryCholesterol         | MyFitnessPal | meal                   |   1864 |
| DietarySodium              | MyFitnessPal | meal                   |   1855 |
| DietarySodium              | MyFitnessPal | Meal                   |   1855 |
| DietaryPotassium           | MyFitnessPal | meal                   |   1847 |
| DietaryPotassium           | MyFitnessPal | Meal                   |   1847 |
| DietaryCarbohydrates       | MyFitnessPal | meal                   |   1834 |
| DietaryCarbohydrates       | MyFitnessPal | Meal                   |   1834 |
| DietaryFiber               | MyFitnessPal | meal                   |   1831 |
| DietaryFiber               | MyFitnessPal | Meal                   |   1831 |
| DietaryVitaminC            | MyFitnessPal | meal                   |   1820 |
| DietaryVitaminC            | MyFitnessPal | Meal                   |   1820 |
| DietaryProtein             | MyFitnessPal | meal                   |   1818 |
| DietaryProtein             | MyFitnessPal | Meal                   |   1818 |
| SleepAnalysis              | SleepWatch   | -                      |   2244 |
| SleepAnalysis              | AutoSleep    | -                      |    279 |
| SleepAnalysis              | AutoSleep    | Recharge               |    140 |
| SleepAnalysis              | AutoSleep    | Asleep                 |    140 |
| SleepAnalysis              | AutoSleep    | Rating                 |    139 |
| SleepAnalysis              | AutoSleep    | Daytime HR             |    139 |
| SleepAnalysis              | AutoSleep    | Average HR             |    139 |
| SleepAnalysis              | AutoSleep    | Deep Sleep             |    124 |
| SleepAnalysis              | AutoSleep    | Lights                 |    109 |
| SleepAnalysis              | AutoSleep    | Energy Threshold       |    109 |
| SleepAnalysis              | AutoSleep    | Nap                    |     31 |
| SleepAnalysis              | AutoSleep    | Edit Slots             |      2 |
| SleepAnalysis              | AutoSleep    | Tags                   |      1 |
| SleepAnalysis              | Sleep++      | HKTimeZone             |     36 |
| SleepAnalysis              | Clock        | HKTimeZone             |      4 |
| DietaryIron                | MyFitnessPal | meal                   |   1811 |
| DietaryIron                | MyFitnessPal | Meal                   |   1811 |
| DietaryCalcium             | MyFitnessPal | meal                   |   1809 |
| DietaryCalcium             | MyFitnessPal | Meal                   |   1809 |
| EnvironmentalAudioExposure | 🕛🕐🕑🕒     | -                      |   2829 |
| BodyMass                   | Renpho       | -                      |    965 |
| BodyMass                   | MyFitnessPal | -                      |    620 |
| BodyMass                   | 🔜           | -                      |      2 |
| VO2Max                     | 🕛🕐🕑🕒     | HKVO2MaxTestType       |    994 |
| BodyMassIndex              | Renpho       | -                      |    965 |
| BodyFatPercentage          | Renpho       | -                      |    955 |
| LeanBodyMass               | Renpho       | -                      |    955 |
| RestingHeartRate           | 🕛🕐🕑🕒     | -                      |    775 |
| WalkingHeartRateAverage    | 🕛🕐🕑🕒     | -                      |    769 |
| HeadphoneAudioExposure     | 🔜           | -                      |    483 |
| HandwashingEvent           | 🕛🕐🕑🕒     | -                      |    183 |
| MindfulSession             | Headspace    | -                      |    117 |
| MindfulSession             | 🕛🕐🕑🕒     | -                      |     20 |
| Height                     | 🔜           | -                      |      2 |
| Height                     | Renpho       | -                      |      1 |
| SixMinuteWalkTestDistance  | 🔜           | AppleDeviceCalibrated  |      3 |


### Record Types
When the records are broken by the count, the histogram looks like:

| Type of Record             | Count of Data |
|----------------------------|--------------:|
| ActiveEnergyBurned         |        737127 |
| DistanceWalkingRunning     |        279203 |
| BasalEnergyBurned          |        277120 |
| HeartRate                  |        257488 |
| StepCount                  |        151665 |
| AppleExerciseTime          |         59054 |
| AppleStandTime             |         24354 |
| AppleStandHour             |         16378 |
| FlightsClimbed             |          8320 |
| HeartRateVariabilitySDNN   |          4181 |
| EnvironmentalAudioExposure |          2829 |
| SleepAnalysis              |          2703 |
| DietarySugar               |          2070 |
| DietaryFatTotal            |          1989 |
| DietaryEnergyConsumed      |          1889 |
| DietaryFatSaturated        |          1884 |
| DietaryFatPolyunsaturated  |          1884 |
| DietaryFatMonounsaturated  |          1874 |
| DietaryCholesterol         |          1864 |
| DietarySodium              |          1855 |
| DietaryPotassium           |          1847 |
| DietaryCarbohydrates       |          1834 |
| DietaryFiber               |          1831 |
| DietaryVitaminC            |          1820 |
| DietaryIron                |          1811 |
| DietaryCalcium             |          1809 |
| BodyMass                   |          1587 |
| VO2Max                     |           994 |
| BodyMassIndex              |           965 |
| LeanBodyMass               |           955 |
| BodyFatPercentage          |           955 |
| RestingHeartRate           |           775 |
| WalkingHeartRateAverage    |           769 |
| HeadphoneAudioExposure     |           483 |
| HandwashingEvent           |           183 |
| MindfulSession             |           137 |
| SixMinuteWalkTestDistance  |             3 |
| Height                     |             3 |
| **Total**                  | **1,856,310** |

### Records by SourceName

The histogram of Record by source is:

| SourceName   | Count of Records |
|--------------|-----------------:|
| AppleWatch   |          1711107 |
| iPhone       |            90629 |
| MyFitnessPal |            28699 |
| Pacer        |            18559 |
| Renpho       |             4496 |
| SleepWatch   |             2244 |
| AutoSleep    |              419 |
| Headspace    |              117 |
| Sleep++      |               36 |
| Clock        |                4 |
| **Total**    |    **1,856,310** |

### Record's SourceName Metadata entries

This histogram is count of all `[]MetadataEntry` entries associated with a `SourceName`.
For example: `AppleWatch` has `268,156 MetadataEntry` elements. 

| SourceName   | Count of Metadata |
|--------------|------------------:|
| AppleWatch   |           268,156 |
| MyFitnessPal |            56,158 |
| AutoSleep    |              1073 |
| Sleep++      |                36 |
| Clock        |                 4 |
| iPhone       |                 3 |
| SleepWatch   |                 0 |
| Pacer        |                 0 |
| Headspace    |                 0 |
| Renpho: 0    |
| **Total**    |       **325,430** |

### SourceName MetaData key types

The `[]MetadataEntry` for a given `SourceName` comprises these counts of 
`MetadataEntry.Key`. That mean **Apple Watch** stored _257,488_ records of _HKMetadataKeyHeartRateMotionContext_, 
and _4,837_ Records of _HKMetadataKeySyncVersion_,

| SourceName   | Metadata Key                        | Count of Metadata Key |
|--------------|-------------------------------------|----------------------:|
| AppleWatch   | HKMetadataKeyHeartRateMotionContext |                257488 |
| MyFitnessPal | meal                                |                 28079 |
| MyFitnessPal | Meal                                |                 28079 |
| AppleWatch   | HKMetadataKeySyncVersion            |                  4837 |
| AppleWatch   | HKMetadataKeySyncIdentifier         |                  4837 |
| AppleWatch   | HKVO2MaxTestType                    |                   994 |
| AutoSleep    | Recharge                            |                   140 |
| AutoSleep    | Asleep                              |                   140 |
| AutoSleep    | "Average HR"                        |                   139 |
| AutoSleep    | Rating                              |                   139 |
| AutoSleep    | "Daytime HR"                        |                   139 |
| AutoSleep    | "Deep Sleep"                        |                   124 |
| AutoSleep    | Lights                              |                   109 |
| AutoSleep    | "Energy Threshold"                  |                   109 |
| Sleep++      | HKTimeZone                          |                    36 |
| AutoSleep    | Nap                                 |                    31 |
| Clock        | HKTimeZone                          |                     4 |
| iPhone       | HKMetadataKeyAppleDeviceCalibrated  |                     3 |
| AutoSleep    | "Edit Slots"                        |                     2 |
| AutoSleep    | Tags                                |                     1 |

### Record HRV 
The example of Record of type `HKQuantityTypeIdentifierHeartRateVariabilitySDNN` is:
```xml
 <Record type="HKQuantityTypeIdentifierHeartRateVariabilitySDNN" sourceName="🕛🕐🕑🕒" sourceVersion="5.0.1" device="&lt;&lt;HKDevice: 0x28347bac0&gt;, name:Apple Watch, manufacturer:Apple, model:Watch, hardware:Watch4,2, software:5.0.1&gt;" unit="ms" creationDate="2018-10-29 10:18:22 -0800" startDate="2018-10-29 10:17:16 -0800" endDate="2018-10-29 10:18:21 -0800" value="49.4602">
  <HeartRateVariabilityMetadataList>
   <InstantaneousBeatsPerMinute bpm="80" time="11:17:20.53 AM"/>
   <InstantaneousBeatsPerMinute bpm="80" time="11:17:21.29 AM"/>
  </HeartRateVariabilityMetadataList>
</Record>
```
