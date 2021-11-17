package main

import (
	"context"
	"go-micro.dev/v4"
	"go-micro.dev/v4/util/log"
	"go-notes/go-micro/registry-etcd/pb"
	"go-notes/go-micro/registry-etcd/reg"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Log("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.hello"),
		micro.Registry(reg.ServiceReg),
	)
	// optionally setup command line usage
	service.Init()

	// Register Handlers
	err :=pb.RegisterSayHandler(service.Server(), new(Say))
	log.Info(err)
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}