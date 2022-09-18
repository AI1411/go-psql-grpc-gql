package repository

import (
	"context"
	"log"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"

	"github.com/AI1411/go-psql_grpc_gql/config"
	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/logger"
)

func initDBForTests(ctx context.Context, t *testing.T, client *db.Client) {
	require.NoError(t, client.Conn(ctx).Exec(`TRUNCATE TABLE public.tests RESTART IDENTITY;`).Error)
	require.NoError(t, client.Conn(ctx).Exec(`TRUNCATE TABLE public.users RESTART IDENTITY;`).Error)
}

func initializeForRepositoryTest(t *testing.T) (context.Context, *db.Client) {
	configPath := "../../../config/config"
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}
	if err := godotenv.Load("../../../env/.env.testing"); err != nil {
		panic("Error loading .env file")
	}

	zapLogger, _ := logger.NewLogger(false)
	client, err := db.NewClient(cfg, zapLogger)
	require.NoError(t, err)

	return context.Background(), client
}
