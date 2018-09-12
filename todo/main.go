package main

import (
	"context"
	"fmt"
	"github.com/globalsign/mgo"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	grpc_runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/nizsheanez/monorepo/todo/model"
	"github.com/nizsheanez/monorepo/todo/sdk/v2"
	"github.com/nizsheanez/monorepo/todo/service"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-client-go/rpcmetrics"
	prometheus_metrics "github.com/uber/jaeger-lib/metrics/prometheus"
	"github.com/urfave/cli"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"runtime"
)

func main() {
	app := cli.NewApp()
	app.Name = path.Base(os.Args[0])
	app.Usage = "Todo app"
	app.Version = "0.0.1"
	app.Flags = commonFlags
	app.Action = start

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

// Panic handler prints the stack trace when recovering from a panic.
var panicHandler = grpc_recovery.RecoveryHandlerFunc(func(p interface{}) error {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	log.Errorf("panic recovered: %+v", string(buf))
	return status.Errorf(codes.Internal, "%s", p)
})

func start(c *cli.Context) {
	// Logrus
	logger := log.NewEntry(log.New())
	grpc_logrus.ReplaceGrpcLogger(logger)
	log.SetLevel(log.DebugLevel)

	tracer, closer, err := initTracer(c, logger)
	if err != nil {
		logger.Fatalf("Cannot initialize Jaeger Tracer %s", err)
	}
	defer closer.Close()

	session, err := initDb(c)
	if err != nil {
		panic(err)
	}

	// Set GRPC Interceptors
	server := grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_prometheus.StreamServerInterceptor,
			grpc_logrus.StreamServerInterceptor(logger),
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(panicHandler)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_logrus.UnaryServerInterceptor(logger),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(panicHandler)),
		)),
	)

	{ // register rpc services

		todoCollection := session.DB("alex").C("todo")

		// todo service
		todoService := &service.TodoService{Model: &model.TodoModel{Collection: todoCollection}}
		todo_sdk.RegisterTodoServiceServer(server, todoService)

		// ... another services ...
	}

	initPrometheus(c)

	log.Println("Starting Grpc service... " + grpcAddr(c))
	lis, err := net.Listen("tcp", grpcAddr(c))
	if err != nil {
		logger.Fatalf("Failed to listen: %v", grpcAddr(c))
	}

	go func() {
		reflection.Register(server)
		err := server.Serve(lis)
		if err != nil {
			logger.Fatalf(err.Error())
		}
	}()

	mux := grpc_runtime.NewServeMux()
	{
		// create grpc client, http gateway will use it
		conn, err := grpc.Dial(grpcAddr(c), grpc.WithInsecure())
		if err != nil {
			logger.Fatalf("Couldn't contact grpc server: " + err.Error())
		}

		err = todo_sdk.RegisterTodoServiceHandler(context.Background(), mux, conn)
		if err != nil {
			logger.Fatalf("Cannot serve http api, " + err.Error())
		}
	}

	grpc_prometheus.Register(server)
	http.ListenAndServe(httpAddr(c), mux)
}

func grpcAddr(c *cli.Context) string {
	return "127.0.0.1:" + c.String("bind-grpc")
}

func httpAddr(c *cli.Context) string {
	return "127.0.0.1:" + c.String("bind-http")
}

func mongoAddr(c *cli.Context) string {
	return c.String("db-host") + ":" + c.String("db-port")
}

func jaegerAddr(c *cli.Context) string {
	return c.String("jaeger-host") + ":" + c.String("jaeger-port")
}

func initTracer(c *cli.Context, logger *log.Entry) (opentracing.Tracer, io.Closer, error) {
	// Prometheus monitoring
	metrics := prometheus_metrics.New()

	// Jaeger tracing
	cfg := config.Configuration{
		ServiceName: "todo",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: c.Float64("jaeger-sampler"),
		},
		Reporter: &config.ReporterConfig{
			LocalAgentHostPort: jaegerAddr(c),
		},
	}
	return cfg.NewTracer(
		config.Logger(jaegerLoggerAdapter{logger}),
		config.Observer(rpcmetrics.NewObserver(metrics.Namespace("todo", nil), rpcmetrics.DefaultNameNormalizer)),
	)
}

func initDb(c *cli.Context) (*mgo.Session, error) {
	sess, err := mgo.Dial(mongoAddr(c))
	if err != nil {
		return nil, errors.New("Can't connect to Mongo on address: " + mongoAddr(c) + ", by reason: " + err.Error())
	}
	return sess, nil
}

func initPrometheus(c *cli.Context) {
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe(c.String("bind-prometheus-http"), mux)
	}()
}

type jaegerLoggerAdapter struct {
	logger *log.Entry
}

func (l jaegerLoggerAdapter) Error(msg string) {
	l.logger.Error(msg)
}

func (l jaegerLoggerAdapter) Infof(msg string, args ...interface{}) {
	l.logger.Info(fmt.Sprintf(msg, args...))
}
