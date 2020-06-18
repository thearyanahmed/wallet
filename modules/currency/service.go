package currency

import (
	"github.com/thearyanahmed/wallet/schema"
)

type Service struct {
	currencyRepository
}

func (service *Service) Currencies() ([]schema.Currency, error) {
	return service.currencyRepository.Currencies()
}

func (service *Service) FindCurrencyByCode(code string) ( *schema.Currency, []error ) {
	return service.currencyRepository.findCurrencyByCode(code)
}