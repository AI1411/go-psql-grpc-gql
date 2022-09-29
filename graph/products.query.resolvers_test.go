package graph_test

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AI1411/go-psql_grpc_gql/db"
)

func TestQueryResolver_Tasks(t *testing.T) {
	ctx := context.Background()
	srv, client := NewGqlServer()

	var tasksResolverTestCases = []struct {
		id      int
		name    string
		query   string
		want    map[string]interface{}
		wantErr error
		setup   func(ctx context.Context, t *testing.T, client *db.Client)
	}{
		{
			id:   1,
			name: "正常系 タスク一覧を取得できる title検索",
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
		{
			id:   2,
			name: "正常系 タスク一覧を取得できる completed 検索",
			query: `query ListTasks {
			  tasks(input: {
				completed: false
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
		{
			id:   3,
			name: "正常系 タスク一覧を取得できる userId 検索",
			query: `query ListTasks {
			  tasks(input: {
				user_id: 2
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
						"id":          float64(2),
						"title":       "task",
						"description": "desc",
						"dueDate":     "2022-09-22T08:47:22Z",
						"completed":   true,
						"user_id":     float64(2),
						"status":      "done",
					},
				},
			},
			setup: func(ctx context.Context, t *testing.T, client *db.Client) {
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('title','test','2022-09-10 08:47:22',false,1, 'waiting','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('task','desc','2022-09-22 08:47:22',true,2, 'done','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			},
		},
		{
			id:   4,
			name: "正常系 タスク一覧を取得できる status 検索",
			query: `query ListTasks {
			  tasks(input: {
				status: "done"
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
						"id":          float64(2),
						"title":       "task",
						"description": "desc",
						"dueDate":     "2022-09-22T08:47:22Z",
						"completed":   true,
						"user_id":     float64(2),
						"status":      "done",
					},
				},
			},
			setup: func(ctx context.Context, t *testing.T, client *db.Client) {
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('title','test','2022-09-10 08:47:22',false,1, 'waiting','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('task','desc','2022-09-22 08:47:22',true,2, 'done','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			},
		},
		{
			id:   5,
			name: "異常系 status 検索 不正な文字列を入力した場合エラーになること",
			query: `query ListTasks {
			  tasks(input: {
				status: 'done'
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
			wantErr: errors.New("http 422: {\"errors\":[{\"message\":\"Unexpected \\u003cInvalid\\u003e\",\"locations\":[{\"line\":3,\"column\":13}],\"extensions\":{\"code\":\"GRAPHQL_PARSE_FAILED\"}}],\"data\":null}"),
			setup: func(ctx context.Context, t *testing.T, client *db.Client) {
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('title','test','2022-09-10 08:47:22',false,1, 'waiting','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('task','desc','2022-09-22 08:47:22',true,2, 'done','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			},
		},
	}

	for _, tt := range tasksResolverTestCases {
		tt := tt

		//tgtIds := []int{1}
		//if !helper.Contains(tgtIds, tt.id) {
		//	continue
		//}

		t.Run(tt.name, func(t *testing.T) {
			initDBForTests(ctx, t, client)
			if tt.setup != nil {
				tt.setup(ctx, t, client)
			}

			var resp map[string]interface{}
			err := srv.Post(tt.query, &resp)

			if tt.wantErr != nil {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}

			if tt.want != nil {
				assert.Equal(t, tt.want, resp)
			}
		})
	}

}
