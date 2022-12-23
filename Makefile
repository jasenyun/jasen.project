GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)

.PHONY: init, build, run

init:
	go get github.com/gin-gonic/gin
build:
	mkdir -p bin/windows_amd64/ && GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X main.Version=$(VERSION)" -o ./bin/windows_amd64/ ./...
	mkdir -p bin/linux_amd64/ && GOOS=linux GOARCH=amd64 go build -ldflags "-s -w -X main.Version=$(VERSION)" -o ./bin/linux_amd64/ ./...
	for exec in ./bin/**/*; do upx -9 $$exec > /dev/null || break ; dones_amd64/ && GOOS=windows GOARCH=amd64 go build -ldflags "-s -w -X main.Version=$(VERSION)" -o ./bin/windows_amd64/ ./...

run:
	go run ./cmd/main
run-gather:
	go run ./cmd/gather.erp/main