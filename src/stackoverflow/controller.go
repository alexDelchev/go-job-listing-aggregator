package stackoverflow

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

	router.HandleFunc("/modules/stackoverflow/scheduler/start", newController.startScheduler).Methods("POST")
	router.HandleFunc("/modules/stackoverflow/scheduler/stop", newController.stopScheduler).Methods("DELETE")

	return newController
}

func (c *controller) startScheduler(writer http.ResponseWriter, request *http.Request) {
	c.scheduler.Start()

	writer.WriteHeader(http.StatusAccepted)
}

func (c *controller) stopScheduler(writer http.ResponseWriter, request *http.Request) {
	c.scheduler.Stop()

	writer.WriteHeader(http.StatusAccepted)
}
