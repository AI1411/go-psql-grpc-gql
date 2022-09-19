package server

import (
	"log"
	"net"

	"github.com/go-redis/redis/v8"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/AI1411/go-psql_grpc_gql/config"
	"github.com/AI1411/go-psql_grpc_gql/db"
	pb "github.com/AI1411/go-psql_grpc_gql/grpc"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/repository"
	"github.com/AI1411/go-psql_grpc_gql/internal/interceptor"
)

type Server struct {
	zapLogger   *zap.Logger
	cfg         *config.Config
	dbClient    *db.Client
	redisClient *redis.Client
}

func NewGPServer(zapLogger *zap.Logger, cfg *config.Config, dbClient *db.Client, redisClient *redis.Client) *Server {
	return &Server{zapLogger: zapLogger, cfg: cfg, dbClient: dbClient, redisClient: redisClient}
}

func (s *Server) Handler() {
	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptor.ZapLoggerInterceptor(),
		grpc_auth.UnaryServerInterceptor(interceptor.AuthFunc),
	))
	dbClient, err := db.NewClient(s.cfg, s.zapLogger)
	testRepo := repository.NewTestRepository(dbClient)
	userRepo := repository.NewUserRepository(dbClient)
	pb.RegisterTestServiceServer(server, NewTestServer(testRepo))
	pb.RegisterUserServiceServer(server, NewUserServer(userRepo))

	if err := server.Serve(lis); err != nil {

		log.Fatalf("failed to serve: %v", err)
	}
}
