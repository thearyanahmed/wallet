package api

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/some-route",controllerOneFunc)
	router.HandleFunc("/some-other-route",anotherEndpont)
}