package repository

import (
	"context"
	"time"

	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/grpc"
)

type Product struct {
	ID            string    `gorm:"id;primaryKey"`
	Name          string    `gorm:"name"`
	Description   *string   `gorm:"description"`
	Price         uint32    `gorm:"price"`
	DiscountPrice *uint32   `gorm:"discount_price"`
	Status        string    `gorm:"status"`
	UserID        uint32    `gorm:"user_id"`
	CreatedAt     time.Time `gorm:"created_at"`
	UpdatedAt     time.Time `gorm:"updated_at"`
}

type ProductRepository struct {
	dbClient *db.Client
}

func NewProductRepository(dbClient *db.Client) *ProductRepository {
	return &ProductRepository{
		dbClient: dbClient,
	}
}

func (r *ProductRepository) ListProducts(
	ctx context.Context, in *grpc.ListProductsRequest,
) (*grpc.ListProductsResponse, error) {
	var products []Product
	baseQuery := r.dbClient.Conn(ctx)
	baseQuery = addWhereEq(baseQuery, "name", in.Name)
	baseQuery = addWhereEq(baseQuery, "status", in.Status)
	baseQuery = addWhereEq(baseQuery, "user_id", in.UserId)
	baseQuery = addWhereGte(baseQuery, "created_at", in.CreatedAtFrom)
	baseQuery = addWhereLte(baseQuery, "created_at", in.CreatedAtTo)
	baseQuery.Find(&products)

	res := make([]*grpc.Product, len(products))
	for i, product := range products {
		res[i] = &grpc.Product{
			Id:            product.ID,
			Name:          product.Name,
			Description:   product.Description,
			Price:         product.Price,
			DiscountPrice: product.DiscountPrice,
			Status:        product.Status,
			UserId:        product.UserID,
			CreatedAt:     product.CreatedAt.String(),
			UpdatedAt:     product.UpdatedAt.String(),
		}
	}

	grpcResponse := &grpc.ListProductsResponse{
		Products: res,
	}

	return grpcResponse, nil
}
