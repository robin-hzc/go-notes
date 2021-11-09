package main

import (
	"github.com/gin-gonic/gin"
	"go-notes/gin-protobuf/pb"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	r := gin.Default()
	r.POST("/protobuf", func(c *gin.Context) {
		body,_ := ioutil.ReadAll(c.Request.Body)
		log.Println(string(body))
		data := &pb.Student{
			Name: "张三",
			Subject:  "数学",
			Scores:60,
		}
		c.ProtoBuf(http.StatusOK, data)
	})
	r.Run(":8080")
}
