package api

import (
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/account-some-route",testFunc)
	http.HandleFunc("/final-some-route",finalTestController)
}