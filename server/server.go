package server

import (
	"context"
	"log"
	"net"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/AI1411/go-psql_grpc_gql/config"
	"github.com/AI1411/go-psql_grpc_gql/db"
	pb "github.com/AI1411/go-psql_grpc_gql/grpc"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/repository"
	"github.com/AI1411/go-psql_grpc_gql/internal/model"
)

type Server struct {
	cfg         *config.Config
	dbClient    *db.Client
	redisClient *redis.Client
}

func NewServer(cfg *config.Config, dbClient *db.Client, redisClient *redis.Client) *Server {
	return &Server{
		cfg:         cfg,
		dbClient:    dbClient,
		redisClient: redisClient,
	}
}

func (s *Server) Handler(c *config.Config, zapLogger *zap.Logger) {
	dbClient, err := db.NewClient(c, zapLogger)
	testRepo := repository.NewTestRepository(dbClient)
	userRepo := repository.NewUserRepository(dbClient)
	sessionRepo := repository.NewSessionRepository(s.redisClient, s.cfg)
	str, err := sessionRepo.CreateSession(context.Background(), &model.Session{
		ID:        1,
		SessionID: uuid.New().String(),
	}, 3600)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(str)

	addr := ":50051"
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterTestServiceServer(server, NewTestServer(testRepo))
	pb.RegisterUserServiceServer(server, NewUserServer(userRepo))

	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
