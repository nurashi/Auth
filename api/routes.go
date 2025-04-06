package api

import (
	"attempt/usecases"
	"github.com/gin-gonic/gin"
)

func ServeRoutes(userService usecases.UserService) {
	router := gin.Default()
	router.GET("/api/users", userService.GetUsers)
	router.Run(":8080")
}
