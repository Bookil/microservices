run:
	go run cmd/main.go
lint:
	golangci-lint run --timeout=20m0s
fmt:
	gofumpt -l -w .

unit_test:
	go test -coverprofile=coverage.out $(shell go list ./... | grep -v /integration)
	
coverage:
	go tool cover -html="coverage.out"