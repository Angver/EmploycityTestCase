package grpc

import (
	"context"

	testcasev1 "github.com/angver/employcitytestcase/internal/api/grpc/gen/employcity/microservice/testcase/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewServerTestCase() *ServerTestCase {
	return &ServerTestCase{}
}

type ServerTestCase struct {
	testcasev1.UnimplementedTestCaseAPIServer
}

func (c ServerTestCase) Get(ctx context.Context, r *testcasev1.GetRequest) (*testcasev1.GetResponse, error) {
	_ = ctx
	_ = r

	return nil, status.Error(codes.Unimplemented, "method is unimplemented, try it later")
}

func (c ServerTestCase) Set(ctx context.Context, r *testcasev1.SetRequest) (*testcasev1.SetResponse, error) {
	_ = ctx
	_ = r

	return nil, status.Error(codes.Unimplemented, "method is unimplemented, try it later")
}

func (c ServerTestCase) Delete(ctx context.Context, r *testcasev1.DeleteRequest) (*testcasev1.DeleteResponse, error) {
	_ = ctx
	_ = r

	return nil, status.Error(codes.Unimplemented, "method is unimplemented, try it later")
}
