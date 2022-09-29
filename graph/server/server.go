package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/cors"

	"github.com/AI1411/go-psql_grpc_gql/config"
	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/graph"
	"github.com/AI1411/go-psql_grpc_gql/graph/generated"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/logger"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/repository"
	"github.com/AI1411/go-psql_grpc_gql/server"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	configPath := "./config/config"
	cfg, err := config.GetConfig(configPath)
	if err != nil {
		log.Fatalf("Loading config: %v", err)
	}
	zapLogger, _ := logger.NewLogger(true)
	client, _ := db.NewClient(cfg, zapLogger)
	tesRepo := repository.NewTestRepository(client)
	userRepo := repository.NewUserRepository(client)
	taskRepo := repository.NewTaskRepository(client)
	productrRepo := repository.NewProductRepository(client)
	testServer := server.NewTestServer(tesRepo)
	userServer := server.NewUserServer(userRepo)
	taskServer := server.NewTaskServer(taskRepo)
	productServer := server.NewProductServer(productrRepo)

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{
		TestServer:    testServer,
		UserServer:    userServer,
		TaskServer:    taskServer,
		ProductServer: productServer,
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:8080"},
		AllowCredentials: true,
	})

	http.Handle("/query", c.Handler(srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
