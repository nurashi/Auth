package api

import (
	"attempt/adapters/httpAuth"
	"attempt/usecases"
	"github.com/gin-gonic/gin"
)

func ServeRoutes(userService usecases.UserService) {
	router := gin.Default()

	RegisterAuthRoutes(router)

	router.GET("api/users", httpAuth.JWTAuthMiddleware(), httpAuth.AdminMiddleWare(), userService.GetUsers)

	router.GET("api/profile", httpAuth.JWTAuthMiddleware(), userService.GetProfile)
	router.PUT("api/profile", httpAuth.JWTAuthMiddleware(), userService.UpdateProfile)

	router.POST("/api/register", userService.Register)
	router.GET("/api/verify", userService.VerifyEmail)

	router.POST("api/login", httpAuth.RateLimitHandler(), userService.Login)

	router.Run(":8080")
}

func RegisterAuthRoutes(r *gin.Engine) {
	auth := r.Group("/auth")
	{
		auth.GET("/google", gin.WrapF(httpAuth.GoogleLogin))
		auth.GET("/google/callback", gin.WrapF(httpAuth.GoogleCallback))
	}
}
