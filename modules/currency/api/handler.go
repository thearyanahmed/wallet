package api

import (
	"fmt"
	"github.com/thearyanahmed/wallet/internal/res"
	"net/http"
)

type handler struct {
	currencyService
}

func NewHandler() *handler {
	return &handler{currencyService{currencyRepository{}}}
}

func (handler *handler) currencies(w http.ResponseWriter,r *http.Request) {
	currencies, err := handler.currencyService.Currencies()

	if err != nil {
		fmt.Println(err)

		res.SendError(w,err.Error(),err.Error(),422)
	}

	res.Send(w,"",currencies,200)
}
