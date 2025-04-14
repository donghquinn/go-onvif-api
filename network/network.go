package network

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"org.donghyuns.com/onvif/ptz/configs"
	"org.donghyuns.com/onvif/ptz/middlewares"
)

func Network() *http.Server {
	router := mux.NewRouter()

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
