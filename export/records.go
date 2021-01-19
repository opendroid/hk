package export

import (
	"encoding/xml"
	"fmt"
	"io"
	"sort"
	"strings"
)

// MetadataEntry Key/Value used across structures.
// Examples of MetadataEntry with Record.SourceName
//   AppleWatch: <MetadataEntry key="HKMetadataKeyHeartRateMotionContext" value="1"/>
//   MyFitnessPal: <MetadataEntry key="meal" value="Lunch"/>
//   MyFitnessPal: <MetadataEntry key="meal" value="Dinner"/>
//   MyFitnessPal: <MetadataEntry key="Meal" value="Dinner"/>
//   MyFitnessPal: <MetadataEntry key="Meal" value="Snacks"/>
//   AppleWatch: <MetadataEntry key="HKMetadataKeySyncVersion" value="1"/>
//   AppleWatch: <MetadataEntry key="HKMetadataKeySyncIdentifier" value="3:1B17A065-24BD-4C42-88C4-2185852AC39B:613580458.67789:613580909.63101:89"/>
//   AppleWatch: <MetadataEntry key="HKVO2MaxTestType" value="2"/>
//   AutoSleep: <MetadataEntry key="Recharge" value="100"/>
//   AutoSleep: <MetadataEntry key="Asleep" value="19020"/>
//   AutoSleep: <MetadataEntry key="Average HR" value="49.63"/>
//   AutoSleep: <MetadataEntry key="Rating" value="11.97"/>
//   AutoSleep: <MetadataEntry key="Daytime HR" value="68.35"/>
//   AutoSleep: <MetadataEntry key="Deep Sleep" value="1020"/>
//   AutoSleep: <MetadataEntry key="Lights" value="NO"/>
//   AutoSleep: <MetadataEntry key="Energy Threshold" value="1200"/>
//   AutoSleep: <MetadataEntry key="Nap" value="YES"/>
//   AutoSleep: <MetadataEntry key="Tags" value="1"/>
//   AutoSleep: <MetadataEntry key="Edit Slots" value="0415,0630,0500,0445,0530,0615,0430,0645,0515,0600,0545"/>
//   Sleep++ and Clock: <MetadataEntry key="HKTimeZone" value="America/Los_Angeles"/>
//   iPhone: <MetadataEntry key="HKMetadataKeyAppleDeviceCalibrated" value="1"/>
type MetadataEntry struct {
	Key   string `xml:"key,attr" json:"key"`
	Value string `xml:"value,attr" json:"value"`
}

// InstantaneousBeatsPerMinute Heart rate
type InstantaneousBeatsPerMinute struct {
	BPM  string `xml:"bpm,attr" json:"bpm"`
	Time string `xml:"time,attr" json:"time"`
}

// HeartRateVariabilityMetadataList  BPM, not like Metadata above
type HeartRateVariabilityMetadataList struct {
	BPM []InstantaneousBeatsPerMinute `xml:"InstantaneousBeatsPerMinute" json:"instantaneous_bpm"`
}

// Date a Export Date element
type Date struct {
	XMLName xml.Name `xml:"ExportDate" json:"-"`
	Value   string   `xml:"value,attr" json:"value"`
}

// FileReference relative path where a WorkoutRoute data is stored
type FileReference struct {
	XMLName xml.Name `xml:"FileReference" json:"-"`
	Path    string   `xml:"path,attr" json:"path"`
}

// Me XML element
type Me struct { // `xml:""`
	XMLName xml.Name `xml:"Me" json:"-"`
	// Attributes
	DateOfBirth   string `xml:"HKCharacteristicTypeIdentifierDateOfBirth,attr" json:"dob"`
	BiologicalSex string `xml:"HKCharacteristicTypeIdentifierBiologicalSex,attr" json:"biological_sex"`
	BloodType     string `xml:"HKCharacteristicTypeIdentifierBloodType,attr" json:"blood_type"`
	SkinType      string `xml:"HKCharacteristicTypeIdentifierFitzpatrickSkinType,attr" json:"skin_type"`
	MedicationUse string `xml:"HKCharacteristicTypeIdentifierCardioFitnessMedicationsUse,attr,omitempty" json:"medication_use,omitempty"`
}

