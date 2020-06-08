# stackoverflow package

This package is responsible for loading data from the stackoverflow jobs rss feed. An rss feed is fetched for each active query, afterwards job listing data is extracted from the xml document, transformed to the common listing model and persisted in the database.

This happens on a scheduled basis with an interval of 1 minnute.

The controller exposes start/stop controls of the scheduler to the application's REST API.

## Contents
```
    ├── stackoverflow
    │   ├── client.go            - Loads the StackOverflow jobs rss feed
    │   ├── controller.go        - Exposes scheduling control
    │   ├── injector.go          - Dependency injection for the stackoverflow package
    │   ├── model.go             - StackOverflow rss job listing model
    │   ├── scheduling.go        - Stackoverflow scraping scheduling
    │   ├── scraper.go           - Transforms the rss feed and persists the data
    │   ├── transform.go         - Transforms the rss model to the common model
    │   └── transform_test.go    - Transform tests
```