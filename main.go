package main

import (
	"log"
	"net"

	"github.com/linqiurong2021/go-web-chat/server"
	"google.golang.org/grpc"
)

// 创建并启动一个 gRPC 服务的过程：创建监听套接字、创建服务端、注册服务、启动服务端。
func main() {
	lis, err := net.Listen("tcp", ":8088")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server.Register(s)
	s.Serve(lis)
}
