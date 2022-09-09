package main

import (
	"context"
	"os"

	"github.com/joho/godotenv"

	"github.com/AI1411/go-pg-ci-example/db"
	"github.com/AI1411/go-pg-ci-example/env"
	"github.com/AI1411/go-pg-ci-example/server"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}
	e := &env.Env{
		Hostname: os.Getenv("GOPG_HOSTNAME"),
		Port:     os.Getenv("GOPG_PORT"),
		User:     os.Getenv("GOPG_USERNAME"),
		Password: os.Getenv("GOPG_PASSWORD"),
		Dbname:   os.Getenv("GOPG_DATABASE"),
	}
	client, err := db.NewClient(e)
	if err != nil {
		panic(err)
	}
	client.Conn(context.Background()).Exec(`SELECT * FROM public.tests;`)

	server.Handler()
}
