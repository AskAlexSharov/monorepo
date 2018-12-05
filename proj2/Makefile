GOOGLE_APIS=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
PROTOBUF=$(GOPATH)/src/github.com/golang/protobuf
protoPaths=$(GOOGLE_APIS):$(PROTOBUF):/usr/local/include

gen:
	rm -rf ./api/*
	mkdir -p ./api/todo/v2
	protoc $(protIncl):./../api/todo/v2 --go_out=plugins=grpc+protoc-gen-grpc-gateway+protoc-gen-swagger:./api/todo/v2 --swagger_out=logtostderr=true:./api/todo/v2 --grpc-gateway_out=logtostderr=true:./api/todo/v2 todo.proto
	mkdir -p ./api/todo/v1
	protoc -I $(protoPaths):./../api/todo/v1 --go_out=plugins=grpc+protoc-gen-grpc-gateway+protoc-gen-swagger:./api/todo/v1 --swagger_out=logtostderr=true:./api/todo/v1 --grpc-gateway_out=logtostderr=true:./api/todo/v1 todo.proto