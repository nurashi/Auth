package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"log"
	"os"
)

var GoogleOAuthConfig *oauth2.Config

func InitGoogleOAuthConfig() {
	clientID := os.Getenv("GOOGLE_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		log.Fatal("Google OAuth credentials are missing. Check your .env or environment variables.")
	}

	GoogleOAuthConfig = &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	log.Println("Google OAuth config initialized.")
}
