package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

const DSN = "host=127.0.0.1 port=15432 user=test password=test dbname=test"

func main() {
	db, err := gorm.Open(
		postgres.Open(DSN),
		&gorm.Config{},
	)

	if err != nil {
		log.Fatal("Unable to open DB connection: ", err)
	}

	db.Migrator().CreateTable("users")
}
