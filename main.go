package main

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
	address = "127.0.0.1:12000"
)

// Init .
func main() {
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("./keys/server.pem", "./keys/server.key")
	if err != nil {
		log.Fatalf("Failed to generate credentials %v", err)
	}

	var opts []grpc.ServerOption
	opts = append(opts, grpc.Creds(creds))

	// 实例化grpc Server
	s := grpc.NewServer(opts...)

	// 注册服务
	proto.RegisterStreamServiceServer(s, controller.StreamService)

	log.Println("Listen on " + address + " with TLS.")

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
