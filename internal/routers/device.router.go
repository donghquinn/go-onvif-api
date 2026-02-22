package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	deviceapi "org.donghyuns.com/onvif/ptz/internal/apis/device"
)

func DeviceRouter(server *mux.Router) {
	deviceService := deviceapi.NewDeviceCtlService()

	server.HandleFunc("/device/service/capa", deviceService.GetServiceCapaCtl).Methods(http.MethodGet)
	server.HandleFunc("/device/info", deviceService.GetDeviceInfoCtl).Methods(http.MethodGet)
	server.HandleFunc("/device/status", deviceService.GetDeviceStatusCtl).Methods(http.MethodGet)
	server.HandleFunc("/device/config", deviceService.GetDeviceConfigCtl).Methods(http.MethodGet)
}
