package usecases

import "github.com/gin-gonic/gin"

type UserService interface {
	GetUsers(c *gin.Context)
}
