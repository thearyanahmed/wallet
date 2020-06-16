package currency

import (
	"github.com/gorilla/mux"
	"github.com/thearyanahmed/wallet/internal/req"
	"github.com/thearyanahmed/wallet/oauth"
)

func RegisterRoutes(router *mux.Router) {
	handler := NewHandler()

	router.HandleFunc("/currencies",oauth.Auth(handler.currencies)).Methods(req.GET)
}