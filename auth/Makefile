protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative pb/auth-service.proto

run:
	go run cmd/auth.go

.PHONY: protoc run