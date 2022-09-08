package main

import (
	"context"

	"github.com/AI1411/go-pg-ci-example/db"
)

func main() {
	client, err := db.NewClient()
	if err != nil {
		panic(err)
	}
	client.Conn(context.Background()).Exec(`SELECT * FROM public.tests;`)
}
