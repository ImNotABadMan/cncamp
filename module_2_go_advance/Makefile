
default:run_http
build_producer:
	go build -x muilteProducer.go

export VERSION=v1.0.0
run_http:
	go build -o ./HttpServer/server -x HttpServer/server.go
	./HttpServer/server
