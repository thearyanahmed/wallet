package user

import (
	request "github.com/thearyanahmed/wallet/internal/req"
	"github.com/thearyanahmed/wallet/internal/res"
	"net/http"
)

type handler struct {
	userService
}

func NewHandler() *handler {
	return &handler{userService{userRepository{}}}
}

func (handler *handler) createNewUser(w http.ResponseWriter,r *http.Request) {

	req := request.Request{}

	if valid := req.Validate(r,w,createNewUserRequest); valid == false {
		return
	}

	res.Send(w,"Reached here",nil,200)
}

