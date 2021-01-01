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

### Record Types
When the records are broken by the count, the histogram looks like:

|  Type of Record | Count of Data |
|-----------------|---------------:|
| ActiveEnergyBurned |  737127  |
| DistanceWalkingRunning |  279203  |
| BasalEnergyBurned |  277120  |
| HeartRate |  257488  |
| StepCount |  151665  |
| AppleExerciseTime |  59054  |
| AppleStandTime |  24354  |
| AppleStandHour |  16378  |
| FlightsClimbed |  8320  |
| HeartRateVariabilitySDNN |  4181  |
| EnvironmentalAudioExposure |  2829  |
| SleepAnalysis |  2703  |
| DietarySugar |  2070  |
| DietaryFatTotal |  1989  |
| DietaryEnergyConsumed |  1889  |
| DietaryFatSaturated |  1884  |
| DietaryFatPolyunsaturated |  1884  |
| DietaryFatMonounsaturated |  1874  |
| DietaryCholesterol |  1864  |
| DietarySodium |  1855  |
| DietaryPotassium |  1847  |
| DietaryCarbohydrates |  1834  |
| DietaryFiber |  1831  |
| DietaryVitaminC |  1820  |
| DietaryIron |  1811  |
| DietaryCalcium |  1809  |
| BodyMass |  1587  |
| VO2Max |  994  |
| BodyMassIndex |  965  |
| LeanBodyMass |  955  |
| BodyFatPercentage |  955  |
| RestingHeartRate |  775  |
| WalkingHeartRateAverage |  769  |
| HeadphoneAudioExposure |  483  |
| HandwashingEvent |  183  |
| MindfulSession |  137  |
| SixMinuteWalkTestDistance |  3  |
| Height |  3  |
| **Total** | **1,856,310** |

### Records by SourceName

The histogram of Record by source is:

| SourceName | Count of Records |
|------------|------------------:|
| AppleWatch | 1711107  |
| iPhone  | 90629  |
| MyFitnessPal |  28699  |
| Pacer |  18559  |
| Renpho |  4496  |
| SleepWatch |  2244  |
| AutoSleep |  419  |
| Headspace |  117  |
| Sleep++ |  36  |
| Clock |  4  |
| **Total** | **1,856,310** |

### Record's SourceName Metadata entries

This histogram is count of all `[]MetadataEntry` entries associated with a `SourceName`.
For example: `AppleWatch` has `268,156 MetadataEntry` elements. 

| SourceName | Count of Metadata |
|------------|------------------:|
| AppleWatch | 268,156 |
| MyFitnessPal | 56,158 |
| AutoSleep | 1073 |
| Sleep++ | 36 |
| Clock | 4 |
| iPhone  | 3 |
| SleepWatch | 0 |
| Pacer | 0 |
| Headspace | 0 |
| Renpho: 0 |
| **Total** |  **325,430** |

### SourceName MetaData key types

The `[]MetadataEntry` for a given `SourceName` comprises these counts of 
`MetadataEntry.Key`. That mean **AppleWatch** stored _257,488_ records of _HKMetadataKeyHeartRateMotionContext_, 
and _4,837_ Records of _HKMetadataKeySyncVersion_,

| SourceName | Metadata Key | Count of Metadata Key |
|------------|---------------|----------------------:|
| AppleWatch | HKMetadataKeyHeartRateMotionContext | 257488 |
| MyFitnessPal | meal | 28079 |
| MyFitnessPal | Meal | 28079 |
| AppleWatch | HKMetadataKeySyncVersion | 4837 |
| AppleWatch | HKMetadataKeySyncIdentifier | 4837 |
| AppleWatch | HKVO2MaxTestType | 994 |
| AutoSleep | Recharge | 140 |
| AutoSleep | Asleep | 140 |
| AutoSleep | "Average HR" | 139 |
| AutoSleep | Rating | 139 |
| AutoSleep | "Daytime HR" | 139 |
| AutoSleep | "Deep Sleep" | 124 |
| AutoSleep | Lights | 109 |
| AutoSleep | "Energy Threshold" | 109 |
| Sleep++ | HKTimeZone | 36 |
| AutoSleep | Nap | 31 |
| Clock | HKTimeZone | 4 |
| iPhone | HKMetadataKeyAppleDeviceCalibrated | 3 |
| AutoSleep | "Edit Slots" | 2 |
| AutoSleep | Tags | 1 |

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
