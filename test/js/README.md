# 使用说明

## 1: 初始化npm 与 安装包
```javascript
  npm init

  npm install @grpc/grpc-js google-protobuf grpc-tools

```

## 2: 获取后端生成的js文件 例: proto/client/v1/hello_grpc_pb.js proto/client/v1/hello_pb.js 

## 2.1 如果后端未生成可让后端提供 .proto文件 后 可自已生成 
```javascript

protoc --js_out=import_style=commonjs,binary:./proto/client/v2/ --plugin=protoc-gen-grpc=./node_modules/grpc-tools/bin/grpc_node_plugin --grpc_out=./proto/client/v2/

```
## 3: 修改 */hello_grpc_pb.js 中的 var grpc = require('grpc'); ==> : var grpc = require('@grpc/grpc-js')


## 4: 调用 例: helloClient.js  userClient.js

