package config

import (
	"github.com/apt-getyou/gtf"
	"github.com/gin-gonic/gin"
)

var (
	HttpAddr = ""
	HttpPort = "80"
	Debug    = "debug"
)

type WebConfig struct {
}

func (WebConfig) initSettingValue() {
	HttpAddr = Env("addr", HttpAddr)
	HttpPort = Env("port", HttpPort)
	Debug = Env("debug", Debug)
}

func (WebConfig) init(engine *gin.Engine) {
	for k, v := range gtf.GtfFuncMap {
		engine.FuncMap[k] = v
	}

	engine.LoadHTMLGlob("resources/view/*/*")
	engine.Static("/static", "resources/static")
}
