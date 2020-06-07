package listing

// Module contains a repository, Service, and a controller,
// with the Service being exported.
type Module struct {
	repository repository
	Service    Service
	controller controller
}
