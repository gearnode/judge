DEP	= dep

bin/judgeserver: cmd/judgeserver/*.go pkg/**/*.go
	@mkdir -p bin
	go build -o bin/judgeserver cmd/judgeserver/main.go

go.sum: go.mod
	go get

.PHONY: test
test: go.sum
	go test -v ./pkg/...
	go test -v ./cmd/...

.PHONY: protoc
protoc:
	protoc -I. --go_out=plugins=grpc:$(GOPATH)/src api/judge/v1alpha1/*.proto

.PHONY: gofmt
gofmt:
	gofmt -w -s pkg/ cmd/

.PHONY: clean
clean:
	rm bin/judgeserver
