# Sample HK Data

Here is sample health kit data points.

```text
Export date: 2020-12-06
Me: DOB:1972-06-19, Sex:Male, BloodType:NotSet, SkinType:NotSet, Medication:None
Records: 1856310, #Metadata: 325430, #HRV: 4145
Clinical: 392
Correlations: 0, #Metadata: 0, #Records: 0
Workouts: 1164, #Metadata: 6127, #Events: 9147, Routes: 78
ActivitiesSummary: 780
```

There are approximately 2 million (1,856,310) *Records* over a period of 4 years.

## Records


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

### SourceName Metadata entries

Records Metadata Entries SourceName:

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

