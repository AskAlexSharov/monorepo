//go:generate protoc -I ./vendor/github.com/golang/protobuf/ptypes -I ./shared   client.proto   --go_out=plugins=grpc:shared

package main