// Record a typical record saved by various devices. A example is:
//   <Record type="HKQuantityTypeIdentifierBodyMassIndex" sourceName="Renpho" sourceVersion="4" unit="count" creationDate="2020-04-19 06:53:46 -0800" startDate="2020-04-14 05:51:27 -0800" endDate="2020-04-14 05:51:27 -0800" value="22.6"/>
type Record struct {
	XMLName xml.Name `xml:"Record" json:"-"`
	// Attributes
	Type          string `xml:"type,attr" json:"type"`
	Unit          string `xml:"unit,attr,omitempty" json:"unit,omitempty"`
	Value         string `xml:"value,attr,omitempty" json:"value,omitempty"`
	SourceName    string `xml:"sourceName,attr" json:"source_name"`
	SourceVersion string `xml:"sourceVersion,attr,omitempty" json:"source_version,omitempty"`
	Device        string `xml:"device,attr,omitempty" json:"device,omitempty"`
	CreationDate  string `xml:"creationDate,attr,omitempty" json:"creation_date,omitempty"`
	StartDate     string `xml:"startDate,attr" json:"start_date"`
	EndDate       string `xml:"endDate,attr" json:"end_date"`
	// Elements
	MetadataEntries []MetadataEntry                    `xml:"MetadataEntry,omitempty" json:"metadata,omitempty"`
	HRV             []HeartRateVariabilityMetadataList `xml:"HeartRateVariabilityMetadataList,omitempty"  json:"hrv,omitempty"`
}

// Audiogram records
type Audiogram struct {
	XMLName xml.Name `xml:"Audiogram" json:"-"`
	// Attributes
	Type          string `xml:"type,attr" json:"type"`
	SourceName    string `xml:"sourceName,attr" json:"source_name"`
	SourceVersion string `xml:"sourceVersion,attr,omitempty" json:"source_version,omitempty"`
	Device        string `xml:"device,attr,omitempty" json:"device,omitempty"`
	CreationDate  string `xml:"creationDate,attr,omitempty" json:"creation_date,omitempty"`
	StartDate     string `xml:"startDate,attr" json:"start_date"`
	EndDate       string `xml:"endDate,attr" json:"end_date"`
}

// SensitivityPoint microphone data
type SensitivityPoint struct {
	XMLName xml.Name `xml:"SensitivityPoint" json:"-"`
	// Attributes
	FrequencyValue string `xml:"frequencyValue" json:"frequency_value"`
	FrequencyUnit  string `xml:"frequencyUnit" json:"frequency_unit"`
	LeftEarValue   string `xml:"leftEarValue,omitempty" json:"left_ear_value,omitempty"`
	LeftEarUnit    string `xml:"leftEarUnit,omitempty" json:"left_ear_unit,omitempty"`
	RightEarValue  string `xml:"rightEarValue,omitempty" json:"right_ear_value,omitempty"`
	RightEarUnit   string `xml:"rightEarUnit,omitempty" json:"right_ear_unit,omitempty"`
}

// ClinicalRecord element data. Contains Source, URL and FilePath data provided.
type ClinicalRecord struct {
	XMLName xml.Name `xml:"ClinicalRecord" json:"-"`
	// Attributes
	Type         string `xml:"type,attr" json:"type"`
	ID           string `xml:"identifier,attr" json:"id"`
	SourceName   string `xml:"sourceName,attr" json:"source_name"`
	SourceURL    string `xml:"sourceURL,attr" json:"source_url"`
	FHIRVersion  string `xml:"fhirVersion,attr" json:"fhir_version"`
	ReceivedDate string `xml:"receivedDate,attr" json:"received_date"`
	FilePath     string `xml:"resourceFilePath,attr" json:"file_path"`
}

// Correlation element groups correlated Records.
type Correlation struct {
	XMLName xml.Name `xml:"Correlation" json:"-"`
	// Attributes
	Type          string `xml:"type,attr" json:"type"`
	SourceName    string `xml:"sourceName,attr" json:"source_name"`
	SourceVersion string `xml:"sourceVersion,attr,omitempty" json:"source_version,omitempty"`
	Device        string `xml:"device,attr,omitempty" json:"device,omitempty"`
	CreationDate  string `xml:"creationDate,attr,omitempty" json:"creation_date,omitempty"`
	StartDate     string `xml:"startDate,attr" json:"start_date"`
	EndDate       string `xml:"endDate,attr" json:"end_date"`
	// Elements
	MetadataEntries []MetadataEntry `xml:"MetadataEntry" json:"metadata"`
	Records         []Record        `xml:"Record" json:"records"`
}

