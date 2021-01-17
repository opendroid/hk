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
		// Push out a funny error page
	}
	logger.Debug("User found", zap.String("user", user))
	go saveHD(user) // Do it in parallel
	index := getIndexPageData(user)
	if err := html.ExecuteTemplate(w, "index.gohtml", index); err != nil {
		logger.Error(err.Error(), zap.String("page", "index.gohtml"))
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}

// records page handler
func records(w http.ResponseWriter, r *http.Request) {
	records := getRecordsPageData()

	if err := html.ExecuteTemplate(w, "records.gohtml", records); err != nil {
		logger.Error(err.Error(), zap.String("page", "records.gohtml"))
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}
