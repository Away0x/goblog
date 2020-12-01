package app

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB
var router *mux.Router

// SetDB 存储数据库 instance 便于项目中使用
func SetDB(d *sqlx.DB) {
	db = d
}

// SetRouter 存储路由 instance 便于项目中使用
func SetRouter(r *mux.Router) {
	router = r
}

// DB 获取数据库
func DB() *sqlx.DB {
	return db
}

// Router 获取路由对象
func Router() *mux.Router {
	return router
}
