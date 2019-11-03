PROTOBUF=$(GOPATH)/src/github.com/golang/protobuf
protoInclude=$(GOOGLE_APIS):$(PROTOBUF):/usr/local/include
GO_CACHE=$(GOCACHE)

gen:
	rm -rf ./api/*
	mkdir -p ./api/todo

	gunk generate github.com/AskAlexSharov/monorepo/api/todo

build:
	ls $(GO_CACHE)
	docker build --build-arg host_cache=$(GO_CACHE) -t "AskAlexSharov/todo" .