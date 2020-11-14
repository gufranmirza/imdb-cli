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

packages = \
	./scrapper\

.PHONY: test
test:
	@$(foreach package,$(packages), \
    		set -e; \
    		go test  $(package);)