package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	//"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"time"

	//"github.com/nizsheanez/monorepo/todo/projects"
	todoV1 "github.com/nizsheanez/monorepo/todo/client"
	"github.com/nizsheanez/monorepo/todo/client/v2"
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

	todoService := todo.NewTodoServiceClient(conn)
	todoServiceV1 := todoV1.NewTodoServiceClient(conn)

	//projectsApi := projects.NewApiClient(conn)

	fmt.Println(todoService.List(context.Background(), &todo.ListRequest{}))
	fmt.Println(todoServiceV1.List(context.Background(), &todoV1.ListRequest{}))

	//projectsApi.List(context.Background(), &projects.ListRequest{})
}
