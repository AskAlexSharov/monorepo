GOOGLE_APIS=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
PROTOBUF=$(GOPATH)/src/github.com/golang/protobuf
protoc=protoc -I ./../api/todo/v2:$(GOOGLE_APIS):$(PROTOBUF):/usr/local/include

gen:
	rm -rf ./api/*
	mkdir -p ./api/todo/v2
	$(protoc) --go_out=plugins=grpc+protoc-gen-grpc-gateway+protoc-gen-swagger:./api/todo/v2 --swagger_out=logtostderr=true:./api/todo/v2 --grpc-gateway_out=logtostderr=true:./api/todo/v2 todo.proto