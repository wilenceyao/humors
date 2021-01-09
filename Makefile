.PHONY: lint

lint:
	golangci-lint run

protoc:
	protoc --go_out=. --go_opt=paths=source_relative humors.proto
	protoc --go_out=. --go_opt=paths=source_relative example/api/api.proto
