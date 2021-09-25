
GOPATH:=$(shell go env GOPATH)
GOOGLEPROTO=proto
MODIFY=Mproto/imports/api.proto=github.com/itzmanish/go-micro/v2/api/proto

.PHONY: proto
proto:

	protoc --proto_path=:. --proto_path=${GOOGLEPROTO} --lint_out=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/slatomate/v1/slatomate.proto

.PHONY: build
build: proto

	go build -o slatomate-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t slatomate-service:latest
