package account

import (
	"github.com/google/uuid"
	"github.com/thearyanahmed/wallet/database"
	"github.com/thearyanahmed/wallet/schema"
	"strings"
)

type accountRepository struct {
	schema.Account
}

func (accountRepo *accountRepository) createNewAccount(userID, orgID uint, currencyCode string) (*schema.Account,[]error) {

	uniqueAccountID := uuid.Must(uuid.NewRandom())

	ref := strings.ReplaceAll(uniqueAccountID.String(),"-","")

	account := schema.Account{
		UserID:                userID,
		RefID:                 ref,
		Type:                  schema.OrganizationAccount,
		OrgID:                 orgID,
		DefaultWalletCurrency: currencyCode,
	}
	errs := database.DB().Create(&account).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &account, nil
}


