package routes

import (
	"goblog/controllers"
	"goblog/routes/middlewares"

	"github.com/gorilla/mux"
)

// InitRoutes 初始化路由
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	registerPage(router, &controllers.Page{})
	registerArticle(router, &controllers.Articles{Router: router})

	// 中间件：强制内容类型为 HTML
	router.Use(middlewares.ForceHTMLMiddleware)
	// 去除请求路径尾部 /
	middlewares.RemoveTrailingSlash(router)

	return router
}
