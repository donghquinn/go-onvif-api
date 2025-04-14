package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"org.donghyuns.com/onvif/ptz/ptz"
)

func PresetRouter(server *mux.Router) {
	server.HandleFunc("/preset/set", ptz.SetPresetCtl).Methods(http.MethodPost)
	server.HandleFunc("/preset/apply", ptz.ApplyPresetCtl).Methods(http.MethodPost)
	server.HandleFunc("/preset/list", ptz.GetPresetListCtl).Methods(http.MethodGet)
}
