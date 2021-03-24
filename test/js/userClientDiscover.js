
//
let grpc = require('@grpc/grpc-js');
let messages = require('./proto/client/v1/user_pb');
let pb = require('./proto/client/v1/user_grpc_pb')
//
let helloMessages = require('./proto/client/v1/hello_pb');
let helloPb = require('./proto/client/v1/hello_grpc_pb')

// 实例
let consul = require("consul")({host:"localhost",port:"8500"});


console.log("\n\n\nuser-service\n\n\n")
// 获取服务发现
consul.health.service({service:"user-service",passing: false,tag: null,dc:null}, function(err, result) {
    if (err) throw err;
    // result 为获取到的services
    if(result.length > 0 ){
        // 调用 负载均衡
        let service = result[0].Service // Service 注意大写
        console.log(service.Address,service.Port,service.Tags)
        let address = service.Address + ":" +service.Port
        //
        let client = new pb.UserServiceClient(
            address,
            grpc.credentials.createInsecure()
        );
        // 调用UserService
        let createRequest = new messages.CreateRequest();
        createRequest.setVersion("v1")
        // 获取需要实体
        let user = new messages.User()
        user.setAccount("123456") // 设置实体的属性值
        user.setPassword("123456") // 设置实体的属性值
        // 设置message
        createRequest.setUser(user)
        // 调用创建
        client.create(createRequest, function(err,data){
            if(err){
                console.error(err);
                return
            }
            console.log("receive",data)

        })

        let loginRequest = new messages.LoginRequest();
        loginRequest.setVersion("v1")
        // 获取需要实体
        loginRequest.setUser(user)
        // 调用登录
        client.login(loginRequest, function(err,data){
            if(err){
                console.error(err);
                return
            }
            console.log("receive",data)
        })
    }
    //
});
console.log("\n\n\nhello-service\n\n\n")
// 获取服务发现
consul.health.service({service:"hello-service",passing: false,tag: null,dc:null}, function(err, result) {
    if (err) throw err;
    // result 为获取到的services
    if(result.length > 0 ){
        // 调用 负载均衡
        let service = result[0].Service // Service 注意大写
        console.log(service.Address,service.Port,service.Tags)
        let address = service.Address + ":" +service.Port
        //
        let helloClient = new helloPb.HelloWorldServiceClient(
            address,
            grpc.credentials.createInsecure()
        );
        // 调用UserService
        let helloRequest = new helloMessages.HelloRequest();
        helloRequest.setName("Hello")
        // 调用创建
        helloClient.hello(helloRequest, function(err,data){
            if(err){
                console.error(err);
                return
            }
            console.log("receive",data)

        })

    }
    //
});
