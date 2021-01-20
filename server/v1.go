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

// errorResponse returns string JSON
func errorResponse(msg string, status int) string {
	e := apiError{msg, status}
	if es, err := json.Marshal(e); err != nil {
		return ""
	} else {
		return string(es)
	}
}

// recordsSummary page handler
func recordsSummary(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, ok := r.Context().Value(sessionIDKey).(string)
	if !ok {
		logger.Error("User not found", zap.String("json", "records"),
			zap.String("method", "recordsSummary"))
		e := errorResponse("User not found", http.StatusBadRequest)
		http.Error(w, e, http.StatusBadRequest)
		return
	}
	// Get user data
	u, ok := users[user]
	if !ok {
		logger.Error("No user data", zap.String("user", user),
			zap.String("method", "recordsSummary"))
		e := errorResponse("No user data", http.StatusBadRequest)
		http.Error(w, e, http.StatusBadRequest)
		return
	}
	ud, err := json.Marshal(u.sources)
	if err != nil {
		logger.Error("Unmarshal error", zap.String("info", err.Error()),
			zap.String("method", "recordsSummary"))
		return
	}

	if _, err := w.Write(ud); err != nil {
		logger.Error("ResponseWriter error", zap.String("info", err.Error()),
			zap.String("method", "recordsSummary"))
	}
}

// recordsTypes page handler
func recordsTypes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	user, ok := r.Context().Value(sessionIDKey).(string)
	if !ok {
		logger.Error("User not found", zap.String("json", "recordsTypes"),
			zap.String("method", "recordsTypes"))
		e := errorResponse("User not found", http.StatusBadRequest)
		http.Error(w, e, http.StatusBadRequest)
		return
	}
	// Get user data
	u, ok := users[user]
	if !ok {
		logger.Error("No user data", zap.String("user", user),
			zap.String("method", "recordsTypes"))
		e := errorResponse("No user data", http.StatusBadRequest)
		http.Error(w, e, http.StatusBadRequest)
		return
	}
	ud, err := json.Marshal(u.recordTypes)
	if err != nil {
		logger.Error("Unmarshal error", zap.String("info", err.Error()),
			zap.String("method", "recordsTypes"))
		return
	}

	if _, err := w.Write(ud); err != nil {
		logger.Error("ResponseWriter error", zap.String("info", err.Error()),
			zap.String("method", "recordsTypes"))
	}
}
