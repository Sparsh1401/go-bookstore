package main

import (
	"database/sql"
	"fmt"
)

var Db *sql.DB

func initDb() *sql.DB {
	Db = connectDb()
	return Db
}

func connectDb() *sql.DB {
	dsn := "root:1234567890@(127.0.0.1:3306)/bookstore?parseTime=true"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Error connnecting to server:", err)
	}

	return db
}
