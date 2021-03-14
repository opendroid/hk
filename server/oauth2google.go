package server

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
)

// auth2ClientConfig stores Web Application OAuth2 client.
// See https://developers.google.com/identity/protocols/oauth2/openid-connect for details
var (
	auth2ClientConfig *oauth2.Config
)

const (
	// userInfoUrl Google API that uses access_token. Try these in https://developers.google.com/oauthplayground
	userInfoUrl  string = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="
	redirectProd string = "https://www.usense.io/__g"
	redirectDev  string = "http://localhost:8080/__g"
)

func init() {
	env := os.Getenv("USENSE_ENVIRONMENT")
	redirectUrl := redirectProd
	if env == "DEVELOPMENT" {
		redirectUrl = redirectDev
	}
	auth2ClientConfig = &oauth2.Config{
		ClientID:     os.Getenv("USENSE_OAUTH2_CLIENT_ID"),
		ClientSecret: os.Getenv("USENSE_OAUTH2_CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  redirectUrl,
		// See all scopes at https://developers.google.com/identity/protocols/oauth2/scopes?authuser=2
		Scopes: []string{"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
	}
}

// gUserInfo user data for UserInfo API call
// https://any-api.com/googleapis_com/oauth2/docs/userinfo/oauth2_userinfo_get
// TODO: save "code" to
type gUserInfo struct {
	Authenticated bool   // true if user is Authenticated. Set locally
	ID            string `json:"id,omitempty"`
	Email         string `json:"email,omitempty"`
	FamilyName    string `json:"family_name,omitempty"`
	GivenName     string `json:"given_name,omitempty"`
	Gender        string `json:"gender,omitempty"`
	Link          string `json:"link,omitempty"`
	Locale        string `json:"locale,omitempty"`
	Name          string `json:"name,omitempty"`
	ProfileURL    string `json:"picture,omitempty"`
	VerifiedEmail bool   `json:"verified_email,omitempty"`
}
