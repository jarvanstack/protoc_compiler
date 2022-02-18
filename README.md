# protoc_compiler 

开箱即用的 protoc 编译环境 + 一键编译的脚本, 开箱即用, 解决编译 pb 文件问题.

> protoc 安装教程: https://blog.csdn.net/jarvan5/article/details/118026721

## QUICK START

### 1. 新建一个 user.proto

```protobuf
syntax = "proto3";
package proto;
option go_package = "/user";
service User {
  rpc GetUser(GetUserRequest) returns (GetUserResponse) {}
}
message GetUserRequest {
  string id = 1;
}
message GetUserResponse {
  UserDTO user = 1;
}
message UserDTO {
  string id = 1 ;
  string name = 2;
}
```

### 2. 编译

```bash
# 为普通 go 文件
protoc --go_out=plugins=grpc,paths=source_relative:. ./user.proto 
# 安装缺失的依赖,
go mod tidy
# 依赖可以无法自动加载可能是 vscode go tools 的问题, 更新下
```
