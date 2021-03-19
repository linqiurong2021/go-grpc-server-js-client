// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('grpc');
var hello_pb = require('./hello_pb.js');

function serialize_v1_HelloRequest(arg) {
  if (!(arg instanceof hello_pb.HelloRequest)) {
    throw new Error('Expected argument of type v1.HelloRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_HelloRequest(buffer_arg) {
  return hello_pb.HelloRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v1_HelloResponse(arg) {
  if (!(arg instanceof hello_pb.HelloResponse)) {
    throw new Error('Expected argument of type v1.HelloResponse');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_HelloResponse(buffer_arg) {
  return hello_pb.HelloResponse.deserializeBinary(new Uint8Array(buffer_arg));
}


// 服务HelloWorld
var HelloWorldServiceService = exports.HelloWorldServiceService = {
  hello: {
    path: '/v1.HelloWorldService/hello',
    requestStream: false,
    responseStream: false,
    requestType: hello_pb.HelloRequest,
    responseType: hello_pb.HelloResponse,
    requestSerialize: serialize_v1_HelloRequest,
    requestDeserialize: deserialize_v1_HelloRequest,
    responseSerialize: serialize_v1_HelloResponse,
    responseDeserialize: deserialize_v1_HelloResponse,
  },
};

exports.HelloWorldServiceClient = grpc.makeGenericClientConstructor(HelloWorldServiceService);
