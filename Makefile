GO    := GO111MODULE=on go
GO_LINUX := GOOS=linux GOARCH=amd64 GO111MODULE=on go
pkgs   = $(shell $(GO) list ./... | grep -v /vendor/)

PREFIX                  ?= $(shell pwd)
BIN_DIR                 ?= $(shell pwd)
DOCKER_IMAGE_NAME       ?= multicase
DOCKER_IMAGE_TAG        ?= 0.1.0

all: clean test build image

test:
	@echo ">> running tests"
	@$(GO) test -short $(pkgs)

build: 
	@echo ">> building binaries"
	@docker run  -v `pwd`:/src -e GOPROXY=https://goproxy.io -e CGO_ENABLED=0 -w /src golang:1.12 go build -ldflags '-extldflags "-static"' multicast.go

build-local: 
	@echo ">> building binaries locally"
	@$(GO) build multicast.go	

image: 
	@echo ">> building docker image"
	@docker build -t "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)" .

publish:
	@echo ">> publishing docker image"
	@docker push "$(DOCKER_IMAGE_NAME):$(DOCKER_IMAGE_TAG)"
clean:
	@echo ">> cleaning previous build"
	@rm ./multicast || echo "no binary to clean"


.PHONY: all build test image publish clean
