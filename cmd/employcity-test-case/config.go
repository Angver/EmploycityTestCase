package main

// Config конфигурация приложения
type Config struct {
	GrpcListen string `long:"grpc-listen" description:"Listening :port for grpc-server" env:"GRPC_LISTEN" required:"true"`
	// MCServer адрес сервера Memcached
	MCServer string `long:"mc-server" description:"Memcached Server" env:"MC_SERVER" required:"true"`
	// StorageEngine используемый движок хранилища: memcached или inmemory
	StorageEngine string `long:"storage-engine" description:"Storage engine" env:"STORAGE_ENGINE"`
}
