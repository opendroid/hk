package export

import (
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
	"sort"
	"strconv"
	"strings"
)

// WalkingDataElement for daily StepCount and DistanceWalkingRunning
type WalkingDataElement struct {
	YYYYMMDD   string  `json:"yyyymmdd"` // CreationDate truncated to CreationDate in format "YYYY-MM-DD"
	SourceName string  `json:"source"`
	Unit       string  `json:"unit,omitempty"`
	Value      float32 `json:"value"`
	Count      int     `json:"count"`
}

// WalkerData holds step count and walking orr running distance
type WalkerData struct {
	StepCount              []WalkingDataElement `json:"step_count"`
	DistanceWalkingRunning []WalkingDataElement `json:"distance_walking_running"`
}

// Get Walker data and send back
func (h *HealthData) WalkerData() *WalkerData {
	if h == nil {
		return nil
	}

	// Save data as sum by date
	sc := make(map[string]WalkingDataElement)
	dwr := make(map[string]WalkingDataElement)
	for _, r := range h.Records {
		if RecordType(r.Type) == StepCount {
			saveWalkElementValue(sc, &r)
		} else if RecordType(r.Type) == DistanceWalkingRunning {
			saveWalkElementValue(dwr, &r)
		}
	}
	var wd WalkerData                            // Prepare data
	wd.StepCount = make([]WalkingDataElement, 0) // Create underlying data slice for Steps
	wd.DistanceWalkingRunning = make([]WalkingDataElement, 0)
	wd.StepCount = mapToWalkingDataElementSlice(wd.StepCount, sc)
	wd.DistanceWalkingRunning = mapToWalkingDataElementSlice(wd.DistanceWalkingRunning, dwr)
	return &wd
}

// saveWalkElementValue adds a record 'r' to map[yyyymmdd]
func saveWalkElementValue(m map[string]WalkingDataElement, r *Record) {
	date := strings.Split(r.StartDate, " ") // Extract "2019-10-23" from "2019-10-23 19:10:11 -0800"
	if len(date) == 0 {
		logger.Warn("Illegal date", zap.String("method", "saveStepCount"), zap.String("date", r.StartDate))
		return
	}
	d := date[0]                                // Valid date extracted
	num, err := strconv.ParseFloat(r.Value, 32) // Convert value to a number
	if err != nil {
		logger.Error("strconv error",
			zap.String("method", "saveStepCount"),
			zap.String("value", r.Value), zap.String("info", err.Error()))
	}
	// Add 'r' record to m["2019-10-23"]
	v, ok := m[d]
	if !ok { // New value for the day
		m[d] = WalkingDataElement{d, r.SourceName, r.Unit, float32(num), 1}
		return
	}
	// Update collected value so far
	if v.Unit != r.Unit { // For now log and continue
		logger.Warn("Non matching unit", zap.String("method", "saveStepCount"),
			zap.String("found", r.Unit), zap.String("new", v.Unit))
	}
	if !strings.Contains(v.SourceName, r.SourceName) {
		v.SourceName += ", " + r.SourceName // For now append all unique sources
	}
	v.Value += float32(num)
	v.Count++ // How many such counts for a day
	m[d] = v  // Update value
}

func mapToWalkingDataElementSlice(wd []WalkingDataElement, m map[string]WalkingDataElement) []WalkingDataElement {
	for _, v := range m {
		wd = append(wd, v)
	}
	sort.Slice(wd, func(i, j int) bool { return wd[i].YYYYMMDD < wd[j].YYYYMMDD })
	return wd
}
