package server

import (
	"context"

	pb "github.com/linqiurong2021/go-web-chat/proto/v1"
	"google.golang.org/grpc"
)

type UserServer struct{}

const Version = "v1"

// 这里实现服务端接口中的方法。
func (s *UserServer) Create(ctx context.Context, in *pb.CreateRequest) (*pb.Response, error) {
	//
	return &pb.Response{Code: 200, Msg: "created", Token: ""}, nil
}

// NewUserServer NewUserServer
func NewUserServer() *UserServer {
	return &UserServer{}
}

// Register Register
func Register(s *grpc.Server) {
	userServer := NewUserServer()
	pb.RegisterUserServiceServer(s, userServer)
	helloServer := NewHelloServer()
	pb.RegisterHelloWorldServiceServer(s, helloServer)
}

// 这里实现服务端接口中的方法。
func (s *UserServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.Response, error) {
	//
	return &pb.Response{Code: 200, Msg: "logined", Token: "123456"}, nil
}

type HelloServer struct{}

// 这里实现服务端接口中的方法。
func (s *HelloServer) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	//
	return &pb.HelloResponse{HelloString: in.Name}, nil
}

func NewHelloServer() *HelloServer {
	return &HelloServer{}
}
