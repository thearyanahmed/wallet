package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	 _ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)
var db *gorm.DB

func Connect() error {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	db, err := gorm.Open("mysql", prepareConnectionString(username, password, database, host, port))

	if err != nil {
		return err
	}

	bootsrap(db)

	return nil
}

func bootsrap(database *gorm.DB) {
	database.SingularTable(false)
}

func DB () *gorm.DB {
	return db
}

func Close() {
	db.Close()
}

func prepareConnectionString(username, password, database, host, port string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,password,host,port,database,
	)
}
