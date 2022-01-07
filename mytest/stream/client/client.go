package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/mytest/stream/proto"
	"log"
	"sync"
	"time"
)

func ServerStream(cc *grpc.ClientConn) error {
	cli := proto.NewServiceClient(cc)
	stream, err := cli.ServerStream(context.TODO(), &proto.StreamRequest{})
	if err != nil {
		return fmt.Errorf("[ServerStream] create stream error: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err != nil {
			return fmt.Errorf("[ServerStream] create recv stream error: %v", err)
		}
		fmt.Printf("[ServerStream] recv from server: %v \n", resp.GetMsg())
	}
}

func ClientStream(cc *grpc.ClientConn) error {
	cli := proto.NewServiceClient(cc)
	stream, err := cli.ClientStream(context.TODO())
	if err != nil {
		return fmt.Errorf("[ClientStream] create stream error: %v", err)
	}

	var i int64
	for {
		if err := stream.Send(&proto.StreamRequest{Msg: fmt.Sprintf("%d", i)}); err != nil {
			return fmt.Errorf("[ClientStream] send error: %v", err)
		}
		i++
		time.Sleep(time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	conn, err := grpc.Dial("127.0.0.1:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}

	go func() {
		defer wg.Done()
		if err := ServerStream(conn); err != nil {
			log.Println(err)
			return
		}
	}()

	go func() {
		defer wg.Done()
		if err := ClientStream(conn); err != nil {
			return
		}
	}()

	// TODO BidirectionalStream

	wg.Wait()
}
