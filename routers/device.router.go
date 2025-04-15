package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"org.donghyuns.com/onvif/ptz/ptz"
)

func DeviceRouter(server *mux.Router) {
	server.HandleFunc("/device/service/capa", ptz.GetServiceCapaCtl).Methods(http.MethodGet)
	server.HandleFunc("/device/info", ptz.GetDeviceInfoCtl).Methods(http.MethodGet)
	server.HandleFunc("/device/status", ptz.GetDeviceStatusCtl).Methods(http.MethodGet)
	server.HandleFunc("/device/config", ptz.GetDeviceConfigCtl).Methods(http.MethodGet)
}
