package network

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"org.donghyuns.com/onvif/ptz/configs"
	"org.donghyuns.com/onvif/ptz/middlewares"
	"org.donghyuns.com/onvif/ptz/routers"
)

func Network() *http.Server {
	router := mux.NewRouter()

	routers.DeviceRouter(router)
	routers.NodeRouter(router)
	routers.PresetRouter(router)
	routers.PtzRouter(router)
	routers.UserRouter(router)

	handler := middlewares.CorsHanlder().Handler(router)

	serving := &http.Server{
		Addr:           fmt.Sprintf(":%s", configs.GlobalConfig.AppPort),
		MaxHeaderBytes: 1 << 30, // 1GB (1 << 30 = 1073741824 bytes)
		Handler:        handler,
		WriteTimeout:   30 * time.Second,
		ReadTimeout:    30 * time.Second,
	}

	return serving
}
