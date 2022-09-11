package server

import (
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/AI1411/go-pg-ci-example/db"
	"github.com/AI1411/go-pg-ci-example/env"
	pb "github.com/AI1411/go-pg-ci-example/grpc"
	"github.com/AI1411/go-pg-ci-example/internal/infra/repository"
)

func Handler(e *env.Env, zapLogger *zap.Logger) {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	dbClient, err := db.NewClient(e, zapLogger)
	testRepo := repository.NewTestRepository(dbClient)
	pb.RegisterTestServiceServer(s, NewTestServer(testRepo))

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
