package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// PageController 普通页面路由
type PageController interface {
	Home(w http.ResponseWriter, r *http.Request)
	About(w http.ResponseWriter, r *http.Request)
	NotFound(w http.ResponseWriter, r *http.Request)
}

func registerPage(router *mux.Router, controller PageController) {
	router.HandleFunc("/", controller.Home).Methods("GET").Name("home")
	router.HandleFunc("/about", controller.About).Methods("GET").Name("about")
	// 自定义 404 页面
	router.NotFoundHandler = http.HandlerFunc(controller.NotFound)
}
