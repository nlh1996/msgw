package utils

import (
	"fmt"

	"github.com/micro/protobuf/proto"
	"google.golang.org/grpc"
)

// 自定义codec类型，
// 实现了grpc.Codec接口中的Marshal和Unmarshal
// 成员变量parentCodec用于当自定义Marshal和Unmarshal失败时的回退codec
type rawCodec struct {
	parentCodec grpc.Codec
}

type frame struct {
	payload []byte
}

// protoCodec实现protobuf的默认的codec
type protoCodec struct{}

func (p *protoCodec) Marshal(v interface{}) ([]byte, error) {
	return proto.Marshal(v.(proto.Message))
}

func (p *protoCodec) Unmarshal(data []byte, v interface{}) error {
	return proto.Unmarshal(data, v.(proto.Message))
}

func (p *protoCodec) String() string {
	return "proto"
}

// Codec 返回了一个grpc.Codec类型的实例，
// 以protobuf原生codec为默认codec，实现了一个透明的Marshal和UnmarshMal
func Codec() grpc.Codec {
	return codecWithParent(&protoCodec{})
}

// 一个协议无感知的codec实现，返回一个grpc.Codec类型的实例
// 该函数尝试将gRPC消息当作raw bytes来实现，当尝试失败后，会有fallback作为一个后退的codec
func codecWithParent(fallback grpc.Codec) grpc.Codec {
	return &rawCodec{fallback}
}

// 序列化函数，
// 尝试将消息转换为*frame类型，并返回frame的payload实现序列化
// 若失败，则采用变量parentCodec中的Marshal进行序列化
func (c *rawCodec) Marshal(v interface{}) ([]byte, error) {
	out, ok := v.(*frame)
	if !ok {
		return c.parentCodec.Marshal(v)
	}
	return out.payload, nil
}

// 反序列化函数，
// 尝试通过将消息转为*frame类型，提取出payload到[]byte，实现反序列化
// 若失败，则采用变量parentCodec中的Unmarshal进行反序列化
func (c *rawCodec) Unmarshal(data []byte, v interface{}) error {
	dst, ok := v.(*frame)
	if !ok {
		return c.parentCodec.Unmarshal(data, v)
	}
	dst.payload = data
	return nil
}

func (c *rawCodec) String() string {
	return fmt.Sprintf("proxy>%s", c.parentCodec.String())
}
