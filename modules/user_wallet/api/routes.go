package api

import "net/http"

func RegisterRoutes() {
	http.HandleFunc("/some-route",controllerOneFunc)
	http.HandleFunc("/some-other-route",anotherEndpont)
	http.HandleFunc("/from-second-controller-some-route",controllerTwoFunc)
}