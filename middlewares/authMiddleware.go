package middlewares

import (
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var token = ctx.GetHeader("authToken")
		if token == "" {
			ctx.JSON(400, "Not found token")
			ctx.Abort()
		}
	}

}
