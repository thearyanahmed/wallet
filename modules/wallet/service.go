package wallet

import "github.com/thearyanahmed/wallet/schema"

type Service struct {
	walletRepository
}

func (service *Service) CreateNewWallet(userID, accountID ,currencyID uint, currencyCode string) (*schema.Wallet,[]error){
	return service.walletRepository.createNewWallet(userID , accountID, currencyID, currencyCode)
}