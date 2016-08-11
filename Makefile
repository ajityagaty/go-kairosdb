PROJECT := go-kairosdb
GITCOMMIT := $(shell git rev-parse --short HEAD)
GIT_BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
GITUNTRACKEDCHANGES := $(shell git status --porcelain --untracked-files=no)
ifneq ($(GITUNTRACKEDCHANGES),)
	GITCOMMIT := $(GITCOMMIT)-dirty
endif

all: build

fmt:
	@gofmt -w ./

proto:
	protoc --proto_path=./kairosdb -I=./vendor --go_out=plugins=grpc:./kairosdb ./kairosdb/*.proto

build: proto
	go test -v  ./client ./kairosdb
	go build ./client ./kairosdb

.PHONY: build run migrate all push
