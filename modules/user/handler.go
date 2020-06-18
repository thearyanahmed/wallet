package user

import (
	request "github.com/thearyanahmed/wallet/internal/req"
	"github.com/thearyanahmed/wallet/internal/res"
	"net/http"
)

type handler struct {
	Service
}

const (
	userCreatedSuccessfully = "User created successfully."
)

func NewHandler() *handler {
	return &handler{Service{userRepository{}}}
}

func (handler *handler) createNewUser(w http.ResponseWriter,r *http.Request) {

	req := request.Request{}

	if valid := req.Validate(r,w,createNewUserRequest); valid == false {
		return
	}

	validated := req.ValidatedFormData(r,[]string{"first_name","last_name","email"})

	user, errs := createNewUser(validated["first_name"],validated["last_name"],validated["email"])

	if len(errs) > 0 {
		res.SendError(w,"Unprocessable entity.",errs,422)
		return
	}

	res.Send(w,userCreatedSuccessfully,user,200)
}

