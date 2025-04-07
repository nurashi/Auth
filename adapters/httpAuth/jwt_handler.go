package httpAuth

import (
	"attempt/utils/jwtAuth"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func LoggerMiddleware(c *gin.Context) {
	log.Printf("Incoming request: %s %s", c.Request.Method, c.Request.URL) // gin does it automaticly
	c.Next()
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return

		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := jwtAuth.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "token issue"})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Next()
	}
}

func RoleMiddleWare(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, err := jwtAuth.ExtractClaims(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No access token"})
			c.Abort()
			return
		}

		role, ok := claims["role"].(string)
		if !ok || role != requiredRole {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No access token"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func AdminMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(401, gin.H{"message": "NO access token 1"})
			c.Abort()
			return
		}
		claims, err := jwtAuth.ParseToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No access token 2"})
			c.Abort()
			return
		}

		fmt.Println("claims: ", claims)

		if claims.Role != "admin" {
			c.JSON(http.StatusUnauthorized, gin.H{"message": "No access token 3"})
			c.Abort()
			return
		}

		c.Next()
	}
}
