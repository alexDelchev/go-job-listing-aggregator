# listing package

This package services the `public.listing` table and defines the common listing model which acts as a generic data model that is produced by the different scrapers. 

This package reads/writes data using the GO PostgreSQL driver, these operations are exported through the package's Service struct.

The package also contains a controller which exposes the read operations to the application's REST API.

## Contents
```
    ├── listing
    │   ├── controller.go        - Read API operations on the listing table
    │   ├── controller_test.go   - Status code tests
    │   ├── injector.go          - Dependency injection for the listing package
    │   ├── model.go             - Common job listing model
    │   ├── repository.go        - Read/write operations on the listing table
    │   ├── service.go           - Exports the read/write operations on the listing table
    │   └── service_test.go      - Write operations tests
```