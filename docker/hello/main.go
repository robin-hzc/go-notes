package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main()  {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/info", func(c *gin.Context) {
		c.JSONP(http.StatusOK,"hello world")
		},
	)
	log.Println("hello world")
	err := router.Run(":8081")
	log.Println("Run err", err)
}