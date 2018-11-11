GO = go
GOFMT = gofmt

.TARGET: all

.PHONY: all
all: go.sum api/judge/v1alpha1/*.pb.go test bin/judgeserver bin/judgectl

bin/judgeserver: cmd/judgeserver/*.go pkg/**/*.go
	@mkdir -p bin
	$(GO) build -o bin/judgeserver cmd/judgeserver/main.go

bin/judgectl: cmd/judgectl/*go
	@mkdir -p bin
	$(GO) build -o bin/judgectl cmd/judgectl/main.go

go.sum: go.mod
	$(GO) get

api/judge/v1alpha1/%.pb.go: api/judge/v1alpha1/*.proto
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src api/judge/v1alpha1/*.proto

coverage.out: pkg/**/*.go
	$(GO) test -coverprofile=coverage.out ./pkg/...

.PHONY: test
test: go.sum coverage.out

.PHONY: gofmt
gofmt:
	$(GOFMT) -w -s pkg/ cmd/

.PHONY: clean
clean:
	rm bin/judgeserver
	rm bin/judgectl
	rm coverage.out
	rm -rf vendor
	rm api/judge/v1alpha1/*.pb.go
