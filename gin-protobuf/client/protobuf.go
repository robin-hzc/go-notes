package client

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-notes/gin-protobuf/pb"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func ExampleTest()  {
	client := &http.Client{}
	data := map[string]string{
		"xx":"1",
		"xx1":"2",
		"xx2":"3",
	}
	b,err:=json.Marshal(data)
	if err!=nil {
		log.Println(err)
		return
	}
	req, err := http.NewRequest("POST", "http://localhost:8080/protobuf", strings.NewReader(string(b)))
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if resp != nil {
		defer resp.Body.Close()
	} else {
		return
	}
	if err != nil {
		log.Println(err)
		return
	} else {
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		} else {
			user := &pb.Student{
				Name: "Test",
			}
			proto.UnmarshalMerge(body, user)
			log.Println(*user)
		}
	}
}
