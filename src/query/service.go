package query

// Service is used to implement query-specific business
// logic and read/write query data through the repository
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
// Contains a query.repository instance
type ServiceImplementation struct {
	repository repository
}
