package middleware

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
)

// AuthToken .
func AuthToken() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			err := grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
			return nil, err
		}
		var (
			appid  string
			appkey string
		)

		if val, ok := md["appid"]; ok {
			appid = val[0]
		}

		if val, ok := md["appkey"]; ok {
			appkey = val[0]
		}

		if appid != "101010" || appkey != "i am key" {
			err := grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
			return nil, err
		}
		return handler(ctx, req)
	}
}

// StreamAuth .
func StreamAuth() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		serverName, ok := grpc.MethodFromServerStream(ss)
		if !ok {
			err := grpc.Errorf(codes.Unauthenticated, "转发失败！")
			return err
		}
		log.Println(serverName)
		return handler(srv, ss)
	}
}
