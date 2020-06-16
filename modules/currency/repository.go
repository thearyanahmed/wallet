package currency

import (
	"fmt"
	"github.com/thearyanahmed/wallet/database"
	"github.com/thearyanahmed/wallet/schema"
)

type currencyRepository struct {
	schema.Currency
}

func (cr *currencyRepository) Currencies() ( []schema.Currency, error ) {

	var records []schema.Currency

	if err:= database.DB().Find(&records).Error; err != nil {
		fmt.Println(err)
		return nil, err
	}

	return records, nil
}