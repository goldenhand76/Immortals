# gRPC 

## Install protobuf (Protocol Buffers)
https://github.com/protocolbuffers/protobuf/releases

## Go gRPC runtime installation : 
```bash
go get google.golang.org/grpc
```
## Generating gRPC go files : 
```bash
protoc .\api\gRPC\immo\immo.proto --go_out=.\api\gRPC\immo --go-grpc_out=.\api\gRPC\immo
```

## Generating gRPC C++ Files : 
```bash
protoc .\api\gRPC\immo\immo.proto --cpp_out=.\ --grpc_out=.\ --plugin=protoc-gen-grpc="C:\msys64\mingw64\bin\grpc_cpp_plugin.exe"
```