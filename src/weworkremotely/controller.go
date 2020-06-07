package weworkremotely

import (
	"github.com/gorilla/mux"
)

type controller struct {
	scheduler *Scheduler
	router    *mux.Router
}

func newContoller(scheduler *Scheduler, router *mux.Router) controller {
	newController := controller{scheduler: scheduler, router: router}

	return newController
}
