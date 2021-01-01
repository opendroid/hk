package export

import (
	"encoding/xml"
	"fmt"
	"strings"
)

// MetadataEntry Key/Value used across structures
type MetadataEntry struct {
	Key   string `xml:"key,attr"`
	Value string `xml:"value,attr"`
}

// InstantaneousBeatsPerMinute Heart rate
type InstantaneousBeatsPerMinute struct {
	BPM  string `xml:"bpm,attr"`
	Time string `xml:"time,attr"`
}

// HeartRateVariabilityMetadataList  BPM, not like Metadata above
type HeartRateVariabilityMetadataList struct {
	BPM []InstantaneousBeatsPerMinute `xml:"InstantaneousBeatsPerMinute"`
}

// Date a Export Date element
type Date struct {
	XMLName xml.Name `xml:"ExportDate"`
	Value   string   `xml:"value,attr"`
}

// FileReference relative path where a WorkoutRoute data is stored
type FileReference struct {
	XMLName xml.Name `xml:"FileReference"`
	Path    string   `xml:"path,attr"`
}

// Me XML element
type Me struct { // `xml:""`
	XMLName xml.Name `xml:"Me"`
	// Attributes
	DateOfBirth   string `xml:"HKCharacteristicTypeIdentifierDateOfBirth,attr"`
	BiologicalSex string `xml:"HKCharacteristicTypeIdentifierBiologicalSex,attr"`
	BloodType     string `xml:"HKCharacteristicTypeIdentifierBloodType,attr"`
	SkinType      string `xml:"HKCharacteristicTypeIdentifierFitzpatrickSkinType,attr"`
	MedicationUse string `xml:"HKCharacteristicTypeIdentifierCardioFitnessMedicationsUse,attr"`
}

// Record a typical record saved by various devices. A example is:
//   <Record type="HKQuantityTypeIdentifierBodyMassIndex" sourceName="Renpho" sourceVersion="4" unit="count" creationDate="2020-04-19 06:53:46 -0800" startDate="2020-04-14 05:51:27 -0800" endDate="2020-04-14 05:51:27 -0800" value="22.6"/>
type Record struct {
	XMLName xml.Name `xml:"Record"`
	// Attributes
	Type          string `xml:"type,attr"`
	Unit          string `xml:"unit,attr"`
	Value         string `xml:"value,attr"`
	SourceName    string `xml:"sourceName,attr"`
	SourceVersion string `xml:"sourceVersion,attr"`
	Device        string `xml:"device,attr"`
	CreationDate  string `xml:"creationDate,attr"`
	StartDate     string `xml:"startDate,attr"`
	EndDate       string `xml:"endDate,attr"`
	// Elements
	MetadataEntries []MetadataEntry                    `xml:"MetadataEntry"`
	HRV             []HeartRateVariabilityMetadataList `xml:"HeartRateVariabilityMetadataList"`
}

// ClinicalRecord element data. Contains Source, URL and FilePath data provided.
type ClinicalRecord struct {
	XMLName xml.Name `xml:"ClinicalRecord"`
	// Attributes
	Type         string `xml:"type,attr"`
	ID           string `xml:"identifier,attr"`
	SourceName   string `xml:"sourceName,attr"`
	SourceURL    string `xml:"sourceURL,attr"`
	FHIRVersion  string `xml:"fhirVersion,attr"`
	ReceivedDate string `xml:"receivedDate,attr"`
	FilePath     string `xml:"resourceFilePath,attr"`
}

// Correlation element groups correlated Records.
type Correlation struct {
	XMLName xml.Name `xml:"Correlation"`
	// Attributes
	Type          string `xml:"type,attr"`
	SourceName    string `xml:"sourceName,attr"`
	SourceVersion string `xml:"sourceVersion,attr,omitempty"`
	Device        string `xml:"device,attr,omitempty"`
	CreationDate  string `xml:"creationDate,attr,omitempty"`
	StartDate     string `xml:"startDate,attr"`
	EndDate       string `xml:"endDate,attr"`
	// Elements
	MetadataEntries []MetadataEntry `xml:"MetadataEntry"`
	Records         []Record        `xml:"Record"`
}

// WorkoutEvent element outlines the Type and Duration of workout
type WorkoutEvent struct {
	XMLName xml.Name `xml:"WorkoutEvent"`
	// Attributes
	Type         string `xml:"type,attr"`
	Date         string `xml:"date,attr"`
	Duration     string `xml:"duration,attr"`
	DurationUnit string `xml:"durationUnit,attr"`
}

// WorkoutRoute gpx data file name with the route
type WorkoutRoute struct {
	MetadataEntries []MetadataEntry `xml:"MetadataEntry"`
	FileReference   []FileReference `xml:"FileReference"`
}

