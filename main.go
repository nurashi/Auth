package main

import (
	"attempt/api"
	"attempt/handlers"
	"attempt/infrastructure/config"
	"attempt/infrastructure/db"
	"attempt/interfaces"
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
	userService := handlers.NewUserService(userRepo)
	api.ServeRoutes(userService, userRepo)

	r := gin.Default()
	api.RegisterAuthRoutes(r, userRepo)

}
