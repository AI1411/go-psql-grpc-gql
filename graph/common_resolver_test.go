package graph_test

import (
	"context"
	"log"
	"os"
	"testing"

	gqclient "github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/require"

	"github.com/AI1411/go-psql_grpc_gql/config"
	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/graph"
	"github.com/AI1411/go-psql_grpc_gql/graph/generated"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/logger"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/repository"
	"github.com/AI1411/go-psql_grpc_gql/server"
)

func initDBForTests(ctx context.Context, t *testing.T, client *db.Client) {
	require.NoError(t, client.Conn(ctx).Exec(`TRUNCATE TABLE public.tests RESTART IDENTITY;`).Error)
	require.NoError(t, client.Conn(ctx).Exec(`TRUNCATE TABLE public.users RESTART IDENTITY;`).Error)
	require.NoError(t, client.Conn(ctx).Exec(`TRUNCATE TABLE public.tasks RESTART IDENTITY;`).Error)
	require.NoError(t, client.Conn(ctx).Exec(`TRUNCATE TABLE public.products RESTART IDENTITY;`).Error)
}

func NewGqlServer() (*gqclient.Client, *db.Client) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	configPath := "../config/config"
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}
	zapLogger, _ := logger.NewLogger(true)
	client, _ := db.NewClient(cfg, zapLogger)
	taskRepo := repository.NewTaskRepository(client)
	productRepo := repository.NewProductRepository(client)
	taskServer := server.NewTaskServer(taskRepo)
	productServer := server.NewProductServer(productRepo)

	srv := gqclient.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		TaskServer:    taskServer,
		ProductServer: productServer,
	}})))

	return srv, client
}
