package oauth2_client

import (
	"golang.org/x/oauth2"
)

type OAuth2Client interface {
	AuthCodeURL() string
}

// oAuth2Client implements the OAuth2Client interface
type oAuth2Client struct {
	config *oauth2.Config
}

var _ = NewOAuth2Client // make compiler happy

func NewOAuth2Client() OAuth2Client {
	return &oAuth2Client{config: nil}
}

func (o *oAuth2Client) AuthCodeURL() string {
	if o == nil {
		return ""
	}

	return ""
}
