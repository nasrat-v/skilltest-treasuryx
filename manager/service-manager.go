package manager

import (
	"skilltest-treasuryx/database"
	"skilltest-treasuryx/router"
	"skilltest-treasuryx/server"
)

type ServiceManager struct {
	_server   server.Server
	_router   router.Router
	_database database.Database
}

// default constructor
func New() ServiceManager {
	return ServiceManager{}
}

func (x *ServiceManager) Create() error {
	if err := x._database.Create(); err != nil {
		return err
	}
	x._server.Create()
	x._router.Create(&x._database)
	return nil
}

func (x *ServiceManager) Start() error {
	x._router.LoadRoutes(x._server.Engine)
	defer x._database.Close() // defer Database closing
	return x._server.Start()
}
