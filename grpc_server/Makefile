build:
	@go build -o bin/blocker

run:
	@./bin/blocker

test:
	@go test -v ./...

proto:
	protoc -I=proto  --go_out=. --go-grpc_out=. proto/main.proto proto/greeter.proto


.PHONY: proto
