# moudle 3 docker practice
FROM golang:1.20.5-alpine3.18 as muiltego
LABEL authors="binz" module="2"

RUN apk add make

ENV GOPATH=/root/go
ENV GO111MODULE=on

RUN mkdir -p ./go/src

COPY HttpServer ./go/src
COPY Makefile ./go/src
COPY ./go.mod.template ./go/src/go.mod

WORKDIR ./go/src

RUN make run_http

FROM debian:stable-slim as debain
COPY --from=muiltego /root/go/src/HttpServer/server ./go/server
RUN chown -R 0777 ./go/server
WORKDIR ./go

ENTRYPOINT ["./go"]