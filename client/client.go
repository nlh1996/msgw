package client

import (
	"context"
	"go-gateway/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:10000"
)

// Init .
func Init() {
	// 连接
	conn, err := grpc.Dial(Address, grpc.WithInsecure())

	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := proto.NewHelloClient(conn)

	// 调用方法
	req := new(proto.HelloRequest)
	req.Name = "gRPC"
	res, err := c.SayHello(context.Background(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Message)
}
