package server

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/AI1411/go-pg-ci-example/grpc"
)

type server struct {
	pb.UnimplementedTestServiceServer
}

func (s *server) ListTest(ctx context.Context, in *pb.ListTestRequest) (*pb.ListTestResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.ListTestResponse{
		Tests: []*pb.GetTestResponse{
			{
				Id:   1,
				Name: in.Name,
			},
		},
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

func Handler() {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterTestServiceServer(s, &server{})

	log.Printf("Listening on %s", addr)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
