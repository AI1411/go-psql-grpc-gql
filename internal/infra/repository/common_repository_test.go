package repository

import (
	"context"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"

	"github.com/AI1411/go-pg-ci-example/db"
	"github.com/AI1411/go-pg-ci-example/env"
	"github.com/AI1411/go-pg-ci-example/internal/infra/logger"
)

func initDBForTests(ctx context.Context, t *testing.T, client *db.Client) {
	require.NoError(t, client.Conn(ctx).Exec(`TRUNCATE TABLE public.tests RESTART IDENTITY;`).Error)
}

func initializeForRepositoryTest(t *testing.T) (context.Context, *db.Client) {
	if err := godotenv.Load("../../../env/.env.testing"); err != nil {
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
	client, err := db.NewClient(e, zapLogger)
	require.NoError(t, err)

	return context.Background(), client
}