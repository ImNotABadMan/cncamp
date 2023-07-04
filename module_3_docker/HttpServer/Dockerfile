FROM golang:1.20.5-alpine3.18 as go-base
LABEL authors="Bin"

RUN apk add make

ENV GOPATH=/go
ENV GO111MODULE=on
ENV GOPROXY="https://mirrors.aliyun.com/goproxy/,direct"

COPY ./server.go ./src
COPY ./Makefile ./src
COPY ./go.mod.tpl ./src/go.mod

WORKDIR ./src

RUN go mod tidy && make build_http

# 约定，方便维护查看
EXPOSE 8422

# 由go基础镜像构造出来的镜像 338 MB
# 多段构建，只需要上一个镜像生成的二进制文件，底层为轻量的debain基础镜像上，最终镜像81.4MB
FROM debian:stable-slim as debain
COPY --from=go-base /go/src/server /usr/local/bin/server

ENTRYPOINT ["/usr/local/bin/server"]
