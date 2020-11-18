export BINARY_NAME = "imdb-cli"
export GO111MODULE=on

# global command
.PHONY: all
all: dependencies build test

.PHONY: dependencies
dependencies:
	go mod download

.PHONY: build
build:
	env GOOS=linux GOARCH=amd64 go build -o ./${BINARY_NAME} .

.PHONY: mock
mock:	
	mockgen -package scrappermock github.com/gufranmirza/imdb-cli/src/scrapper Scrapper > scrapper/scrappermock/mock.go

packages = \
	./scrapper\

.PHONY: test
test:
	@$(foreach package,$(packages), \
    		set -e; \
    		go test  $(package);)