# Consul 服务发现

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

## 获取健康服务
```go
// services 可做负载均衡
services, _ ,err = client.Health().Service("user-service","",false,nil)
return services, err
```

##  获取某个服务的服务器地址与端口
```go
service := services[0].Service
address := service.Address  + ":" + strconv.Itoa(service.Port)
// 拨号
con, err := grpc.Dial(address, grpc.WithInsecure())
```