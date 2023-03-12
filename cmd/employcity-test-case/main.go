package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	articlegrpc "github.com/angver/employcitytestcase/internal/api/grpc"
	articlev1 "github.com/angver/employcitytestcase/internal/api/grpc/gen/employcity/microservice/article/v1"
	"github.com/angver/employcitytestcase/internal/inmemory"
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

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		err := startGRPCServer(ctx, cfg.GrpcListen)
		if err != nil {
			log.Println(fmt.Errorf("can't start gRPC server or server return error while working: %w", err))
		}
	}()

	wg.Wait()
}

func startGRPCServer(
	ctx context.Context,
	listen string,
) error {
	log.Println("gRPC started", listen)

	lis, err := net.Listen("tcp", listen)
	if err != nil {
		return fmt.Errorf("failed to listen GRPC server: %w", err)
	}

	var opts []grpc.ServerOption
	s := grpc.NewServer(opts...)
	articlev1.RegisterArticleAPIServer(s, articlegrpc.NewServerTestCase(
		inmemory.NewArticleStorage(),
		articlegrpc.NewArticleToPbMapper(),
	))

	reflection.Register(s)

	go func() {
		<-ctx.Done()
		s.GracefulStop()
	}()

	return s.Serve(lis)
}
