package httpAuth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func RateLimitHandler() gin.HandlerFunc {
	requests := make(chan int, 5)

	for i := 0; i < cap(requests); i++ {
		requests <- i
	}

	close(requests)

	// its just stoper to avoid to many requests
	limiter := time.Tick(5 * time.Second)

	return func(c *gin.Context) {
		select {
		case <-limiter:
			// if limiter works just give access
			c.Next()

		default:
			// like else
			c.JSON(http.StatusTooManyRequests, gin.H{
				"message": "too many requests",
			})
			c.Abort()
		}
	}
}
