GOOGLE_APIS=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
PROTOBUF=$(GOPATH)/src/github.com/golang/protobuf
protoc=protoc -I ./../api/todo/v2:$(GOOGLE_APIS):$(PROTOBUF):/usr/local/include

gen:
	mkdir -p ./api/todo/v2
	$(protoc) --js_out=import_style=typescript,binary:./api/todo/v2 --grpc-web_out=import_style=typescript,mode=grpcwebtext:./api/todo/v2 todo.proto