package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"org.donghyuns.com/onvif/ptz/ptz"
)

func PtzRouter(server *mux.Router) {
	server.HandleFunc("/ptz/move/relative", ptz.RelativeMoveCtl).Methods(http.MethodPost)
	server.HandleFunc("/ptz/move/continous", ptz.ContinouseMoveCtl).Methods(http.MethodPost)
}
