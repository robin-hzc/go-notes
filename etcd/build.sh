#!/bin/bash
wget https://github.com/etcd-io/etcd/releases/download/v3.5.0/etcd-v3.5.0-linux-amd64.tar.gz

tar xf etcd-v3.5.0-linux-amd64.tar.gz

cd etcd-v3.5.0-linux-amd64

echo "FROM alpine

ADD etcd /usr/local/bin/
ADD etcdctl /usr/local/bin/
ADD etcdutl /usr/local/bin/
RUN mkdir -p /var/etcd/
RUN mkdir -p /var/lib/etcd/

# https://github.com/etcd-io/etcd/blob/main/Dockerfile-release.amd64
RUN echo 'hosts: files mdns4_minimal [NOTFOUND=return] dns mdns4' >> /etc/nsswitch.conf

EXPOSE 2379 2380

CMD [\"/usr/local/bin/etcd\"]">>Dockerfile

docker build -t etcd_v3.5.0 .

cd ../

docker-compose up -d