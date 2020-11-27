package models

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// DB client
var DB *sqlx.DB

func init() {
	var err error
	config := mysql.Config{
		User:                 "root",
		Passwd:               "a12345678",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	DB, err = sqlx.Connect("mysql", config.FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}

	// create tables
	createArticleTable()
}
