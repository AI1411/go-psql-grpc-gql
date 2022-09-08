package main

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/require"

	"github.com/AI1411/go-pg-ci-example/db"
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
	client, err := db.NewClient()
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
