package api

import (
	"github.com/thearyanahmed/wallet/schema"
)

type currencyService struct {
	currencyRepository
}

func (service *currencyService) Currencies() ([]schema.Currency, error) {
	return service.currencyRepository.Currencies()
}