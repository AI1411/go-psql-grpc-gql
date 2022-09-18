package main

import (
	"context"
	"log"

	"github.com/AI1411/go-psql_grpc_gql/config"
	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/logger"
	"github.com/AI1411/go-psql_grpc_gql/server"
)

func main() {
	configPath := "./config/config"
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}
	log.Printf("config: %+v", cfg)
	// get env
	zapLogger, _ := logger.NewLogger(true)
	client, _ := db.NewClient(cfg, zapLogger)

	client.Conn(context.Background()).Exec(`SELECT * FROM public.tests;`)

	server.Handler(cfg, zapLogger)
}
