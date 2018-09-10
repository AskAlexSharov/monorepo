package main

import (
	"context"
	"github.com/nizsheanez/monorepo/todo/todo"
	"google.golang.org/grpc"
	"log"
)

var serverAddr = "127.0.0.1"

func main() {
	conn, err := grpc.Dial(serverAddr)
	if err != nil {
		log.Fatalf("can't connect todo: %s", err)
	}

	todoApi := todo.NewApiClient(conn)
	projectsApi := projects.NewApiClient(conn)

	todoApi.List(context.Background(), projectsApi)
}
