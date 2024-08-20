run:
	go run cmd/main.go
lint:
	golangci-lint run --timeout=20m0s
fmt:
	gofumpt -l -w .

integration_test:
	go test -coverprofile=coverage.out -coverpkg ./internal/adapters/db/mysql_adapter/... ./integration/db

coverage:
	go tool cover -html="coverage.out"