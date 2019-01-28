package controllers

import (
	"github.com/apt-getyou/calibre-go-web/app/http/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
}

func (HomeController) Home(c *gin.Context) {
	var user model.User
	_ = c.Bind(&user)
	c.HTML(http.StatusOK, "index.html", nil)
}
