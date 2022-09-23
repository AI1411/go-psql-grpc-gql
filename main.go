package main

import (
	"log"

	"github.com/AI1411/go-psql_grpc_gql/config"
	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/logger"
	"github.com/AI1411/go-psql_grpc_gql/pkg/redis"
	"github.com/AI1411/go-psql_grpc_gql/server"
)

func main() {
	configPath := "./config/config"
	cfg, err := config.GetConfig(configPath)
	log.Printf("config: %+v", cfg)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}

	zapLogger, _ := logger.NewLogger(true)
	client, _ := db.NewClient(cfg, zapLogger)
	redisClient := redis.NewRedisClient(cfg)
	defer redisClient.Close()

	s := server.NewGPServer(zapLogger, cfg, client, redisClient)

	s.Handler()
}
