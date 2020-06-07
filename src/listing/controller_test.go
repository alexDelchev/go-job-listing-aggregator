package listing

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

type serviceMock struct{}

func (s serviceMock) GetListingByID(id uint64) (Listing, error) {
	if id == 0 {
		return Listing{}, errors.New("Sample error")
	}

	return Listing{}, nil
}

func (s serviceMock) GetLatestListingsBySourceName(sourceName string) ([]Listing, error) {
	if sourceName == "" {
		return []Listing{}, errors.New("Sample error")
	}

	return []Listing{}, nil
}

func (s serviceMock) GetListingsByQueryID(id uint64) ([]Listing, error) {
	if id == 0 {
		return []Listing{}, errors.New("Sample error")
	}

	return []Listing{}, nil
}

func (s serviceMock) GetListingsByQueryIDAndSourceName(queryID uint64, sourceName string) ([]Listing, error) {
	if queryID == 0 && sourceName == "" {
		return []Listing{}, errors.New("Sample error")
	}

	return []Listing{}, nil
}

func (s serviceMock) GetSourceNames() ([]string, error) {
	return []string{}, errors.New("Sample error")
}

func (s serviceMock) ListingExistsInDB(listing Listing) (bool, error) {
	if listing.ID == 0 {
		return false, errors.New("Sample error")
	}

	return false, nil
}

func (s serviceMock) CreateListing(listing Listing) (Listing, error) {
	if listing.ExternalID == "" {
		return listing, errors.New("Sample error")
	}

	return listing, nil
}

func (s serviceMock) CreateListings(listings []Listing) {
	return
}

func constructController() controller {
	listingService := serviceMock{}
	router := mux.NewRouter()

	return newController(listingService, router)
}

func testResponseStatus(t *testing.T, writer *httptest.ResponseRecorder, expectedStatusCode int) {
	actualStatusCode := writer.Result().StatusCode

	if expectedStatusCode != actualStatusCode {
		t.Errorf("%s failed: Expected status code %d, got %d",
			t.Name(), expectedStatusCode, actualStatusCode)
	}
}

func TestGetListingByID(t *testing.T) {
	listingController := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/listings?id=0", nil)

	listingController.getListingByID(writer, request)

	testResponseStatus(t, writer, http.StatusInternalServerError)
}

func TestGetListingByIDBadRequest(t *testing.T) {
	listingController := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/listings?id=a", nil)

	listingController.getListingByID(writer, request)

	testResponseStatus(t, writer, http.StatusBadRequest)
}

func TestGetSourceNames(t *testing.T) {
	listingController := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/sourceNames", nil)

	listingController.getSourceNames(writer, request)

	testResponseStatus(t, writer, http.StatusInternalServerError)
}

func TestGetLatestListingsBySourceName(t *testing.T) {
	listingController := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/listings/", nil)

	listingController.getLatestListingsBySourceName(writer, request)

	testResponseStatus(t, writer, http.StatusInternalServerError)
}

func TestGetListingsByQueryID(t *testing.T) {
	listingController := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/listings/query?id=0", nil)

	listingController.getListingsByQueryID(writer, request)

	testResponseStatus(t, writer, http.StatusInternalServerError)
}

func TestGetListingsByQueryIDBadRequest(t *testing.T) {
	listingController := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/listings/query?id=a", nil)

	listingController.getListingsByQueryID(writer, request)

	testResponseStatus(t, writer, http.StatusBadRequest)
}

func TestGetListingsByQueryIDAndSourceName(t *testing.T) {
	listingController := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/listings//query?id=0", nil)

	listingController.getListingsByQueryIDAndSourceName(writer, request)

	testResponseStatus(t, writer, http.StatusInternalServerError)
}

func TestGetListingsByQueryIDAndSourceNameBadRequest(t *testing.T) {
	listingController := constructController()

	writer := httptest.NewRecorder()
	request := httptest.NewRequest("GET", "/listings//query?id=a", nil)

	listingController.getListingsByQueryIDAndSourceName(writer, request)

	testResponseStatus(t, writer, http.StatusBadRequest)
}
