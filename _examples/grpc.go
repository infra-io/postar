package main

import (
	"context"
	"fmt"
	"github.com/avinoplan/postar/api"
	"google.golang.org/grpc"
	"os"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5897", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	req := &api.SendEmailRequest{
		Email: &api.Email{
			Receivers: []string{os.Getenv("POSTAR_RECEIVER")},
			Subject:   "测试邮件",
			BodyType:  "text/html",
			Body:      "<p>邮件内容</p>",
		},
		Options: nil,
	}
	fmt.Printf("client req: %+v\n", req)

	client := api.NewPostarServiceClient(conn)
	rsp, err := client.SendEmail(context.Background(), req)
	if err != nil {
		panic(err)
	}

	fmt.Printf("server rsp: %+v\n", rsp)
}
