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

func TestListProduct(t *testing.T) {
	ctx, client := initializeForRepositoryTest(t)

	var listProductTestcases = []struct {
		id    int
		name  string
		in    *grpc.ListProductsRequest
		want  *grpc.ListProductsResponse
		setup func(ctx context.Context, t *testing.T, client *db.Client)
	}{
		{
			id:   1,
			name: "Product 一覧正常系<TID:1>",
			in:   &grpc.ListProductsRequest{},
			want: &grpc.ListProductsResponse{
				Products: []*grpc.Product{
					{
						Id:            "8c2ca258-8b16-437b-9de6-f5650c3e385e",
						Name:          "test",
						Description:   helper.StringToPtr("remarks"),
						Price:         1000,
						DiscountPrice: helper.Uint32ToPtr(900),
						Status:        "sale",
						UserId:        uint32(1),
						CreatedAt:     "2022-09-16 08:47:22.182 +0000 UTC",
						UpdatedAt:     "2022-09-16 08:47:22.182 +0000 UTC",
					},
					{
						Id:            "ecdde875-0d2a-454d-ace4-8eb613bdda87",
						Name:          "name",
						Description:   helper.StringToPtr("description"),
						Price:         2000,
						DiscountPrice: helper.Uint32ToPtr(1800),
						Status:        "sold",
						UserId:        uint32(1),
						CreatedAt:     "2022-09-16 08:47:22.182 +0000 UTC",
						UpdatedAt:     "2022-09-16 08:47:22.182 +0000 UTC",
					},
				},
			},
			setup: func(ctx context.Context, t *testing.T, client *db.Client) {
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.products ("id", "name", "description","price","discount_price","status","user_id", "created_at", "updated_at") VALUES ('8c2ca258-8b16-437b-9de6-f5650c3e385e','test', 'remarks', 1000, 900, 'sale',1,'2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.products ("id", "name", "description","price","discount_price","status","user_id", "created_at", "updated_at") VALUES ('ecdde875-0d2a-454d-ace4-8eb613bdda87','name', 'description', 2000, 1800, 'sold',1,'2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			},
		},
		{
			id:   2,
			name: "Product 一覧正常系 Name検索<TID:2>",
			in: &grpc.ListProductsRequest{
				Name: helper.StringToPtr("test"),
			},
			want: &grpc.ListProductsResponse{
				Products: []*grpc.Product{
					{
						Id:            "8c2ca258-8b16-437b-9de6-f5650c3e385e",
						Name:          "test",
						Description:   helper.StringToPtr("remarks"),
						Price:         1000,
						DiscountPrice: helper.Uint32ToPtr(900),
						Status:        "sale",
						UserId:        uint32(1),
						CreatedAt:     "2022-09-16 08:47:22.182 +0000 UTC",
						UpdatedAt:     "2022-09-16 08:47:22.182 +0000 UTC",
					},
				},
			},
			setup: func(ctx context.Context, t *testing.T, client *db.Client) {
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.products ("id", "name", "description","price","discount_price","status","user_id", "created_at", "updated_at") VALUES ('8c2ca258-8b16-437b-9de6-f5650c3e385e','test', 'remarks', 1000, 900, 'sale',1,'2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.products ("id", "name", "description","price","discount_price","status","user_id", "created_at", "updated_at") VALUES ('ecdde875-0d2a-454d-ace4-8eb613bdda87','name', 'description', 2000, 1800, 'sold',1,'2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			},
		},
		{
			id:   3,
			name: "Product 一覧正常系 UserID 検索<TID:3>",
			in: &grpc.ListProductsRequest{
				UserId: helper.Uint32ToPtr(2),
			},
			want: &grpc.ListProductsResponse{
				Products: []*grpc.Product{
					{
						Id:            "ecdde875-0d2a-454d-ace4-8eb613bdda87",
						Name:          "name",
						Description:   helper.StringToPtr("description"),
						Price:         2000,
						DiscountPrice: helper.Uint32ToPtr(1800),
						Status:        "sold",
						UserId:        uint32(2),
						CreatedAt:     "2022-09-16 08:47:22.182 +0000 UTC",
						UpdatedAt:     "2022-09-16 08:47:22.182 +0000 UTC",
					},
				},
			},
			setup: func(ctx context.Context, t *testing.T, client *db.Client) {
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.products ("id", "name", "description","price","discount_price","status","user_id", "created_at", "updated_at") VALUES ('8c2ca258-8b16-437b-9de6-f5650c3e385e','test', 'remarks', 1000, 900, 'sale',1,'2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.products ("id", "name", "description","price","discount_price","status","user_id", "created_at", "updated_at") VALUES ('ecdde875-0d2a-454d-ace4-8eb613bdda87','name', 'description', 2000, 1800, 'sold',2,'2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			},
		},
	}

	for _, tt := range listProductTestcases {
		t.Run(
			tt.name, func(t *testing.T) {
				initDBForTests(context.Background(), t, client)
				if tt.setup != nil {
					tt.setup(ctx, t, client)
				}

				repo := NewProductRepository(client)
				got, err := repo.ListProducts(ctx, tt.in)
				require.NoError(t, err)

				assert.Equal(t, tt.want, got)
			},
		)

	}
}
