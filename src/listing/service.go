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
