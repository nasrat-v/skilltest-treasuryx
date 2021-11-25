package router

import (
	"skilltest-treasuryx/controller"
	"skilltest-treasuryx/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	_controller controller.Controller
}

// default constructor
func New() Router {
	return Router{}
}

func (x *Router) LoadRoutes(server *gin.Engine) {
	converter := server.Group("/api")
	{
		converter.POST("/payment", middleware.BasicAuth, x._controller.Payment)
	}
}
