package api

import (
	"fmt"
	"github.com/thearyanahmed/wallet/internal/res"
	"net/http"
)

func controllerOneFunc(w http.ResponseWriter,r *http.Request) {
	res.Send(w,"Hello world",nil,200)
}

func anotherEndpont(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Hello world")
}