package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
	_port  string
}

// default constructor
func New() Server {
	return Server{}
}

func (x *Server) Create() {
	x.Engine = gin.Default()
	x._port = os.Getenv("PORT_SRV")
	if x._port == "" {
		x._port = "4242" // default port
	}
}

func (x *Server) Start() error {
	return x.Engine.Run(":" + x._port)
}
