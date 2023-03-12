package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/angver/employcitytestcase/internal"
	articlegrpc "github.com/angver/employcitytestcase/internal/api/grpc"
	articlev1 "github.com/angver/employcitytestcase/internal/api/grpc/gen/employcity/microservice/article/v1"
	"github.com/angver/employcitytestcase/internal/inmemory"
	"github.com/angver/employcitytestcase/internal/memcached"
	"github.com/jessevdk/go-flags"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var cfg Config
	parser := flags.NewParser(&cfg, flags.Default)
	_, err := parser.Parse()
	if err != nil {
		log.Fatal(fmt.Errorf("failed to parse config: %w", err))
	}

	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()

	var storage internal.ArticleStorage
	switch cfg.StorageEngine {
	case "memcached":
		storage = memcached.NewArticleStorage(cfg.MCServer)
	case "inmemory":
		fallthrough
	default:
		storage = inmemory.NewArticleStorage()
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := startGRPCServer(ctx, cfg.GrpcListen, storage)
		if err != nil {
			log.Println(fmt.Errorf("can't start gRPC server or server return error while working: %w", err))
		}
	}()

	wg.Wait()
}

func startGRPCServer(
	ctx context.Context,
	listen string,
	storage internal.ArticleStorage,
) error {
	log.Println("gRPC started", listen)

	lis, err := net.Listen("tcp", listen)
	if err != nil {
		return fmt.Errorf("failed to listen GRPC server: %w", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	articlev1.RegisterArticleAPIServer(s, articlegrpc.NewServerTestCase(
		storage,
		articlegrpc.NewArticleToPbMapper(),
	))

	reflection.Register(s)

	go func() {
		<-ctx.Done()
		s.GracefulStop()
	}()

	return s.Serve(lis)
}
