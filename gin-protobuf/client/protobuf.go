package client

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-study/gin-protobuf/pb"
	"io/ioutil"
	"net/http"
)

func ExampleTest()  {
	resp, err := http.Get("http://localhost:8080/protobuf")
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			user := &pb.Student{}
			proto.UnmarshalMerge(body, user)
			fmt.Println(*user)
		}

	}
}
