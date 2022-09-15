package repository

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/grpc"
	"github.com/AI1411/go-psql_grpc_gql/internal/helper"
)

var listTestTestcases = []struct {
	id    int
	name  string
	in    *grpc.ListTestRequest
	want  []*grpc.GetTestResponse
	setup func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "テスト一覧正常系<TID:1>",
		in:   &grpc.ListTestRequest{},
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
	{
		id:   2,
		name: "IDで検索<TID:2>",
		in: &grpc.ListTestRequest{
			Id: helper.Uint32ToPtr(1),
		},
		want: []*grpc.GetTestResponse{
			{
				Id:   1,
				Name: "test1",
			},
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test1');`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test2');`).Error)
		},
	},
	{
		id:   3,
		name: "Nameで検索<TID:3>",
		in: &grpc.ListTestRequest{
			Name: "test2",
		},
		want: []*grpc.GetTestResponse{
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

func testListTest(t *testing.T) {
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
				got, err := repo.ListTest(ctx, tt.in)
				require.NoError(t, err)

				assert.Equal(t, tt.want, got)
			},
		)

	}
}

var getTestTestcases = []struct {
	id        int
	name      string
	in        *grpc.GetTestRequest
	want      *grpc.GetTestResponse
	wantError error
	setup     func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "テスト詳細正常系<TID:1>",
		in: &grpc.GetTestRequest{
			Id: 1,
		},
		want: &grpc.GetTestResponse{
			Id:   1,
			Name: "test1",
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test1');`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test2');`).Error)
		},
	},
	{
		id:   2,
		name: "テスト詳細異常系 対象が見つからない場合、NotFoundエラーになること<TID:2>",
		in: &grpc.GetTestRequest{
			Id: 3,
		},
		wantError: status.Error(codes.NotFound, "test not found"),
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test1');`).Error)
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test2');`).Error)
		},
	},
}

func testGetTest(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	for _, tt := range getTestTestcases {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				initDBForTests(context.Background(), t, client)
				if tt.setup != nil {
					tt.setup(ctx, t, client)
				}

				repo := NewTestRepository(client)
				got, err := repo.GetTest(ctx, tt.in)

				if tt.wantError != nil {
					assert.Equal(t, tt.wantError, err)
				}

				if tt.want != nil {
					assert.Equal(t, tt.want, got)
				}
			},
		)

	}
}

var createTestTestcases = []struct {
	id        int
	name      string
	in        *grpc.CreateTestRequest
	want      *grpc.CreateTestResponse
	wantError error
	setup     func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "テスト作成正常系<TID:1>",
		in: &grpc.CreateTestRequest{
			Name: "test",
		},
		want: &grpc.CreateTestResponse{
			Id:   1,
			Name: "test",
		},
	},
}

func testCreateTest(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	for _, tt := range createTestTestcases {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				initDBForTests(context.Background(), t, client)
				if tt.setup != nil {
					tt.setup(ctx, t, client)
				}

				repo := NewTestRepository(client)
				got, err := repo.CreateTest(ctx, tt.in)

				if tt.wantError != nil {
					assert.Equal(t, tt.wantError, err)
				}

				if tt.want != nil {
					assert.Equal(t, tt.want, got)
				}
			},
		)

	}
}

var updateTestTestcases = []struct {
	id        int
	name      string
	in        *grpc.UpdateTestRequest
	want      *grpc.UpdateTestResponse
	wantError error
	setup     func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "テスト更新正常系<TID:1>",
		in: &grpc.UpdateTestRequest{
			Id:   1,
			Name: "updated",
		},
		want: &grpc.UpdateTestResponse{
			Id:   1,
			Name: "updated",
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test');`).Error)
		},
	},
	{
		id:   2,
		name: "テスト更新異常系 更新対象が見当たらない場合、NotFoundエラーが返ること<TID:2>",
		in: &grpc.UpdateTestRequest{
			Id:   3,
			Name: "updated",
		},
		wantError: status.Error(codes.NotFound, "test not found"),
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test');`).Error)
		},
	},
}

func testUpdateTest(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	for _, tt := range updateTestTestcases {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				initDBForTests(context.Background(), t, client)
				if tt.setup != nil {
					tt.setup(ctx, t, client)
				}

				repo := NewTestRepository(client)
				got, err := repo.UpdateTest(ctx, tt.in)

				if tt.wantError != nil {
					assert.Error(t, tt.wantError, err)
				}

				if tt.want != nil {
					assert.Equal(t, tt.want, got)
				}
			},
		)

	}
}

var deleteTestTestcases = []struct {
	id        int
	name      string
	in        *grpc.DeleteTestRequest
	want      *grpc.DeleteTestResponse
	wantError error
	setup     func(ctx context.Context, t *testing.T, client *db.Client)
}{
	{
		id:   1,
		name: "テスト削除正常系<TID:1>",
		in: &grpc.DeleteTestRequest{
			Id: 1,
		},
		want: &grpc.DeleteTestResponse{
			Id: 1,
		},
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test');`).Error)
		},
	},
	{
		id:   2,
		name: "テスト削除異常系 削除対象が見当たらない場合、NotFoundエラーが返ること<TID:2>",
		in: &grpc.DeleteTestRequest{
			Id: 3,
		},
		wantError: status.Error(codes.NotFound, "test not found"),
		setup: func(ctx context.Context, t *testing.T, client *db.Client) {
			require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.tests ("id", "name") VALUES (DEFAULT, 'test');`).Error)
		},
	},
}

func testDeleteTest(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	for _, tt := range deleteTestTestcases {
		tt := tt
		t.Run(
			tt.name, func(t *testing.T) {
				initDBForTests(context.Background(), t, client)
				if tt.setup != nil {
					tt.setup(ctx, t, client)
				}

				repo := NewTestRepository(client)
				got, err := repo.DeleteTest(ctx, tt.in)

				if tt.wantError != nil {
					assert.Error(t, tt.wantError, err)
				}

				if tt.want != nil {
					assert.Equal(t, tt.want, got)
				}
			},
		)

	}
}

func TestAllTestcaseOfTest(t *testing.T) {
	testListTest(t)
	testGetTest(t)
	testCreateTest(t)
	testUpdateTest(t)
	testDeleteTest(t)
}
