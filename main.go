package main

import (
	"github.com/apt-getyou/calibre-go-web/app/http/middleware"
	"github.com/apt-getyou/calibre-go-web/config"
	"github.com/apt-getyou/calibre-go-web/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	config.Bootstrap()
	// 初始化引擎
	engine := gin.Default()
	// 注册一个路由和处理函数
	routes.BuildRoutes(engine)
	engine.Use(middleware.RequestMiddleware())

	// 绑定端口，然后启动应用
	engine.Run(config.HttpAddr + ":" + config.HttpPort)
}

/**
* 根请求处理函数
* 所有本次请求相关的方法都在 context 中，完美
* 输出响应 hello, world
*/
func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}
