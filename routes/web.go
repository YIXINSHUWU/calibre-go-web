package routes

import (
	"github.com/apt-getyou/calibre-go-web/app/http/controllers"
	"github.com/gin-gonic/gin"
)

func BuildRoutes(engine *gin.Engine) {
	engine.GET("/", controllers.Home)
}
