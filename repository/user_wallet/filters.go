package user_wallet

import "github.com/thearyanahmed/wallet/database"

func (ua *UserWallet) SomeFilter() *UserWallet {
	database.DB().Where("id = ?",1).First(&ua)

	return ua
}
