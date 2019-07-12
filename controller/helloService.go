package controller

import (
	"context"
	pb "go-gateway/proto"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService .
var HelloService = new(helloService)

func (h *helloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := new(pb.HelloReply)
	resp.Message = "Hello" + req.Name + "."
	return resp, nil
}
