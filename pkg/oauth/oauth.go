package oauth

import (
	"github.com/ryvasa/go-super-farmer-auth-service/pkg/env"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func NewOauth(env *env.Env) *oauth2.Config {
	// Konfigurasi OAuth 2.0
	googleOauthConfig := &oauth2.Config{
		ClientID:     env.Google.ClientID,
		ClientSecret: env.Google.ClientSecret,
		RedirectURL:  env.Google.RedirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}

	return googleOauthConfig
}
