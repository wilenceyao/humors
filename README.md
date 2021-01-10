# humors

![](https://travis-ci.com/wilenceyao/humors.svg?branch=main)

基于 MQTT + Protobuf 的RPC框架实现，具体用法可参考 example

## 安装protoc和protoc-gen-go
参考：https://developers.google.com/protocol-buffers/docs/gotutorial

## 安装protoc-gen-go-humors
```
go get github.com/wilenceyao/humors/cmd/protoc-gen-go-humors
```

## 定义Protobuf服务
```
syntax = "proto3";
package api;

option go_package = "github.com/wilenceyao/humors/example/api";

message SumRequest {
    int32 a = 1;
    int32 b = 2;
}

message SumResponse {
    int32 sum = 1;
}

service ExampleService {
    rpc Sum(SumRequest) returns (SumResponse);
}
```

## 生成stub代码
```
protoc --go_out=. --go_opt=paths=source_relative --go-humors_out=. --go-humors_opt=paths=source_relative example/api/api.proto
```

## Server 注册服务

[server代码示例](/example/server/main.go)

## Client调用服务

[client代码示例](/example/client/main.go)

## 问题咨询和反馈
请提交issue