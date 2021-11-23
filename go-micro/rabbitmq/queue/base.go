package queue

import (
	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	"go-micro.dev/v4/broker"
)
var BrKrSub broker.Broker

func init()  {
	rabbitmq.DefaultRabbitURL = "amqp://guest:guest@localhost:5672/"
	BrKrSub = broker.NewBroker(
	)
}
