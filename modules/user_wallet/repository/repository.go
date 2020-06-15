package repository

import (
	"fmt"
	"github.com/thearyanahmed/wallet/schema"
)

type UserWalletRepository struct {
	schema.UserWallet
}

func (uwr *UserWalletRepository) Test() {
	fmt.Println("ello world")
}

