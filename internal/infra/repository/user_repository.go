package repository

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/grpc"
)

type User struct {
	ID        uint32
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type UserRepository struct {
	dbClient *db.Client
}

func NewUserRepository(dbClient *db.Client) *UserRepository {
	return &UserRepository{
		dbClient: dbClient,
	}
}

func (r *UserRepository) ListUsers(ctx context.Context, in *grpc.ListUsersRequest,
) (*grpc.ListUsersResponse, error) {
	var users []User
	baseQuery := r.dbClient.Conn(ctx)
	baseQuery = addWhereLike(baseQuery, "name", in.Name)
	baseQuery = addWhereEq(baseQuery, "email", in.Email)
	baseQuery = addWhereGte(baseQuery, "created_at", in.CreatedAtFrom)
	baseQuery = addWhereLte(baseQuery, "created_at", in.CreatedAtTo)
	baseQuery.Find(&users)

	res := make([]*grpc.User, len(users))
	for i, user := range users {
		res[i] = &grpc.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		}
	}

	grpcResponse := &grpc.ListUsersResponse{
		Users: res,
	}
	return grpcResponse, nil
}

func (r *UserRepository) GetUser(ctx context.Context, in *grpc.GetUserRequest,
) (*grpc.GetUserResponse, error) {
	var user User
	if err := r.dbClient.Conn(ctx).First(&user, in.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	grpcResponse := &grpc.GetUserResponse{
		User: &grpc.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		},
	}
	return grpcResponse, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, in *grpc.CreateUserRequest,
) (*grpc.CreateUserResponse, error) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	user := User{
		Name:      in.Name,
		Email:     in.Email,
		Password:  string(hash),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	r.dbClient.Conn(ctx).Create(&user)

	grpcResponse := &grpc.CreateUserResponse{
		User: &grpc.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}
	return grpcResponse, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, in *grpc.UpdateUserRequest,
) (*grpc.UpdateUserResponse, error) {
	var user User
	if err := r.dbClient.Conn(ctx).First(&user, in.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	user = User{
		ID:        user.ID,
		Name:      in.Name,
		Email:     in.Email,
		Password:  string(hash),
		UpdatedAt: time.Now(),
	}
	if err := r.dbClient.Conn(ctx).Save(&user).Error; err != nil {
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	grpcResponse := &grpc.UpdateUserResponse{
		User: &grpc.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}
	return grpcResponse, nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, in *grpc.DeleteUserRequest,
) (*grpc.DeleteUserResponse, error) {
	var user User
	if err := r.dbClient.Conn(ctx).First(&user, in.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	if err := r.dbClient.Conn(ctx).Delete(&user).Error; err != nil {
		return nil, status.Error(codes.Internal, "failed to delete user")
	}

	grpcResponse := &grpc.DeleteUserResponse{
		User: &grpc.User{
			Id:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
	}

	return grpcResponse, nil
}
