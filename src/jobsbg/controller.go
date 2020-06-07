package jobsbg

import (
	"net/http"

	"github.com/gorilla/mux"
)

type controller struct {
	scheduler *Scheduler
	router    *mux.Router
}

func newContoller(scheduler *Scheduler, router *mux.Router) controller {
	newController := controller{scheduler: scheduler, router: router}

	router.HandleFunc("/modules/jobsbg/scheduler/start", newController.startScheduler).Methods("POST")

	return newController
}

func (c *controller) startScheduler(writer http.ResponseWriter, request *http.Request) {
	c.scheduler.Start()

	writer.WriteHeader(http.StatusAccepted)
}
