package api

import (
	"fmt"
	"net/http"
)

func controllerTwoFunc(w http.ResponseWriter,r *http.Request) {
	fmt.Println("Controller two func actually.")
}