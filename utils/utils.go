package utils

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CheckError log error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// GetRouteVariable get var from router
func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

// RouteName2URL 通过路由名称来获取 URL
func RouteName2URL(router *mux.Router, routeName string, pairs ...string) string {
	url, err := router.Get(routeName).URL(pairs...)
	if err != nil {
		CheckError(err)
		return ""
	}

	return url.String()
}

// Int64ToString 将 int64 转换为 string
func Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}
