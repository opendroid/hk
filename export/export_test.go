package export

import (
	"fmt"
	"testing"
)

// TestParse tests the Parse
// go test -run TestParse -v
func TestParse(t *testing.T) {
	t.Run("No file", func(t *testing.T) {
		file := "./whatever.xml"
		_, err := Parse(file)
		if err == nil {
			t.Fail()
		}
		t.Logf("Expected: %v", err)
	})

	t.Run("File exists", func(t *testing.T) {
		file := "../public/sampleexport.xml"
		d, err := Parse(file)
		if err != nil {
			t.Fail()
			return
		}
		if d == nil {
			t.Fail()
			return
		}
	})

}

// Parse reads and parses data from Apple Export file. Describe summary of data.
func ExampleParse() {
	file := "../public/sampleexport.xml"
	d, err := Parse(file)
	if err != nil {
		return
	}
	d.Describe()
	// Output:
	// Export date: 2020-12-06
	// Me: DOB:1990-12-06, Sex:Male, BloodType:NotSet, SkinType:NotSet, Medication:None
	// Records: 5, #Metadata: 0, #HRV: 1
	// Records by Type:
	// 	BodyMassIndex: 4
	// 	HeartRateVariabilitySDNN: 1
	// Total Count: 5
	// Records by SourceName:
	// 	Renpho: 4
	// 	🕛🕐🕑🕒: 1
	// Total Count: 5
	// Records Metadata Entries SourceName:
	// 	Renpho: 0
	// 	🕛🕐🕑🕒: 0
	// Total Count: 0
	// Types (4) by SourceName "Renpho"
	// 	BodyMassIndex: 4
	// Total Count: 4
	// Types (1) by SourceName "🕛🕐🕑🕒"
	// 	HeartRateVariabilitySDNN: 1
	// Total Count: 1
	// SourceName (1) by Type "BodyMassIndex"
	// 	Renpho: 4
	// Total Count: 4
	// SourceName (1) by Type "HeartRateVariabilitySDNN"
	// 	🕛🕐🕑🕒: 1
	// Total Count: 1
	// Clinical: 6
	// Correlations: 0, #Metadata: 0, #Records: 0
	// Workouts: 1, #Metadata: 1, #Events: 8, Routes: 1
	// ActivitiesSummary: 17
}

// ActivitySummary String provides description of Activity Summary
func ExampleActivitySummary() {
	file := "../public/sampleexport.xml"
	d, err := Parse(file)
	if err != nil {
		return
	}
	for _, a := range d.ActivitiesSummary {
		fmt.Printf("%s\n", a)
	}
	// Output:
	// ActivitySummary:2020-11-20 (Energy: 487.231/300 Cal, Move: 0/0, Exercise: 74/30 Minutes, Stand: 16/11 Hours)
	// ActivitySummary:2020-11-21 (Energy: 551.951/300 Cal, Move: 0/0, Exercise: 79/30 Minutes, Stand: 16/11 Hours)
	// ActivitySummary:2020-11-22 (Energy: 534.026/300 Cal, Move: 0/0, Exercise: 80/30 Minutes, Stand: 18/11 Hours)
	// ActivitySummary:2020-11-23 (Energy: 585.547/300 Cal, Move: 0/0, Exercise: 83/30 Minutes, Stand: 17/11 Hours)
	// ActivitySummary:2020-11-24 (Energy: 604.823/300 Cal, Move: 0/0, Exercise: 83/30 Minutes, Stand: 19/11 Hours)
	// ActivitySummary:2020-11-25 (Energy: 606.726/300 Cal, Move: 0/0, Exercise: 85/30 Minutes, Stand: 18/11 Hours)
	// ActivitySummary:2020-11-26 (Energy: 599.972/300 Cal, Move: 0/0, Exercise: 88/30 Minutes, Stand: 18/11 Hours)
	// ActivitySummary:2020-11-27 (Energy: 648.394/300 Cal, Move: 0/0, Exercise: 100/30 Minutes, Stand: 18/11 Hours)
	// ActivitySummary:2020-11-28 (Energy: 758.974/300 Cal, Move: 0/0, Exercise: 141/30 Minutes, Stand: 18/11 Hours)
	// ActivitySummary:2020-11-29 (Energy: 585.901/300 Cal, Move: 0/0, Exercise: 102/30 Minutes, Stand: 19/11 Hours)
	// ActivitySummary:2020-11-30 (Energy: 695.816/300 Cal, Move: 0/0, Exercise: 109/30 Minutes, Stand: 19/11 Hours)
	// ActivitySummary:2020-12-01 (Energy: 655.492/300 Cal, Move: 0/0, Exercise: 95/30 Minutes, Stand: 16/11 Hours)
	// ActivitySummary:2020-12-02 (Energy: 634.471/300 Cal, Move: 0/0, Exercise: 90/30 Minutes, Stand: 18/11 Hours)
	// ActivitySummary:2020-12-03 (Energy: 659.84/300 Cal, Move: 0/0, Exercise: 96/30 Minutes, Stand: 17/11 Hours)
	// ActivitySummary:2020-12-04 (Energy: 647.78/300 Cal, Move: 0/0, Exercise: 93/30 Minutes, Stand: 18/11 Hours)
	// ActivitySummary:2020-12-05 (Energy: 583.725/300 Cal, Move: 0/0, Exercise: 91/30 Minutes, Stand: 16/11 Hours)
	// ActivitySummary:2020-12-06 (Energy: 610.788/300 Cal, Move: 0/0, Exercise: 113/30 Minutes, Stand: 14/11 Hours)
}
