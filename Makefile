APP=$(shell basename $(shell git remote get-url origin))
REGISTRY=pontarr
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
#TARGETARCH=amd64
PROJECT_ID := pontarr
IMAGE_NAME := mytelebot
#TARGETOS := linux
IMAGE := ghcr.io/$(PROJECT_ID)/$(IMAGE_NAME):${VERSION}-${TARGETOS}-${TARGETARCH}

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

get: 
	go get

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o mytelebot -ldflags "-X 'github.com/pontarr/mytelebot/cmd.appVersion=${VERSION}'"
# shell dpkg --print-architecture
image:
	docker build . -t ${IMAGE}

push:
	docker push ${IMAGE}

clean:
	rm -rf mytelebot
	docker rmi ${IMAGE}

test:
	go test -v
