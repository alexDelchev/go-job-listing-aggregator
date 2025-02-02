package listing

import (
	"log"
)

const listingsBySourceNameLimit uint16 = 20

// Service is used to implement listing-specific business
// logic and to read/write data using the repository.
type Service interface {
	GetListingByID(id uint64) (Listing, error)

	GetLatestListingsBySourceName(sourceName string) ([]Listing, error)

	GetListingsByQueryID(id uint64) ([]Listing, error)

	GetListingsByQueryIDAndSourceName(queryID uint64, sourceName string) ([]Listing, error)

	GetSourceNames() ([]string, error)

	ListingExistsInDB(listing Listing) (bool, error)

	CreateListing(listing Listing) (Listing, error)

	CreateListings(listings []Listing)
}

// ServiceImplementation implements the Service interface.
// Contains a repository instance which is not exported."
type ServiceImplementation struct {
	repository repository
}

// NewService returns a new service implementation which
// consists of a pointer to a ServiceImplementation struct.
func NewService(repository repository) Service {
	return &ServiceImplementation{repository: repository}
}

// GetListingByID return the Listing for the given internal ID.
func (s *ServiceImplementation) GetListingByID(id uint64) (Listing, error) {
	return s.repository.getListingByID(id)
}

// GetLatestListingsBySourceName returns the last 100 listings for the given
// source name.
func (s *ServiceImplementation) GetLatestListingsBySourceName(sourceName string) ([]Listing, error) {
	return s.repository.getLatestListingsBySourceName(sourceName, listingsBySourceNameLimit)
}

// GetListingsByQueryID returns all of the listings for the given query ID.
func (s *ServiceImplementation) GetListingsByQueryID(id uint64) ([]Listing, error) {
	return s.repository.getListingsByQueryID(id)
}

// GetListingsByQueryIDAndSourceName returns all of the listings for the given
// query ID and source name.
func (s *ServiceImplementation) GetListingsByQueryIDAndSourceName(queryID uint64, sourceName string) ([]Listing, error) {
	return s.repository.getListingsByQueryIDAndSourceName(queryID, sourceName)
}

// GetSourceNames returns all of distinct source names in the listing table.
func (s *ServiceImplementation) GetSourceNames() ([]string, error) {
	return s.repository.getSourceNames()
}

// ListingExistsInDB returns whether a Listing exists with the same extrenal ID
// and source name.
func (s *ServiceImplementation) ListingExistsInDB(listing Listing) (bool, error) {
	return s.repository.listingExists(listing.ExternalID, listing.SourceName)
}

// CreateListing persists the given listing and returns the resulting model.
func (s *ServiceImplementation) CreateListing(listing Listing) (Listing, error) {
	log.Printf("Persisting listing %+v", listing)
	createdID, err := s.repository.insertListing(&listing)
	if err != nil {
		return Listing{}, err
	}
	return s.GetListingByID(createdID)
}

// CreateListings checks for the existence of each listing and, if not found, persits it.
func (s *ServiceImplementation) CreateListings(listings []Listing) {
	for _, listing := range listings {
		exists, err := s.ListingExistsInDB(listing)
		if err != nil {
			continue
		}

		if exists {
			log.Printf(
				"Listing with external ID %s for source %s already exists in the database",
				listing.ExternalID, listing.SourceName)
			continue
		} else {
			s.CreateListing(listing)
		}
	}
}
