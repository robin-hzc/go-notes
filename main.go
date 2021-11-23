package main

import (
	"go-notes/channel"
	"time"
)

func main()  {
	factory := channel.NewFactory()
	go factory.Consumer()
	go factory.Producer()
	time.Sleep(21 *time.Second)
	factory.Stop()
}
