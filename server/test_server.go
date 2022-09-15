package server

import (
	"context"

	"github.com/AI1411/go-psql_grpc_gql/grpc"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/repository"
)

type TestServer struct {
	grpc.UnimplementedTestServiceServer
	r *repository.TestRepository
}

func NewTestServer(r *repository.TestRepository) *TestServer {
	return &TestServer{
		r: r,
	}
}

func (s *TestServer) ListTest(ctx context.Context, in *grpc.ListTestRequest) (*grpc.ListTestResponse, error) {
	res, err := s.r.ListTest(ctx, in)
	if err != nil {
		return nil, err
	}
	return &grpc.ListTestResponse{
		Tests: res,
	}, nil
}

func (s *TestServer) GetTest(ctx context.Context, in *grpc.GetTestRequest) (*grpc.GetTestResponse, error) {
	test, err := s.r.GetTest(ctx, in)
	if err != nil {
		return nil, err
	}
	res := &grpc.GetTestResponse{
		Id:   test.Id,
		Name: test.Name,
	}
	return res, nil
}

func (s *TestServer) CreateTest(ctx context.Context, in *grpc.CreateTestRequest) (*grpc.CreateTestResponse, error) {
	test, err := s.r.CreateTest(ctx, in)
	if err != nil {
		return nil, err
	}
	return &grpc.CreateTestResponse{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) UpdateTest(ctx context.Context, in *grpc.UpdateTestRequest) (*grpc.UpdateTestResponse, error) {
	test, err := s.r.UpdateTest(ctx, in)
	if err != nil {
		return nil, err
	}

	return &grpc.UpdateTestResponse{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) DeleteTest(ctx context.Context, in *grpc.DeleteTestRequest) (*grpc.DeleteTestResponse, error) {
	test, err := s.r.DeleteTest(ctx, in)
	if err != nil {
		return nil, err
	}

	return &grpc.DeleteTestResponse{
		Id: test.Id,
	}, nil
}
