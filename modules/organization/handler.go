package organization

import (
	request "github.com/thearyanahmed/wallet/internal/req"
	"github.com/thearyanahmed/wallet/internal/res"
	"net/http"
	"strconv"
)

type handler struct {
	organizationService
}


const (
	orgCreatedSuccessfully = "Organization created successfully."
)


func NewHandler() *handler {
	return &handler{organizationService{organizationRepository{}}}
}

func (handler *handler) createNewOrganization(w http.ResponseWriter,r *http.Request) {

	req := request.Request{}

	if valid := req.Validate(r,w,createNewOrganizationRequest); valid == false {
		return
	}

	validated := req.ValidatedFormData(r,[]string{"user_id","name","currency_code","org_id"})

	// check if user exists
	// check if currency exists

	userID, _ := strconv.Atoi(validated["user_id"])

	org, errs := createOrganization(uint(userID),validated["name"])

	if len(errs) > 0 {
		res.SendError(w,"Unprocessable entity.",errs,422)
		return
	}

	response := createdResponse{
		ID:   org.ID,
		Name: org.Name,
		CreatedAt: org.CreatedAt,
	}

	res.Send(w,orgCreatedSuccessfully,response,200)

}
