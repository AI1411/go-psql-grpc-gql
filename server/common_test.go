package server

import (
	"os"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"

	"github.com/AI1411/go-pg-ci-example/db"
	"github.com/AI1411/go-pg-ci-example/env"
	pb "github.com/AI1411/go-pg-ci-example/grpc"
	"github.com/AI1411/go-pg-ci-example/internal/infra/logger"
	"github.com/AI1411/go-pg-ci-example/internal/infra/repository"
)

func initializeForServerTest() {
	lis = bufconn.Listen(BufSize)
	if err := godotenv.Load(".env.testing"); err != nil {
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
	s := grpc.NewServer()
	pb.RegisterTestServiceServer(s, &server{
		r: repository.NewTestRepository(client),
	})
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
}
