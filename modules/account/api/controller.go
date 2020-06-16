package api

import (
	"fmt"
	"github.com/thearyanahmed/wallet/modules/account/service"
	"net/http"
)

func testFunc(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Controller two func actually.")
	service.TestAccountServic()
}

func finalTestController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Miau")
	service.FinalTest()
}