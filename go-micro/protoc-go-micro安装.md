protoc-go-micro安装

由于官方安装go install go-micro.dev/v4/cmd/protoc-gen-micro@latest不可用

可以在配置GOBIN后源码编译方式

git clone https://github.com/asim/go-micro.git

cd go-micro

git switch -c v4.3.0

cd cmd/protoc-gen-micro

go build -o protoc-gen-micro main.go