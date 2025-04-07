package jwtAuth

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"
)

func ExtractClaims(c *gin.Context) (jwt.MapClaims, error) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		return nil, errors.New("Authorization header is empty")
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, errors.New("Invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Invalid parsing")
	}
	return claims, nil
}
