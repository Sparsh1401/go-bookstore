package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

func GetDB() *gorm.DB {
	Db = connectDb()
	return Db
}

func connectDb() *gorm.DB {
	dsn := "root:1234567890@(127.0.0.1:3306)/bookstore?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	return db
}
