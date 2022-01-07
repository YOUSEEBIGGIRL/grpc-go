package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/mytest/interceptor/proto"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	cli := proto.NewHelloServiceClient(conn)
	resp, err := cli.Hello(context.TODO(), &proto.HelloRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(resp.GetMessage())
}
