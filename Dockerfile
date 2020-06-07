FROM golang:1.14
 
RUN mkdir -p /go-job-listing-aggregator
 
WORKDIR /go-job-listing-aggregator
 
COPY . /go-job-listing-aggregator

COPY go.mod .

COPY go.sum .

RUN go mod download

RUN go test ./... --tags=test
 
RUN go build ./src/main.go
 
CMD ["./main"]