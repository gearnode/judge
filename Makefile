GO = go
GOFMT = gofmt
PROTOC = protoc

.TARGET: all

.PHONY: all
all: go.sum api/judge/v1alpha1/*.pb.go test bin/judgeserver bin/judgectl priv/server.crt

bin:
	mkdir -p bin

bin/judgeserver: bin cmd/judgeserver/*.go pkg/**/*.go
	$(GO) build -o bin/judgeserver cmd/judgeserver/main.go

bin/judgectl: bin cmd/judgectl/*go pkg/**/*.go
	$(GO) build -o bin/judgectl cmd/judgectl/main.go

bin/mkcert:
	$(GO) get -u github.com/FiloSottile/mkcert
	ln -fs $(GOPATH)/bin/mkcert bin/mkcert

priv:
	mkdir -p priv

priv/server.crt: bin/mkcert priv
	mkcert 127.0.0.1
	mv ./127.0.0.1.pem ./priv/server.crt
	mv ./127.0.0.1-key.pem ./priv/server.key

go.sum: go.mod
	$(GO) get
	$(GO) mod vendor

api/judge/v1alpha1/%.pb.go: api/judge/v1alpha1/*.proto
	$(PROTOC) -I. --go_out=plugins=grpc:$(GOPATH)/src api/judge/v1alpha1/*.proto

coverage.out: pkg/**/*.go
	$(GO) test -coverprofile=coverage.out ./pkg/...

.PHONY: test
test: go.sum coverage.out

.PHONY: gofmt
gofmt:
	$(GOFMT) -w -s pkg/ cmd/

.PHONY: clean
clean:
	rm coverage.out
	rm -rf vendor
	rm -rf priv
	rm -rf bin
	rm api/judge/v1alpha1/*.pb.go
