package user

import (
	"github.com/gorilla/mux"
	"github.com/thearyanahmed/wallet/internal/req"
)

func RegisterRoutes(router *mux.Router) {
	handler := NewHandler()

	router.HandleFunc("/users",handler.createNewUser).Methods(req.POST)
}