// WorkoutEvent element outlines the Type and Duration of workout
type WorkoutEvent struct {
	XMLName xml.Name `xml:"WorkoutEvent" json:"-"`
	// Attributes
	Type         string `xml:"type,attr" json:"type"`
	Date         string `xml:"date,attr" json:"date"`
	Duration     string `xml:"duration,attr,omitempty" json:"duration,omitempty"`
	DurationUnit string `xml:"durationUnit,attr,omitempty" json:"duration_unit,omitempty"`
}

// WorkoutRoute gpx data file name with the route
type WorkoutRoute struct {
	XMLName xml.Name `xml:"WorkoutRoute" json:"-"`
	// Attributes
	SourceName    string `xml:"sourceName,attr" json:"source_name"`
	SourceVersion string `xml:"sourceVersion,attr,omitempty" json:"source_version,omitempty"`
	Device        string `xml:"device,attr,omitempty" json:"device,omitempty"`
	CreationDate  string `xml:"creationDate,attr,omitempty" json:"creation_date,omitempty"`
	StartDate     string `xml:"startDate,attr" json:"start_date"`
	EndDate       string `xml:"endDate,attr" json:"end_date"`
	// Elements
	MetadataEntries []MetadataEntry `xml:"MetadataEntry" json:"metadata"`
	FileReference   []FileReference `xml:"FileReference" json:"file_reference"`
}

// Workout key workout element.
type Workout struct {
	XMLName xml.Name `xml:"Workout"  json:"-"`
	// attributes
	ActivityType          string `xml:"workoutActivityType,attr" json:"activity_type"`
	Duration              string `xml:"duration,attr,omitempty" json:"duration,omitempty"`
	DurationUnit          string `xml:"durationUnit,attr,omitempty" json:"duration_unit,omitempty"`
	TotalDistance         string `xml:"totalDistance,attr,omitempty" json:"total_distance,omitempty"`
	TotalDistanceUnit     string `xml:"totalDistanceUnit,attr,omitempty" json:"total_distance_unit,omitempty"`
	TotalEnergyBurned     string `xml:"totalEnergyBurned,attr,omitempty" json:"total_energy_burned,omitempty"`
	TotalEnergyBurnedUnit string `xml:"totalEnergyBurnedUnit,attr,omitempty" json:"total_energy_burned_unit,omitempty"`
	SourceName            string `xml:"sourceName,attr" json:"source_name"`
	SourceVersion         string `xml:"sourceVersion,attr,omitempty" json:"source_version,omitempty"`
	Device                string `xml:"device,attr,omitempty" json:"device,omitempty"`
	CreationDate          string `xml:"creationDate,attr,omitempty" json:"creation_date,omitempty"`
	StartDate             string `xml:"startDate,attr" json:"start_date"`
	EndDate               string `xml:"endDate,attr" json:"end_date"`
	// Elements
	MetadataEntries []MetadataEntry `xml:"MetadataEntry,omitempty" json:"metadata,omitempty"`
	Events          []WorkoutEvent  `xml:"WorkoutEvent,omitempty" json:"events,omitempty"`
	Routes          []WorkoutRoute  `xml:"WorkoutRoute,omitempty" json:"routes,omitempty"`
}

// ActivitySummary daily work summary.
type ActivitySummary struct {
	XMLName xml.Name `xml:"ActivitySummary" json:"-"`
	// Attributes
	YYYYMMDD         string `xml:"dateComponents,attr,omitempty" json:"yyyymmdd,omitempty"`
	EnergyBurned     string `xml:"activeEnergyBurned,attr,omitempty" json:"energy_burned,omitempty"`
	EnergyBurnedGoal string `xml:"activeEnergyBurnedGoal,attr,omitempty" json:"energy_burned_goal,omitempty"`
	EnergyBurnedUnit string `xml:"activeEnergyBurnedUnit,attr,omitempty" json:"energy_burned_unit,omitempty"`
	MoveTime         string `xml:"appleMoveTime,attr,omitempty" json:"move_time,omitempty"`
	MoveTimeGoal     string `xml:"appleMoveTimeGoal,attr,omitempty" json:"move_time_goal,omitempty"`
	ExerciseTime     string `xml:"appleExerciseTime,attr,omitempty" json:"exercise_time,omitempty"`
	ExerciseTimeGoal string `xml:"appleExerciseTimeGoal,attr,omitempty" json:"exercise_time_goal,omitempty"`
	StandHours       string `xml:"appleStandHours,attr,omitempty" json:"stand_hours,omitempty"`
	StandHoursGoal   string `xml:"appleStandHoursGoal,attr,omitempty" json:"stand_hours_goal,omitempty"`
}

