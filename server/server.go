package server

import (
	"go-gateway/controller"
	"go-gateway/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	// gRPC服务地址
	address = "127.0.0.1:11000"
)

// Init .
func Init() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("./keys/server.pem", "./keys/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	// 实例化grpc Server
	s := grpc.NewServer(grpc.Creds(creds))

	// 注册服务
	proto.RegisterHelloServer(s, controller.HelloService)
	proto.RegisterStreamServiceServer(s, controller.StreamService)

	log.Println("Listen on " + address + " with TLS.")

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
