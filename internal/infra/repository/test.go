package repository

import (
	"context"

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
		return nil, err
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
