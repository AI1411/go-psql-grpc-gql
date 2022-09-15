package repository

import (
	"context"
	"time"

	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/grpc"
)

type User struct {
	ID        uint32    `gorm:"primaryKey"`
	Name      string    `gorm:"type:varchar(255);not null"`
	Email     string    `gorm:"type:varchar(255);not null"`
	Password  string    `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
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
	r.dbClient.Conn(ctx).Find(&users)

	res := make([]*grpc.User, len(users))
	for i, user := range users {
		res[i] = &grpc.User{
			Id:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Password:  user.Password,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
		}
	}

	grpcResponse := &grpc.ListUsersResponse{
		Users: res,
	}
	return grpcResponse, nil
}
