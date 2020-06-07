package stackoverflow

import (
	"log"
	"time"
)

// Scheduler calls Scraper.RunForActiveQueries at a
// given interval.
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

// Start starts the scheduling.
func (s *Scheduler) Start() {
	log.Println("Starting StackOverflow Scheduler")
	action := s.Scraper.RunForActiveQueries
	interval := 1 * time.Minute

	channel := schedule(action, interval)

	if s.stopChannel != nil {
		s.stopChannel <- true
	}
	s.stopChannel = channel
}

// Stop stops the scheduling.
func (s *Scheduler) Stop() {
	if s.stopChannel != nil {
		log.Println("Stopping StackOverflow Scheduler")
		s.stopChannel <- true
		s.stopChannel = nil
	}
}
