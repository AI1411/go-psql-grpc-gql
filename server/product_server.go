package server

import (
	"context"

	"github.com/AI1411/go-psql_grpc_gql/grpc"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/repository"
)

type ProductServer struct {
	grpc.UnimplementedProductServiceServer
	r *repository.ProductRepository
}

func NewProductServer(r *repository.ProductRepository) *ProductServer {
	return &ProductServer{
		r: r,
	}
}

func (s *ProductServer) ListProducts(ctx context.Context, in *grpc.ListProductsRequest) (*grpc.ListProductsResponse, error) {
	res, err := s.r.ListProducts(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
