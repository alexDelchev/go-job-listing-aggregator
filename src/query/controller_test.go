package query

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type serviceMock struct{}

const errorQueryID uint64 = 99

func (s serviceMock) GetQueryByID(id uint64) (Query, error) {
	if id == 0 {
		return Query{}, errors.New("Sample error")
	}

	return Query{ID: id}, nil
}

func (s serviceMock) GetActiveQueries() ([]Query, error) {
	return []Query{}, errors.New("Sample error")
}

func (s serviceMock) GetInactiveQueries() ([]Query, error) {
	return []Query{}, errors.New("Sample error")
}

func (s serviceMock) GetAllQueries() ([]Query, error) {
	return []Query{}, errors.New("Sample error")
}

func (s serviceMock) CreateQuery(query Query) (Query, error) {
	if query.ID == errorQueryID {
		return query, errors.New("Sample error")
	}

	return query, nil
}

func (s serviceMock) UpdateQuery(query Query) (Query, error) {
	if query.ID == errorQueryID {
		return query, errors.New("Sample error")
	}

	return query, nil
}

func (s serviceMock) ActivateQuery(id uint64) (Query, error) {
	if id == errorQueryID {
		return Query{}, errors.New("Sample error")
	}

	return Query{}, nil
}

func (s serviceMock) DeactivateQuery(id uint64) (Query, error) {
	if id == errorQueryID {
		return Query{}, errors.New("Sample error")
	}

	return Query{}, nil
}

func constructController() controller {
	mock := serviceMock{}
	router := mux.NewRouter()

	return newController(mock, router)
}

func testStatusCode(t *testing.T, writer *httptest.ResponseRecorder, expectedStatusCode int) {
	actualStatusCode := writer.Result().StatusCode

	if expectedStatusCode != actualStatusCode {
		t.Errorf("%s failed: Expected status code %d, got %d",
			t.Name(), expectedStatusCode, actualStatusCode)
	}
}

func TestGetQueryByID(t *testing.T) {
	controller := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/queries?id=0", nil)

	controller.getQueryByID(writer, request)

	testStatusCode(t, writer, http.StatusInternalServerError)
}

func TestGetQueryByIDBadRequest(t *testing.T) {
	controller := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/queries?id=a", nil)

	controller.getQueryByID(writer, request)

	testStatusCode(t, writer, http.StatusBadRequest)
}

func TestGetAllQueries(t *testing.T) {
	controller := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/queries/all", nil)

	controller.getAllQueries(writer, request)

	testStatusCode(t, writer, http.StatusInternalServerError)
}

func TestGetActiveQueries(t *testing.T) {
	controller := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/queries/active", nil)

	controller.getActiveQueries(writer, request)

	testStatusCode(t, writer, http.StatusInternalServerError)
}

func TestGetInactiveQueries(t *testing.T) {
	controller := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/queries/inactive", nil)

	controller.getInactiveQueries(writer, request)

	testStatusCode(t, writer, http.StatusInternalServerError)
}

func TestControllerActivateQuery(t *testing.T) {
	controller := constructController()

	writer := httptest.NewRecorder()
	target := fmt.Sprintf("/queries/activate?id=%d", errorQueryID)
	request := httptest.NewRequest("PATCH", target, nil)

	controller.activateQuery(writer, request)

	testStatusCode(t, writer, http.StatusInternalServerError)
}

func TestControllerActivateQueryBadRequest(t *testing.T) {
	controller := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("PATCH", "/queries/activate?id=a", nil)

	controller.activateQuery(writer, request)

	testStatusCode(t, writer, http.StatusBadRequest)
}
