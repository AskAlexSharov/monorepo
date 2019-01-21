#PROTOBUF=$(GOPATH)/src/github.com/golang/protobuf
#protoInclude=$(GOOGLE_APIS):$(PROTOBUF):/usr/local/include
GO_CACHE=$(GOCACHE)

#gen:
#	rm -rf ./api/*
#	mkdir -p ./api/todo
#	protoc -I $(protoInclude):./../../api/todo \
#		--go_out=plugins=grpc:./api/todo \
#		todo.proto

build:
	ls $(GO_CACHE)
	docker build --build-arg host_cache=$(GO_CACHE) -t "nizsheanez/opengolang" .

