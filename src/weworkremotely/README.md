# weworkremotely package

This package is responsible for loading data from the stackoverflow jobs rss feed. An rss feed is fetched for each active query, afterwards job listing data is extracted from the xml document. The different job listings are filtered based on the values of the current active query. The relevant listings for said query are then transformed to the common model and persisted in the database.

This happens on a scheduled basis with an interval of 1 minnute.

The controller exposes start/stop controls of the scheduler to the application's REST API.

## Contents
```
    └── weworkremotely
        ├── client.go            - Loads the WeWorkRemotely rss feed
        ├── controller.go        - Exposes scheduling control
        ├── injector.go          - Dependency injection for the weworkremotely package
        ├── model.go             - WeWorkRemotely rss job listing model
        ├── scheduling.go        - WeWorkRemotely scraping scheduling
        ├── scraper.go           - Transforms the rss feed and persists the data
        ├── scraper_test.go      - Job listing filtering test
        ├── transform.go         - Transforms the rss model to the common model
        └── transform_test.go    - Transform tests
```