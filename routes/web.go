package routes

import (
	"github.com/apt-getyou/calibre-go-web/app/http/controllers"
	"github.com/gin-gonic/gin"
)

type Web struct {
}

func (Web) BuildRoutes(engine *gin.Engine) {
	engine.GET("/", controllers.HomeController{}.Home)
}
