package query

type repositoryMock struct {
	updateQueryParameter Query
}

func (r *repositoryMock) getQueryByID(id uint64) (Query, error) {
	return Query{}, nil
}

func (r *repositoryMock) getAllQueries() ([]Query, error) {
	return []Query{}, nil
}

func (r *repositoryMock) getActiveQueries() ([]Query, error) {
	return []Query{}, nil
}

func (r *repositoryMock) getInactiveQueries() ([]Query, error) {
	return []Query{}, nil
}

func (r *repositoryMock) insertQuery(query Query) (uint64, error) {
	return 1, nil
}

func (r *repositoryMock) updateQuery(query Query) (bool, error) {
	r.updateQueryParameter = query
	return true, nil
}
