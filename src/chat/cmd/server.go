package cmd

// ServerCommand with command line flags and env
type ServerCommand struct {
	Mongo MongoGroup `group:"mongo" namespace:"mongo" env-namespace:"MONGO"`
	Grpc  GrpcGroup  `group:"grpc" namespace:"grpc" env-namespace:"GRPC"`

	CommonOpts
}

// StoreGroup defines options group for store params
type GrpcGroup struct {
	Url string `long:"url" env:"TYPE" description:"host and port to grpc bind" choice:"bolt" choice:"mongo" default:"127.0.0.1:2339"`
}

// MongoGroup holds all mongo params, used by store, avatar and cache
type MongoGroup struct {
	URL string `long:"url" env:"URL" description:"mongo url"`
	DB  string `long:"db" env:"DB" default:"todo" description:"mongo database"`
}
