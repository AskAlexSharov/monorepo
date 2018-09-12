package main

import (
	"context"
	"flag"
	"fmt"
	//"github.com/nizsheanez/monorepo/todo/projects"
	//"github.com/nizsheanez/monorepo/todo/todo"
	"log"

	//"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"time"

	//"github.com/nizsheanez/monorepo/todo/projects"
	todoV1 "github.com/nizsheanez/monorepo/todo/client"
	todo "github.com/nizsheanez/monorepo/todo/client/v2"
)

var (
	grpcAddr = flag.String("grpc.addr", "localhost:8002", "Address for gRPC server")	)

func main() {
	//root := context.Background()

	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	todoApi := todo.NewTodoServiceClient(conn)
	todoApiV1 := todoV1.NewTodoServiceClient(conn)

	//projectsApi := projects.NewApiClient(conn)

	fmt.Println(todoApi.ListTodo(context.Background(), &todo.ListTodoRequest{}))
	fmt.Println(todoApiV1.ListTodo(context.Background(), &todoV1.ListTodoRequest{}))

	//projectsApi.List(context.Background(), &projects.ListRequest{})
}
