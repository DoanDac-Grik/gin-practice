package middlewares

import "github.com/gin-gonic/gin"

func DummyMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		print("OK OK OK GO THROUGH DUMMY MIDDLEWARE")

	}
}
