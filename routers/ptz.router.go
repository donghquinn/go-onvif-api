package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"org.donghyuns.com/onvif/ptz/ptz"
)

func PtzRouter(server *mux.Router) {
	server.HandleFunc("/ptz/move/relative", ptz.RelativeMoveCtl).Methods(http.MethodPut)
	server.HandleFunc("/ptz/move/continuous", ptz.ContinouseMoveCtl).Methods(http.MethodPut)
	server.HandleFunc("/ptz/default/set", ptz.SetDefaultPositionCtl).Methods(http.MethodPost)
	server.HandleFunc("/ptz/default/move", ptz.MoveToDefaultPositionCtl).Methods(http.MethodPut)
}
