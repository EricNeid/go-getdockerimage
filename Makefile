# SPDX-FileCopyrightText: 2021 Eric Neidhardt
# SPDX-License-Identifier: CC0-1.0

.PHONY: all
all: test build

.PHONY: build
build:
	go build ./cmd/getdockerimage/

.PHONY: build-linux-amd64
build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build ./cmd/getdockerimage/

.PHONY: build-linux-arm
build-linux-arm:
	GOOS=linux GOARCH=arm go build ./cmd/getdockerimage/

.PHONY: build-windows-amd64
build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build ./cmd/getdockerimage/

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	go install golang.org/x/lint/golint@latest
	staticcheck ./...
