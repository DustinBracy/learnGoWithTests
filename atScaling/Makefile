build:
# 	golangci-lint run
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./adapters/grpcserver/greet.proto
	go test ./...

unit-tests:
	go test -short ./...

gen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./adapters/grpcserver/greet.proto