package routers

import (
	"net/http"

	"github.com/gorilla/mux"
	"org.donghyuns.com/onvif/ptz/ptz"
)

func UserRouter(server *mux.Router) {
	server.HandleFunc("/user/create", ptz.CreateUserCtl).Methods(http.MethodPost)
	server.HandleFunc("/user/create/profile", ptz.CreateProfileCtl).Methods(http.MethodPost)
	server.HandleFunc("/user/profile", ptz.GetProfile).Methods(http.MethodGet)
	server.HandleFunc("/user/list", ptz.GetUserListCtl).Methods(http.MethodGet)
}
