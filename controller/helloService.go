package controller

import (
	"context"
	"go-gateway/proto"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService .
var HelloService = new(helloService)

func (h *helloService) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	resp := new(proto.HelloReply)
	resp.Message = "Hello " + req.Name + "."

	return resp, nil
}

