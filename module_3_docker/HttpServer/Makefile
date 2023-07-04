
default:run_http

export VERSION=v1.0.0
build_http:
	go build -o ./server -x server.go

run_http: build_http
	./HttpServer/server
