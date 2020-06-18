package account

import (
	"github.com/thearyanahmed/wallet/database"
	"github.com/thearyanahmed/wallet/schema"
)

type accountRepository struct {
	schema.Account
}

func (accountRepo *accountRepository) createNewAccount(userID, orgID uint, currencyCode string) (*schema.Account,error) {

	//uniqueAccountID := uuid.Must(uuid.NewRandom())

	//ref := strings.ReplaceAll(uniqueAccountID.String(),"-","")

	account := schema.Account{
		UserID:                userID,
		RefID:                 "asdasdasdasdasdasdasdasdadasdasdassasdasd",
		Type:                  schema.OrganizationAccount,
		OrgID:                 orgID,
		DefaultWalletCurrency: currencyCode,
	}
	err := database.DB().Create(&account).Error

	if err != nil {
		return nil, err
	}
	return &account, nil
}


