package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"

	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/env"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/logger"
	"github.com/AI1411/go-psql_grpc_gql/server"
)

func main() {
	// get env
	if err := godotenv.Load("env/.env"); err != nil {
		panic("Error loading .env file")
	}
	e := &env.Env{
		Hostname: os.Getenv("GOPG_HOSTNAME"),
		Port:     os.Getenv("GOPG_PORT"),
		User:     os.Getenv("GOPG_USERNAME"),
		Password: os.Getenv("GOPG_PASSWORD"),
		Dbname:   os.Getenv("GOPG_DATABASE"),
	}
	zapLogger, _ := logger.NewLogger(true)
	client, _ := db.NewClient(e, zapLogger)

	client.Conn(context.Background()).Exec(`SELECT * FROM public.tests;`)

	server.Handler(e, zapLogger)
}
