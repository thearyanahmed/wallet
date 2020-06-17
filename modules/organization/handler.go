package account

import "net/http"

type handler struct {
	organizationService
}

func NewHandler() *handler {
	return &handler{organizationService{organizationRepository{}}}
}

func (handler *handler) createNewOrganization(w http.ResponseWriter,r *http.Request) {

}
