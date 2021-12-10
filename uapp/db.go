package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

// var url = "host=localhost user=postgres password=admin dbname=localdb port=5432 sslmode=disable"

var url = "host=db user=postgres password=admin dbname=udb port=5432 sslmode=disable"
var err error

func DbConnection() {
	Database, err = gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		fmt.Print(err.Error())
	}

	fmt.Println("DB Connected!")

	Database.AutoMigrate(&User{})
}
