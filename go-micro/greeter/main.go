package main

import (
	"context"
	"go-micro.dev/v4"
	pb2 "go-notes/go-micro/greeter/pb"
	"log"
)

type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *pb2.Request, rsp *pb2.Response) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("helloworld"),
	)
	service.Init()
	pb2.RegisterGreeterHandler(service.Server(), new(Greeter))
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}