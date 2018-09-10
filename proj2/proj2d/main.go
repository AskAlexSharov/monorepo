package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/sony/gobreaker"
	"google.golang.org/grpc"
	"os"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/nizsheanez/monorepo/todo/projects"
	"github.com/nizsheanez/monorepo/todo/todo"
)

var (
	grpcAddr = flag.String("grpc.addr", "localhost:8002", "Address for gRPC server")
)

func main() {

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.With(logger, "caller", log.DefaultCaller)
	logger = log.With(logger, "transport", "grpc")

	root := context.Background()

	cc, err := grpc.Dial(*grpcAddr)
	if err != nil {
		_ = logger.Log("err", err)
		os.Exit(1)
	}
	defer cc.Close()

	//conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	conn, err := grpc.Dial("127.0.0.1:8083", grpc.WithInsecure(), grpc.WithTimeout(time.Second))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	todoApi := todo.NewApiService("todo")
	projectsApi := projects.NewApiService("todo", service.Client())

	todoApi.List(context.Background(), &todo.ListRequest{})
	projectsApi.List(context.Background(), &projects.ListRequest{})

}
