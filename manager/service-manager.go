package manager

import (
	"skilltest-treasuryx/router"
	"skilltest-treasuryx/server"
)

type ServiceManager struct {
	_server server.Server
	_router router.Router
}

// default constructor
func New() ServiceManager {
	return ServiceManager{}
}

func (x *ServiceManager) Create() {
	x._server.Create()
}

func (x *ServiceManager) Start() error {
	x._router.LoadRoutes(x._server.Engine)
	return x._server.Start()
}
