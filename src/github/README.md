# github package

This package is responsible for loading data from the github jobs api based on the active queries. Data from the API response is transformed to the common model and persited in the database.

This happens on a scheduled basis with an interval of 1 minnute.

The controller exposes start/stop controls of the scheduler to the application's REST API.

## Contents
```
    ├── github
    │   ├── client.go            - HTTP client for the github jobs API
    │   ├── controller.go        - Exposes scheduling control
    │   ├── injector.go          - Dependency injection for the github pacakge
    │   ├── model.go             - GitHub API model
    │   ├── scheduling.go        - Github scraping scheduling
    │   ├── scraper.go           - Loads and persists data from the API
    │   ├── transform.go         - Transforms the GitHub API model to the common model
    │   └── transform_test.go    - Transform tests
```