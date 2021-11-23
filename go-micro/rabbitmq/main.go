package main

import (
	"context"
	"fmt"
	"github.com/asim/go-micro/plugins/broker/rabbitmq/v4"
	"go-micro.dev/v4"
	"go-micro.dev/v4/broker"
	"go-micro.dev/v4/server"
	"log"
	"time"
)

var (
	topic = "go.micro.topic.foo"
)



func main() {
	rabbitmq.DefaultRabbitURL = "amqp://guest:guest@localhost:5672"
	brkrSub := broker.NewSubscribeOptions(
		broker.Queue(topic),
		broker.DisableAutoAck(),
		rabbitmq.DurableQueue(),
	)

	b := rabbitmq.NewBroker()
	b.Init()
	if err := b.Connect(); err != nil {
		log.Println("cant conect to broker, skip: %v", err)
	}

	s := server.NewServer(server.Broker(b))

	service := micro.NewService(
		micro.Server(s),
		micro.Broker(b),
	)
	h := &Example{}
	// Register a subscriber
	micro.RegisterSubscriber(
		"topic",
		service.Server(),
		h.Handler,
		server.SubscriberContext(brkrSub.Context),
		server.SubscriberQueue(topic),
	)
	//service.Init()
	go Sub(b)
	go Pub(b)
	if err := service.Run(); err != nil {
		log.Println(err)
	}
}
type Example struct{}

func (e *Example) Handler(ctx context.Context, r interface{}) error {
	return nil
}


/*Sub
 *@Description: 消费者
 *@param b
 */
func Sub( b broker.Broker)  {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		b.Subscribe(topic, func(event broker.Event) error {
			msg := event.Message()
			event.Ack()
			log.Println("Subscribe",string(msg.Body))
			return nil
		},broker.Queue(topic),
			broker.DisableAutoAck(),
			rabbitmq.DurableQueue())
		i++
	}
	return
}
/*Pub
 *@Description: 生产者
 *@param b
 */
func Pub(b broker.Broker)  {
	tick := time.NewTicker(time.Second)
	i := 0
	for _ = range tick.C {
		msg := &broker.Message{
			Header: map[string]string{
				"id": fmt.Sprintf("%d", i),
			},
			Body: []byte(fmt.Sprintf("%d: %s", i, time.Now().String())),
		}
		if err := b.Publish(topic, msg); err != nil {
			log.Printf("[pub] failed: %v", err)
		} else {
			fmt.Println("[pub] pubbed message:", string(msg.Body))
		}
		i++
	}
}