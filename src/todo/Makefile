PROTOBUF=$(GOPATH)/src/github.com/golang/protobuf
protoInclude=:$(GOOGLE_APIS):$(PROTOBUF):/usr/local/include

gen:
	rm -rf ./api/*
	mkdir -p ./api/todo
	protoc -I $(protoInclude):./../../api/todo \
		--go_out=plugins=grpc:./api/todo \
		todo.proto

build:
	docker build -t "nizsheanez/todo" .