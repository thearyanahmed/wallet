package user

import (
	"fmt"
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

	validated := req.ValidatedFormData(r,[]string{"first_name","last_name","email"})
	// create user
	user, errs := createNewUser(validated["first_name"],validated["last_name"],validated["email"])

	if len(errs) > 0 {
		res.SendError(w,"Unprocessable entity.",errs,422)
		return
	}

	fmt.Println("Created_user",user)
	//err := createNewUser()
	// create org if needed
	// create account
	// create defalt user_wallet


	res.Send(w,"Reached here",nil,200)
}

