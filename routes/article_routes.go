package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ArticleController article 路由
type ArticleController interface {
	Show(w http.ResponseWriter, r *http.Request)
	Index(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Store(w http.ResponseWriter, r *http.Request)
	Edit(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

func registerArticle(router *mux.Router, controller ArticleController) {
	router.HandleFunc("/articles/{id:[0-9]+}", controller.Show).Methods("GET").Name("articles.show")
	router.HandleFunc("/articles", controller.Index).Methods("GET").Name("articles.index")
	router.HandleFunc("/articles", controller.Store).Methods("POST").Name("articles.store")
	router.HandleFunc("/articles/create", controller.Create).Methods("GET").Name("articles.create")
	router.HandleFunc("/articles/{id:[0-9]+}/edit", controller.Edit).Methods("GET").Name("articles.edit")
	router.HandleFunc("/articles/{id:[0-9]+}", controller.Update).Methods("POST").Name("articles.update")
	router.HandleFunc("/articles/{id:[0-9]+}/delete", controller.Delete).Methods("POST").Name("articles.delete")
}
