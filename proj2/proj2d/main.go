package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/nizsheanez/monorepo/todo/api/todo/v2"
	"github.com/nizsheanez/monorepo/todo/sdk/v2"
	"log"
	"strings"

	//"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"time"

	//"github.com/nizsheanez/monorepo/todo/projects"
	"github.com/nizsheanez/monorepo/proj2/api/todo/v1"
	todo_sdk_v1 "github.com/nizsheanez/monorepo/proj2/api/todo/v1"
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

	strings.Replace()
	strings.NewReplacer()
	todoService := todo.NewTodoServiceClient(conn)
	todoServiceV1 := todo_sdk_v1.NewTodoServiceClient(conn)

	//projectsApi := projects.NewApiClient(conn)

	fmt.Println(todoService.List(context.Background(), &todo.ListRequest{}))
	fmt.Println(todoServiceV1.List(context.Background(), &todo_sdk_v1.ListRequest{}))

	//projectsApi.List(context.Background(), &projects.ListRequest{})
}
