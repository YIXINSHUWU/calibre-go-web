package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type HomeController struct {
}

func (HomeController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func (HomeController) Login(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}
