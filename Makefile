# Prefer running recipes over files with same names
.PHONY: lint test cover build run start

help:
	@echo
	@echo "gofrog"
	@echo
	@echo "  Commands: "
	@echo
	@echo "    help - Show this message."
	@echo "    lint - Run configured linters."
	@echo "    test - Test application."
	@echo "    cover - Get unit test coverage report."
	@echo "    build - Build application."
	@echo "    run - Run application."
	@echo "    start - Build and run application."

lint:
	golangci-lint run

test:
	# -race requires cgo; hence CGO_ENABLED=1
	CGO_ENABLED=1 go test -coverprofile=coverage.out -v -race -cover ./...

cover:
	CGO_ENABLED=0 go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out

build:
	CGO_ENABLED=0 go build -o gofrog

run:
	./gofrog -config config.toml

start:
	make build 
	make run
