SOURCES := $(shell find . -type f -name '*.go')

GITHASH=$(shell git rev-list -1 HEAD)
DATE=$(shell date +'%FT%TZ%:z')

FLAGS="-X komunalno/cdc/cmd.GitCommitHash=$(GITHASH) -X komunalno/cdc/cmd.BuildDate=$(DATE)"

cdc: $(SOURCES) go.sum
	go build -ldflags $(FLAGS)

.PHONY: test
test: $(SOURCES) go.sum
	go test ./... -v

.PHONY: quickcover
quickcover: $(SOURCES) go.sum
	go test ./... -cover

.PHONY: cover
cover: $(SOURCES) go.sum coverage.out
	go tool cover -html=coverage.out 

coverage.out: $(SOURCES)
	go test ./... -coverprofile=coverage.out

go.sum: go.mod
	go mod tidy

.PHONY: install
install: cdc
	go install -ldflags $(FLAGS)
