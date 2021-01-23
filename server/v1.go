package server

import (
	"compress/gzip"
	"encoding/json"
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
	"net/http"
	"strings"
)

type apiError struct {
	Text   string `json:"text"`
	Status int    `json:"status"`
}

// jsonErrorResponse returns string JSON
func jsonErrorResponse(w http.ResponseWriter, msg string, status int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	e, _ := json.Marshal(apiError{msg, status})
	_, _ = w.Write(e)
}

// recordsData API return a JSON body for a records category of data. One of "Sources", "Types" or "All"
func recordsData(cat RecordsDataCategory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		user, ok := r.Context().Value(sessionIDKey).(string)
		if !ok {
			logger.Error("User not found", zap.String("method", "recordsData"), zap.String("cat", string(cat)))
			jsonErrorResponse(w, "User not found", http.StatusBadRequest)
			return
		}
		// Get user data
		u, ok := users[user]
		if !ok {
			logger.Error("No user data", zap.String("user", user),
				zap.String("method", "recordsData"), zap.String("cat", string(cat)))
			jsonErrorResponse(w, "No user data", http.StatusBadRequest)
			return
		}
		ud, err := getRecordCategoryData(u, cat)
		if err != nil {
			logger.Error("Unmarshal error", zap.String("info", err.Error()),
				zap.String("method", "recordsData"), zap.String("cat", string(cat)))
			jsonErrorResponse(w, "No user data", http.StatusBadRequest)
			return
		}
		w.Header().Set("Cache-Control", "private, max-age:120")
		ae := r.Header.Get("Accept-Encoding")
		if strings.Contains(ae, "gzip") {
			if err := gzipResponse(w, ud); err != nil {
				logger.Error("gzip write error", zap.String("info", err.Error()),
					zap.String("method", "recordsData"), zap.String("cat", string(cat)))
			} else {
				return // No error all good
			}
		}
		// Send uncompressed response
		if _, err := w.Write(ud); err != nil {
			logger.Error("ResponseWriter error", zap.String("info", err.Error()),
				zap.String("method", "recordsData"), zap.String("cat", string(cat)))
		}
	}
}

// gzipResponse applies gzip compression to d, returns error if it fails
func gzipResponse(w http.ResponseWriter, d []byte) error {
	w.Header().Set("Content-Encoding", "gzip")
	gz := gzip.NewWriter(w)
	if _, err := gz.Write(d); err != nil {
		w.Header().Set("Content-Encoding", "identity") // Adjust content encoding
		return err
	}
	if err := gz.Close(); err != nil { // Reset coding
		w.Header().Set("Content-Encoding", "identity") // Adjust content encoding
		return err
	}
	return nil
}

// getRecordCategoryData returns Marshal for Source, Type or All data.
func getRecordCategoryData(u userData, cat RecordsDataCategory) ([]byte, error) {
	switch cat {
	case recordsSource:
		return json.Marshal(u.sources)
	case recordsTypes:
		return json.Marshal(u.recordTypes)
	case activitySummary:
		return json.Marshal(u.health.Summary())
	case bodyMass:
		return json.Marshal(u.mass)
	case exposure:
		return json.Marshal(u.audioExposure)
	case recordsAll:
		fallthrough
	default:
		return json.Marshal(u.summary)
	}
}
