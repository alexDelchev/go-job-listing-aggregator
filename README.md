# go-job-listing-aggregator

A web scraper written in go which periodically collects job position listings from 4 different sources and saves the data in a database. The application connects to a PostgreSQL database, where it creates its own versioned objects.

The application exposes its data through a RESTful API, which is also used to control scraping for the different sources.

The application is packaged into a docker container, docker-compose is used to create a system with an 
empty PostgreSQL database accessible through the docker network.

## Dependencies

 - `github.com/PuerkitoBio/goquery` - `v1.5.1` - used for text extraction from html documents.
 - `github.com/gorilla/mux` - `v1.7.4` - used for routing incoming http requests to handler functions.
 - `github.com/lib/pq` - `v1.6.0` - PostgreSQL GO driver.
 - `golang.org/x/net` - `v0.0.0-20200528225125-3c3fba18258b` - used for html parsing in combination with goquery.

## Structure

Check each package's README for details:
 - [database](src/config/database)
 - [main](src)
 - [github](src/github)
 - [jobsbg](src/jobsbg)
 - [listing](src/listing)
 - [query](src/query)
 - [stackoverflow](src/stackoverflow)
 - [testutils](src/testutils)
 - [weworkremotely](src/weworkremotely)


```
.
├── resources
│   ├── db
│   │   └── migration            - Database versioning
│   └── properties               - Application properties
└── src
    ├── config
    │   └── database             - Database connection configuration & migration validtion
    ├── github                   - GitHub Jobs API scraping
    ├── jobsbg                   - jobs.bg website scraping
    ├── listing                  - Common job listing model, read/write operations
    ├── query                    - Common query model, read/write operations
    ├── stackoverflow            - StackOverflow jobs rss feed scraping
    ├── testutils                - Common test utility functions
    └── weworkremotely           - WeWorkRemotely rss feed scraping
```

## API

`GET /listings?id=`
 - Returns the listing model for the given id.

`GET /listings/query?id=`
 - Returns all listings for the given query id.

`GET /listings/{sourceName}`
 - Returns the last 100 listings for the given source name.

`GET /listings/{sourceName}/query?id=`
 - Returns all the listings for the given query id and source name.

`GET /sourceNames`
 - Returns all distinct source names.
 
`GET /queries?id=`
 - Returns the query for the given id.

`PUT /queries`
 - Replaces the values for the given query. Returns the resulting model.

`POST /queries`
 - Creates a new query. Returns the resulting model.

`GET /queries/all`
 - Returns all queries.

`GET /queries/active`
 - Returns all active queries.

`GET /queries/inactive`
 - Returns all inactive queries.

`PATCH /queries/activate?id=`
 - Updates the Active field to true of the query for the given id. Returns the resulting model.

`PATCH /queries/deactivate?id=`
 - Updates the Active field to true of the query for the given id. Returns the resulting model.

`POST /modules/github/scheduler/start`
 - Starts the scheduled scraping of the github website. Returns status Accepted.

`DELETE /modules/github/scheduler/stop`
 - Stops the scheduled scraping of the github website. Returns status Accepted.

`POST /modules/jobsbg/scheduler/start`
 - Starts the scheduled scraping of the jobs.bg website. Returns status Accepted.

`DELETE /modules/jobsbg/scheduler/stop`
 - Stops the scheduled scraping of the jobs.bg website. Returns status Accepted.

`POST /modules/stackoverflow/scheduler/start`
 - Starts the scheduled scraping of the StackOverflow rss feed. Returns status Accepted.

`DELETE /modules/stackoverflow/scheduler/stop`
 - Stops the scheduled scraping of the StackOverflow rss feed. Returns status Accepted.

`POST /modules/weworkremotely/scheduler/start`
 - Starts the scheduled scraping of the WeWorkRemotely rss feed. Returns status Accepted.

`DELETE /modules/weworkremotely/scheduler/stop`
 - Stops the scheduled scraping of the WeWorkRemotely rss feed. Returns status Accepted.

## How to run

 - clone project
 - `docker-compose up --build` to start the system
 - `./create-queries.sh` to create some sample queries

