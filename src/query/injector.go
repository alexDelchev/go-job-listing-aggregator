package query

// Module contains a query repository, Service, controller,
// with the service being exported.
type Module struct {
	repository repository
	Service    Service
	controller controller
}
