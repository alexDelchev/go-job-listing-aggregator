package jobsbg

import (
	"time"
)

// Scheduler executes Service.RunForActiveQueries at a given time period.Scheduler
// This scheduled execution can be also be started/stopped.
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
