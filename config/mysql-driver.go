package config

import (
	"fmt"
	"os"

	_migrationBooks "rest-api-clean-arch/features/books/data"
	_migrationUsers "rest-api-clean-arch/features/users/data"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {

	config := map[string]string{
		"DB_USERNAME": os.Getenv("DB_USER"),
		"DB_PASSWORD": os.Getenv("DB_PASSWORD"),
		"DB_HOST":     os.Getenv("DB_HOST"),
		"DB_PORT":     os.Getenv("DB_PORT"),
		"DB_NAME":     os.Getenv("DB_NAME"),
	}

	connString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=UTC",
		config["DB_USERNAME"],
		config["DB_PASSWORD"],
		config["DB_HOST"],
		config["DB_PORT"],
		config["DB_NAME"],
	)

	db, err := gorm.Open(mysql.Open(connString), &gorm.Config{})

	if err != nil {
		panic(err.Error())
	}

	InitMigrate(db)

	return db

}

func InitMigrate(db *gorm.DB) {
	db.AutoMigrate(&_migrationUsers.User{})
	db.AutoMigrate(&_migrationBooks.Book{})
}
