package main

import (
	"attempt/api"
	"attempt/infrastructure/db"
	"attempt/interfaces"
	"attempt/usecases"
)

func main() {
	db.ConnectDB()

	userRepo := interfaces.NewUserRepository(db.DB)
	userService := usecases.NewUserService(userRepo)
	api.ServeRoutes(userService)
}
