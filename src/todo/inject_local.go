//+build wireinject

package main

import (
	"time"

	"github.com/globalsign/mgo"
	"github.com/google/go-cloud/requestlog"
	"github.com/google/go-cloud/runtimevar"
	"github.com/google/go-cloud/runtimevar/filevar"
	"github.com/google/go-cloud/server"
	"github.com/google/go-cloud/wire"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"github.com/urfave/cli"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
)

func setupLocal(ctx *cli.Context) (*application, func(), error) {
	panic(wire.Build(
		wire.InterfaceValue(new(requestlog.Logger), requestlog.Logger(nil)),
		wire.InterfaceValue(new(trace.Exporter), trace.Exporter(nil)),
		server.Set,
		applicationSet,
		localGrpc,
		localRuntimeVar,
		localDb,
	))
}

func localDb(ctx *cli.Context) (*mgo.Session, error) {
	addr := mongoAddr(ctx)
	sess, err := mgo.Dial(addr)
	if err != nil {
		return nil, errors.New("Can't connect to Mongo on address: " + addr + ", by reason: " + err.Error())
	}
	return sess, nil
}

func localGrpc(ctx *cli.Context) *grpc.Server {
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
func localRuntimeVar(ctx *cli.Context) (*runtimevar.Variable, func(), error) {
	v, err := filevar.New("message_of_the_day", runtimevar.StringDecoder, &filevar.Options{
		WaitDuration: time.Minute,
	})
	if err != nil {
		return nil, nil, err
	}
	return v, func() { v.Close() }, nil
}
