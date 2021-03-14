package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/opendroid/hk/logger"
	"go.uber.org/zap"
	"html/template"
	"io/ioutil"
	"math"
	"net/http"
)

// vc are View Controllers for the handlers
// Save the health data in a map
var (
	html *template.Template
	file string
)

// init initializes the templates
func init() {
	helpers := template.FuncMap{"numWithComma": numWithComma}
	html = template.Must(template.New(root).Funcs(helpers).ParseGlob(htmlGlob))
}

// Index page handler
func index(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(contextKeyUserID).(string)
	if !ok {
		logger.Error("User not found", zap.String("page", "index"))
		displayErrorMessage(w, umNoData)
		return
	}

	// If data exists return. No need to re-analyze
	if d, ok := users[user]; ok && d.health != nil {
		logger.Debug("User health data exists", zap.String("user", user))
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

// authHandler initiates OAuth2 with Google. Redirects the user to Google's "oauth2.Endpoint".
// Once the auth is successful the user is re-directed to URL auth2ClientConfig.RedirectURL.
// The allowed callback URLs need to configured in GCP Project.
func authHandler(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(contextKeyUserID).(string)
	if !ok {
		logger.Error("User required for authentication", zap.String("page", "records"))
		displayErrorMessage(w, umNotProcessed)
		return
	}
	url := auth2ClientConfig.AuthCodeURL(user)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// loginSuccess routes successful requests to user
func loginSuccess(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(contextKeyUserID).(string)
	if !ok {
		logger.Error("User not found", zap.String("page", "records"))
		displayErrorMessage(w, umNotProcessed)
		return
	}
	// Get the state i.e. {userCookie}:{redirectURL}.
	userState, code := r.FormValue("state"), r.FormValue("code")
	if user != userState {
		logger.Error("Invalid user", zap.String("user", user), zap.String("state", userState))
	}
	// Use the code an userState to get the user info
	if data, err := googleUserInfo(userState, code); err != nil {
		logger.Error(err.Error(), zap.String("page", "loginSuccess"))
		displayErrorMessage(w, umNotProcessed)
		return
	} else {
		updateAuth(user, data)
		logger.Info("Updated User", zap.String("user", user))
	}
	http.Redirect(w, r, "/records-xhr-all", http.StatusTemporaryRedirect)
}

// records page handler
func records(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(contextKeyUserID).(string)
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

// execTemplate Executes similar looking "tplName" templates with "menu" PageHeader data.
func execTemplate(menu PageHeader, tplName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sid, ok := r.Context().Value(contextKeyUserID).(string)
		if !ok {
			logger.Error("User not found", zap.String("page", tplName))
			displayErrorMessage(w, umNotProcessed)
			return
		}
		menu.NP.State = "data" // Indicator for NV to show or not to show user data
		authenticated := isAuthenticated(sid)
		if !authenticated { // If not Authenticated redirect to login page
			menu.NP.State = "login"
			logger.Info("User not authenticated", zap.String("page", tplName),
				zap.String("user", sid), zap.String("redirect", "/login"))
			http.Redirect(w, r, "/login", http.StatusTemporaryRedirect)
			return
		}

		if err := html.ExecuteTemplate(w, tplName, menu); err != nil {
			logger.Error(err.Error(), zap.String("page", tplName))
			_, _ = w.Write([]byte(err.Error()))
			return
		}
	}
}

// execPublicTemplate a public page eg privacy, contactus or support
func execPublicTemplate(menu PageHeader, tplName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sid, _ := r.Context().Value(contextKeyUserID).(string)
		menu.NP.State = "data" // Indicator for NV to show or not to show user data
		authenticated := isAuthenticated(sid)
		if !authenticated { // If not Authenticated redirect to login page
			menu.NP.State = "login"
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
	displayErrorMessage(w, umNoError)
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

// googleUserInfo returns user data from Google
func googleUserInfo(state, code string) (*gUserInfo, error) {
	if state == "" {
		return nil, fmt.Errorf("invalid oauth state")
	}
	token, err := auth2ClientConfig.Exchange(context.TODO(), code)
	if err != nil {
		return nil, fmt.Errorf("exchange error: %s", err.Error())
	}
	r, err := http.Get(userInfoUrl + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("userinfo error: %s", err.Error())
	}
	defer func() { _ = r.Body.Close() }()
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return nil, fmt.Errorf("parse error: %s", err.Error())
	}
	var userInfo gUserInfo
	if err := json.Unmarshal(data, &userInfo); err != nil {
		return nil, fmt.Errorf("unmarshal error: %s", err.Error())
	}
	userInfo.Authenticated = true // Set the map
	logger.Debug("User info", zap.Bool("auth", userInfo.Authenticated),
		zap.String("name", userInfo.Name), zap.Bool("email_verified", userInfo.VerifiedEmail),
		zap.String("email", userInfo.Email), zap.String("gender", userInfo.Gender),
		zap.String("locale", userInfo.Locale), zap.String("id", userInfo.ID))
	return &userInfo, nil
}
