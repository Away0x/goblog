package main

import (
	"fmt"
	"goblog/app"
	"goblog/database"
	"goblog/routes"
	"goblog/utils"
	"net/http"
)

func main() {
	db := database.InitDB()
	defer db.Close()

	router := routes.InitRoutes()

	app.SetDB(db)
	app.SetRouter(router)

	err := http.ListenAndServe(":3001", app.Router())
	utils.CheckError(err)
	fmt.Printf("server on 3001")
}
