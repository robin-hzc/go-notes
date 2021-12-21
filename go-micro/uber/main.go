package main

import (
	"github.com/afex/hystrix-go/hystrix"
	"go-micro.dev/v4"
)
func main(){
	// New Service
	service := micro.NewService(
		micro.Name("com.foo.breaker.example"),
		micro.WrapClient(),
	)

	// Initialise service
	service.Init()
	hystrix.DefaultMaxConcurrent = 3//change concurrrent to 3
	hystrix.DefaultTimeout = 200 //change timeout to 200 milliseconds...

}


