// Package database sets up the db
package database

import (
	"log"

	"github.com/adrr-dev/inventory-app/internal/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Database struct {
	DataFile string
}

func (d Database) InitializeDB() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open(d.DataFile), &gorm.Config{})
	check(err)

	err = DB.AutoMigrate(&repository.User{}, &repository.Inventory{})
	check(err)
	return DB
}
