package client

import (
	"context"
	"go-gateway/proto"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

// StreamClientInit .
func StreamClientInit() {
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

	client := proto.NewStreamServiceClient(conn)

	req := &proto.StreamRequest{Pt: &proto.StreamPoint{Name: "gRPC Stream Client: List", Value: 0}}
	err = printLists(client, req)
	errNotNil("printLists", err)

	req = &proto.StreamRequest{Pt: &proto.StreamPoint{Name: "gRPC Stream Client: Record", Value: 0}}
	err = printRecord(client, req)
	errNotNil("printRecord", err)

	req = &proto.StreamRequest{Pt: &proto.StreamPoint{Name: "gRPC Stream Client: Route", Value: 0}}
	err = printRoute(client, req)
	errNotNil("printRoute", err)
}

func errNotNil (funcName string, err error) {
	if err != nil {
		log.Fatalf("%s.err: %v",funcName, err)
	}
}

func printLists(client proto.StreamServiceClient, req *proto.StreamRequest) error {
	stream, err := client.List(context.Background(), req)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}
	log.Println("stream end!!!")
	return nil
}

func printRecord(client proto.StreamServiceClient, req *proto.StreamRequest) error {
	return nil
}

func printRoute(client proto.StreamServiceClient, req *proto.StreamRequest) error {
	return nil
}
