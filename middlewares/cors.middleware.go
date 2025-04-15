package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

var originList = []string{
	"https://itscontrol.local",
	"https://its-control-dev.zetra.kr",
	"https://onvif-api-dev.zetra.kr",
}

func CorsHanlder() *cors.Cors {
	corHandler := cors.New(cors.Options{
		AllowedOrigins:   originList,
		AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodOptions},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
		MaxAge:           86400,
		Debug:            false,
	})

	return corHandler
}
