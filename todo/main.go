package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"runtime"
	"sync"

	"github.com/globalsign/mgo"
	"github.com/google/go-cloud/health"
	"github.com/google/go-cloud/runtimevar"
	"github.com/google/go-cloud/server"
	"github.com/google/go-cloud/wire"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	grpc_runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/nizsheanez/monorepo/todo/api/todo/v2"
	"github.com/nizsheanez/monorepo/todo/model"
	"github.com/nizsheanez/monorepo/todo/service"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	//"github.com/uber/jaeger-client-go/config"
	//"github.com/uber/jaeger-client-go/rpcmetrics"
	//prometheus_metrics "github.com/uber/jaeger-lib/metrics/prometheus"
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

var applicationSet = wire.NewSet(
	newApplication,
	appHealthChecks,
	trace.AlwaysSample,
)

type fakeHealthChecker struct{}

func (*fakeHealthChecker) CheckHealth() error {
	return nil
}

// appHealthChecks returns a health check for the database. This will signal
// to Kubernetes or other orchestrators that the server should not receive
// traffic until the server is able to connect to its database.
func appHealthChecks(db *mgo.Session) ([]health.Checker, func()) {
	//dbCheck := sqlhealth.New(db)
	c := &fakeHealthChecker{}
	list := []health.Checker{c}
	return list, func() {
		//dbCheck.Stop()
	}
}

// application is the main server struct for Guestbook. It contains the state of
// the most recently read message of the day.
type application struct {
	srv        *server.Server
	logger     *log.Entry
	grpcServer *grpc.Server
	db         *mgo.Session

	// The following fields are protected by mu:
	mu   sync.RWMutex
	motd string // message of the day
}

// newApplication creates a new application struct based on the backends
func newApplication(
	srv *server.Server,
	logger *log.Entry,
	db *mgo.Session,
	grpcServer *grpc.Server,
	motdVar *runtimevar.Variable) *application {
	app := &application{
		srv:        srv,
		logger:     logger,
		grpcServer: grpcServer,
		db:         db,
	}
	go app.watchMOTDVar(motdVar)
	return app
}

// watchMOTDVar listens for changes in v and updates the app's message of the
// day. It is run in a separate goroutine.
func (app *application) watchMOTDVar(v *runtimevar.Variable) {
	ctx := context.Background()
	for {
		snap, err := v.Watch(ctx)
		if err != nil {
			log.Printf("watch MOTD variable: %v", err)
			continue
		}
		log.Println("updated MOTD to", snap.Value)
		app.mu.Lock()
		app.motd = snap.Value.(string)
		app.mu.Unlock()
	}
}

func start(c *cli.Context) {
	//tracer, closer, err := initTracer(c, logger)
	//if err != nil {
	//	logger.Fatalf("Cannot initialize Jaeger Tracer %s", err)
	//}
	//defer closer.Close()

	var app *application
	var cleanup func()
	var err error
	switch c.String("env") {
	case "gcp":
		//app, cleanup, err = setupGCP(ctx, cf)
	case "aws":
		//app, cleanup, err = setupAWS(ctx, cf)
	case "local":
		app, cleanup, err = setupLocal(c)
	default:
		log.Fatalf("unknown -env=%s", c.String("env"))
	}
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	{ // register rpc services

		todoCollection := app.db.DB("alex").C("todo")

		// todo service
		todoService := &service.TodoService{Model: &model.TodoModel{Collection: todoCollection}}
		todo.RegisterTodoServiceServer(app.grpcServer, todoService)

		// ... another services ...
	}

	initPrometheus(c)

	log.Println("Starting Grpc service... " + grpcAddr(c))
	lis, err := net.Listen("tcp", grpcAddr(c))
	if err != nil {
		app.logger.Fatalf("Failed to listen: %v", grpcAddr(c))
	}

	go func() {
		reflection.Register(app.grpcServer)
		err := app.grpcServer.Serve(lis)
		if err != nil {
			app.logger.Fatalf(err.Error())
		}
	}()

	mux := grpc_runtime.NewServeMux()
	{
		// create grpc client, http gateway will use it
		conn, err := grpc.Dial(grpcAddr(c), grpc.WithInsecure())
		if err != nil {
			app.logger.Fatalf("Couldn't contact grpc server: " + err.Error())
		}

		err = todo.RegisterTodoServiceHandler(context.Background(), mux, conn)
		if err != nil {
			app.logger.Fatalf("Cannot serve http api, " + err.Error())
		}
	}

	grpc_prometheus.Register(app.grpcServer)
	log.Println("Starting HTTP service... " + httpAddr(c))
	http.ListenAndServe(httpAddr(c), mux)
}

func grpcAddr(c *cli.Context) string {
	return "127.0.0.1:" + c.String("bind-grpc")
}

func httpAddr(c *cli.Context) string {
	return "127.0.0.1:" + c.String("bind-http")
}

func mongoAddr(ctx *cli.Context) string {
	return ctx.String("db-host") + ":" + ctx.String("db-port")
}

func jaegerAddr(c *cli.Context) string {
	return c.String("jaeger-host") + ":" + c.String("jaeger-port")
}

//func initTracer(c *cli.Context, logger *log.Entry) (opentracing.Tracer, io.Closer, error) {
//	// Prometheus monitoring
//	metrics := prometheus_metrics.New()
//
//	// Jaeger tracing
//	cfg := config.Configuration{
//		ServiceName: "todo",
//		Sampler: &config.SamplerConfig{
//			Type:  "const",
//			Param: c.Float64("jaeger-sampler"),
//		},
//		Reporter: &config.ReporterConfig{
//			LocalAgentHostPort: jaegerAddr(c),
//		},
//	}
//	return cfg.NewTracer(
//		config.Logger(jaegerLoggerAdapter{logger}),
//		config.Observer(rpcmetrics.NewObserver(metrics.Namespace("todo", nil), rpcmetrics.DefaultNameNormalizer)),
//	)
//}

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
