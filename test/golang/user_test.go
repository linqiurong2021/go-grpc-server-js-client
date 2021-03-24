package test

import (
	"context"
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"net"
	"strconv"
	"testing"

	pb "github.com/linqiurong2021/go-web-chat/proto/v1"
	"github.com/linqiurong2021/go-web-chat/server"
	"google.golang.org/grpc"
)

const port = ":8089"

func initGrpcServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	server.Register(s)
	s.Serve(lis)
}

// 服务发现
func discoverServices()  (services []*api.ServiceEntry , err error) {
	// 默认配置
	config := api.DefaultConfig()
	// 创建Client
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	// 获取健康的服务
	services, _ ,err = client.Health().Service("user-service","",false,nil)
	return services, err
}

// 服务发现测试
func TestDiscoverServices(t *testing.T)  {
	//
	services , err := discoverServices()
	if err != nil {
		log.Fatalf("discoverServices err %s \n", err)
	}
	//
	if len(services) > 0 {
		service := services[0].Service
		address := service.Address  + ":" + strconv.Itoa(service.Port)
		con, err := grpc.Dial(address, grpc.WithInsecure())
		if err != nil {
			t.Errorf("connect %s error : %s\n", port, err)
		}
		defer con.Close()
		//
		client := pb.NewUserServiceClient(con)
		resp, err := client.Create(context.Background(), &pb.CreateRequest{})
		if err != nil {
			// log.Printf(" request failure : error %s\n", err)
			t.Errorf(" request failure : error %s\n", err)
		}
		if resp.Code == 200 && resp.Msg != "created" {
			t.Errorf(" create failure : error %s\n", resp.Msg)
		}
	}else{
		log.Fatalf("services not found")
	}



}

//
func TestUserCreate(t *testing.T) {
	//
	con, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		t.Errorf("connect %s error : %s\n", port, err)
	}
	defer con.Close()
	//
	client := pb.NewUserServiceClient(con)
	resp, err := client.Create(context.Background(), &pb.CreateRequest{})
	if err != nil {
		// log.Printf(" request failure : error %s\n", err)
		t.Errorf(" request failure : error %s\n", err)
	}
	if resp.Code == 200 && resp.Msg != "created" {
		t.Errorf(" create failure : error %s\n", resp.Msg)
	}
}

//
func TestUserLogin(t *testing.T) {
	//
	con, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		t.Errorf("connect %s error : %s\n", port, err)
	}
	defer con.Close()
	//
	client := pb.NewUserServiceClient(con)
	resp, err := client.Login(context.Background(), &pb.LoginRequest{Version: "v1", User: &pb.User{Account: "123456", Password: "123456"}})
	if err != nil {
		// log.Printf(" request failure : error %s\n", err)
		t.Errorf(" request failure : error %s\n", err)
	}
	if resp.Code == 200 && resp.Msg != "logined" {
		t.Errorf(" login failure : error %s\n", resp.Msg)
	}
}

// 初始化
func TestMain(m *testing.M) {
	fmt.Println("init main")
	// 启用服务
	go initGrpcServer()
	m.Run()
}
