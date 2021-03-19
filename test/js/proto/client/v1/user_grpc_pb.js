// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var user_pb = require('./user_pb.js');

function serialize_v1_CreateRequest(arg) {
  if (!(arg instanceof user_pb.CreateRequest)) {
    throw new Error('Expected argument of type v1.CreateRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_CreateRequest(buffer_arg) {
  return user_pb.CreateRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v1_LoginRequest(arg) {
  if (!(arg instanceof user_pb.LoginRequest)) {
    throw new Error('Expected argument of type v1.LoginRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_LoginRequest(buffer_arg) {
  return user_pb.LoginRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_v1_Response(arg) {
  if (!(arg instanceof user_pb.Response)) {
    throw new Error('Expected argument of type v1.Response');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_v1_Response(buffer_arg) {
  return user_pb.Response.deserializeBinary(new Uint8Array(buffer_arg));
}


// 用户服务
var UserServiceService = exports.UserServiceService = {
  create: {
    path: '/v1.UserService/Create',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.CreateRequest,
    responseType: user_pb.Response,
    requestSerialize: serialize_v1_CreateRequest,
    requestDeserialize: deserialize_v1_CreateRequest,
    responseSerialize: serialize_v1_Response,
    responseDeserialize: deserialize_v1_Response,
  },
  login: {
    path: '/v1.UserService/Login',
    requestStream: false,
    responseStream: false,
    requestType: user_pb.LoginRequest,
    responseType: user_pb.Response,
    requestSerialize: serialize_v1_LoginRequest,
    requestDeserialize: deserialize_v1_LoginRequest,
    responseSerialize: serialize_v1_Response,
    responseDeserialize: deserialize_v1_Response,
  },
};

exports.UserServiceClient = grpc.makeGenericClientConstructor(UserServiceService);
