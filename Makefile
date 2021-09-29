# SPDX-FileCopyrightText: 2021 Eric Neidhardt
# SPDX-License-Identifier: CC0-1.0

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
