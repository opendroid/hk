package export

import (
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
	"sort"
	"strconv"
)

// BodyMassElement elements in various types of body-mass data
type BodyMassElement struct {
	CreationDate int64   `json:"creation_timestamp_sec"`
	SourceName   string  `json:"source"`
	Unit         string  `json:"unit,omitempty"`
	Value        float32 `json:"value"`
}

// BodyMassComposition data slices related to body mass.
type BodyMassComposition struct {
	Mass          []BodyMassElement `json:"mass"`
	LeanMass      []BodyMassElement `json:"lean_body_mass"`
	BMI           []BodyMassElement `json:"bmi"`
	FatPercentage []BodyMassElement `json:"fat_percent"`
}

// BodyMassData get the body mass data
func (h *HealthData) BodyMassData() *BodyMassComposition {
	if h == nil {
		return nil
	}
	var m BodyMassComposition
	m.Mass = make([]BodyMassElement, 0)
	m.BMI = make([]BodyMassElement, 0)
	m.FatPercentage = make([]BodyMassElement, 0)
	for _, d := range h.Records {
		switch RecordType(d.Type) {
		case BodyMass:
			m.Mass = appendMassData(m.Mass, d)
		case BodyMassIndex:
			m.BMI = appendMassData(m.BMI, d)
		case BodyFatPercentage:
			m.FatPercentage = appendMassData(m.FatPercentage, d)
		case LeanBodyMass:
			m.LeanMass = appendMassData(m.LeanMass, d)
		}
	}
	// Sort by creation date
	sort.Slice(m.Mass, func(i, j int) bool { return m.Mass[i].CreationDate < m.Mass[j].CreationDate })
	sort.Slice(m.LeanMass, func(i, j int) bool { return m.LeanMass[i].CreationDate < m.LeanMass[j].CreationDate })
	sort.Slice(m.BMI, func(i, j int) bool { return m.BMI[i].CreationDate < m.BMI[j].CreationDate })
	sort.Slice(m.FatPercentage, func(i, j int) bool { return m.FatPercentage[i].CreationDate < m.FatPercentage[j].CreationDate })
	return &m
}

// appendMassData helper appends relevant components slice
func appendMassData(d []BodyMassElement, r Record) []BodyMassElement {
	num, err := strconv.ParseFloat(r.Value, 32)
	if err != nil {
		logger.Error("appendData: strconv error",
			zap.String("method", "appendMassData"),
			zap.String("value", r.Value), zap.String("info", err.Error()))
	}
	e := BodyMassElement{
		CreationDate: recordTime(r.CreationDate),
		SourceName:   r.SourceName,
		Unit:         r.Unit,
		Value:        float32(num),
	}
	return append(d, e)
}
