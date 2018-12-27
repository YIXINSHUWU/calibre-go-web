package config

var (
	HttpAddr string
	HttpPort string
	Debug    string
)

func initSettingValue() {
	HttpAddr = Env("addr", "127.0.0.1")
	HttpPort = Env("port", "8080")
	Debug = Env("debug", "false")
}