// HealthData is composite of exports.xml
type HealthData struct {
	// Attribute
	Locale string `xml:"locale,attr" json:"locate"`
	// Elements
	Exported          Date               `xml:"ExportDate" json:"exported_date"`
	Me                Me                 `xml:"Me" json:"me"`
	Records           []Record           `xml:"Record"  json:"records,omitempty"`
	ClinicalRecords   []ClinicalRecord   `xml:"ClinicalRecord" json:"clinical_records,omitempty"`
	Correlations      []Correlation      `xml:"Correlation" json:"correlations,omitempty"`
	Workouts          []Workout          `xml:"Workout" json:"workouts,omitempty"`
	ActivitiesSummary []ActivitySummary  `xml:"ActivitySummary" json:"activities_summary,omitempty"`
	Audiograms        []Audiogram        `xml:"Audiogram,omitempty,omitempty"  json:"audiograms,omitempty"`
	SensitivityPoints []SensitivityPoint `xml:"SensitivityPoint,omitempty" json:"sensitivity_points,omitempty"`
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
	h.DescribeRecordsSourceName()
	h.DescribeRecordsTypesBySourceName()
	h.DescribeRecordsSourceNameByTypes()
}

// DescribeRecordsSourceName Metadata Keys. Sorts keys by alphanumeric
func (h *HealthData) DescribeRecordsSourceName() {
	// Get all Metadata keys and stores them in a map whose key is SourceName
	// [SourceName] = Map {"key-1": 5, "key-2": 10}
	md := make(map[string]map[string]int)
	sm := make(map[string]int) // SourceName count of r.MetadataEntries map
	for _, r := range h.Records {
		if _, ok := md[r.SourceName]; !ok {
			md[r.SourceName] = make(map[string]int) // Create map if a new key is seen
		}
		// Count all occurrences of a "key-1"
		if _, ok := sm[r.SourceName]; !ok {
			sm[r.SourceName] = len(r.MetadataEntries)
		} else {
			sm[r.SourceName] += len(r.MetadataEntries)
		}
		// Create map of r.MetadataEntries
		for _, m := range r.MetadataEntries {
			if _, ok := md[r.SourceName][m.Key]; !ok {
				md[r.SourceName][m.Key] = 1
			} else {
				md[r.SourceName][m.Key]++
			}
		}
	}
	// Convert sm to []KeyCount
	// snmc: SourceName metadata entries count: # of 'r.MetadataEntries' for 'SourceName'
	snmc := make(KeyCounts, 0, len(sm))
	for k, v := range sm {
		snmc = append(snmc, KeyCount{Key: k, Count: v})
	}
	// Sort 'SourceName' that  has highest count of 'MetadataEntries' first
	sort.Sort(snmc)
	for _, s := range snmc {
		if s.Count > 0 {
			h := fmt.Sprintf("MetadataEntries (%d) recorded by %q", s.Count, s.Key)
			printStringInts(h, md[s.Key])
		}
	}
}

// DescribeRecordsTypesBySourceName lists all types reported by a source
func (h *HealthData) DescribeRecordsTypesBySourceName() {
	if h == nil {
		return
	}
	// Record SourceName Type Map of ["MyFitnessPal"] = Map {"type-1": 10, "type-2": 12 }
	rstm := make(map[string]map[string]int)
	rs := make(map[string]int) // How many types w/ each SourceName
	for _, r := range h.Records {
		if _, ok := rstm[r.SourceName]; !ok { // Create map for a unseen type
			rstm[r.SourceName] = make(map[string]int)
		}
		if _, ok := rs[r.SourceName]; !ok {
			rs[r.SourceName] = 1
		} else {
			rs[r.SourceName]++ // Another SourceName seen
		}
		if _, ok := rstm[r.SourceName][r.Type]; !ok { // Add new type seen
			rstm[r.SourceName][r.Type] = 1
		} else {
			rstm[r.SourceName][r.Type]++
		}
	}
	// Sort SourceName with highest types first
	snc := make(KeyCounts, 0, len(rs))
	for k, v := range rs {
		snc = append(snc, KeyCount{Key: k, Count: v})
	}
	sort.Sort(snc)
	for _, s := range snc {
		if s.Count > 0 {
			h := fmt.Sprintf("Types (%d) by SourceName %q", s.Count, s.Key)
			printStringInts(h, rstm[s.Key])
		}
	}
}

