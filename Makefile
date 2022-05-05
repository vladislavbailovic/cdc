SOURCES := $(shell find . -type f -name '*.go')

GITHASH=$(shell git rev-list -1 HEAD)
DATE=$(shell date +'%FT%TZ%:z')

FLAGS="-X komunalno/cdc/pkg/opts.GitCommitHash=$(GITHASH) -X komunalno/cdc/pkg/opts.BuildDate=$(DATE)"

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

cdc: $(SOURCES) go.sum
	go build -ldflags $(FLAGS)

.PHONY: install
install: cdc
	go install -ldflags $(FLAGS)
