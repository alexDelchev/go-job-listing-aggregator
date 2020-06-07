// +build test

package query

import "testing"

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

func TestRepositoryActivateQuery(t *testing.T) {
	mock := repositoryMock{}
	var mockInterface repository = &mock

	service := NewService(mockInterface)

	service.ActivateQuery(1)

	expectedActiveStatus := true
	actualActiveStatus := mock.updateQueryParameter.Active

	if expectedActiveStatus != actualActiveStatus {
		t.Errorf("%s failed: Expected active status %t, got %t",
			t.Name(), expectedActiveStatus, actualActiveStatus)
	}
}

func TestRepositoryDeactivateQuery(t *testing.T) {
	mock := repositoryMock{}
	var mockInterface = &mock

	service := NewService(mockInterface)

	service.DeactivateQuery(1)

	expectedActiveStatus := false
	actualActiveStatus := mock.updateQueryParameter.Active

	if expectedActiveStatus != actualActiveStatus {
		t.Errorf("%s failed: Expected active status %t, got %t",
			t.Name(), expectedActiveStatus, actualActiveStatus)
	}
}
