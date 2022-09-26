package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/grpc"
	"github.com/AI1411/go-psql_grpc_gql/internal/helper"
)

var listTaskTestcases = []struct {
	id    int
	name  string
	in    *grpc.ListTasksRequest
	want  *grpc.ListTasksResponse
	setup func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "タスク一覧正常系<TID:1>",
		in:   &grpc.ListTasksRequest{},
		want: &grpc.ListTasksResponse{
			Tasks: []*grpc.Task{
				{
					Id:          1,
					Title:       "title",
					Description: "test",
					DueDate:     "2022-09-10T08:47:22Z",
					Completed:   false,
					UserId:      1,
					Status:      "waiting",
					CreatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
				},
				{
					Id:          2,
					Title:       "task",
					Description: "desc",
					DueDate:     "2022-09-22T08:47:22Z",
					Completed:   true,
					UserId:      2,
					Status:      "done",
					CreatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
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
		name: "タスク一覧正常系 Title検索<TID:2>",
		in: &grpc.ListTasksRequest{
			Title: "title",
		},
		want: &grpc.ListTasksResponse{
			Tasks: []*grpc.Task{
				{
					Id:          1,
					Title:       "title",
					Description: "test",
					DueDate:     "2022-09-10T08:47:22Z",
					Completed:   false,
					UserId:      1,
					Status:      "waiting",
					CreatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
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
		name: "タスク一覧正常系 Title検索 部分一致が機能していること<TID:3>",
		in: &grpc.ListTasksRequest{
			Title: "ti",
		},
		want: &grpc.ListTasksResponse{
			Tasks: []*grpc.Task{
				{
					Id:          1,
					Title:       "title",
					Description: "test",
					DueDate:     "2022-09-10T08:47:22Z",
					Completed:   false,
					UserId:      1,
					Status:      "waiting",
					CreatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
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
		name: "タスク一覧正常系 due_date_from 検索<TID:4>",
		in: &grpc.ListTasksRequest{
			DueDateFrom: "2022-09-12 08:47:22.182",
		},
		want: &grpc.ListTasksResponse{
			Tasks: []*grpc.Task{
				{
					Id:          2,
					Title:       "task",
					Description: "desc",
					DueDate:     "2022-09-22T08:47:22Z",
					Completed:   true,
					UserId:      2,
					Status:      "done",
					CreatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
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
		name: "タスク一覧正常系 due_date_to 範囲検索<TID:5>",
		in: &grpc.ListTasksRequest{
			DueDateTo: "2022-09-16 08:47:22.182",
		},
		want: &grpc.ListTasksResponse{
			Tasks: []*grpc.Task{
				{
					Id:          1,
					Title:       "title",
					Description: "test",
					DueDate:     "2022-09-10T08:47:22Z",
					Completed:   false,
					UserId:      1,
					Status:      "waiting",
					CreatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('title','test','2022-09-10 08:47:22',false,1, 'waiting','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('task','desc','2022-09-22 08:47:22',true,2, 'done','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
		},
	},
	{
		id:   6,
		name: "タスク一覧正常系 completed 検索<TID:6>",
		in: &grpc.ListTasksRequest{
			Completed: helper.BoolToPtr(false),
		},
		want: &grpc.ListTasksResponse{
			Tasks: []*grpc.Task{
				{
					Id:          1,
					Title:       "title",
					Description: "test",
					DueDate:     "2022-09-10T08:47:22Z",
					Completed:   false,
					UserId:      1,
					Status:      "waiting",
					CreatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('title','test','2022-09-10 08:47:22',false,1, 'waiting','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('task','desc','2022-09-22 08:47:22',true,2, 'done','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
		},
	},
	{
		id:   7,
		name: "タスク一覧正常系 user_id 検索<TID:7>",
		in: &grpc.ListTasksRequest{
			UserId: helper.Uint32ToPtr(2),
		},
		want: &grpc.ListTasksResponse{
			Tasks: []*grpc.Task{
				{
					Id:          2,
					Title:       "task",
					Description: "desc",
					DueDate:     "2022-09-22T08:47:22Z",
					Completed:   true,
					UserId:      2,
					Status:      "done",
					CreatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('title','test','2022-09-10 08:47:22',false,1, 'waiting','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('task','desc','2022-09-22 08:47:22',true,2, 'done','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
		},
	},
	{
		id:   8,
		name: "タスク一覧正常系 status 検索<TID:8>",
		in: &grpc.ListTasksRequest{
			Status: "done",
		},
		want: &grpc.ListTasksResponse{
			Tasks: []*grpc.Task{
				{
					Id:          2,
					Title:       "task",
					Description: "desc",
					DueDate:     "2022-09-22T08:47:22Z",
					Completed:   true,
					UserId:      2,
					Status:      "done",
					CreatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('title','test','2022-09-10 08:47:22',false,1, 'waiting','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('task','desc','2022-09-22 08:47:22',true,2, 'done','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
		},
	},
	{
		id:   9,
		name: "タスク一覧正常系 created_at_from 検索<TID:9>",
		in: &grpc.ListTasksRequest{
			CreatedAtFrom: "2022-09-18 08:47:22.182",
		},
		want: &grpc.ListTasksResponse{
			Tasks: []*grpc.Task{
				{
					Id:          2,
					Title:       "task",
					Description: "desc",
					DueDate:     "2022-09-22T08:47:22Z",
					Completed:   true,
					UserId:      2,
					Status:      "done",
					CreatedAt:   "2022-09-22 08:47:22.182 +0000 UTC",
					UpdatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('title','test','2022-09-10 08:47:22',false,1, 'waiting','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('task','desc','2022-09-22 08:47:22',true,2, 'done','2022-09-22 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
		},
	},
	{
		id:   10,
		name: "タスク一覧正常系 created_at_to 範囲検索<TID:10>",
		in: &grpc.ListTasksRequest{
			CreatedAtTo: "2022-09-16 08:47:22.182",
		},
		want: &grpc.ListTasksResponse{
			Tasks: []*grpc.Task{
				{
					Id:          1,
					Title:       "title",
					Description: "test",
					DueDate:     "2022-09-10T08:47:22Z",
					Completed:   false,
					UserId:      1,
					Status:      "waiting",
					CreatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt:   "2022-09-16 08:47:22.182 +0000 UTC",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('title','test','2022-09-10 08:47:22',false,1, 'waiting','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tasks ("title","description","due_date","completed","user_id", "status", "created_at", "updated_at") VALUES ('task','desc','2022-09-22 08:47:22',true,2, 'done','2022-09-22 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
		},
	},
}

func TestListTask(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	for _, tt := range listTaskTestcases {
		tt := tt

		//tgtIds := []int{1}
		//if !helper.Contains(tgtIds, tt.id) {
		//	continue
		//}

		t.Run(
			tt.name, func(t *testing.T) {
				initDBForTests(context.Background(), t, client)
				if tt.setup != nil {
					tt.setup(ctx, t, client)
				}

				repo := NewTaskRepository(client)
				got, err := repo.ListTasks(ctx, tt.in)
				require.NoError(t, err)

				assert.Equal(t, tt.want, got)
			},
		)

	}
}

func TestAllTaskTest(t *testing.T) {
	TestListTask(t)
}
