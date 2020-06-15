package repository

import "github.com/thearyanahmed/wallet/database"

func (urwf *UserWalletRepository) SomeFilter() *UserWalletRepository {
	database.DB().Where("id = ?",1).First(&urwf)

	return urwf
}
