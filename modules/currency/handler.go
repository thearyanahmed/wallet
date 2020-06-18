package currency

import (
	"github.com/thearyanahmed/wallet/internal/res"
	"net/http"
)

type handler struct {
	Service
}

func NewHandler() *handler {
	return &handler{Service{currencyRepository{}}}
}

func (handler *handler) currencies(w http.ResponseWriter,r *http.Request) {

	currencies, err := handler.Service.Currencies()

	if err != nil {
		res.SendError(w,err.Error(),err.Error(),422)
	}

	res.Send(w,"",currencies,200)
}
