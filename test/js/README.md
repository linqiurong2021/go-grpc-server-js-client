# 使用说明

## 1: 初始化npm 与 安装包
```javascript
  npm init

  npm install @grpc/grpc-js google-protobuf grpc-tools
```

## 2: 获取后端生成的js文件 
```javascript

例: proto/client/v1/hello_grpc_pb.js proto/client/v1/hello_pb.js 

```

## 2.1 如果后端未生成可让后端提供 .proto文件 后 可自已生成 
```javascript

protoc --js_out=import_style=commonjs,binary:./proto/client/v2/ --plugin=protoc-gen-grpc=./node_modules/grpc-tools/bin/grpc_node_plugin --grpc_out=./proto/client/v2/

```
## 3: 修改文件
```javascript

修改 */hello_grpc_pb.js 中的 var grpc = require('grpc'); ==> : var grpc = require('@grpc/grpc-js')

```

## 4: 调用 
```javascript
例: helloClient.js  userClient.js
```


## 服务发现
```javascript
// 实例
let consul = require("consul")({host:"localhost",port:"8500"});
// 获取服务发现
// 获取健康的 user-service 服务
consul.health.service({service:"user-service",passing: false,tag: null,dc:null}, function(err, result) {
    if (err) throw err;
    // 健康的服务
});
```

