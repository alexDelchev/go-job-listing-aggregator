package listing

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

	router.HandleFunc("/listings", newController.getListingByID).Methods("GET")

	return newController
}

func writeResponse(writer http.ResponseWriter, data interface{}, code int) {
	responseBody, err := json.Marshal(data)
	if err != nil {
		log.Println(err)
		http.Error(writer, err.Error(), code)
		return
	}

	writer.WriteHeader(http.StatusOK)
	writer.Header().Add("Content-type", "application/json")
	writer.Write(responseBody)
}

func (c *controller) getListingByID(writer http.ResponseWriter, request *http.Request) {
	idString := request.URL.Query()["id"][0]
	ID, err := strconv.ParseUint(idString, 0, 64)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	listing, err := c.service.GetListingByID(ID)
	if err != nil {
		writeResponse(writer, err.Error, http.StatusInternalServerError)
	}

	writeResponse(writer, listing, http.StatusOK)
}
