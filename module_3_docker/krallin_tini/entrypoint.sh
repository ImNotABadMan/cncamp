#!/bin/sh

php-fpm --nodaemonize &
nginx -g 'daemon off;'


