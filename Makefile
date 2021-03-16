.PHONY: deps vendor test cover clean build

init:
	go mod init

deps:
	go get -u ./...

vendor:
	go mod vendor

test:
	go test -mod vendor -cover -coverprofile=coverage.out ./...

cover: test
	go tool cover -html=coverage.out

clean: vendor
	go mod tidy

build: clean
	go vet -mod vendor ./...
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod vendor ./...
