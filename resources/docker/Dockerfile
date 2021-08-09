FROM golang:1.16.3-alpine AS builder

RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

RUN git clone https://github.com/PabloGilvan/transaction.git /data/app

WORKDIR /data/app

RUN go env -w GOPROXY=https://proxy.golang.org && \
    go env -w CGO_ENABLED="0" && \
    go env -w GO111MODULE='on' && \
    sh build.sh

FROM alpine:latest
RUN  mkdir -p /go/bin && \
     mkdir /logs && \
     chmod -R 755 /logs
WORKDIR /go/bin
COPY --from=builder /data/app/bin/api .
EXPOSE 8080
ENTRYPOINT /go/bin/api