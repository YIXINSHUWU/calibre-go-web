package controllers

import (
	"github.com/apt-getyou/calibre-go-web/app/http/manager"
	"github.com/apt-getyou/calibre-go-web/app/http/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
}

func (HomeController) Home(context *gin.Context) {
	var user model.User
	_ = context.Bind(&user)
	manager.HomeManager{}.ShiShi(&user)
	context.JSON(http.StatusOK, "hello, world")
}
