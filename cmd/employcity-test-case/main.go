package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"

	testcase_grpc "github.com/angver/employcitytestcase/internal/api/grpc"
	testcasev1 "github.com/angver/employcitytestcase/internal/api/grpc/gen/employcity/microservice/testcase/v1"
	"github.com/jessevdk/go-flags"
	"google.golang.org/grpc"
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
	testcasev1.RegisterTestCaseAPIServer(s, testcase_grpc.NewServerTestCase())

	go func() {
		<-ctx.Done()
		s.GracefulStop()
	}()

	return s.Serve(lis)
}
