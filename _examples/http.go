package main

import (
	"bytes"
	"fmt"
	"github.com/avinoplan/postar/api"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"net/http"
)

func main() {
	url := "http://127.0.0.1:5897/sendEmail"

	emailReq := &api.SendEmailRequest{
		Email:   &api.Email{
			Receivers: nil,
			Subject:   "测试邮件",
			BodyType:  "text/html",
			Body:      "<p>邮件内容</p>",
		},
		Options: nil,
	}
	fmt.Printf("client req: %+v\n", emailReq)

	marshaled, err := proto.Marshal(emailReq)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(url, "application/octet-stream", bytes.NewReader(marshaled))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	emailRsp := new(api.SendEmailResponse)
	err = proto.Unmarshal(body, emailRsp)
	if err != nil {
		panic(err)
	}

	fmt.Printf("server rsp: %+v\n", emailRsp)
}
