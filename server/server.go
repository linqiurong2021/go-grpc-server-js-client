package server

import (
	"context"
	"log"

	"github.com/hashicorp/consul/api"
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

// RegisterToConsul 注册到consul
func (s *UserServer) RegisterToConsul() error {
	// 获取配置信息
	config := api.DefaultConfig()
	// 创建client
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}
	// 配置服务信息
	reg := api.AgentServiceRegistration{
		ID: "go-grpc-server-js-client",// 可用配置文件
		Name: "user-service",// 可用配置文件
		Tags: []string{"User","Grpc"},// 可用配置文件
		Address: "192.168.110.66", // 可用配置文件
		Port: 8088,// 可用配置文件
		Check: &api.AgentServiceCheck{
			Name: "user-service-heart-beats",// 可用配置文件
			Interval: "10s",// 可用配置文件
			Timeout: "5s",// 可用配置文件
			TCP: "192.168.110.66:8088",// 可用配置文件
		},
	}
	// 注册
	err = client.Agent().ServiceRegister(&reg)
	return err
}

// NewUserServer NewUserServer
func NewUserServer() *UserServer {
	return &UserServer{}
}

// Register Register
func Register(s *grpc.Server) {
	userServer := NewUserServer()
	// 把服务注册到Consul
	err := userServer.RegisterToConsul()
	if err != nil {
		log.Fatalf("RegisterToConsul Error %s \n", err)
	}

	pb.RegisterUserServiceServer(s, userServer)
	helloServer := NewHelloServer()
	helloServer.RegisterToConsul()
	pb.RegisterHelloWorldServiceServer(s, helloServer)
}

// 这里实现服务端接口中的方法。
func (s *UserServer) Login(ctx context.Context, in *pb.LoginRequest) (*pb.Response, error) {
	//
	return &pb.Response{Code: 200, Msg: "logined", Token: "123456"}, nil
}

type HelloServer struct{}

// 这里实现服务端接口中的方法。
func (h *HelloServer) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	//
	return &pb.HelloResponse{HelloString: in.Name}, nil
}
// RegisterToConsul 注册到consul
func (h *HelloServer) RegisterToConsul() error {
	// 获取配置信息
	config := api.DefaultConfig()
	// 创建client
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}
	// 配置服务信息
	reg := api.AgentServiceRegistration{
		ID: "go-grpc-server-js-client-hello",// 可用配置文件
		Name: "hello-service",// 可用配置文件
		Tags: []string{"Hello","Grpc"},// 可用配置文件
		Address: "192.168.110.66", // 可用配置文件
		Port: 8088,// 可用配置文件
		Check: &api.AgentServiceCheck{
			Name: "hello-service-heart-beats",// 可用配置文件
			Interval: "10s",// 可用配置文件
			Timeout: "5s",// 可用配置文件
			TCP: "192.168.110.66:8088",// 可用配置文件
		},
	}
	// 注册
	err = client.Agent().ServiceRegister(&reg)
	return err
}

func NewHelloServer() *HelloServer {
	return &HelloServer{}
}
