package user

import (
	"net/http"
)

func RegisterRoutes() {

	handler := NewHandler()

	http.HandleFunc("/users",handler.CreateNewUser)
}