package middlewares

import (
	"net/http"

	"github.com/rs/cors"
)

var originList = []string{
	"http://localhost:3000",
	"http://localhost:9852",
	"https://toonizia.local",
	"https://www.toonizia.com",
	"https://toonizia.donghyuns.com",
	"unknown",
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
