# jobsbg package

This package is responsible for loading data from the jobs.bg website. Data is scraped from the html page for a given search query and transformed to the commong job listing model, after which it is persited in the database.

This happens on a scheduled basis with an interval of 1 minnute.

The controller exposes start/stop controls of the scheduler to the application's REST API.

## Contents
```
    ├── jobsbg
    │   ├── client.go            - HTTP client for the jobs.bg website
    │   ├── controller.go        - Exposes scheduling control
    │   ├── injector.go          - Dependency injection for the jobsbg package
    │   ├── model.go             - jobs.bg website data models
    │   ├── scheduling.go        - jobs.bg scraping scheduling
    │   ├── scraper.go           - Extracts and persists data from html 
    │   ├── scraper_test.go      - Model extraction from html tests
    │   ├── transform.go         - Transforms the jobsbg model to the common model
    │   └── transform_test.go    - Transform tests
```