// DescribeRecordsSourceNameByTypes lists all Sources that provided a specific type.
func (h *HealthData) DescribeRecordsSourceNameByTypes() {
	if h == nil {
		return
	}
	// Record Type SourceName Map of ["BodyMass"] = Map {"MyFitnessPal": 10, "Renpho": 12 }
	rtsm := make(map[string]map[string]int)
	rt := make(map[string]int) // How many types w/ each Type
	for _, r := range h.Records {
		if _, ok := rtsm[r.Type]; !ok { // Create map for a unseen type
			rtsm[r.Type] = make(map[string]int)
		}
		if _, ok := rt[r.Type]; !ok {
			rt[r.Type] = 1
		} else {
			rt[r.Type]++ // Another Type seen
		}
		if _, ok := rtsm[r.Type][r.SourceName]; !ok { // Add new SourceName seen
			rtsm[r.Type][r.SourceName] = 1
		} else {
			rtsm[r.Type][r.SourceName]++
		}
	}
	// Sort SourceName with highest types first
	snc := make(KeyCounts, 0, len(rt))
	for k, v := range rt {
		snc = append(snc, KeyCount{Key: k, Count: v})
	}
	sort.Sort(snc)
	for _, s := range snc {
		if s.Count > 0 {
			h := fmt.Sprintf("SourceName (%d) by Type %q", len(rtsm[s.Key]), shortenHKKey(s.Key))
			printStringInts(h, rtsm[s.Key])
		}
	}
}

// DescribeRecordsTable writes a summary to a ioWrite, in table columns as
// Type .. device .. metadata tag .. count
func (h *HealthData) DescribeRecordsTable(w io.Writer, sep string) {
	if h == nil {
		return
	}
	// cm, combined type map for Type -> SourceName -> MetaDataEntry -> Count
	cm := make(map[string]map[string]map[string]int)
	tm := make(map[string]int)             // [Type] map for  sorting by Highest to Lowest
	tsm := make(map[string]map[string]int) // [Type][SourceName] map for  sorting by Highest to Lowest
	for _, r := range h.Records {
		t, s := r.Type, r.SourceName
		if _, ok := cm[t]; !ok { // Is cm['Type'] entry not seen
			cm[r.Type] = make(map[string]map[string]int) // Create a  map
		}
		if _, ok := cm[t][s]; !ok { // Is cm['Type']['sourceName'] entry not seen
			cm[t][s] = make(map[string]int) // Create a map
		}
		if _, ok := tsm[t]; !ok {
			tsm[t] = make(map[string]int)
		}
		if _, ok := cm[t][s]; !ok { // Is key seen before
			cm[t][s] = make(map[string]int)
		}
		// Go through Metadata Entries
		if len(r.MetadataEntries) == 0 {
			tm[t]++
			tsm[t][s]++
			cm[t][s]["-"]++
			continue
		}
		for _, md := range r.MetadataEntries {
			k := md.Key
			cm[t][s][k]++
			tm[t]++
			tsm[t][s]++
		}
	}

	tms, err := maps2SortedSlice(tm) // Sort from highest count of MDKeys to lowest #
	if err != nil {
		return
	}

	for _, v := range tms {
		tsms, err := maps2SortedSlice(tsm[v.Key])
		if err != nil {
			continue
		}
		for _, s := range tsms {
			if len(cm[v.Key][s.Key]) == 0 { //  No MetadataEntry with this [t][s]
				_, _ = fmt.Fprintf(w, "%s%s%s%s%s%s%d\n", shortenHKKey(v.Key), sep, shortenHKKey(s.Key), sep, "-", sep, s.Count)
				continue
			}
			writeSortedRows(w, v.Key, s.Key, cm[v.Key][s.Key], sep)
		}
	}
}
