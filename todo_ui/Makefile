GOOGLE_APIS=$(GOPATH)/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis
PROTOBUF=$(GOPATH)/src/github.com/golang/protobuf
protoc=protoc -I ./../api/todo/v2:$(GOOGLE_APIS):$(PROTOBUF):/usr/local/include

gen:
	rm -rf ./src/api/*
	mkdir -p ./src/api/todo/v2
#	$(protoc) --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts --js_out=import_style=commonjs,binary:./api/todo/v2 --ts_out=service=true:./src/api/todo/v2 --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./src/api/todo/v2 todo.proto
	$(protoc) --plugin=protoc-gen-ts=protoc-gen-grpc-web=/usr/local/bin/protoc-gen-grpc-web --grpc-web_out=import_style=typescript,mode=grpcwebtext:./src/api/todo/v2 todo.proto
#	echo ${CURDIR}
	#hack for https://github.com/improbable-eng/ts-protoc-gen/issues/99
#	grep -v 'google_api_annotations_pb' ./api/todo/v2/todo_pb.d.ts > ./api/todo/v2/todo_pb.d.ts
#	grep -v 'google_api_annotations_pb' ./api/todo/v2/todo_grpc_web_pb.js > ./api/todo/v2/todo_grpc_web_pb.js
#	grep -v 'google_api_annotations_pb' ./api/todo/v2/todo_pb.js > ./api/todo/v2/todo_pb.js
