package main

import (
	"github.com/joho/godotenv"
	"github.com/thearyanahmed/wallet/oauth"
	"log"
	"net/http"
	"os"
	_ "time"
)

const (
	ENV_UNAVAILABLE = 1001
)

func main() {

	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(ENV_UNAVAILABLE)
	}

	oauth.Boot()

	http.HandleFunc("/protected", oauth.Auth(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm protected"))
	}, oauth.Server))

	log.Fatal(http.ListenAndServe(":9096", nil))
}