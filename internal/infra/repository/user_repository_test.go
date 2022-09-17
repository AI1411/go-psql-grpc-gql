package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

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

func TestListUser(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	for _, tt := range listUserTestcases {
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

				repo := NewUserRepository(client)
				got, err := repo.ListUsers(ctx, tt.in)
				require.NoError(t, err)

				assert.Equal(t, tt.want, got)
			},
		)

	}
}

var getUserTestcases = []struct {
	id        int
	name      string
	in        *grpc.GetUserRequest
	want      *grpc.GetUserResponse
	wantError error
	setup     func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "ユーザ詳細正常系<TID:1>",
		in: &grpc.GetUserRequest{
			Id: 1,
		},
		want: &grpc.GetUserResponse{
			User: &grpc.User{
				Id:        1,
				Name:      "test",
				Email:     "test@gmail.com",
				CreatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
				UpdatedAt: "2022-09-16 08:47:22.182 +0000 UTC",
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('akira','akira@gmail.com','$2a$10$tyNZh8SbBwP1rvGzHWzSPeNATz/N24wTXKDc1FS53waJzFlEeTWl6','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
	{
		id:   2,
		name: "ユーザ詳細異常系 レコードが見つからない場合、NotFoundエラーになること<TID:2>",
		in: &grpc.GetUserRequest{
			Id: 3,
		},
		wantError: status.Error(codes.NotFound, "user not found"),
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('akira','akira@gmail.com','$2a$10$tyNZh8SbBwP1rvGzHWzSPeNATz/N24wTXKDc1FS53waJzFlEeTWl6','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
}

func TestGetUser(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	for _, tt := range getUserTestcases {
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

				repo := NewUserRepository(client)
				got, err := repo.GetUser(ctx, tt.in)

				if tt.wantError != nil {
					require.Equal(t, tt.wantError, err)
				}

				if tt.want != nil {
					assert.Equal(t, tt.want, got)
				}
			},
		)

	}
}

var createUserTestcases = []struct {
	id        int
	name      string
	in        *grpc.CreateUserRequest
	want      *grpc.CreateUserResponse
	wantError error
	setup     func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "ユーザ作成正常系<TID:1>",
		in: &grpc.CreateUserRequest{
			Name:     "test",
			Email:    "test@gmail.com",
			Password: "password",
		},
		want: &grpc.CreateUserResponse{
			User: &grpc.User{
				Id:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
		},
	},
}

func TestCreateUser(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	for _, tt := range createUserTestcases {
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

				repo := NewUserRepository(client)
				got, err := repo.CreateUser(ctx, tt.in)

				if tt.wantError != nil {
					require.Equal(t, tt.wantError, err)
				}

				if tt.want != nil {
					assert.Equal(t, tt.want, got)
				}
			},
		)

	}
}

var updateUserTestcases = []struct {
	id        int
	name      string
	in        *grpc.UpdateUserRequest
	want      *grpc.UpdateUserResponse
	wantError error
	setup     func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "ユーザ更新正常系<TID:1>",
		in: &grpc.UpdateUserRequest{
			Id:       1,
			Name:     "update",
			Email:    "update@gmail.com",
			Password: "update",
		},
		want: &grpc.UpdateUserResponse{
			User: &grpc.User{
				Id:    1,
				Name:  "update",
				Email: "update@gmail.com",
			},
		},

		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
	{
		id:   2,
		name: "ユーザ更新異常系 レコードが見つからない場合、NotFoundエラーになること<TID:2>",
		in: &grpc.UpdateUserRequest{
			Id:       4,
			Name:     "update",
			Email:    "update@gmail.com",
			Password: "update",
		},
		wantError: status.Error(codes.NotFound, "user not found"),

		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
}

func TestUpdateUser(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	for _, tt := range updateUserTestcases {
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

				repo := NewUserRepository(client)
				got, err := repo.UpdateUser(ctx, tt.in)

				if tt.wantError != nil {
					require.Equal(t, tt.wantError, err)
				}

				if tt.want != nil {
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

var deleteUserTestcases = []struct {
	id        int
	name      string
	in        *grpc.DeleteUserRequest
	want      *grpc.DeleteUserResponse
	wantError error
	setup     func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "ユーザ削除正常系<TID:1>",
		in: &grpc.DeleteUserRequest{
			Id: 1,
		},
		want: &grpc.DeleteUserResponse{
			User: &grpc.User{
				Id:    1,
				Name:  "test",
				Email: "test@gmail.com",
			},
		},

		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
	{
		id:   2,
		name: "ユーザ削除異常系 レコードが見つからない場合、NotFoundエラーになること<TID:2>",
		in: &grpc.DeleteUserRequest{
			Id: 4,
		},
		wantError: status.Error(codes.NotFound, "user not found"),

		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$n4h5tHioqRmJjm/2MQyHYOCehdG1OjfV9VzH8YXWZ/LHH93rQjWiK','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
}

func TestDeleteUser(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	for _, tt := range deleteUserTestcases {
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

				repo := NewUserRepository(client)
				got, err := repo.DeleteUser(ctx, tt.in)

				if tt.wantError != nil {
					require.Equal(t, tt.wantError, err)
				}

				if tt.want != nil {
					assert.Equal(t, tt.want, got)
				}
			},
		)
	}
}

var changePasswordTestcases = []struct {
	id           int
	name         string
	in           *grpc.ChangePasswordRequest
	wantPassword string
	wantError    error
	setup        func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "ユーザパスワード変更正常系<TID:1>",
		in: &grpc.ChangePasswordRequest{
			Id:          1,
			OldPassword: "test",
			NewPassword: "password",
		},
		wantPassword: "password",
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$QON8RQ5kMr.JGRtPpB5RB.QbpjOxjZoIfUP3SBntExHXcGPnTUy5y','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
	{
		id:   2,
		name: "ユーザパスワード変更異常系 古いパスワードが現在のものと誤っている場合、エラーになること<TID:1>",
		in: &grpc.ChangePasswordRequest{
			Id:          1,
			OldPassword: "invalid",
			NewPassword: "password",
		},
		wantPassword: "password",
		wantError:    status.Error(codes.InvalidArgument, "invalid old password"),
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$QON8RQ5kMr.JGRtPpB5RB.QbpjOxjZoIfUP3SBntExHXcGPnTUy5y','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
	{
		id:   3,
		name: "ユーザパスワード変更異常系 新しいパスワードが入力されていない場合、エラーになること<TID:1>",
		in: &grpc.ChangePasswordRequest{
			Id:          1,
			OldPassword: "test",
			NewPassword: "",
		},
		wantPassword: "password",
		wantError:    status.Error(codes.InvalidArgument, "new password is required"),
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.users ("name","email","password","created_at","updated_at") VALUES ('test','test@gmail.com','$2a$10$QON8RQ5kMr.JGRtPpB5RB.QbpjOxjZoIfUP3SBntExHXcGPnTUy5y','2022-09-16 08:47:22.182','2022-09-16 08:47:22.182')`).Error)
		},
	},
}

func TestChangePassword(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	for _, tt := range changePasswordTestcases {
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

				repo := NewUserRepository(client)
				got, err := repo.ChangePassword(ctx, tt.in)

				if tt.wantError != nil {
					require.Equal(t, tt.wantError, err)
				}

				if tt.wantPassword != "" && tt.wantError == nil {
					err := bcrypt.CompareHashAndPassword([]byte(got.NewPassword), []byte(tt.wantPassword))
					assert.NoError(t, err)
				}
			},
		)
	}
}

func TestAllUserTest(t *testing.T) {
	TestListUser(t)
	TestGetUser(t)
	TestCreateUser(t)
	TestUpdateUser(t)
	TestDeleteUser(t)
	TestChangePassword(t)
}
