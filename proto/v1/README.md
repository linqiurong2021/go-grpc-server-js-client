
## 生成js 

```shell


# 1: 前端需要安装 
"@grpc/grpc-js"
"google-protobuf"
"grpc-tools"

# 2: 执行命令 --plugin=protoc-gen-grpc=xxxx 其中xxx是安装grpc-tools的路径
# 注意点  Go package "." has inconsistent names hello # 如果有多个.proto 则package v1;//包名需要一致
protoc --js_out=import_style=commonjs,binary:. --plugin=protoc-gen-grpc=../../test/js/node_modules/grpc-tools/bin/grpc_node_plugin --grpc_out=.  --go_out=plugins=grpc:. ./*.proto 

# 3: 生成的 xxx.grpc_pb.js 中的 var grpc = require('grpc'); 修改为 : var grpc = require('@grpc/grpc-js')

```

# 调用 1

```javascript

let grpc = require('@grpc/grpc-js');
// 注意路径
let messages = require('./proto/client/v1/hello_pb');
let services = require('./proto/client/v1/hello_grpc_pb')
let request = new messages.HelloRequest();
request.setName("Hello World")
let client = new services.HelloWorldServiceClient(
  '192.168.110.66:8088', // server的 地址
  grpc.credentials.createInsecure()
);
//
console.log(request,'request')
// 调用
client.hello(request, function(err,data){
  // console.log(err,data)
  if(err){
    console.error(err);
  }
  console.log("receive",data)
  // console.log(data);
  console.log(data.getHellostring());
})

console.log("end")

```

# 调用 1

```javascript


let grpc = require('@grpc/grpc-js');
let messages = require('./proto/client/v1/user_pb');
let services = require('./proto/client/v1/user_grpc_pb')
let request = new messages.CreateRequest();
// 
request.setVersion("v1")
// 获取需要实体
let user = new messages.User()
user.setAccount("123456") // 设置实体的属性值
user.setPassword("123456") // 设置实体的属性值
// 设置message
request.setUser(user)
//
let client = new services.UserServiceClient(
  '192.168.110.66:8088', // grpc的服务地址
  grpc.credentials.createInsecure()
);

//
console.log(request,'request')
// 调用
client.create(request, function(err,data){
  // console.log(err,data)
  if(err){
    console.error(err);
  }
  console.log("receive",data)
  // console.log(data);
  console.log(data.getCode(), 'code');
  console.log(data.getMsg(),'msg');
  // console.log(data.getData(),'data');
  // console.log(data.getCode());
})

//
client.login(request, function(err,data){
  // console.log(err,data)
  if(err){
    console.error(err);
  }
  console.log("receive",data)
  // console.log(data);
  // console.log(data.getCode(), 'code');
  // console.log(data.getMsg(),'msg');
  // console.log(data.getData(),'data');
  // console.log(data.getToken());
})

console.log("end")


```

