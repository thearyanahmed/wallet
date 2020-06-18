package wallet

import (
	"github.com/thearyanahmed/wallet/database"
	"github.com/thearyanahmed/wallet/schema"
)

type walletRepository struct {
	schema.Wallet
}

func (repo *walletRepository) createNewWallet(userID, accountID ,currencyID uint, currencyCode string) (*schema.Wallet,error) {
	wallet := schema.Wallet{
		UserID:           userID,
		AccountID:        accountID,
		CurrencyCode:     currencyCode,
		CurrencyID:       currencyID,
		AvailableBalance: 0,
		TotalBalance:     0,
	}
	err := database.DB().Create(&wallet).Error

	if err != nil {
		return nil, err
	}
	return &wallet, nil
}