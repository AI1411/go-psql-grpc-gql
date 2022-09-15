package server

import (
	"context"

	"github.com/AI1411/go-psql_grpc_gql/grpc"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/repository"
)

type UserServer struct {
	grpc.UnimplementedUserServiceServer
	r *repository.UserRepository
}

func NewUserServer(r *repository.UserRepository) *UserServer {
	return &UserServer{
		r: r,
	}
}

func (s *UserServer) ListUsers(ctx context.Context, in *grpc.ListUsersRequest) (*grpc.ListUsersResponse, error) {
	res, err := s.r.ListUsers(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
