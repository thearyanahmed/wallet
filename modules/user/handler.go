package user

import "net/http"

type handler struct {
	userService
}

func NewHandler() *handler {
	return &handler{userService{userRepository{}}}
}

func (handler *handler) CreateNewUser(w http.ResponseWriter,r *http.Request) {

}