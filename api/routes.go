package api

import (
	"attempt/adapters/httpAuth"
	"attempt/handlers"
	"attempt/interfaces"

	"github.com/gin-gonic/gin"
)

func ServeRoutes(userService handlers.UserService, userRepo interfaces.UserRepository) {
	router := gin.Default()

	RegisterAuthRoutes(router, userRepo)

	router.GET("api/users", httpAuth.JWTAuthMiddleware(), httpAuth.AdminMiddleWare(), userService.GetUsers)
	router.GET("api/profile", httpAuth.JWTAuthMiddleware(), userService.GetProfile)
	router.PUT("api/profile", httpAuth.JWTAuthMiddleware(), userService.UpdateProfile)

	router.POST("/api/register", userService.Register)
	router.GET("/api/verify", userService.VerifyEmail)

	router.POST("api/login", httpAuth.RateLimitHandler(), userService.Login)

	router.GET("/welcome", func(c *gin.Context) {
		email := c.Query("email")
		name := c.Query("name")
		picture := c.Query("picture")

		html := `<html>
			<head><title>Welcome</title></head>
			<body style="font-family:sans-serif; text-align:center; padding:40px;">
				<h1>name , ` + name + `!</h1>
				<p>email of use: ` + email + `</p>
				<img src="` + picture + `" style="width:100px; height:100px;">
			</body>
		</html>`

		c.Data(200, "text/html; charset=utf-8", []byte(html))
	})

	router.Run(":8080")
}

func RegisterAuthRoutes(r *gin.Engine, userRepo interfaces.UserRepository) {
	auth := r.Group("/auth")
	{
		auth.GET("/google", gin.WrapF(httpAuth.GoogleLogin))
		auth.GET("/google/callback", httpAuth.GoogleCallback(userRepo))

	}
}
