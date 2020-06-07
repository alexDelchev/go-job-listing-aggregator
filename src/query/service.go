package query

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

type ServiceImplementation struct {
	repository repository
}
