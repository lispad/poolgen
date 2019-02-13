GOPATH := $(shell go env GOPATH)
GOBIN := $(GOPATH)/bin

default: build install

build:
	go build -i

install: build
	cp ./poolgen $(GOBIN)

.PHONY: build install
