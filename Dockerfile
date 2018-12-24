FROM golang:1.11-alpine AS builder

WORKDIR /go/src/github.com/gearnode/judge

RUN apk add --no-cache make

COPY . .

RUN make build

FROM alpine:3.8

LABEL org.judge.version="v1alpha1" \
      org.judge.release-date="2018-11-25" \
      maintainer="Judge Authors"

RUN apk add --no-cache ca-certificates

COPY --from=builder /go/src/github.com/gearnode/judge/bin/judgectl /bin/judgectl
COPY --from=builder /go/src/github.com/gearnode/judge/bin/judgeserver /bin/judgeserver

RUN addgroup -g 1000 -S judge && \
    adduser -u 1000 -S judge -G judge

USER judge

WORKDIR /home/judge
