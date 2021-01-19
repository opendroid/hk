package server

import (
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
	"html/template"
	"net/http"
)

// vc are View Controllers for the handlers
// Save the health data in a map
var (
	html  *template.Template
	users map[string]userData
	file  string
)

// init initializes the templates
func init() {
	html = template.Must(template.ParseGlob(htmlGlob))
	users = make(map[string]userData)
}

// Index page handler
func index(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(sessionIDKey).(string)
	if !ok {
		logger.Error("User not found", zap.String("page", "index"))
		displayErrorMessage(w, umNoData)
		return
	}
	logger.Debug("User found", zap.String("user", user))
	go processHealthData(user) // Do it in parallel
	index := getIndexPageData(user)
	if err := html.ExecuteTemplate(w, "index.gohtml", index); err != nil {
		logger.Error(err.Error(), zap.String("page", "index.gohtml"))
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}

// records page handler
func records(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(sessionIDKey).(string)
	if !ok {
		logger.Error("User not found", zap.String("page", "records"))
		displayErrorMessage(w, umNotProcessed)
		return
	}
	// Get user data
	records := getRecordsPageData(user)
	if records == nil {
		displayErrorMessage(w, umNotProcessed)
		return
	}
	if err := html.ExecuteTemplate(w, "records.gohtml", records); err != nil {
		logger.Error(err.Error(), zap.String("page", "records.gohtml"))
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}

// errorsPage page handler
func errorsPage(w http.ResponseWriter, _ *http.Request) {
	displayErrorMessage(w, wmNoError)
}

// displayErrorMessage helper to display a error message
func displayErrorMessage(w http.ResponseWriter, info string) {
	data := getErrorPageData(info)

	if err := html.ExecuteTemplate(w, "error.gohtml", data); err != nil {
		logger.Error(err.Error(), zap.String("page", "error.gohtml"), zap.String("info", info))
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}
