package listing

import (
	"database/sql"

	"github.com/gorilla/mux"
)

// Module contains a repository, Service, and a controller,
// with the Service being exported.
type Module struct {
	repository repository
	Service    Service
	controller controller
}

// NewDefaultModule returns a new Module instance with the
// repository, Service, and controller having the default
// dependencies injected.
func NewDefaultModule(database *sql.DB, router *mux.Router) Module {
	repository := newRepository(database)
	service := NewService(repository)
	controller := newController(service, router)

	return Module{repository: repository, Service: service, controller: controller}
}
