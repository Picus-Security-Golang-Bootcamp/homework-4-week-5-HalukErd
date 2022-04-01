package postgres

import (
	"fmt"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DbInstance *gorm.DB
var once sync.Once

func NewPsqlDB() *gorm.DB {
	once.Do(func() {
		db, err := connectPsqlDB()
		if err != nil {
			panic(err)
		}
		DbInstance = db
	})
	return DbInstance
}

func connectPsqlDB() (*gorm.DB, error) {
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("LIBRARY_DB_HOST"),
		os.Getenv("LIBRARY_DB_PORT"),
		os.Getenv("LIBRARY_DB_USERNAME"),
		os.Getenv("LIBRARY_DB_NAME"),
		os.Getenv("LIBRARY_DB_PASSWORD"),
	)

	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("database could not be opened. Cause: %v", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDb.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("database ping has no error")

	return db, nil
}
