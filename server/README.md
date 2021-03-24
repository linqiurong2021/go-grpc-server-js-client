# Consul服务注册
## consul 下载
```shell
go get github.com/hashicorp/consul/api
```
## 获取配置
```go
config := api.DefaultConfig()
```
## 创建consul Client
```go
client, err := api.NewClient(config)
if err != nil {
    return err
}
```

## 配置服务信息
```go
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
```

## 注册服务
```go
err = client.Agent().ServiceRegister(&reg)
return err
```