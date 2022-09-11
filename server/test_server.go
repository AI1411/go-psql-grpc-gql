package server

import (
	"context"
	"log"

	pb "github.com/AI1411/go-pg-ci-example/grpc"
	"github.com/AI1411/go-pg-ci-example/internal/infra/repository"
)

type testServer struct {
	pb.UnimplementedTestServiceServer
	r *repository.TestRepository
}

func NewTestServer(r *repository.TestRepository) *testServer {
	return &testServer{
		r: r,
	}
}

func (s *testServer) ListTest(ctx context.Context, in *pb.ListTestRequest) (*pb.ListTestResponse, error) {
	res, err := s.r.ListTest(ctx, in)
	if err != nil {
		return nil, err
	}
	return &pb.ListTestResponse{
		Tests: res,
	}, nil
}

func (s *testServer) GetTest(ctx context.Context, in *pb.GetTestRequest) (*pb.GetTestResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.GetTestResponse{
		Id: in.Id,
	}, nil
}

func (s *testServer) CreateTest(ctx context.Context, in *pb.CreateTestRequest) (*pb.CreateTestResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.CreateTestResponse{
		Name: in.Name,
	}, nil
}

func (s *testServer) UpdateTest(ctx context.Context, in *pb.UpdateTestRequest) (*pb.UpdateTestResponse, error) {
	return &pb.UpdateTestResponse{
		Name: in.Name,
	}, nil
}

func (s *testServer) DeleteTest(ctx context.Context, in *pb.DeleteTestRequest) (*pb.DeleteTestResponse, error) {
	return &pb.DeleteTestResponse{
		Id: in.Id,
	}, nil
}