package query

// Service is used to implement query-specific business
// logic and read/write query data through the repository.
type Service interface {
	GetQueryByID(id uint64) (Query, error)

	GetActiveQueries() ([]Query, error)

	GetInactiveQueries() ([]Query, error)

	GetAllQueries() ([]Query, error)

	CreateQuery(query Query) (Query, error)

	UpdateQuery(query Query) (Query, error)

	ActivateQuery(id uint64) (Query, error)

	DeactivateQuery(id uint64) (Query, error)
}

// ServiceImplementation implements the service interface.
// Contains a query.repository instance.
type ServiceImplementation struct {
	repository repository
}

// GetQueryByID returns the Query for the given ID. The returned error
// is not nil if there was an error in the repository layer.
func (s *ServiceImplementation) GetQueryByID(id uint64) (Query, error) {
	return s.repository.getQueryByID(id)
}

// GetActiveQueries returns a slice of Queries where Query.Active == true.
// The returned error is not nil if there was an error in the repository
// layer.
func (s *ServiceImplementation) GetActiveQueries() ([]Query, error) {
	return s.repository.getActiveQueries()
}

// GetInactiveQueries returns a slice of Queries where Query.Active == false.
// The returned error is not nil if there was an error in the repository layer.
func (s *ServiceImplementation) GetInactiveQueries() ([]Query, error) {
	return s.repository.getInactiveQueries()
}

// GetAllQueries returns a slice containing all queries. The returned error
// is not nil if there was an error in the repoistory layer.
func (s *ServiceImplementation) GetAllQueries() ([]Query, error) {
	return s.repository.getAllQueries()
}
