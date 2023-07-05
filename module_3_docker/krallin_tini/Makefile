default: run

build:
	docker build -f Dockerfile . -t use_krallin_tini:v1.0.0
run: build
	docker run -d --name krallin_tini use_krallin_tini:v1.0.0
down:
	docker rm -f krallin_tini