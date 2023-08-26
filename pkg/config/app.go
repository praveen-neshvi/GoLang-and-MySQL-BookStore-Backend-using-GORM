package config

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Connect() {
	dsn := "root:root@tcp(localhost:3306)/bookstore_go?charset=utf8mb4&parseTime=true&loc=Local"

	connection, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	db = connection
}

func GetDB() *gorm.DB {
	return db
}
