BINARY_NAME=myapp

lint:
	golangci-lint run --fix --config=.golangci.yaml

run:
	go run ./cmd/main.go

test:
	go test -v ./...

build:
	CGO_ENABLED=0 go build -o ./out/${BINARY_NAME} ./cmd/main.go

clean:
	go clean
	rm -f ./out

.PHONY: lint