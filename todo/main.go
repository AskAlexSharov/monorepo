//go:generate protoc -I ./vendor/github.com/golang/protobuf/ptypes -I ./todo     todo.proto     --micro_out=./todo --go_out=./todo
//go:generate protoc -I ./vendor/github.com/golang/protobuf/ptypes -I ./projects projects.proto --micro_out=./projects --go_out=./projects

package main
