package account

import "github.com/thearyanahmed/wallet/schema"

type Service struct {
	accountRepository
}

func (service *Service) CreateNewAccount(userID, orgID uint, currencyCode string) (*schema.Account,[]error) {
	return service.accountRepository.createNewAccount(userID,orgID,currencyCode)
}