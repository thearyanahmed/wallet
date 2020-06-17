package account

import (
	"github.com/gorilla/mux"
	"github.com/thearyanahmed/wallet/internal/req"
)

func RegisterRoutes(router *mux.Router) {
	handler := NewHandler()

	router.HandleFunc("/organizations",handler.createNewOrganization).Methods(req.POST)
}