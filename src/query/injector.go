package query

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// Module contains a query repository, Service, controller,
// with the service being exported.
type Module struct {
	repository repository
	Service    Service
	controller controller
}

// NewDefaultModule returns a new Module instance where the,
// repository, Service, and contoller have the default dependencies
// injected.
func NewDefaultModule(database *sql.DB, router *mux.Router) Module {
	repository := newRepository(database)
	service := NewService(repository)
	controller := newController(service, router)

	return Module{repository: repository, Service: service, controller: controller}
}
