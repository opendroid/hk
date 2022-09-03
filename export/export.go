// Package export provided parsing and analysis methods for Apple Health Exports Data.
package export

import (
	"encoding/xml"
	"os"
)

// Parse Apple HK exported XML fileName and returns pointer to HealthData, or error if any
func Parse(fileName string) (*HealthData, error) {
	// Read XML file
	health, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var data HealthData
	if err = xml.Unmarshal(health, &data); err != nil {
		return nil, err
	}
	return &data, nil
}
