build:
	@go build -o bin/blocker

run:
	@./bin/blocker

test:
	@go test -v ./...

proto:
	protoc -go_out=. --go-grpc_out=. proto/main.proto


.PHONY: proto
