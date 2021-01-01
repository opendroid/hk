package export

import (
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
		t.Logf("%v", err)
	})

	t.Run("File exists", func(t *testing.T) {
		file := "./sampleexport.xml"
		d, err := Parse(file)
		if err != nil {
			t.Fail()
			return
		}
		d.Describe()
	})

}

// ExampleParse reading data from Apple Export file
func ExampleParse() {
	file := "./sampleexport.xml"
	d, err := Parse(file)
	if err != nil {
		return
	}
	d.Describe()
	// Output:
	// Export date: 2020-12-06
	// Me: DOB:1990-12-06, Sex:Male, BloodType:NotSet, SkinType:NotSet, Medication:None
	// Records: 5, #Metadata: 0, #HRV: 1
	// Clinical: 6
	// Correlations: 0, #Metadata: 0, #Records: 0
	// Workouts: 1, #Metadata: 1, #Events: 8, Routes: 1
	// ActivitiesSummary: 17
}
