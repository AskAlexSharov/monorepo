package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	//"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"time"

	"github.com/nizsheanez/monorepo/todo/projects"
	"github.com/nizsheanez/monorepo/todo/todo"
)

var (
	grpcAddr = flag.String("grpc.addr", "localhost:8002", "Address for gRPC server")
)

func main() {
	//root := context.Background()

	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	todoApi := todo.NewApiClient(conn)
	projectsApi := projects.NewApiClient(conn)

	fmt.Println(todoApi.List(context.Background(), &todo.ListRequest{}))
	projectsApi.List(context.Background(), &projects.ListRequest{})
}
