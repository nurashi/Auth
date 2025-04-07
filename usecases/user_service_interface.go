package usecases

import "github.com/gin-gonic/gin"

type UserService interface {
	GetUsers(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
	GetProfile(c *gin.Context)
	UpdateProfile(c *gin.Context)
}
