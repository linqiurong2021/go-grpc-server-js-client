
let grpc = require('@grpc/grpc-js');
let messages = require('./proto/client/v1/user_pb');
let services = require('./proto/client/v1/user_grpc_pb')
let request = new messages.CreateRequest();
request.setVersion("v1")
// 获取需要实体
let user = new messages.User()
user.setAccount("123456") // 设置实体的属性值
user.setPassword("123456") // 设置实体的属性值
// 设置message
request.setUser(user)
//
let client = new services.UserServiceClient(
  '192.168.110.66:8088',
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
