package api

import (
	"github.com/thearyanahmed/wallet/oauth"
	"net/http"
)

func RegisterRoutes() {
	handler := NewHandler()

	http.HandleFunc("/currencies",oauth.Auth(handler.currencies))
}