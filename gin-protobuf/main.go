package main

import (
	"github.com/gin-gonic/gin"
	"go-notes/gin-protobuf/pb"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/protobuf", func(c *gin.Context) {
		data := &pb.Student{
			Name: "张三",
			Subject:  "数学",
			Scores:60,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	r.Run(":8080")
}
