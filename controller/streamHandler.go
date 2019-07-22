package controller

import (
	"context"
	"go-gateway/proto"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// Handler .
func Handler(srv interface{}, stream grpc.ServerStream) error {
	const address = "127.0.0.1:12000"
	var opts []grpc.DialOption
	creds, err := credentials.NewClientTLSFromFile("./keys/server.pem", "yinghuo2018")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))

	// 转发
	conn, err := grpc.Dial(address, opts...)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := proto.NewStreamServiceClient(conn)
	type streamService struct{}
	req := &proto.StreamRequest{Pt: &proto.StreamPoint{Name: "gRPC Stream Client: List", Value: 0}}

	stream2, err := client.List(context.Background(), req)
	if err != nil {
		log.Println(err)
	}
	for {
		resp, err := stream2.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if err := stream.SendMsg(resp); err != nil {
			log.Println(err)
		}
	}

	return nil
}
