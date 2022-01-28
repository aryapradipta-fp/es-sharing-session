Requirements
- `docker`
- `go`
- [ES chrome extension](https://chrome.google.com/webstore/detail/elasticsearch-head/ffmkiejjmecolpfloofpjologoblkegm), optional

How to run
1. `go mod download` to download all of the dependencies
2. `docker-compose up` to start up ES cluster
3. `go run main.go` to seed up the data

The data will be in `movie` index
