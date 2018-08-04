DEP	= dep

Gopkg.lock: Gopkg.toml
	$(DEP) ensure -update

.PHONY: test
test:
	go test -v ./pkg/...
	go test -v ./cmd/...

.PHONY: protoc
protoc:
	protoc -I. --go_out=plugins=grpc:$GOPATH/src api/judge/v1alpha1/*.proto

.PHONY: gofmt
gofmt:
	gofmt -w -s pkg/ cmd/

