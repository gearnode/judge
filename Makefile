GO = go
GOFMT = gofmt
PROTOC = protoc
DOCKER_NAME = gearnode/judge
VERSION = v1alpha1

.TARGET = all

.PHONY: all
all: go.sum pkg/apiserver/$(VERSION)/*.pb.go test bin/judgeserver bin/judgectl priv/server.crt docker-build

.PHONY: build
build: vendor bin/judgeserver bin/judgectl

bin:
	mkdir -p bin

bin/judgeserver: bin cmd/judgeserver/*.go pkg/**/*.go
	$(GO) build -o bin/judgeserver cmd/judgeserver/main.go

bin/judgectl: bin cmd/judgectl/*go pkg/**/*.go
	$(GO) build -o bin/judgectl cmd/judgectl/main.go

priv:
	mkdir -p priv

priv/server.crt: priv
	mkcert 127.0.0.1
	mv ./127.0.0.1.pem ./priv/server.crt
	mv ./127.0.0.1-key.pem ./priv/server.key

go.sum: go.mod
	$(GO) get
	$(GO) mod vendor

pkg/apiserver/$(VERSION)/%.pb.go: api/judge/$(VERSION)/*.proto
	$(PROTOC) -I. --go_out=plugins=grpc:$(GOPATH)/src api/judge/$(VERSION)/*.proto

coverage.out: pkg/**/*.go
	$(GO) test -coverprofile=coverage.out ./pkg/...

.PHONY: test
test: go.sum coverage.out

.PHONY: gofmt
gofmt:
	$(GOFMT) -w -s pkg/ cmd/

.PHONY: vendor
vendor:
	$(GO) mod vendor

.PHONY: clean
clean:
	rm -rf coverage.out
	rm -rf vendor
	rm -rf priv
	rm -rf bin

.PHONY: docker-build
docker-build:
	docker build . -t $(DOCKER_NAME):$(VERSION)
