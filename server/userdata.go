package server

import (
	"compress/gzip"
	"encoding/gob"
	"github.com/opendroid/hk/export"
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
	"os"
	"time"
)

// userData saves health data summary for the user.
type userData struct {
	auth          *GoogleUserInfo // GoogleUserInfo for this user
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

// processHealthData parses the user data from xmlHDFile and update it in users data table
func processHealthData(uID string) {

	if tempFileName == "" { // create temp file name
		f, err := os.CreateTemp(tempDir, tempFilePattern)
		if err != nil { // TODO: what we should do here?
			logger.Error("error creating temp file", zap.String("user", uID), zap.String("error", err.Error()))
			return
		}
		defer func() { _ = f.Close() }()
		tempFileName = f.Name()
		start := time.Now()
		health, err := export.Parse(xmlHDFile)
		if err != nil {
			logger.Info("Error parsing xmlHDFile", zap.String("user", uID), zap.String("error", err.Error()))
			return
		}
		end := time.Since(start)
		logger.Debug("Analysis done", zap.String("user", uID), zap.Int64("ms", end.Milliseconds()))
		saveHDtoGZ(health, tempFileName, uID)
	} else {
		logger.Debug("using cache", zap.String("user", uID))
	}
	health, err := readHDFromFile(tempFileName, uID)
	if err != nil {
		return
	}

	// Save data to memory
	summary := health.RecordsSummary()
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

// updateAuth adds an auth object to users[uID] if not available
func updateAuth(uID string, auth *GoogleUserInfo) {
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

// saveHDtoGZ saves the health data to a xmlHDFile. See example: https://go.dev/play/p/8g1CJEgVVCw
func saveHDtoGZ(hd *export.HealthData, fName, uID string) {
	start := time.Now()
	gout, err := os.Create(fName)
	if err != nil {
		logger.Error("GOB Create", zap.String("user", uID), zap.String("error", err.Error()), zap.String("file", fName))
		return
	}
	defer func() { _ = gout.Close() }()
	gz := gzip.NewWriter(gout)
	defer func() { _ = gz.Close() }()
	enc := gob.NewEncoder(gz)
	if err := enc.Encode(hd); err != nil {
		logger.Error("GOB Write", zap.String("user", uID), zap.String("error", err.Error()), zap.String("file", fName))
		return
	}
	end := time.Since(start)
	logger.Debug("saveHDtoGZ", zap.String("user", uID), zap.Int64("ms", end.Milliseconds()), zap.String("file", fName))
}

// readHDFromFile read data from a xmlHDFile
func readHDFromFile(file, uID string) (*export.HealthData, error) {
	// Read from .gz
	start := time.Now()
	fr, err := os.Open(file)
	defer func() { _ = fr.Close() }()
	if err != nil {
		logger.Info("readHDFromFile GOB Open", zap.String("user", uID), zap.String("error", err.Error()))
		return nil, err
	}
	gz, err := gzip.NewReader(fr)
	defer func() { _ = gz.Close() }()
	if err != nil {
		logger.Info("readHDFromFile GOB Read", zap.String("user", uID), zap.String("error", err.Error()))
		return nil, err
	}
	g := gob.NewDecoder(gz)
	var tmpHD export.HealthData
	if err := g.Decode(&tmpHD); err != nil {
		logger.Info("readHDFromFile GOB Decode", zap.String("user", uID), zap.String("error", err.Error()))
		return nil, err
	}
	end := time.Since(start)
	logger.Debug("readHDFromFile", zap.String("user", uID), zap.Int64("ms", end.Milliseconds()))
	return &tmpHD, nil
}
