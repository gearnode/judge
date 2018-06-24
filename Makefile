.PHONY: all
all:
	bazel build //...

.PHONY: test
test:
	bazel test //...

.PHONY: gofmt
gofmt:
	gofmt -w -s pkg/ cmd/

.PHONY: gazelle
gazelle:
	bazel run //:gazelle
