package graph_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AI1411/go-psql_grpc_gql/db"
)

var tasksResolverTestCases = []struct {
	id    int
	name  string
	query string
	want  map[string]interface{}
	setup func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "正常系",
		query: `query ListTasks {
			  tasks(input: {
				title: "title"
			  }) {
				id
				title
				description
				dueDate
				completed
				user_id
				status
			  }
			}
		`,
		want: map[string]interface{}{
			"tasks": []interface{}{
				map[string]interface{}{
					"id":          float64(1),
					"title":       "title",
					"description": "test",
					"dueDate":     "2022-09-10T08:47:22Z",
					"completed":   false,
					"user_id":     float64(1),
					"status":      "waiting",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('title','test','2022-09-10 08:47:22',false,1, 'waiting','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('task','desc','2022-09-22 08:47:22',true,2, 'done','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
		},
	},
}

func TestQueryResolver_Tasks(t *testing.T) {
	ctx := context.Background()
	srv, client := NewGqlServer()
	initDBForTests(context.Background(), t, client)

	for _, tt := range tasksResolverTestCases {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			if tt.setup != nil {
				tt.setup(ctx, t, client)
			}

			var resp map[string]interface{}
			srv.MustPost(tt.query, &resp)
			_, exist := resp["tasks"]
			assert.True(t, exist)

			if tt.want != nil {
				assert.Equal(t, tt.want, resp)
			}
		})
	}

}
