package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/AskAlexSharov/monorepo/src/proj2/api/todo"
	"github.com/hashicorp/logutils"
	"google.golang.org/grpc"
)

var (
	grpcAddr = flag.String("grpc.addr", "localhost:8002", "Address for gRPC server")
)

func main() {
	//root := context.Background()

	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)

	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	todoService := todo.NewTodoServiceClient(conn)

	//projectsApi := projects.NewApiClient(conn)

	fmt.Println(todoService.List(context.Background(), &todo.ListRequest{}))

	//projectsApi.List(context.Background(), &projects.ListRequest{})
}
