package currency

import (
	"github.com/thearyanahmed/wallet/database"
	"github.com/thearyanahmed/wallet/schema"
)

type currencyRepository struct {
	schema.Currency
}

func (cr *currencyRepository) Currencies() ( []schema.Currency, error ) {

	var records []schema.Currency

	if err:= database.DB().Find(&records).Error; err != nil {
		return nil, err
	}

	return records, nil
}

func (cr *currencyRepository) findCurrencyByCode(code string) ( *schema.Currency, []error ) {
	var currency schema.Currency

	errs := database.DB().Where("code = ?",code).First(&currency).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}

	return &currency, nil
}