package server

import (
	"fmt"
	"github.com/opendroid/hk/export"
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
	"time"
)

// userData saves health data summary for the user.
type userData struct {
	health        *export.HealthData
	summary       export.NameTypeKeyCounts
	sources       export.KeyCounts
	recordTypes   export.KeyCounts
	mass          *export.BodyMassComposition
	audioExposure *export.AudioExposure
}

func processHealthData(uID string) {
	// Read XML file
	start := time.Now()
	if health, err := export.Parse(file); err != nil {
		fmt.Printf("Error parsing %q, %v\n", file, err)
	} else {
		summary := health.RecordsSummary() // Pre-process data
		sources := health.RecordsSources()
		recordTypes := health.RecordsTypes()
		mass := health.BodyMassData()
		ae := health.AudioExposure()
		users[uID] = userData{
			health:        health,
			summary:       summary,
			sources:       sources,
			recordTypes:   recordTypes,
			mass:          mass,
			audioExposure: ae,
		}
	}
	end := time.Since(start)
	logger.Debug("Analysis done", zap.String("user", uID), zap.Int64("ms", end.Milliseconds()))
}
