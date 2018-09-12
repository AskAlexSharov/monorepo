//go:generate protoc -I$(pwd)/:$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis:$GOPATH/src/github.com/golang/protobuf:$GOPATH/src:/usr/local/include --go_out=plugins=grpc:. todo.proto

package todo_sdk
