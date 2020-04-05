all: test build 

build:
	cd cmd/getdockerimage && go build

build-windows:
	cd cmd/getdockerimage && GOOS=windows GOARCH=amd64 go build

test:
	go test ./...
