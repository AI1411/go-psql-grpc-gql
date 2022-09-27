package graph_test

import (
	"log"
	"os"
	"testing"

	gqclient "github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/stretchr/testify/assert"

	"github.com/AI1411/go-psql_grpc_gql/config"
	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/graph"
	"github.com/AI1411/go-psql_grpc_gql/graph/generated"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/logger"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/repository"
	"github.com/AI1411/go-psql_grpc_gql/server"
)

func TestQueryResolver_Tasks(t *testing.T) {
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
	tesRepo := repository.NewTestRepository(client)
	userRepo := repository.NewUserRepository(client)
	taskRepo := repository.NewTaskRepository(client)
	testServer := server.NewTestServer(tesRepo)
	userServer := server.NewUserServer(userRepo)
	taskServer := server.NewTaskServer(taskRepo)

	srv := gqclient.New(handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		TestServer: testServer,
		UserServer: userServer,
		TaskServer: taskServer,
	}})))

	q := `query ListTasks {
			  tasks(input: {
				title: "title"
				status: ""
			  }) {
				title
			  }
			}
		`

	var resp map[string]interface{}
	srv.MustPost(q, &resp)
	_, exist := resp["tasks"]
	assert.True(t, exist)
	assert.Equal(t, "title", resp["tasks"].([]interface{})[0].(map[string]interface{})["title"])
}