// Workout key workout element.
type Workout struct {
	XMLName xml.Name `xml:"Workout"`
	// attributes
	ActivityType          string `xml:"workoutActivityType,attr"`
	Duration              string `xml:"duration,attr"`
	DurationUnit          string `xml:"durationUnit,attr"`
	TotalDistance         string `xml:"totalDistance,attr"`
	TotalDistanceUnit     string `xml:"totalDistanceUnit,attr"`
	TotalEnergyBurned     string `xml:"totalEnergyBurned,attr"`
	TotalEnergyBurnedUnit string `xml:"totalEnergyBurnedUnit,attr"`
	SourceName            string `xml:"sourceName,attr"`
	SourceVersion         string `xml:"sourceVersion,attr"`
	Device                string `xml:"device,attr"`
	CreationDate          string `xml:"creationDate,attr"`
	StartDate             string `xml:"startDate,attr"`
	EndDate               string `xml:"endDate,attr"`
	// Elements
	MetadataEntries []MetadataEntry `xml:"MetadataEntry"`
	Events          []WorkoutEvent  `xml:"WorkoutEvent"`
	Routes          []WorkoutRoute  `xml:"WorkoutRoute"`
}

type ActivitySummary struct {
	XMLName xml.Name `xml:"ActivitySummary"`
	// Attributes
	YYYYMMDD         string `xml:"dateComponents,attr"`
	EnergyBurned     string `xml:"activeEnergyBurned,attr"`
	EnergyBurnedGoal string `xml:"activeEnergyBurnedGoal,attr"`
	EnergyBurnedUnit string `xml:"activeEnergyBurnedUnit,attr"`
	MoveTime         string `xml:"appleMoveTime,attr"`
	MoveTimeGoal     string `xml:"appleMoveTimeGoal,attr"`
	ExerciseTime     string `xml:"appleExerciseTime,attr"`
	ExerciseTimeGoal string `xml:"appleExerciseTimeGoal,attr"`
	StandHours       string `xml:"appleStandHours,attr"`
	StandHoursGoal   string `xml:"appleStandHoursGoal,attr"`
}

type HealthData struct {
	Exported          Date              `xml:"ExportDate"`
	Me                Me                `xml:"Me"`
	Records           []Record          `xml:"Record"`
	ClinicalRecords   []ClinicalRecord  `xml:"ClinicalRecord"`
	Correlations      []Correlation     `xml:"Correlation"`
	Workouts          []Workout         `xml:"Workout"`
	ActivitiesSummary []ActivitySummary `xml:"ActivitySummary"`
}

// String describes ActivitySummary
func (a ActivitySummary) String() string {
	description := ""
	if a.XMLName.Space != "" {
		description += a.XMLName.Space
	}
	if a.XMLName.Local != "" {
		if a.XMLName.Space != "" {
			description += "."
		}
		description += a.XMLName.Local + ":"
	}
	if a.YYYYMMDD != "" {
		description += a.YYYYMMDD + " ("
	}

	if a.EnergyBurned != "" {
		description += "Energy: " + a.EnergyBurned
		if a.EnergyBurnedGoal != "" {
			description += "/" + a.EnergyBurnedGoal
		}
		if a.EnergyBurnedUnit != "" {
			description += " " + a.EnergyBurnedUnit
		}
		description += ", "
	}

	if a.MoveTime != "" {
		description += "Move: " + a.MoveTime
		if a.MoveTimeGoal != "" {
			description += "/" + a.MoveTimeGoal
		}
		description += ", "
	}

	if a.ExerciseTime != "" {
		description += "Exercise: " + a.ExerciseTime
		if a.ExerciseTimeGoal != "" {
			description += "/" + a.ExerciseTimeGoal
		}
		description += " Minutes, "
	}

	if a.StandHours != "" {
		description += "Stand: " + a.StandHours
		if a.StandHoursGoal != "" {
			description += "/" + a.StandHoursGoal
		}
		description += " Hours"
	}

	return description + ")"
}

// String describes Me
func (m Me) String() string {
	description := ""
	if m.XMLName.Space != "" {
		description += m.XMLName.Space
	}
	if m.XMLName.Local != "" {
		if m.XMLName.Space != "" {
			description += "."
		}
		description += m.XMLName.Local + ": "
	}

	if m.DateOfBirth != "" {
		description += "DOB:" + m.DateOfBirth + ", "
	}
	if m.BiologicalSex != "" {
		description += "Sex:" + strings.Replace(m.BiologicalSex, "HKBiologicalSex", "", 1) + ", "
	}

	if m.BloodType != "" {
		description += "BloodType:" + strings.Replace(m.BloodType, "HKBloodType", "", 1) + ", "
	}

	if m.SkinType != "" {
		description += "SkinType:" + strings.Replace(m.SkinType, "HKFitzpatrickSkinType", "", 1) + ", "
	}

	if m.MedicationUse != "" {
		description += "Medication:" + m.MedicationUse
	}

	return description
}

