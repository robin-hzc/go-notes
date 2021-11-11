package main

import (
	"context"
	"github.com/asim/go-micro/plugins/registry/consul/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/util/log"
	"go-notes/go-micro/hello/pb"
	"google.golang.org/grpc"
	"time"
)

type Say struct{}

func (s *Say) Hello(ctx context.Context, req *pb.Request, rsp *pb.Response) error {
	log.Log("Received Say.Hello request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

func main() {
	go func() {
		for {
			grpc.DialContext(context.TODO(), "127.0.0.1:9091")
			time.Sleep(time.Second)
		}
	}()
	reg := consul.NewRegistry(
		func(op *registry.Options) {
			op.Addrs = []string{
				"127.0.0.1:8500",
			}
		},
	)
	service := micro.NewService(
		micro.Name("go.micro.srv.hello"),
		micro.Version("1"),
		micro.Registry(reg),
	)
	//micro.RegisterSubscriber
	// optionally setup command line usage
	service.Init()

	// Register Handlers
	err := pb.RegisterSayHandler(service.Server(), new(Say))
	log.Info(err)
	// Run server
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
