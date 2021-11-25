package router

import (
	"skilltest-treasuryx/src/controller"
	"skilltest-treasuryx/src/database"
	"skilltest-treasuryx/src/middleware"

	"github.com/gin-gonic/gin"
)

type Router struct {
	_controller controller.Controller
}

// default constructor
func New() Router {
	return Router{}
}

func (x *Router) Create(database *database.Database) {
	x._controller.Create(database)
}

// Initialise routes with controllers handlers
func (x *Router) LoadRoutes(server *gin.Engine) {
	converter := server.Group("/api")
	{
		converter.POST("/payment", middleware.BasicAuth, x._controller.Payment)
	}
}
