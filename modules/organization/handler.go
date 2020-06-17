package organization

import (
	request "github.com/thearyanahmed/wallet/internal/req"
	"net/http"
)

type handler struct {
	organizationService
}

func NewHandler() *handler {
	return &handler{organizationService{organizationRepository{}}}
}

func (handler *handler) createNewOrganization(w http.ResponseWriter,r *http.Request) {

	req := request.Request{}

	if valid := req.Validate(r,w,createNewOrganizationRequest); valid == false {
		return
	}

	//validated := req.ValidatedFormData(r,[]string{"user_id","name","currency_code"})



	// find currency code
	// find user
	// create org
	// create account -> org-id = id
	// create wallet ->
}
