package server

import (
	"log"
	"net"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/env"
	pb "github.com/AI1411/go-psql_grpc_gql/grpc"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/repository"
	"github.com/AI1411/go-psql_grpc_gql/internal/interceptor"
)

func Handler(e *env.Env, zapLogger *zap.Logger) {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptor.ZapLoggerInterceptor(),
		grpc_auth.UnaryServerInterceptor(interceptor.AuthFunc),
	))
	dbClient, err := db.NewClient(e, zapLogger)
	testRepo := repository.NewTestRepository(dbClient)
	userRepo := repository.NewUserRepository(dbClient)
	pb.RegisterTestServiceServer(s, NewTestServer(testRepo))
	pb.RegisterUserServiceServer(s, NewUserServer(userRepo))

	if err := s.Serve(lis); err != nil {

		log.Fatalf("failed to serve: %v", err)
	}
}
