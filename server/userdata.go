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
	auth          *gUserInfo // gUserInfo for this user
	health        *export.HealthData
	summary       export.NameTypeKeyCounts
	sources       export.KeyCounts
	recordTypes   export.KeyCounts
	mass          *export.BodyMassComposition
	audioExposure *export.AudioExposure
	walkData      *export.WalkerData
}

// users main datastore for users data
var (
	users map[string]userData
)

func init() {
	users = make(map[string]userData)
}

// processHealthData parses the user data from file and update it in users data table
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
		wd := health.WalkerData()
		updateHealthData(uID, &userData{
			health:        health,
			summary:       summary,
			sources:       sources,
			recordTypes:   recordTypes,
			mass:          mass,
			audioExposure: ae,
			walkData:      wd,
		})
	}
	end := time.Since(start)
	logger.Debug("Analysis done", zap.String("user", uID), zap.Int64("ms", end.Milliseconds()))
}

// updateAuth adds a auth object to users[uID] if not available
func updateAuth(uID string, auth *gUserInfo) {
	u, ok := users[uID]
	if !ok {
		u = userData{auth: auth}
	} else {
		u.auth = auth
	}
	users[uID] = u
}

// updateHealthData in users[uID] while keeping users[uID].auth intact.
func updateHealthData(uID string, h *userData) {
	u, ok := users[uID]
	if !ok {
		u = *h
	} else {
		a := u.auth // Save auth
		u = *h
		u.auth = a
	}
	users[uID] = u
}

// isAuthenticated returns if user has been successfully authenticated with Google auth.
func isAuthenticated(uID string) bool {
	if uID == "" {
		return false
	}
	if u, ok := users[uID]; ok && u.auth != nil {
		return u.auth.Authenticated
	}
	return false
}