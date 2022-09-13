package repository

import (
	"context"
	"errors"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/AI1411/go-pg-ci-example/db"
	"github.com/AI1411/go-pg-ci-example/grpc"
)

type Test struct {
	ID   uint32 `gorm:"id;primaryKey"`
	Name string `gorm:"name"`
}

type TestRepository struct {
	dbClient *db.Client
}

func NewTestRepository(dbClient *db.Client) *TestRepository {
	return &TestRepository{
		dbClient: dbClient,
	}
}

func (r *TestRepository) ListTest(
	ctx context.Context, in *grpc.ListTestRequest,
) ([]*grpc.GetTestResponse, error) {
	var tests []Test
	baseQuery := r.dbClient.Conn(ctx)
	baseQuery = addWhereLike(baseQuery, "name", in.Name)
	baseQuery = addWhereEq(baseQuery, "id", in.Id)
	if err := baseQuery.
		Find(&tests).Error; err != nil {
		return nil, status.Error(codes.Internal, "failed to get test list")
	}

	res := make([]*grpc.GetTestResponse, len(tests))
	for i, t := range tests {
		res[i] = &grpc.GetTestResponse{
			Id:   t.ID,
			Name: t.Name,
		}
	}
	return res, nil
}

func (r *TestRepository) GetTest(ctx context.Context, request *grpc.GetTestRequest) (*grpc.GetTestResponse, error) {
	var test Test
	if err := r.dbClient.Conn(ctx).Where("id = ?", request.Id).First(&test).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "test not found")
		}
		return nil, status.Error(codes.Internal, "failed to get test")
	}
	return &grpc.GetTestResponse{
		Id:   test.ID,
		Name: test.Name,
	}, nil
}
