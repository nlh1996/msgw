package main

import (
	"go-gateway/client"
	"go-gateway/server"
)

// func write(data *[]byte) {
// 	c1 := &test.Class{
// 		Num: 1,
// 		Students: []*model.Student{
// 			{Name: "xiaoming", Age: 21, Sex: test.Sex_MAN},
// 			{Name: "xiaohua", Age: 21, Sex: test.Sex_WOMAN},
// 			{Name: "xiaojin", Age: 21, Sex: test.Sex_MAN},
// 		},
// 	}

// 	// 使用protobuf工具把struct数据类型格式化成字节数组（压缩和编码）
// 	*data, _ = proto.Marshal(c1)
// }

// func read(data []byte) {
// 	class := new(model.Class)

// 	// 使用protobuf工具把字节数组解码成struct(解码)
// 	proto.Unmarshal(data, class)

// 	log.Println(class.Num)
// 	for _, v := range class.Students {
// 		log.Println(v.Name, v.Age, v.Sex)
// 	}
// }

func main() {
	go client.Init()
	go client.StreamClientInit()
	server.Init()
}
