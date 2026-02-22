package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	deviceapi "org.donghyuns.com/onvif/ptz/internal/apis/device"
	"org.donghyuns.com/onvif/ptz/pkg/database"
)

func DeviceRouter(server *mux.Router, dbCon *database.PostgresService) {
	deviceService := deviceapi.NewDeviceCtlService(dbCon)

	server.HandleFunc("/device/service/capa", deviceService.GetServiceCapaCtl).Methods(http.MethodGet)
	server.HandleFunc("/device/info", deviceService.GetDeviceInfoCtl).Methods(http.MethodGet)
	server.HandleFunc("/device/status", deviceService.GetDeviceStatusCtl).Methods(http.MethodGet)
	server.HandleFunc("/device/config", deviceService.GetDeviceConfigCtl).Methods(http.MethodGet)
}
