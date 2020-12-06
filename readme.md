## 双向流gRPC
一个简单的Go语言的双向流RPC示例
### 项目目录
```shell
go-grpc-example
|   go.mod
|   go.sum
|   readme.md
|
+---client
|       stream.go
|
+---proto
|       counter.proto
|
+---server
        stream.go
```
### Protoc版本
```shell
$ protoc --version
libprotoc 3.14.0
```
### 运行
```shell
protoc -I proto/ --go-grpc_out=proto/ proto/counter.proto
go run server/stream.go
go run client/stream.go
```
