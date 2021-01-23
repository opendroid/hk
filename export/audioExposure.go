package export

import (
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
	"sort"
	"strconv"
	"time"
)

// Exposure elements for exposure
type Exposure struct {
	CreationDate int64   `json:"creation_timestamp_sec"`
	EndDate      int64   `json:"end_date"`
	SourceName   string  `json:"source"`
	Unit         string  `json:"unit,omitempty"`
	Value        float32 `json:"value"`
}

// AudioExposure slice of all data
type AudioExposure struct {
	Exposure []Exposure `json:"exposure"`
}

// AudioExposure processes and returns Audio exposure data
func (h *HealthData) AudioExposure() *AudioExposure {
	if h == nil {
		return nil
	}
	var d AudioExposure
	d.Exposure = make([]Exposure, 0)
	for _, r := range h.Records {
		if RecordType(r.Type) == EnvironmentalAudioExposure {
			num, err := strconv.ParseFloat(r.Value, 32)
			if err != nil {
				logger.Error("appendData: strconv error",
					zap.String("method", "AudioExposure"),
					zap.String("value", r.Value), zap.String("info", err.Error()))
			}
			n := float32(num)
			e1 := Exposure{recordTime(r.CreationDate), recordTime(r.EndDate),
				r.SourceName, r.Unit, n}
			d.Exposure = append(d.Exposure, e1)
		}
	}
	// Sort data first
	sort.Slice(d.Exposure, func(i, j int) bool { return d.Exposure[i].CreationDate < d.Exposure[j].CreationDate })
	return &d
}

// recordTime converts a Record format to time object.
// Safer to convert timestamp to Unix MS so it can be manipulated by various browsers.
// eg: https://stackoverflow.com/questions/60308048/why-chart-js-charts-are-not-plotting-data-in-safari-works-in-chrome
func recordTime(ts string) int64 {
	// Sample"creation_date": "2019-10-23 19:10:11 -0800" from Health Export
	format := "2006-01-02 15:04:05 -0700"
	if t, err := time.Parse(format, ts); err != nil {
		logger.Error("Bad timestamp", zap.String("method", "recordTime"),
			zap.String("value", ts), zap.String("info", err.Error()))
		return 0
	} else {
		return t.Unix()
	}
}
