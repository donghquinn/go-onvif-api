package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"org.donghyuns.com/onvif/ptz/ptz"
)

func NodeRouter(server *mux.Router) {
	server.HandleFunc("/node/list", ptz.GetNodeListCtl).Methods(http.MethodGet)
	server.HandleFunc("/node/detail/{nodeProfile}", ptz.GetNodeDetailCtl).Methods(http.MethodGet)
}
