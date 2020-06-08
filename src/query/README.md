# query package

This package is responsible for servicing the `public.query` table. Also defines the commong query model which is used as a parameter to the different scrapers.

This package reads/writes data using the Go PostgreSQL driver and exposes these operations through the exported Service struct.

The controller also exposes the read/write operations to the application's REST API.

## Contents
```
    ├── query
    │   ├── controller.go        - Read/write API operations on the query table
    │   ├── controller_test.go   - Status code tests
    │   ├── injector.go          - Dependency injetion for the query package
    │   ├── model.go             - Commong query model
    │   ├── repository.go        - Read/write operation on the query table
    │   ├── service.go           - Exports the read/write operations on the query table
    │   └── service_test.go      - Write operation tests
```