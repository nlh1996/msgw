package controller

import (
	"context"
	"fmt"
	"go-gateway/proto"
	"log"

	"github.com/micro/go-micro/metadata"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

// 定义helloService并实现约定的接口
type helloService struct{}

// HelloService .
var HelloService = new(helloService)

func (h *helloService) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	// 解析metadata中的信息并验证
	md, ok := metadata.FromContext(ctx)
	if !ok {
		return nil, grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
	}
	var (
		appid  string
		appkey string
	)

	if val, ok := md["appid"]; ok {
		log.Println(val[0])
	}

	if val, ok := md["appkey"]; ok {
		log.Println(val[0])
	}

	if appid != "101010" || appkey != "i am key" {
		return nil, grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
	}

	resp := new(proto.HelloReply)
	resp.Message = fmt.Sprintf("Hello %s.\nToken info: appid=%s,appkey=%s", req.Name, appid, appkey)

	return resp, nil
}
