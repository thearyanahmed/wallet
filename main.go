package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/thearyanahmed/wallet/database"
	"github.com/thearyanahmed/wallet/internal/reminder"
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

	db, err := database.Connect()

	if err != nil {
		reminder.Remind("[Set log]. Failed to connect to database.")
		log.Fatal(err)
		return
	}
	lock.Unlock()

	fmt.Println("Migrating.")

	migrate(db)
	fmt.Println("Migration completed.")

	defer db.Close()


	registerRoutes()
	log.Fatal(http.ListenAndServe(":9096", nil))
}

func registerRoutes() {
	http.HandleFunc("/protected", oauth.Auth(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, I'm protected."))
	}))

}

func migrate(db *gorm.DB) {
	//db.AutoMigrate(&model.User{})
	//db.AutoMigrate(&model2.Account{})
	//db.AutoMigrate(&model3.UserWallet{})
}


func loadEnvOrExit() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(ENV_UNAVAILABLE)
	}
}

