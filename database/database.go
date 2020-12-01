package database

import (
	"goblog/utils"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

// InitDB 连接数据库
func InitDB() *sqlx.DB {
	var (
		err error
		db  *sqlx.DB
	)
	config := mysql.Config{
		User:                 "root",
		Passwd:               "a12345678",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		DBName:               "goblog",
		AllowNativePasswords: true,
	}

	db, err = sqlx.Connect("mysql", config.FormatDSN())
	utils.CheckError(err)

	// 设置最大连接数
	db.SetMaxOpenConns(25)
	// 设置最大空闲连接数
	db.SetMaxIdleConns(25)
	// 设置每个链接的过期时间
	db.SetConnMaxLifetime(5 * time.Minute)

	// 尝试连接，失败会报错
	err = db.Ping()
	utils.CheckError(err)

	// create tables
	createTables(db)

	return db
}

func createTables(db *sqlx.DB) {
	createArticleTable(db)
}

func createArticleTable(db *sqlx.DB) {
	schema := `CREATE TABLE IF NOT EXISTS articles(
		id bigint(20) PRIMARY KEY AUTO_INCREMENT NOT NULL,
		title varchar(255) COLLATE utf8mb4_unicode_ci NOT NULL,
		body longtext COLLATE utf8mb4_unicode_ci)`

	_, err := db.Exec(schema)
	utils.CheckError(err)
}
