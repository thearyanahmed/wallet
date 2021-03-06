package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/thearyanahmed/wallet/database"
	"github.com/thearyanahmed/wallet/internal/reminder"
	"github.com/thearyanahmed/wallet/modules/currency"
	"github.com/thearyanahmed/wallet/modules/organization"
	user "github.com/thearyanahmed/wallet/modules/user"
	"github.com/thearyanahmed/wallet/modules/wallet"
	"github.com/thearyanahmed/wallet/oauth"
	"github.com/thearyanahmed/wallet/schema"
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

	lock.Lock()

	db, err := database.Connect()

	if err != nil {
		reminder.Remind("[Set log]. Failed to connect to database.")
		log.Fatal(err)
		return
	}
	lock.Unlock()

	fmt.Println("Migrating.")

	schema.Migrate()

	fmt.Println("Migration completed.")

	defer db.Close()

	router := mux.NewRouter()

	oauth.Boot(router)
	registerRoutes(router)

	log.Fatal(http.ListenAndServe(":9096", router))
}

func registerRoutes(router *mux.Router) {
	currency.RegisterRoutes(router)
	wallet.RegisterRoutes(router)
	user.RegisterRoutes(router)
	organization.RegisterRoutes(router)
}

func loadEnvOrExit() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(ENV_UNAVAILABLE)
	}
}

