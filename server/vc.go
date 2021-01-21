package server

import (
	"fmt"
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
	"html/template"
	"math"
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
	helpers := template.FuncMap{"numWithComma": numWithComma}
	html = template.Must(template.New(root).Funcs(helpers).ParseGlob(htmlGlob))
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

	// If data exists return. No need to re-analyze
	if _, ok := users[user]; ok {
		logger.Debug("User data exists", zap.String("user", user))
		records(w, r)
		return
	}
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

// recordsPage Executes similar looking "tplName" templates with "menu" PageHeader data.
func recordsPage(menu PageHeader, tplName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, ok := r.Context().Value(sessionIDKey).(string)
		if !ok {
			logger.Error("User not found", zap.String("page", tplName))
			displayErrorMessage(w, umNotProcessed)
			return
		}

		if err := html.ExecuteTemplate(w, tplName, menu); err != nil {
			logger.Error(err.Error(), zap.String("page", tplName))
			_, _ = w.Write([]byte(err.Error()))
			return
		}
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

// numWithComma a template helper function. Converts -1234567890 to -1,234,567,890.
// Used to print pretty on webpage
func numWithComma(n int) string {
	c := "" // String with ,
	orig := n
	for n != 0 {
		r := int(math.Abs(float64(n % 1000)))
		n /= 1000
		s := ""
		if n != 0 {
			s = fmt.Sprintf(",%03d", r) // Middle part of number
		} else {
			s = fmt.Sprintf("%d", r)
		}
		c = s + c
	}
	if orig < 0 {
		c = "-" + c
	}
	return c
}
