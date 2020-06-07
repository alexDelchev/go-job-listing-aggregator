package jobsbg

// Module contains a Scraper, Scheduler, and a controller instance,
// with the Scheduler and Scraper exported.
type Module struct {
	Scraper    Scraper
	Scheduler  Scheduler
	controller controller
}
