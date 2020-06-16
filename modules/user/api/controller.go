package api

import (
	"fmt"
	"github.com/thearyanahmed/wallet/modules/user/service"
	"net/http"
)

func testFunc(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Controller two func actually.")
	service.TestUserService()
}
func callFunc(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Controller two func actually.")
	service.SomethingElse()
}