
default:run_http
build_producer:
	go build -x muilteProducer.go

export VERSION=v1.0.0
build_http:
	ls -lh *
	pwd
	go build -o ./HttpServer/server -x ./HttpServer/server.go
run_http: build_http
	./HttpServer/server
