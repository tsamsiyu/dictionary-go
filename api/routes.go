package api

import (
	"dict/api/controllers"
	routing "github.com/buaazp/fasthttprouter"
)

var router *routing.Router

// Router applies mapping between url string and controllers
func Router() *routing.Router {
	if router == nil {
		router = routing.New()
		router.GET("/users", controllers.UsersList)
	}
	return router
}
