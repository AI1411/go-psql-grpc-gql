package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AI1411/go-pg-ci-example/db"
	"github.com/AI1411/go-pg-ci-example/grpc"
)

var listTestTestcases = []struct {
	id    int
	name  string
	want  []*grpc.GetTestResponse
	setup func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "テストの例<TID:1>",
		want: []*grpc.GetTestResponse{
			{
				Id:   1,
				Name: "test1",
			},
			{
				Id:   2,
				Name: "test2",
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test1');`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test2');`).Error)
		},
	},
}

func TestListTest(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	for _, tt := range listTestTestcases {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				initDBForTests(context.Background(), t, client)
				if tt.setup != nil {
					tt.setup(ctx, t, client)
				}

				repo := NewTestRepository(client)
				in := &grpc.ListTestRequest{}
				got, err := repo.ListTest(ctx, in)
				require.NoError(t, err)

				assert.Equal(t, tt.want, got)
			},
		)

	}
}
