package main

import (
	"github.com/apt-getyou/calibre-go-web/app/http/middleware"
	"github.com/apt-getyou/calibre-go-web/config"
	"github.com/apt-getyou/calibre-go-web/routes"
	"github.com/gin-gonic/gin"
	"os"
)

var web = []gin.HandlerFunc{
	middleware.VerifyCsrfToken{}.Handle(),
}

func main() {
	engine := gin.Default()

	e := config.Bootstrap(engine)
	if e != nil {
		os.Exit(1)
	}
	// 初始化引擎
	engine.Group("/", web...)
	{
		routes.Web{}.BuildRoutes(engine)
	}

	// 绑定端口，然后启动应用
	engine.Run(config.HttpAddr + ":" + config.HttpPort)
}
