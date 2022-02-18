#!/bin/bash
# --with-http_stub_status_module -> Nginx性能统计模块
# --with-http_ssl_module -> ssl加密模块
# --add-module=/ngx_http_image_filter_module -> 图片裁剪模块
./configure \
--prefix=/usr/local/nginx \
--user=nginx \
--group=nginx \
--sbin-path=/usr/local/nginx/sbin/nginx \
--modules-path=/usr/local/nginx/modules \
--conf-path=/usr/local/nginx/conf/nginx.conf \
--error-log-path=/usr/local/nginx/logs/error.log \
--http-log-path=/usr/local/nginx/logs/access.log \
--pid-path=/usr/local/nginx/logs/nginx.pid \
--lock-path=/usr/local/nginx/logs/nginx.lock \
--with-http_stub_status_module \
--with-http_ssl_module \
&&
make && make install
cd /usr/local/nginx&&mkdir ssl&&cd ssl
chmod -R +rwx ssl
openssl genrsa -out ca.key 2048
openssl req -new -x509 -days 365 -key ca.key -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=localhost" -out ca.crt
openssl req -newkey rsa:2048 -nodes -keyout server.key -subj "/C=CN/ST=GD/L=SZ/O=Acme, Inc./CN=localhost" -out server.csr
# shellcheck disable=SC2169
openssl x509 -req -extfile <(printf "subjectAltName=DNS:localhost") -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt
