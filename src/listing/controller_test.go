package listing

import (
	"errors"

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
