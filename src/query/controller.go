package query

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type controller struct {
	service Service
	router  *mux.Router
}

func newController(service Service, router *mux.Router) controller {
	newController := controller{service: service, router: router}

	router.HandleFunc("/queries", newController.getQueryByID).Methods("GET")

	return newController
}

func writeResponse(writer http.ResponseWriter, data interface{}, code int) {
	responseBody, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(code)
	writer.Header().Add("Content-type", "application/json")
	writer.Write(responseBody)
}

func (c *controller) getQueryByID(writer http.ResponseWriter, request *http.Request) {
	idString := request.URL.Query().Get("id")
	queryID, err := strconv.ParseUint(idString, 0, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	query, err := c.service.GetQueryByID(queryID)
	if err != nil {
		writeResponse(writer, err.Error, http.StatusInternalServerError)
	}

	writeResponse(writer, query, http.StatusOK)
}
