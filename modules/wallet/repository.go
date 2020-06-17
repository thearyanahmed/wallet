package wallet

import (
	"fmt"
	"github.com/thearyanahmed/wallet/schema"
)

type UserWalletRepository struct {
	schema.Wallet
}

func (uwr *UserWalletRepository) Test() {
	fmt.Println("ello world")
}

