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

// ActivitySummary daily work summary.
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
			h := fmt.Sprintf("SourceName (%d) by Type %q", len(rtsm[s.Key]), s.Key)
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
