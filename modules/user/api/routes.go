package api

import (
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/user-some-route",testFunc)
	http.HandleFunc("/call-some-route",callFunc)
}