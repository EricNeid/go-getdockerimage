all: test build

build:
	go build ./cmd/getdockerimage/

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build ./cmd/getdockerimage/

build-linux-arm:
	GOOS=linux GOARCH=arm go build ./cmd/getdockerimage/

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build ./cmd/getdockerimage/

test:
	go test ./...
