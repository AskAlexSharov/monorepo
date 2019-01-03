PROTOBUF=$(GOPATH)/src/github.com/golang/protobuf
protInclude=$(PROTOBUF):/usr/local/include

gen:
	rm -rf ./api/*

	mkdir -p ./api/todo
	protoc -I $(protInclude):./../../api/todo \
		--go_out=plugins=grpc:./api/todo \
		todo.proto

	go generate ./...

