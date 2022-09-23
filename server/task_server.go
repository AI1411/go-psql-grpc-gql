package server

import (
	"context"

	"github.com/AI1411/go-psql_grpc_gql/grpc"
	"github.com/AI1411/go-psql_grpc_gql/internal/infra/repository"
)

type TaskServer struct {
	grpc.UnimplementedTaskServiceServer
	r *repository.TaskRepository
}

func NewTaskServer(r *repository.TaskRepository) *TaskServer {
	return &TaskServer{
		r: r,
	}
}

func (s *TaskServer) ListTasks(ctx context.Context, in *grpc.ListTasksRequest) (*grpc.ListTasksResponse, error) {
	res, err := s.r.ListTasks(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (s *TaskServer) CreateTask(ctx context.Context, in *grpc.CreateTaskRequest) (*grpc.CreateTaskResponse, error) {
	res, err := s.r.CreateTask(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}
