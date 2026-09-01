.PHONY: dev build run test test-verbose test-cover

dev:
	air -c .air.toml

build:
	go build -o ./tmp/main .

run:
	go run .

test:
	go test ./...

# Run the test suite with verbose output.
test-verbose:
	go test -v ./...
