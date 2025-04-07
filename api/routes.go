package api

import (
	"attempt/adapters/httpAuth"
	"attempt/usecases"
	"github.com/gin-gonic/gin"
)

func ServeRoutes(userService usecases.UserService) {
	router := gin.Default()

	/* JWTAuthMiddleware() -> its a first stage of security, used to every endpoint
	*  AdminMiddleWare() -> its a second stage of security, used to endpoint which are for admin
	 */
	router.GET("api/users", httpAuth.JWTAuthMiddleware(), httpAuth.AdminMiddleWare(), userService.GetUsers)

	// only after auth
	router.GET("api/profile", httpAuth.JWTAuthMiddleware(), userService.GetProfile)
	router.PUT("api/profile", httpAuth.JWTAuthMiddleware(), userService.UpdateProfile)

	// public router(just access to endpoint), if someone try to get access to other endpoint, programm will change their link to this two ones.
	router.POST("/api/register", userService.Register)
	router.GET("/api/verify", userService.VerifyEmail)

	router.POST("api/login", httpAuth.RateLimitHandler(), userService.Login)

	router.Run(":8080")
}