// String describes a Record
func (r Record) String() string {
	description := ""
	if r.XMLName.Space != "" {
		description += r.XMLName.Space
	}
	if r.XMLName.Local != "" {
		if r.XMLName.Space != "" {
			description += "."
		}
		description += r.XMLName.Local + ": "
	}
	// Print Record "Type" and "value unit"
	if r.Type != "" {
		description += strings.Replace(r.Type, "HKQuantityTypeIdentifier", "", 1) + " "
		if r.Value != "" {
			description += r.Value + " "
			if r.Unit != "" {
				description += r.Unit
			}
		}
	}

	if r.SourceName != "" {
		description += " on " + r.SourceName
	}

	if r.SourceVersion != "" {
		description += " version: " + r.SourceVersion
	}

	if r.Device != "" {
		description += ", Device: " + r.Device
	}

	if r.StartDate != "" {
		description += " (S:" + yyyymmdd(r.StartDate) + ", "
	}

	if r.EndDate != "" {
		if r.StartDate == "" {
			description += " ("
		}
		description += "E:" + yyyymmdd(r.EndDate) + ", "
	}

	if r.CreationDate != "" {
		if r.StartDate == "" && r.EndDate != "" {
			description += " ("
		}
		description += "C:" + yyyymmdd(r.CreationDate)
	}

	return description + ")"
}

// Describe prints summary health data
func (h *HealthData) Describe() {
	if h == nil {
		return
	}

	fmt.Printf("Export date: %s\n", yyyymmdd(h.Exported.Value))
	fmt.Printf("%s\n", h.Me)
	// Records summary data
	h.DescribeRecords()

	// Clinical summary data
	fmt.Printf("Clinical: %d\n", len(h.ClinicalRecords))

	// Correlations summary data
	nm, nR := 0, 0
	for _, c := range h.Correlations {
		m, r := len(c.MetadataEntries), len(c.Records)
		nm += m
		nR += r
	}
	fmt.Printf("Correlations: %d, #Metadata: %d, #Records: %d\n", len(h.Correlations), nm, nR)

	// Workouts summary data
	nm, nR, ne := 0, 0, 0
	for _, w := range h.Workouts {
		m, e, r := len(w.MetadataEntries), len(w.Events), len(w.Routes)
		ne += e
		nm += m
		nR += r
	}
	fmt.Printf("Workouts: %d, #Metadata: %d, #Events: %d, Routes: %d\n", len(h.Workouts), nm, ne, nR)
	fmt.Printf("ActivitiesSummary: %d\n", len(h.ActivitiesSummary))
}

// UnseenCheck Prints out any unseen keys
func (h *HealthData) UnseenCheck() {
	if h == nil {
		fmt.Printf("No health data found.")
		return
	}
	// Workouts: Check for metadata first
	seen := workoutKeys(h.Workouts)
	unknown := unknownWorkoutKeys(seen)
	printUnknownKeys("New Workout Metadata Keys:", unknown)
	seen = workoutActivityTypes(h.Workouts)
	unknown = unknownWorkoutActivityTypes(seen)
	printUnknownKeys("New Workout Activity Types:", unknown)

	// Check the records for new data types
	seen = recordKeys(h.Records)
	unknown = unknownRecordKeys(seen)
	printUnknownKeys("New Records Metadata Keys:", unknown)
	seen = recordTypes(h.Records)
	unknown = unknownRecordTypes(seen)
	printUnknownKeys("New Records Metadata Types:", unknown)
	seen = recordSources(h.Records)
	unknown = unknownRecordSources(seen)
	printUnknownKeys("New Records Sources:", unknown)
	at := recordSourcesTypes(h.Records)
	ut := unknownSourcesTypes(at)
	printUnknownSourcesTypes("New Records Sources Types:", ut)
}

// DescribeRecords prints a summary data by source type
func (h *HealthData) DescribeRecords() {
	if h == nil {
		return
	}
	nm, nHRV := 0, 0
	for _, r := range h.Records {
		m, hrv := len(r.MetadataEntries), len(r.HRV)
		nm += m
		nHRV += hrv
	}
	fmt.Printf("Records: %d, #Metadata: %d, #HRV: %d\n", len(h.Records), nm, nHRV)

	// Describe summary by Type
	t := make(map[string]int)
	for _, r := range h.Records {
		if _, ok := t[r.Type]; !ok {
			t[r.Type] = 1
		} else {
			t[r.Type]++
		}
	}
	printStringInts("Records by Type:", t)

	// Describe record by SourceName
	t = make(map[string]int)
	for _, r := range h.Records {
		if _, ok := t[r.SourceName]; !ok {
			t[r.SourceName] = 1
		} else {
			t[r.SourceName]++
		}
	}
	printStringInts("Records by SourceName:", t)

	// Meta data for each SourceName, init count to 0
	for k := range t {
		t[k] = 0
	}
	for _, r := range h.Records {
		t[r.SourceName] += len(r.MetadataEntries)
	}
	printStringInts("Records Metadata Entries SourceName:", t)

	// Print Metadata keys by SourceName
	md := make(map[string]map[string]int)
	for _, r := range h.Records {
		if _, ok := md[r.SourceName]; !ok {
			md[r.SourceName] = make(map[string]int)
		}
		for _, m := range r.MetadataEntries {
			if _, ok := md[r.SourceName][m.Key]; !ok {
				md[r.SourceName][m.Key] = 1
			} else {
				md[r.SourceName][m.Key]++
			}
		}
	}
	for k, v := range md {
		for dk, dv := range v {
			fmt.Printf("%q:%q:%d\n", k, dk, dv)
		}
	}
}
