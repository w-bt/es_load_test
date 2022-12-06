.PHONY: all
export GO111MODULE=on

all: clean build test

APP=es_load_test
MANUAL_TOOLS_SCRIPT=manual-tools
APP_VERSION:=$(shell cat .version)
APP_COMMIT:=$(shell git rev-parse HEAD)
APP_EXECUTABLE="./out/$(APP)"
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
DB_HOST=$(shell cat application.yml | grep -i -w DB_HOST | cut -d " " -f2)
DB_NAME=$(shell cat application.yml | grep -i -w DB_NAME | cut -d " " -f2)
DB_USER=$(shell cat application.yml | grep -i -w DB_USER | cut -d " " -f2)

assign-vars = $(if $(1),$(1),$(shell grep '$(2):' application.yml | tail -n1| cut -d':' -f2))

build-deps:
	go mod download

compile:
	go fmt ./...
	mkdir -p out/
	go build -o $(APP_EXECUTABLE) -ldflags "-X main.version=$(APP_VERSION) -X main.commit=$(APP_COMMIT)"
	chmod +x out/*

fmt:
	go fmt $(ALL_PACKAGES)

vet:
	go vet $(ALL_PACKAGES)

lint:
	@for p in $(ALL_PACKAGES); do \
		echo "==> Linting $$p"; \
		golint $$p | { grep -vwE "exported (var|function|method|type|const) \S+ should have comment" || true; }& \
	done

build: build-deps compile

build-local: copy-config build

clean:
	rm -rf out/

test-cover:
	go get -u github.com/jokeyrhyme/go-coverage-threshold/cmd/go-coverage-threshold
	ENVIRONMENT=test go-coverage-threshold

test-cover-html:
	mkdir -p out/
	go test -covermode=count  -coverprofile=coverage-all.out  ./...
	@go tool cover -html=coverage-all.out -o out/coverage.html
	@go tool cover -func=coverage-all.out

copy-config:
	cp application.sample.yml application.yml

ci: clean test

build-debug: build-deps
	mkdir -p out/
	go build -o $(APP_EXECUTABLE) -ldflags "-X main.version=$(APP_VERSION) -X main.commit=$(APP_COMMIT)" -gcflags "all=-N -l" cmd/*.go

build-debug-server: copy-config build-debug
	dlv --listen=:2345 --headless=true --api-version=2 exec out/es_load_test server

build-run-server: copy-config build
	out/es_load_test server
