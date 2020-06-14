package main

import (
	"github.com/thearyanahmed/wallet/oauth"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	_ "time"

)


func main() {

	oauth.Boot()

	http.HandleFunc("/protected", oauth.Auth(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm protected"))
	}, oauth.Server))

	log.Fatal(http.ListenAndServe(":9096", nil))
}