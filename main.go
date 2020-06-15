package main

import (
	"github.com/joho/godotenv"
	"github.com/thearyanahmed/wallet/database"
	"github.com/thearyanahmed/wallet/oauth"
	"log"
	"net/http"
	"os"
	"sync"
)

const (
	ENV_UNAVAILABLE = 1001
)

var (
	lock = &sync.Mutex{}
)

func main() {
	loadEnvOrExit()
	oauth.Boot()

	lock.Lock()

	err := database.Connect()
	if err != nil {
		log.Fatal(err)
		return
	}

	lock.Unlock()
	defer database.Close()

	http.HandleFunc("/protected", oauth.Auth(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm protected"))
	}))

	log.Fatal(http.ListenAndServe(":9096", nil))
}

func loadEnvOrExit() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(ENV_UNAVAILABLE)
	}
}
