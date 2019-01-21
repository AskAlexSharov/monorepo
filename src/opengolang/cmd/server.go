package cmd

import (
	"context"
	"github.com/AskAlexSharov/monorepo/src/todo/api/todo"
	"github.com/AskAlexSharov/monorepo/src/todo/model"
	"github.com/AskAlexSharov/monorepo/src/todo/service"
	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

// ServerCommand with command line flags and env
type ServerCommand struct {
	Cache  CacheGroup  `group:"cache" namespace:"cache" env-namespace:"CACHE"`
	DGraph DGraphGroup `group:"dgraph" namespace:"dgraph" env-namespace:"DGRAPH"`
	Grpc   GrpcGroup   `group:"grpc" namespace:"grpc" env-namespace:"GRPC"`

	CommonOpts
}

// StoreGroup defines options group for store params
type GrpcGroup struct {
	Url string `long:"url" env:"TYPE" description:"host and port to grpc bind" choice:"bolt" choice:"mongo" default:"127.0.0.1:2339"`
}

// CacheGroup defines options group for cache params
type CacheGroup struct {
	Type string `long:"type" env:"TYPE" description:"type of cache" choice:"mem" choice:"mongo" choice:"none" default:"mem"`
	Max  struct {
		Items int   `long:"items" env:"ITEMS" default:"1000" description:"max cached items"`
		Value int   `long:"value" env:"VALUE" default:"65536" description:"max size of cached value"`
		Size  int64 `long:"size" env:"SIZE" default:"50000000" description:"max size of total cache"`
	} `group:"max" namespace:"max" env-namespace:"MAX"`
}

// DGraphGroup holds all mongo params, used by store, avatar and cache
type DGraphGroup struct {
	URL string `long:"url" env:"URL" description:"dgraph url"`
	DB  string `long:"db" env:"DB" default:"todo" description:"dgraph database"`
}

// serverApp holds all active objects
type serverApp struct {
	*ServerCommand
	grpcPortListener net.Listener
	grpcServer       *grpc.Server
	db               *dgo.Dgraph

	terminated chan struct{}
}

// Execute is the entry point for "server" command, called by flag parser
func (s *ServerCommand) Execute(args []string) error {

	//log.Printf("[INFO] start server on port %d", s.Port)
	resetEnv("SECRET", "AUTH_GOOGLE_CSEC", "AUTH_GITHUB_CSEC", "AUTH_FACEBOOK_CSEC", "AUTH_YANDEX_CSEC")

	ctx, cancel := context.WithCancel(context.Background())
	go func() { // catch signal and invoke graceful termination
		stop := make(chan os.Signal, 1)
		signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
		<-stop
		log.Print("[WARN] interrupt signal")
		cancel()
	}()

	app, err := s.newServerApp()
	if err != nil {
		return err
	}
	if err = app.run(ctx); err != nil {
		return err
	}
	log.Printf("[INFO] terminated")
	return nil
}

func (a *serverApp) run(ctx context.Context) error {
	go func() {
		// shutdown on context cancellation
		<-ctx.Done()
		log.Print("[INFO] shutdown initiated")
		a.grpcServer.GracefulStop()
		//if a.db != nil {
		//	err := a.db.Disconnect(ctx)
		//	if err != nil {
		//		log.Printf("[WARN] $%s\n", err)
		//	}
		//}
		log.Print("[INFO] shutdown completed")
	}()

	reflection.Register(a.grpcServer)
	go func() {
		err := a.grpcServer.Serve(a.grpcPortListener)
		if err != nil {
			log.Print(err.Error())
		}
	}()

	err := a.grpcServer.Serve(a.grpcPortListener)
	log.Printf("[WARN] http server terminated, %s", err)
	close(a.terminated)
	return nil
}

// Panic handler prints the stack trace when recovering from a panic.
var panicHandler = grpc_recovery.RecoveryHandlerFunc(func(p interface{}) error {
	buf := make([]byte, 1<<16)
	runtime.Stack(buf, true)
	log.Printf("panic recovered: %+v", string(buf))
	return status.Errorf(codes.Internal, "%s", p)
})

type fakeHealthChecker struct{}

func (*fakeHealthChecker) CheckHealth() error {
	return nil
}

// appHealthChecks returns a health check for the database. This will signal
// to Kubernetes or other orchestrators that the server should not receive
// traffic until the server is able to connect to its database.
//func appHealthChecks(db *dgo.Dgraph) ([]health.Checker, func()) {
//	//dbCheck := sqlhealth.New(db)
//	c := &fakeHealthChecker{}
//	list := []health.Checker{c}
//	return list, func() {
//		//dbCheck.Stop()
//	}
//}

func (s *ServerCommand) makeDb() (*dgo.Dgraph, error) {
	// Dial a gRPC connection. The address to dial to can be configured when
	// setting up the dgraph cluster.
	d, err := grpc.Dial(s.DGraph.URL, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	return dgo.NewDgraphClient(
		api.NewDgraphClient(d),
	), nil
}

func (s *ServerCommand) makeGrpc() *grpc.Server {
	return grpc.NewServer(
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_ctxtags.StreamServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			//grpc_opentracing.StreamServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_prometheus.StreamServerInterceptor,
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(panicHandler)),
		)),
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_ctxtags.UnaryServerInterceptor(grpc_ctxtags.WithFieldExtractor(grpc_ctxtags.CodeGenRequestFieldExtractor)),
			//grpc_opentracing.UnaryServerInterceptor(grpc_opentracing.WithTracer(tracer)),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(panicHandler)),
		)),
	)
}

// localRuntimeVar is a Wire provider function that returns the Message of the
// Day variable based on a local file.
//func (s *ServerCommand) makeRuntimeVar() (*runtimevar.Variable, func(), error) {
//	v, err := filevar.New("message_of_the_day", runtimevar.StringDecoder, &filevar.Options{
//		WaitDuration: time.Minute,
//	})
//	if err != nil {
//		return nil, nil, err
//	}
//	return v, func() { v.Close() }, nil
//}

// newServerApp prepares application and return it with all active parts
// doesn't start anything
func (s *ServerCommand) newServerApp() (*serverApp, error) {
	db, err := s.makeDb()
	if err != nil {
		return nil, errors.Wrap(err, "failed to make db")
	}

	grpcServer := s.makeGrpc()

	{ // register rpc services

		//todoCollection := db.Database(s.Mongo.DB).Collection("todo")

		// todo service
		todoService := &service.TodoService{Model: &model.TodoModel{Db: db}}
		todo.RegisterTodoServiceServer(grpcServer, todoService)

		// ... another services ...
	}

	s.makePrometheus()
	grpc_prometheus.Register(grpcServer)

	lis, err := s.makeGrpcPortListener()
	if err != nil {
		return nil, errors.Wrap(err, "Failed to listen: "+s.Grpc.Url)
	}

	return &serverApp{
		grpcPortListener: lis,
		grpcServer:       grpcServer,
		db:               db,
	}, nil
}

func (s *ServerCommand) makePrometheus() {
	go func() {
		mux := http.NewServeMux()
		mux.Handle("/metrics", promhttp.Handler())
		http.ListenAndServe("127.0.0.1:8081", mux)
	}()
}

func (s *ServerCommand) makeGrpcPortListener() (net.Listener, error) {
	log.Println("[INFO] Starting Grpc service... " + s.Grpc.Url)
	return net.Listen("tcp", s.Grpc.Url)

}
