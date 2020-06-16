package service

import (
	"fmt"
	"github.com/thearyanahmed/wallet/modules/account/service"
)

func TestUserService () {
	fmt.Println("Calling testing user service")
}

func TestHumanService() {
	fmt.Println("Testing human service")
}

type UserService struct {
	service.AccountService
}

func SomethingElse() {
	userSvc := UserService{}
	userSvc.SomeAccountServiceFunction()
	service.CallThisAccountService()
}
