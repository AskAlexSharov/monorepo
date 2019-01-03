PROTOBUF=$(GOPATH)/src/github.com/golang/protobuf
protoInclude=:$(GOOGLE_APIS):$(PROTOBUF):/usr/local/include

gen:
	rm -rf ./api/*
	mkdir -p ./api/todo
	protoc -I $(protoInclude):./../../api/todo \
		--go_out=plugins=grpc:./api/todo \
		todo.proto

    GO111MODULE=on wire

build:
	docker build -t "nizsheanez/todo:1.0.0" -t nizsheanez/todo:latest .