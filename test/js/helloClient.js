
let grpc = require('@grpc/grpc-js');
let messages = require('./proto/client/v1/hello_pb');
let services = require('./proto/client/v1/hello_grpc_pb')
let request = new messages.HelloRequest();
request.setName("Hello World")
let client = new services.HelloWorldServiceClient(
  '192.168.110.66:8088',
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
