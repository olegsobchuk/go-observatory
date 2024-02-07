lint:
	@echo "Running linter"
	${GOPATH}/bin/golangci-lint run
test:
	go test ./...
all: | lint test
