package httpAuth

import (
	"attempt/infrastructure/config"
	"attempt/interfaces"
	"attempt/utils/jwtAuth"
	"context"
	"net/http"
	"net/url"

	"github.com/gin-gonic/gin"
	oauth2service "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

// OAuth2
func GoogleLogin(w http.ResponseWriter, r *http.Request) {
	url := config.GoogleOAuthConfig.AuthCodeURL("random-state")
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleCallback(userRepo interfaces.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()

		code := c.Query("code")
		if code == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Code not found"})
			return
		}

		token, err := config.GoogleOAuthConfig.Exchange(ctx, code)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Token exchange failed: " + err.Error()})
			return
		}

		client := config.GoogleOAuthConfig.Client(ctx, token)
		service, err := oauth2service.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service: " + err.Error()})
			return
		}

		userinfo, err := service.Userinfo.Get().Do()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get user info: " + err.Error()})
			return
		}

		user, err := userRepo.FindOrCreateUser(userinfo.Email, userinfo.Name, userinfo.Picture)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "FindOrCreateUser failed: " + err.Error()})
			return
		}

		jwtToken, err := jwtAuth.GenerateToken(user.Email, user.Role)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "JWT generation failed: " + err.Error()})
			return
		}

		c.SetCookie("Authorization", jwtToken, 3600, "/", "", false, true)

		redirectURL := "/welcome?email=" + url.QueryEscape(user.Email)
		
		if user.Name != "" {
			redirectURL += "&name=" + url.QueryEscape(user.Name)
		}
		
		if user.Picture != "" {
			redirectURL += "&picture=" + url.QueryEscape(user.Picture)
		}

		c.Redirect(http.StatusSeeOther, redirectURL)
	}
}

