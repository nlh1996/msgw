package main

import (
	"context"
	"go-gateway/proto"
	"go-gateway/utils"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

const (
	//gRPC服务地址
	address = "127.0.0.1:11000"

	//是否开启TLS认证
	openTLS = true
)

// customCredential 自定义认证
type customCredential struct{}

func (c *customCredential) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"appid":  "101010",
		"appkey": "i am key",
	}, nil
}

func (c *customCredential) RequireTransportSecurity() bool {
	if openTLS {
		return true
	}
	return false
}

// Init .
func main() {
	var opts []grpc.DialOption
	if openTLS {
		// TLS连接
		creds, err := credentials.NewClientTLSFromFile("./keys/server.pem", "yinghuo2018")
		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}

	// 使用自定义认证
	opts = append(opts, grpc.WithPerRPCCredentials(new(customCredential)))

	// 连接
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	client := proto.NewHelloClient(conn)

	req := &proto.HelloRequest{Name: "gRPC"}
	// 调用方法
	res, err := client.SayHello(utils.GetCtx(), req)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(res.Message)
}
