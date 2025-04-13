package main

import (
	"attempt/api"
	"attempt/infrastructure/config"
	"attempt/infrastructure/db"
	"attempt/interfaces"
	"attempt/usecases"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db.ConnectDB()
	config.InitGoogleOAuthConfig()

	userRepo := interfaces.NewUserRepository(db.DB)
	userService := usecases.NewUserService(userRepo)
	api.ServeRoutes(userService)

	r := gin.Default()
	api.RegisterAuthRoutes(r)

}
