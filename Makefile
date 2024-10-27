.PHONY: setup
setup:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: proto
proto:
	protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		proto/user/user.proto

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: run
run:
	go run cmd/server/main.go

.PHONY: all
all: setup tidy proto run
