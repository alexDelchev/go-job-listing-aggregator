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
	router.HandleFunc("/queries", newController.updateQuery).Methods("PUT").Headers("Content-type", "application/json")
	router.HandleFunc("/queries", newController.createQuery).Methods("POST").Headers("Content-type", "application/json")
	router.HandleFunc("/queries/all", newController.getAllQueries).Methods("GET")
	router.HandleFunc("/queries/active", newController.getActiveQueries).Methods("GET")
	router.HandleFunc("/queries/inactive", newController.getInactiveQueries).Methods("GET")
	router.HandleFunc("/queries/activate", newController.activateQuery).Methods("PATCH")
	router.HandleFunc("/queries/deactivate", newController.deactivateQuery).Methods("PATCH")

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

func (c *controller) createQuery(writer http.ResponseWriter, request *http.Request) {
	var query Query

	err := json.NewDecoder(request.Body).Decode(&query)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	query, err = c.service.CreateQuery(query)
	if err != nil {
		writeResponse(writer, err.Error, http.StatusInternalServerError)
	}

	writeResponse(writer, query, http.StatusOK)
}

func (c *controller) updateQuery(writer http.ResponseWriter, request *http.Request) {
	var query Query

	err := json.NewDecoder(request.Body).Decode(&query)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	if query.ID == 0 {
		http.Error(writer, "Invalid Query model", http.StatusBadRequest)
		return
	}

	query, err = c.service.UpdateQuery(query)
	if err != nil {
		writeResponse(writer, err.Error, http.StatusInternalServerError)
	}

	writeResponse(writer, query, http.StatusOK)
}

func (c *controller) getAllQueries(writer http.ResponseWriter, request *http.Request) {
	queries, err := c.service.GetAllQueries()
	if err != nil {
		writeResponse(writer, err.Error, http.StatusInternalServerError)
	}

	writeResponse(writer, queries, http.StatusOK)
}

func (c *controller) getActiveQueries(writer http.ResponseWriter, request *http.Request) {
	queries, err := c.service.GetActiveQueries()
	if err != nil {
		writeResponse(writer, err.Error, http.StatusInternalServerError)
	}

	writeResponse(writer, queries, http.StatusOK)
}

func (c *controller) getInactiveQueries(writer http.ResponseWriter, request *http.Request) {
	queries, err := c.service.GetInactiveQueries()
	if err != nil {
		writeResponse(writer, err.Error, http.StatusInternalServerError)
	}

	writeResponse(writer, queries, http.StatusOK)
}

func (c *controller) activateQuery(writer http.ResponseWriter, request *http.Request) {
	idString := request.URL.Query().Get("id")
	queryID, err := strconv.ParseUint(idString, 0, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	query, err := c.service.ActivateQuery(queryID)
	if err != nil {
		writeResponse(writer, err.Error, http.StatusInternalServerError)
	}

	writeResponse(writer, query, http.StatusOK)
}

func (c *controller) deactivateQuery(writer http.ResponseWriter, request *http.Request) {
	idString := request.URL.Query().Get("id")
	queryID, err := strconv.ParseUint(idString, 0, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	query, err := c.service.DeactivateQuery(queryID)
	if err != nil {
		writeResponse(writer, err.Error, http.StatusInternalServerError)
	}

	writeResponse(writer, query, http.StatusOK)
}
