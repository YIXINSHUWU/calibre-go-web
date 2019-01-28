package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type VerifyCsrfToken struct {

}

func (VerifyCsrfToken) Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("before request")
		ctx.Next()
		fmt.Println("after request")
	}
}
