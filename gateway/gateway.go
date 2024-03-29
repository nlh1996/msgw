package main

import (
	"go-gateway/controller"
	"go-gateway/middleware"
	"go-gateway/utils"
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
	// 添加中间件，拦截验证token
	opts = append(opts, grpc.UnaryInterceptor(middleware.AuthToken()))
	opts = append(opts, grpc.StreamInterceptor(middleware.StreamAuth()))

	//自定义编解码
	opts = append(opts, grpc.CustomCodec(utils.Codec()))

	opts = append(opts, grpc.UnknownServiceHandler(controller.Handler))

	// 实例化grpc Server
	s := grpc.NewServer(opts...)

	log.Println("Listen on " + address + " with TLS.")

	if err := s.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
