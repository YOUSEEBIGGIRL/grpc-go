package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/mytest/interceptor/proto"
	"log"
	"net"
)

type HelloService struct {
	proto.UnimplementedHelloServiceServer
}

func (h *HelloService) Hello(
	ctx context.Context,
	req *proto.HelloRequest,
) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Message: "hello"}, nil
}

func interceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	log.Printf("before handling. Info: %+v", info)
	resp, err = handler(ctx, req)
	log.Printf("after handling. resp: %+v", resp)
	return resp, err
}

func main() {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(interceptor), // 单向调用拦截
	)
	proto.RegisterHelloServiceServer(server, &HelloService{})

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	if err := server.Serve(l); err != nil {
		log.Fatalln(err)
	}
}
