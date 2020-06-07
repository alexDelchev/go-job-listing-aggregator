package weworkremotely

import (
	"time"
)

// Scheduler calls Service.RunForActiveQueries at a given
// interval/
type Scheduler struct {
	Scraper     *Scraper
	stopChannel chan<- bool
}

type task func()

func schedule(action task, interval time.Duration) chan<- bool {
	ticker := time.NewTicker(interval)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case <-ticker.C:
				action()
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}
