package google

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func GoogleOauthConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("GCP_CLIENT_ID"),
		ClientSecret: os.Getenv("GCP_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_AUTH_CALLBACK"),
		Scopes: []string{
			os.Getenv("GCP_SCOPE_EMAIL"),
			os.Getenv("GCP_SCOPE_PROFILE"),
		},
		Endpoint: google.Endpoint,
	}
}
