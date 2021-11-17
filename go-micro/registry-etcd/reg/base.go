package reg

import (
	"github.com/asim/go-micro/plugins/registry/etcd/v4"
	"go-micro.dev/v4/registry"
)

var ServiceReg registry.Registry

func init()  {
	ServiceReg = etcd.NewRegistry(
		func(op *registry.Options) {
			op.Addrs = []string{
				"127.0.0.1:23791",
			}
		},
	)
}
