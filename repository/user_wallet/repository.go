package user_wallet

import (
	"fmt"
	"github.com/thearyanahmed/wallet/schema"
)

type UserWallet struct {
	schema.UserWallet
}

func (uw *UserWallet) Test() {
	fmt.Println("ello world")
}

