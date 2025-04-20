package handlers

import "github.com/gin-gonic/gin"

type UserService interface {
	GetUsers(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	LoginLogic(email, password string) (string, error)
	GetProfile(c *gin.Context)
	UpdateProfile(c *gin.Context)
	VerifyEmail(c *gin.Context)
	Welcome(c *gin.Context)
}
