//go:generate protoc -I ./vendor/github.com/golang/protobuf/ptypes -I ./todo     todo.proto     --go_out=plugins=grpc:todo
//go:generate protoc -I ./vendor/github.com/golang/protobuf/ptypes -I ./projects projects.proto --go_out=plugins=grpc:projects

package main
