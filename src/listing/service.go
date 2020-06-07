package listing

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

// GetListingByID return the Listing for the given internal ID.
func (s *ServiceImplementation) GetListingByID(id uint64) (Listing, error) {
	return s.repository.getListingByID(id)
}

// GetLatestListingsBySourceName returns the last 100 listings for the given
// source name.
func (s *ServiceImplementation) GetLatestListingsBySourceName(sourceName string) ([]Listing, error) {
	return s.repository.getLatestListingsBySourceName(sourceName, listingsBySourceNameLimit)
}
