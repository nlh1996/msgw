package controller

import (
	"context"
	pb "go-gateway/proto"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService .
var HelloService = new(helloService)

func (h *helloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "Hello " + req.Name + "."
	for i := 0; i < 4; i++ {
		if ctx.Err() == context.Canceled {
			return nil, status.Errorf(codes.Canceled, "SearchService.Search canceled")
		}
		time.Sleep(1 * time.Second)
	}
	return resp, nil
}
