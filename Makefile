.PHONY: protoc

protoc:
	protoc --go_out=. --go_opt=paths=source_relative humors.proto
	protoc --go_out=. --go_opt=paths=source_relative example/api/api.proto

lint:
	golangci-lint run