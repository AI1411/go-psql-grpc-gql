package graph_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AI1411/go-psql_grpc_gql/db"
)

func TestQueryResolver_Products(t *testing.T) {
	ctx := context.Background()
	srv, client := NewGqlServer()

	var productsResolverTestCases = []struct {
		id      int
		name    string
		query   string
		want    map[string]interface{}
		wantErr error
		setup   func(ctx context.Context, t *testing.T, client *db.Client)
	}{
		{
			id:   1,
			name: "正常系 product 一覧を取得できる name 検索",
			query: `query ListProducts {
			  products(input: {
				name: "sale"
			  }) {
				id
				name
				description
				status
				user_id
				price
				discountPrice
			  }
			}`,
			want: map[string]interface{}{
				"products": []interface{}{
					map[string]interface{}{
						"id":            "8c2ca258-8b16-437b-9de6-f5650c3e385e",
						"name":          "sale",
						"description":   "remarks",
						"status":        "sale",
						"user_id":       "1",
						"price":         float64(1000),
						"discountPrice": float64(900),
					},
				},
			},
			setup: func(ctx context.Context, t *testing.T, client *db.Client) {
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.products ("id", "name", "description","price","discount_price","status","user_id", "created_at", "updated_at") VALUES ('8c2ca258-8b16-437b-9de6-f5650c3e385e','sale', 'remarks', 1000, 900, 'sale',1,'2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.products ("id", "name", "description","price","discount_price","status","user_id", "created_at", "updated_at") VALUES ('ecdde875-0d2a-454d-ace4-8eb613bdda87','name', 'description', 2000, 1800, 'sold',1,'2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			},
		},
		{
			id:   2,
			name: "正常系 product 一覧を取得できる status 検索",
			query: `query ListProducts {
			  products(input: {
				status: "sold"
			  }) {
				id
				name
				description
				status
				user_id
				price
				discountPrice
			  }
			}`,
			want: map[string]interface{}{
				"products": []interface{}{
					map[string]interface{}{
						"id":            "ecdde875-0d2a-454d-ace4-8eb613bdda87",
						"name":          "name",
						"description":   "description",
						"status":        "sold",
						"user_id":       "1",
						"price":         float64(2000),
						"discountPrice": float64(1800),
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
			name: "正常系 product 一覧を取得できる createdAtFrom 検索",
			query: `query ListProducts {
			  products(input: {
				createdAtFrom: "2022-09-17 08:47:22.182000"
			  }) {
				id
				name
				description
				status
				user_id
				price
				discountPrice
			  }
			}`,
			want: map[string]interface{}{
				"products": []interface{}{
					map[string]interface{}{
						"id":            "ecdde875-0d2a-454d-ace4-8eb613bdda87",
						"name":          "name",
						"description":   "description",
						"status":        "sold",
						"user_id":       "1",
						"price":         float64(2000),
						"discountPrice": float64(1800),
					},
				},
			},
			setup: func(ctx context.Context, t *testing.T, client *db.Client) {
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.products ("id", "name", "description","price","discount_price","status","user_id", "created_at", "updated_at") VALUES ('8c2ca258-8b16-437b-9de6-f5650c3e385e','test', 'remarks', 1000, 900, 'sale',1,'2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.products ("id", "name", "description","price","discount_price","status","user_id", "created_at", "updated_at") VALUES ('ecdde875-0d2a-454d-ace4-8eb613bdda87','name', 'description', 2000, 1800, 'sold',1,'2022-11-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			},
		},
		{
			id:   4,
			name: "正常系 product 一覧を取得できる createdAtTo 検索",
			query: `query ListProducts {
			  products(input: {
				createdAtTo: "2022-09-17 08:47:22.182000"
			  }) {
				id
				name
				description
				status
				user_id
				price
				discountPrice
			  }
			}`,
			want: map[string]interface{}{
				"products": []interface{}{
					map[string]interface{}{
						"id":            "8c2ca258-8b16-437b-9de6-f5650c3e385e",
						"name":          "test",
						"description":   "remarks",
						"status":        "sale",
						"user_id":       "1",
						"price":         float64(1000),
						"discountPrice": float64(900),
					},
				},
			},
			setup: func(ctx context.Context, t *testing.T, client *db.Client) {
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.products ("id", "name", "description","price","discount_price","status","user_id", "created_at", "updated_at") VALUES ('8c2ca258-8b16-437b-9de6-f5650c3e385e','test', 'remarks', 1000, 900, 'sale',1,'2022-09-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
				require.NoError(t, client.Conn(ctx).Exec(`INSERT INTO public.products ("id", "name", "description","price","discount_price","status","user_id", "created_at", "updated_at") VALUES ('ecdde875-0d2a-454d-ace4-8eb613bdda87','name', 'description', 2000, 1800, 'sold',1,'2022-11-16 08:47:22.182','2022-09-16 08:47:22.182' )`).Error)
			},
		},
	}

	for _, tt := range productsResolverTestCases {
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
