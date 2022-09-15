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

func (s *UserServer) GetUser(ctx context.Context, in *grpc.GetUserRequest) (*grpc.GetUserResponse, error) {
	res, err := s.r.GetUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *UserServer) CreateUser(ctx context.Context, in *grpc.CreateUserRequest) (*grpc.CreateUserResponse, error) {
	res, err := s.r.CreateUser(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
