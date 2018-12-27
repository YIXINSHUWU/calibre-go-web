package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Home(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}
