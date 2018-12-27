package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func RequestMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("before request")
		ctx.Next()
		fmt.Println("after request")
	}
}
