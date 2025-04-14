package configs

import "os"

type GlobalConf struct {
	AppPort string
	AppHost string
}

var GlobalConfig GlobalConf

func SetGlobalConfig() {
	GlobalConfig.AppHost = os.Getenv("APP_HOST")
	GlobalConfig.AppPort = os.Getenv("APP_PORT")
}
