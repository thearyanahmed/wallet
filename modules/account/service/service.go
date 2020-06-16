package service

import (
	"fmt"
	"github.com/thearyanahmed/wallet/modules/account/repository"
	"github.com/thearyanahmed/wallet/modules/user/service"
)

func TestAccountServic() {
	fmt.Println("Account service.")
}

func CallThisAccountService(){
	fmt.Println("Call this account service")
}

func FinalTest() {
	service.TestHumanService()
}

type AccountService struct {
	repository.Repository
}

func (as *AccountService) SomeAccountServiceFunction() {
	fmt.Println("Some account service function.")
}