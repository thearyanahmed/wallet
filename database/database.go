package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"os"
)

type manager struct {
	db *gorm.DB
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
	Manager = manager{db: db}
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
// does not work.

func (manager *manager) CreateRecord(value interface{}, table string) (*interface{},[]error)  {
	errs := manager.db.Create(&value).GetErrors()

	if len(errs) > 0 {
		return nil, errs
	}
	return &value, nil
}