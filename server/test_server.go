package server

import (
	"context"

	pb "github.com/AI1411/go-psql_grpc_gql/grpc/test"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/repository"
)

type TestServer struct {
	pb.UnimplementedTestServiceServer
	r *repository.TestRepository
}

func NewTestServer(r *repository.TestRepository) *TestServer {
	return &TestServer{
		r: r,
	}
}

func (s *TestServer) ListTest(ctx context.Context, in *pb.ListTestRequest) (*pb.ListTestResponse, error) {
	res, err := s.r.ListTest(ctx, in)
	if err != nil {
		return nil, err
	}
	return &pb.ListTestResponse{
		Tests: res,
	}, nil
}

func (s *TestServer) GetTest(ctx context.Context, in *pb.GetTestRequest) (*pb.GetTestResponse, error) {
	test, err := s.r.GetTest(ctx, in)
	if err != nil {
		return nil, err
	}
	res := &pb.GetTestResponse{
		Id:   test.Id,
		Name: test.Name,
	}
	return res, nil
}

func (s *TestServer) CreateTest(ctx context.Context, in *pb.CreateTestRequest) (*pb.CreateTestResponse, error) {
	test, err := s.r.CreateTest(ctx, in)
	if err != nil {
		return nil, err
	}
	return &pb.CreateTestResponse{
		Id:   test.Id,
		Name: test.Name,
	}, nil
}

func (s *TestServer) UpdateTest(ctx context.Context, in *pb.UpdateTestRequest) (*pb.UpdateTestResponse, error) {
	return &pb.UpdateTestResponse{
		Name: in.Name,
	}, nil
}

func (s *TestServer) DeleteTest(ctx context.Context, in *pb.DeleteTestRequest) (*pb.DeleteTestResponse, error) {
	return &pb.DeleteTestResponse{
		Id: in.Id,
	}, nil
}
