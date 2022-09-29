package repository

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/AI1411/go-psql_grpc_gql/config"
	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/logger"
)

func initDBForTests(ctx context.Context, t *testing.T, client *db.Client) {
	require.NoError(t, client.Conn(ctx).Exec(`TRUNCATE TABLE public.tests RESTART IDENTITY;`).Error)
	require.NoError(t, client.Conn(ctx).Exec(`TRUNCATE TABLE public.users RESTART IDENTITY;`).Error)
	require.NoError(t, client.Conn(ctx).Exec(`TRUNCATE TABLE public.tasks RESTART IDENTITY;`).Error)
	require.NoError(t, client.Conn(ctx).Exec(`TRUNCATE TABLE public.products RESTART IDENTITY;`).Error)
}

func initializeForRepositoryTest(t *testing.T) (context.Context, *db.Client) {
	configPath := "../../../config/config"
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}
	zapLogger, _ := logger.NewLogger(false)
	client, err := db.NewClient(cfg, zapLogger)
	require.NoError(t, err)

	return context.Background(), client
}
