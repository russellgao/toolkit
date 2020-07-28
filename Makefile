APP=tkctl
VERSION=`cat VERSION`
GITBRANCH=`git symbolic-ref --short -q HEAD`
GITREVISION=`git log -n1 --format=%H`
BUILDUSER=gaoweizong@hd123.com
BUILDDATE=`date "+%Y-%m-%d %H:%M:%S"`

PACKAGES=`go list ./... | grep -v /vendor/`
VETPACKAGES=`go list ./... | grep -v /vendor/ | grep -v /examples/`
GOFILES=`find . -name "*.go" -type f -not -path "./vendor/*"`

## list: list packages and go files
list:
	@echo $(PACKAGES)
	@echo $(VETPACKAGES)
	@echo $(GOFILES)

## vet: static check
vet:
	@go vet $(VETPACKAGES)

## fmt: format code
fmt:
	@gofmt -s -w $(GOFILES)

## fmt-check: format result check
fmt-check:
	@diff=$$(gofmt -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
	  echo "Please run 'make fmt' and commit the result:"; \
	  echo "$${diff}"; \
	  exit 1; \
	fi;

## clean: cleans the binary
clean:
	@echo "Cleaning"
	@go clean
	@if [ -d target ] ; then rm -rf target ; fi

## test: runs go test with default values
test:
	@go test ./...

## build: build the application to registry
build: clean
	@go build -ldflags="-X 'github.com/prometheus/common/version.Version=$(VERSION)' -X 'github.com/prometheus/common/version.BuildUser=$(BUILDUSER)' -X 'github.com/prometheus/common/version.BuildDate=$(BUILDDATE)' -X 'github.com/prometheus/common/version.Branch=$(GITBRANCH)' -X 'github.com/prometheus/common/version.Revision=$(GITREVISION)'" -o target/$(APP)

## run: runs go run main.go
run:
	@go run main.go

## compile: build the application of different operating systems
compile:
	@# 32-Bit Systems
	@# FreeBDS
	@GOOS=freebsd GOARCH=386 go build -o target/$(APP)-freebsd-386 main.go
	@# MacOS
	@GOOS=darwin GOARCH=386 go build -o target/$(APP)-darwin-386 main.go
	@# Linux
	@GOOS=linux GOARCH=386 go build -o target/$(APP)-linux-386 main.go
	@# Windows
	@GOOS=windows GOARCH=386 go build -o target/$(APP)-windows-386 main.go
	@# 64-Bit
	@# FreeBDS
	@GOOS=freebsd GOARCH=amd64 go build -o target/$(APP)-freebsd-amd64 main.go
	@# MacOS
	@GOOS=darwin GOARCH=amd64 go build -o target/$(APP)-darwin-amd64 main.go
	@# Linux
	@GOOS=linux GOARCH=amd64 go build -o target/$(APP)-linux-amd64 main.go
	@# Windows
	@GOOS=windows GOARCH=amd64 go build -o target/$(APP)-windows-amd64 main.go

## docker-build: build the docker image
docker-build:
	@docker build -t russellgao/toolkit:$(VERSION) --build-arg VERSION=$(VERSION) --build-arg BUILDUSER=$(BUILDUSER) .

## docker-push: push the docker image to registry
docker-push:
	@docker push russellgao/toolkit:$(VERSION)

help:
	@echo "Usage: \n"
	@sed -n 's/^##//p' Makefile | column -t -s ':' |  sed -e 's/^/ /'

## all: execut test、build、docker-build、docker-push targets
all: test build docker-build docker-push

.PHONY: list vet fmt fmt-check clean bindata test build run docker-build docker-push help all