package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
	"sync"
)

type manager struct {
	db *gorm.DB
	mux sync.Mutex
	useTransaction bool
}

var Manager manager

func Connect() ( *gorm.DB, error ) {
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")

	db, err := gorm.Open("mysql", prepareConnectionString(username, password, database, host, port))

	if err != nil {
		return nil,err
	}

	bootsrap(db)
	Manager = manager{db: db,useTransaction: false}
	return db, nil
}

func bootsrap(database *gorm.DB) {
	database.SingularTable(false)
}

func DB () *gorm.DB {
	return Manager.db
}

func Close() {
	Manager.db.Close()
}

func prepareConnectionString(username, password, database, host, port string) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		username,password,host,port,database,
	)
}

func (manager *manager) Transaction(callback func() error) error {

	//manager.BeginTransaction()
	//
	//err := callback()
	//
	//if err != nil {
	//	fmt.Println("Got error from callback")
	//	manager.Rollback()
	//	return err
	//} else {
	//	manager.Commit()
	//	return nil
	//}

	manager.db = manager.db.Begin()

	err := callback()

	if err != nil {
		manager.db.Begin().Rollback()
		return err
	}

	manager.db.Commit()
	return nil
}

func (manager *manager) BeginTransaction() {
	DB().Exec("START TRANSACTION")
}

func (manager *manager) Commit() {
	DB().Exec("COMMIT")
}

func (manager *manager) Rollback() {
	DB().Exec("ROLLBACK")
}
