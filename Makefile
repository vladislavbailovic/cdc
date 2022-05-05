SOURCES := $(shell find . -type f -name '*.go')

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

grpc/loader.pb.go: grpc/loader.proto
	protoc --go_out=. grpc/loader.proto

grpc/loader_grpc.pb.go: grpc/loader.proto
	protoc --go-grpc_out=. grpc/loader.proto

.PHONY: serve
serve: build/kog-server
	./build/kog-server

build/kog-server: $(SOURCES) grpc/loader.pb.go grpc/loader_grpc.pb.go
	go build -o build/kog-server cmd/grpc/server/main.go

.PHONY: ping
ping: build/kog-client build/kog-server
	./build/kog-server &
	./build/kog-client
	killall kog-server

build/kog-client: $(SOURCES) grpc/loader.pb.go grpc/loader_grpc.pb.go
	go build -o build/kog-client cmd/grpc/client/main.go
