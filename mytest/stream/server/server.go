package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/mytest/stream/proto"
	"log"
	"net"
	"time"
)

var _ proto.ServiceServer = &StreamServer{}

type StreamServer struct {
	proto.UnimplementedServiceServer
}

func (s *StreamServer) ClientStream(server proto.Service_ClientStreamServer) error {
	for {
		msg, err := server.Recv()
		if err != nil {
			return err
		}
		fmt.Printf("recv message from client: %v \n", msg.GetMsg())
	}
}

func (s *StreamServer) ServerStream(
	request *proto.StreamRequest,
	server proto.Service_ServerStreamServer,
) error {
	var i int64
	for {
		if err := server.Send(
			&proto.StreamResponse{
				Msg: fmt.Sprintf("msg: %d", i),
			}); err != nil {
			return err
		}
		i++
		time.Sleep(time.Second)
	}
}

func (s *StreamServer) BidirectionalStream(server proto.Service_BidirectionalStreamServer) error {
	// TODO
	return nil
}

func main() {
	server := grpc.NewServer()
	proto.RegisterServiceServer(server, &StreamServer{})

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	if err := server.Serve(l); err != nil {
		log.Fatalln(err)
	}
}
