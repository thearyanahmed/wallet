package organization

import (
	"time"
)

type (
	accountResponse struct {
		RefID string `json:"ref_id"`
		DefaultCurrencyForWallet string `json:"default_currency_for_wallet"`
	}

	orgResponse struct {
		ID uint `json:"id"`
		Name string `json:"name"`
		CreatedAt time.Time `json:"created_at"`
	}

	walletResponse struct {
		AccountReference string `json:"account_reference"`
		CurrencyCode string `json:"currency_code"`
		AvailableBalance int64 `json:"available_balance"`
		TotalBalance int64 `json:"total_balance"`
	}

	createdResponse struct {
		Organization orgResponse `json:"organization"`
		Account accountResponse `json:"account"`
		Wallet walletResponse `json:"wallet"`
	}
)
