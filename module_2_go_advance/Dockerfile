# moudle 3 docker practice
FROM golang:1.20.5-alpine3.18 as muiltego
LABEL authors="binz" module="2"

RUN apk add make

ENV GOPATH=/go
ENV GO111MODULE=on

RUN mkdir -p ./src

# 拷贝目录会递归到目录下的文件，需要指定目录
COPY ./HttpServer ./src/HttpServer
COPY Makefile ./src
COPY ./go.mod.template ./src/go.mod

WORKDIR ./src

RUN make build_http

FROM debian:stable-slim as debain
RUN mkdir -p ./go
COPY --from=muiltego /go/src/HttpServer/server ./go/server
RUN chown -R 0777 ./go/server
WORKDIR ./go

ENTRYPOINT ["./server"]