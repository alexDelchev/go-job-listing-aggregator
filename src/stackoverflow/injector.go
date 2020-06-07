package stackoverflow

// Module contains a Scraper, Scheduler, and a controller instance,
// with the Scraper and Scheduler being exported.
type Module struct {
	Scraper    Scraper
	Scheduler  Scheduler
	controller controller
}
