lint:
	@echo "Running linter"
	${GOPATH}/bin/golangci-lint run

test:
	go test ./...

check: | lint test
