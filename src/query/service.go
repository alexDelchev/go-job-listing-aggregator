package query

import "log"

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

// NewService returns a Service interface which holds the pointer
// to a new ServiceImplementation struct.
func NewService(repository repository) Service {
	return &ServiceImplementation{repository: repository}
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

// CreateQuery persists the given Query model into the database. The resulting
// query is returned, along with an error if one has occured in the repository
// layer.
func (s *ServiceImplementation) CreateQuery(query Query) (Query, error) {
	log.Printf("Persisting Query %+v", query)
	createdID, err := s.repository.insertQuery(query)
	if err != nil {
		return Query{}, err
	}
	return s.GetQueryByID(createdID)
}

// UpdateQuery updates the given Query model.  The resulting query is returned,
// along with an error if one has occured in the repository layer.
func (s *ServiceImplementation) UpdateQuery(query Query) (Query, error) {
	log.Printf("Updating Query %+v", query)
	_, err := s.repository.updateQuery(query)
	if err != nil {
		return Query{}, err
	}
	return s.GetQueryByID(query.ID)
}

// ActivateQuery gets the Query for the given ID, sets Query.Active to true,
// and persists the new state in the datase. The resuling query is returned,
// along with an error if one has occurred in the repository layer.
func (s *ServiceImplementation) ActivateQuery(id uint64) (Query, error) {
	log.Printf("Activating Query %d", id)
	query, err := s.GetQueryByID(id)
	if err != nil {
		return query, err
	}
	query.Active = true
	return s.UpdateQuery(query)
}

// DeactivateQuery gets the Query for the given ID, sets Query.Active to false,
// and persists the new state in the datase. The resuling query is returned,
// along with an error if one has occurred in the repository layer.
func (s *ServiceImplementation) DeactivateQuery(id uint64) (Query, error) {
	log.Printf("Deactivating Query %d", id)
	query, err := s.GetQueryByID(id)
	if err != nil {
		return query, err
	}
	query.Active = false
	return s.UpdateQuery(query)
}
