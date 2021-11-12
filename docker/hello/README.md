运行命令：

GOOS=linux GOARCH=amd64 go build -o hello main.go

docker build -t test:tag .

docker-compose up -d


监控日志

docker logs -f 容器名字

docker logs -f -t --since="2021-11-12" --tail=10 容器名字