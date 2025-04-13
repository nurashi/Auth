package httpAuth

import (
	"attempt/infrastructure/config"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// GoogleLogin initiates the OAuth2 flow by redirecting to Google's authentication page
func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	// Generate the authorization URL with a random state for CSRF protection
	url := config.GoogleOAuthConfig.AuthCodeURL("random-state")
	// Redirect the user to Google's consent page
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

// GoogleCallback handles the callback from Google after user authentication
func GoogleCallback(w http.ResponseWriter, r *http.Request) {
	// Get the authorization code from the query parameters
	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Code not found", http.StatusBadRequest)
		return
	}

	// Exchange the authorization code for an access token
	token, err := config.GoogleOAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Failed to exchange token: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Create an HTTP client with the token
	client := config.GoogleOAuthConfig.Client(context.Background(), token)

	// Fetch the user info from Google API
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Failed to get user info", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the JSON response
	var userInfo map[string]interface{}
	if err := json.Unmarshal(body, &userInfo); err != nil {
		http.Error(w, "Failed to parse user info: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Write a simple response back to the user
	fmt.Fprintf(w, "Hello, %v!", userInfo["email"])
}
