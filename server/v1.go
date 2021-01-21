package server

import (
	"encoding/json"
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
	"net/http"
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
			return
		}

		if _, err := w.Write(ud); err != nil {
			logger.Error("ResponseWriter error", zap.String("info", err.Error()),
				zap.String("method", "recordsData"), zap.String("cat", string(cat)))
		}
	}
}

// getRecordCategoryData returns Marshal for Source, Type or All data.
func getRecordCategoryData(u userData, cat RecordsDataCategory) ([]byte, error){
	switch cat {
	case RecordsSource:
		return json.Marshal(u.sources)
	case RecordsTypes:
		return json.Marshal(u.recordTypes)
	case RecordsAll:
		fallthrough
	default:
		return json.Marshal(u.summary)
	}
}