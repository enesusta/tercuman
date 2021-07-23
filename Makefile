.PHONY: default

default: clean mock test

clean:
	rm -rf ./build
	rm -rf mocks

security:
	gosec -quiet ./...

test:
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build: clean test
	CGO_ENABLED=0 go build -ldflags="-w -s" -o main.go

mock:
	go generate ./...

