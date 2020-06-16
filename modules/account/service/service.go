package service

import (
	"fmt"
	"github.com/thearyanahmed/wallet/modules/account/repository"
	"github.com/thearyanahmed/wallet/modules/user/api"
)

func TestAccountServic() {
	fmt.Println("Account service.")
}

func CallThisAccountService(){
	fmt.Println("Call this account service")
}

func FinalTest() {
	api.TestHumanService()
}

type AccountService struct {
	repository.Repository
}

func (as *AccountService) SomeAccountServiceFunction() {
	fmt.Println("Some account service function.")
}