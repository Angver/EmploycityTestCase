package main

// Config конфигурация приложения
type Config struct {
	GrpcListen string `long:"grpc-listen" description:"Listening :port for grpc-server" env:"GRPC_LISTEN" required:"true"`
}
