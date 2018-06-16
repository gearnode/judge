.PHONY: proto
proto:
	protoc -I svc/ \
  -I $GOPATH/src \
  --go_out=plugins=grpc:svc \
  svc/api.proto
