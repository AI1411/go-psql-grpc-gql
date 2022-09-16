package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/grpc"
)

var listUserTestcases = []struct {
	id    int
	name  string
	in    *grpc.ListUsersRequest
	want  *grpc.ListUsersResponse
	setup func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "ユーザ一覧正常系<TID:1>",
		in:   &grpc.ListUsersRequest{},
		want: &grpc.ListUsersResponse{
			Users: []*grpc.User{
				{
					Id:        1,
					Name:      "test",
					Email:     "test@gmail.com",
					CreatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
				},
				{
					Id:        2,
					Name:      "akira",
					Email:     "akira@gmail.com",
					CreatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('akira','akira@gmail.com','$2a$10$tyNZh8SbBwP1rvGzHWzSPeNATz/N24wTXKDc1FS53waJzFlEeTWl6','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
	{
		id:   2,
		name: "ユーザ一覧正常系 Name検索<TID:2>",
		in: &grpc.ListUsersRequest{
			Name: "test",
		},
		want: &grpc.ListUsersResponse{
			Users: []*grpc.User{
				{
					Id:        1,
					Name:      "test",
					Email:     "test@gmail.com",
					CreatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('akira','akira@gmail.com','$2a$10$tyNZh8SbBwP1rvGzHWzSPeNATz/N24wTXKDc1FS53waJzFlEeTWl6','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
	{
		id:   3,
		name: "ユーザ一覧正常系 Name検索 部分一致が機能していること<TID:3>",
		in: &grpc.ListUsersRequest{
			Name: "te",
		},
		want: &grpc.ListUsersResponse{
			Users: []*grpc.User{
				{
					Id:        1,
					Name:      "test",
					Email:     "test@gmail.com",
					CreatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('akira','akira@gmail.com','$2a$10$tyNZh8SbBwP1rvGzHWzSPeNATz/N24wTXKDc1FS53waJzFlEeTWl6','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
	{
		id:   4,
		name: "ユーザ一覧正常系 Email検索<TID:4>",
		in: &grpc.ListUsersRequest{
			Email: "akira@gmail.com",
		},
		want: &grpc.ListUsersResponse{
			Users: []*grpc.User{
				{
					Id:        2,
					Name:      "akira",
					Email:     "akira@gmail.com",
					CreatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
					UpdatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('akira','akira@gmail.com','$2a$10$tyNZh8SbBwP1rvGzHWzSPeNATz/N24wTXKDc1FS53waJzFlEeTWl6','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
	{
		id:   5,
		name: "ユーザ一覧正常系 Email検索　部分一致では検索できないこと<TID:5>",
		in: &grpc.ListUsersRequest{
			Email: "aki",
		},
		want: &grpc.ListUsersResponse{
			Users: []*grpc.User{},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('akira','akira@gmail.com','$2a$10$tyNZh8SbBwP1rvGzHWzSPeNATz/N24wTXKDc1FS53waJzFlEeTWl6','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
	{
		id:   6,
		name: "ユーザ一覧正常系 created_at_from検索<TID:6>",
		in: &grpc.ListUsersRequest{
			CreatedAtFrom: "2022-09-15 08:47:22.182",
		},
		want: &grpc.ListUsersResponse{
			Users: []*grpc.User{
				{
					Id:        2,
					Name:      "akira",
					Email:     "akira@gmail.com",
					CreatedAt: "2022-09-20 08:47:22.182 +0000 UTC",
					UpdatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-10 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('akira','akira@gmail.com','$2a$10$tyNZh8SbBwP1rvGzHWzSPeNATz/N24wTXKDc1FS53waJzFlEeTWl6','2022-09-20 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
	{
		id:   7,
		name: "ユーザ一覧正常系 created_at_to 検索<TID:7>",
		in: &grpc.ListUsersRequest{
			CreatedAtTo: "2022-09-12 08:47:22.182",
		},
		want: &grpc.ListUsersResponse{
			Users: []*grpc.User{
				{
					Id:        1,
					Name:      "test",
					Email:     "test@gmail.com",
					CreatedAt: "2022-09-10 08:47:22.182 +0000 UTC",
					UpdatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-10 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('akira','akira@gmail.com','$2a$10$tyNZh8SbBwP1rvGzHWzSPeNATz/N24wTXKDc1FS53waJzFlEeTWl6','2022-09-20 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
	{
		id:   8,
		name: "ユーザ一覧正常系 created_at_from created_at_to 範囲検索<TID:8>",
		in: &grpc.ListUsersRequest{
			CreatedAtFrom: "2022-09-12 08:47:22.182",
			CreatedAtTo:   "2022-09-16 08:47:22.182",
		},
		want: &grpc.ListUsersResponse{
			Users: []*grpc.User{
				{
					Id:        3,
					Name:      "ishii",
					Email:     "ishii@gmail.com",
					CreatedAt: "2022-09-15 08:47:22.182 +0000 UTC",
					UpdatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
				},
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-10 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('akira','akira@gmail.com','$2a$10$tyNZh8SbBwP1rvGzHWzSPeNATz/N24wTXKDc1FS53waJzFlEeTWl6','2022-09-20 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('ishii','ishii@gmail.com','$2a$10$tyNZh8SbBwP1rvGzHWzSPeNATz/N24wTXKDc1FS53waJzFlEeTWl6','2022-09-15 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
}

func testListUser(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	for _, tt := range listUserTestcases {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				initDBForTests(context.Background(), t, client)
				if tt.setup != nil {
					tt.setup(ctx, t, client)
				}

				repo := NewUserRepository(client)
				got, err := repo.ListUsers(ctx, tt.in)
				require.NoError(t, err)

				assert.Equal(t, tt.want, got)
			},
		)

	}
}

func TestAllUserTest(t *testing.T) {
	testListUser(t)
}
