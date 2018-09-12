package main

import (
	"github.com/urfave/cli"
)

var commonFlags = []cli.Flag{
	cli.UintFlag{
		Name:   "bind-http",
		Usage:  "bind address for HTTP",
		EnvVar: "BIND_HTTP",
		Value:  8080,
	},
	cli.UintFlag{
		Name:   "bind-grpc",
		Usage:  "bind address for gRPC",
		EnvVar: "BIND_GRPC",
		Value:  2339,
	},
	cli.StringFlag{
		Name:   "bind-prometheus-http",
		Usage:  "bind prometheus address for HTTP",
		EnvVar: "BIND_PROMETHEUS_HTTP",
		Value:  ":8081",
	},
	cli.StringFlag{
		Name:   "db-name",
		Usage:  "database name",
		EnvVar: "DB_NAME",
		Value:  "todo",
	},
	cli.StringFlag{
		Name:   "db-user",
		Usage:  "database username",
		EnvVar: "DB_USER",
		Value:  "postgres",
	},
	cli.StringFlag{
		Name:   "db-password",
		Usage:  "database password",
		EnvVar: "DB_PASSWORD",
	},
	cli.StringFlag{
		Name:   "db-host",
		Usage:  "postgres host",
		EnvVar: "DB_HOST",
		Value:  "127.0.0.1",
	},
	cli.UintFlag{
		Name:   "db-port",
		Usage:  "database port",
		EnvVar: "DB_PORT",
		Value:  27017,
	},
	cli.StringFlag{
		Name:   "jaeger-host",
		Usage:  "Jaeger hostname",
		EnvVar: "JAEGER_HOST",
		Value:  "127.0.0.1",
	},
	cli.UintFlag{
		Name:   "jaeger-port",
		Usage:  "Jaeger port",
		EnvVar: "JAEGER_PORT",
		Value:  5775,
	},
	cli.Float64Flag{
		Name:   "jaeger-sampler",
		Usage:  "Jaeger sampler",
		EnvVar: "JAEGER_SAMPLER",
		Value:  0.05,
	},
	cli.StringFlag{
		Name:   "jaeger-tags",
		Usage:  "Jaeger tags",
		EnvVar: "JAEGER_TAGS",
		Value:  "todo",
	},
}
