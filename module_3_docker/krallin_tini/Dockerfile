#FROM debian:stable-slim
FROM php:7.4.33-fpm-alpine
#FROM ubuntu:20.04

RUN apk add --no-cache tini
RUN apk add --no-cache nginx

#Add init
ENV TINI_VERSION v0.19.0
LABEL authors="87650"

ADD https://github.com/krallin/tini/releases/download/${TINI_VERSION}/tini /tini
RUN chmod +x /tini

ENTRYPOINT ["/sbin/tini", "--"]

COPY ./entrypoint.sh /entrypoint.sh
RUN chmod 0777 /entrypoint.sh

CMD ["/entrypoint.sh"]

