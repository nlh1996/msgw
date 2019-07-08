package server

import (
	"go-gateway/controller"
	"go-gateway/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:10000"
)

// Init .
func Init() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer()

	// 注册Hello
	proto.RegisterHelloServer(s, controller.HelloService)

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
