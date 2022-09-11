package server

import (
	"context"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/AI1411/go-pg-ci-example/db"
	"github.com/AI1411/go-pg-ci-example/env"
	pb "github.com/AI1411/go-pg-ci-example/grpc"
	"github.com/AI1411/go-pg-ci-example/internal/infra/repository"
)

type server struct {
	pb.UnimplementedTestServiceServer
	r *repository.TestRepository
}

func (s *server) ListTest(ctx context.Context, in *pb.ListTestRequest) (*pb.ListTestResponse, error) {
	res, err := s.r.ListTest(ctx, in)
	if err != nil {
		return nil, err
	}
	return &pb.ListTestResponse{
		Tests: res,
	}, nil
}

func (s *server) GetTest(ctx context.Context, in *pb.GetTestRequest) (*pb.GetTestResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.GetTestResponse{
		Id: in.Id,
	}, nil
}

func (s *server) CreateTest(ctx context.Context, in *pb.CreateTestRequest) (*pb.CreateTestResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.CreateTestResponse{
		Name: in.Name,
	}, nil
}

func (s *server) UpdateTest(ctx context.Context, in *pb.UpdateTestRequest) (*pb.UpdateTestResponse, error) {
	return &pb.UpdateTestResponse{
		Name: in.Name,
	}, nil
}

func (s *server) DeleteTest(ctx context.Context, in *pb.DeleteTestRequest) (*pb.DeleteTestResponse, error) {
	return &pb.DeleteTestResponse{
		Id: in.Id,
	}, nil
}

func Handler(e *env.Env, zapLogger *zap.Logger) {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	dbClient, err := db.NewClient(e, zapLogger)
	testRepo := repository.NewTestRepository(dbClient)
	pb.RegisterTestServiceServer(s, &server{
		r: testRepo,
	})

	log.Printf("Listening on %s", addr)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
