package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"org.donghyuns.com/onvif/ptz/ptz"
)

func DeviceRouter(server *mux.Router) {
	server.HandleFunc("/device/capa", ptz.GetServiceCapaCtl).Methods(http.MethodGet)
	server.HandleFunc("/device/info", ptz.GetDeviceInfoCtl).Methods(http.MethodGet)
}
