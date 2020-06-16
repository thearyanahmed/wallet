package api

import "net/http"

func RegisterRoutes() {
	handler := NewHandler()

	http.HandleFunc("/currencies",handler.currencies)
}