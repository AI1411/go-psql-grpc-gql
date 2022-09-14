package main

import (
	"context"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/require"

	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/env"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/logger"
)

type test struct {
	ID   int
	Name string
}

func initDBForTests(ctx context.Context, t *testing.T, client *db.Client) {
	require.NoError(t, client.Conn(ctx).Exec(`TRUNCATE TABLE public.tests RESTART IDENTITY;`).Error)
}

var exampleTestCases = []struct {
	id    int
	name  string
	want  []test
	setup func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "テストの例<TID:1>",
		want: []test{
			{
				ID:   1,
				Name: "test1",
			},
			{
				ID:   2,
				Name: "test2",
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test1');`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test2');`).Error)
		},
	},
}

func TestExample(t *testing.T) {
	if err := godotenv.Load("env/.env.testing"); err != nil {
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
	ctx := context.Background()
	for _, tt := range exampleTestCases {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				initDBForTests(context.Background(), t, client)
				if tt.setup != nil {
					tt.setup(ctx, t, client)
				}

				var got []test
				client.Conn(ctx).Find(&got)
				if !cmp.Equal(got, tt.want) {
					t.Errorf("diff %s", cmp.Diff(got, tt.want))
				}
			},
		)

	}
}
