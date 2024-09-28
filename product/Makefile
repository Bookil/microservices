lint:
	golangci-lint run --timeout=20m0s

fmt:
	gofumpt -l -w .

#test
unit_test:
	go test -coverprofile=coverage.out $(shell go list ./... | grep -v /integration)

db_integration_test:
	go test -coverprofile=coverage.out -coverpkg ./internal/adapters/db/mysql_adapter/... ./integration/db

coverage:
	go tool cover -html="coverage.out"