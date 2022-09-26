package repository

import (
	"context"
	"errors"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"github.com/AI1411/go-psql_grpc_gql/db"
	"github.com/AI1411/go-psql_grpc_gql/grpc"
)

type Task struct {
	ID          uint32
	Title       string
	Description string
	DueDate     string
	Completed   bool
	UserID      uint32
	Status      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TaskRepository struct {
	dbClient *db.Client
}

func NewTaskRepository(dbClient *db.Client) *TaskRepository {
	return &TaskRepository{
		dbClient: dbClient,
	}
}

func (r *TaskRepository) ListTasks(
	ctx context.Context, in *grpc.ListTasksRequest,
) (*grpc.ListTasksResponse, error) {
	var tasks []Task
	baseQuery := r.dbClient.Conn(ctx)
	baseQuery = addWhereLike(baseQuery, "title", in.Title)
	baseQuery = addWhereLte(baseQuery, "due_date", in.DueDateTo)
	baseQuery = addWhereGte(baseQuery, "due_date", in.DueDateFrom)
	baseQuery = addWhereEq(baseQuery, "completed", in.Completed)
	baseQuery = addWhereEq(baseQuery, "user_id", in.UserId)
	baseQuery = addWhereEq(baseQuery, "status", in.Status)
	baseQuery = addWhereGte(baseQuery, "created_at", in.CreatedAtFrom)
	baseQuery = addWhereLte(baseQuery, "created_at", in.CreatedAtTo)
	baseQuery.Find(&tasks)

	res := make([]*grpc.Task, len(tasks))
	for i, task := range tasks {
		res[i] = &grpc.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			DueDate:     task.DueDate,
			Completed:   task.Completed,
			UserId:      task.UserID,
			Status:      task.Status,
			CreatedAt:   task.CreatedAt.String(),
			UpdatedAt:   task.UpdatedAt.String(),
		}
	}

	grpcResponse := &grpc.ListTasksResponse{
		Tasks: res,
	}

	return grpcResponse, nil
}

func (r *TaskRepository) GetTask(
	ctx context.Context, in *grpc.GetTaskRequest,
) (*grpc.GetTaskResponse, error) {
	var task Task
	if err := r.dbClient.Conn(ctx).First(&task, in.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	grpcResponse := &grpc.GetTaskResponse{
		Task: &grpc.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			DueDate:     task.DueDate,
			Completed:   task.Completed,
			UserId:      task.UserID,
			Status:      task.Status,
		},
	}

	return grpcResponse, nil
}

func (r *TaskRepository) CreateTask(
	ctx context.Context, in *grpc.CreateTaskRequest,
) (*grpc.CreateTaskResponse, error) {
	task := Task{
		Title:       in.Title,
		Description: in.Description,
		DueDate:     in.DueDate,
		UserID:      in.UserId,
		Status:      in.Status,
	}

	r.dbClient.Conn(ctx).Create(&task)

	grpcResponse := &grpc.CreateTaskResponse{
		Task: &grpc.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			DueDate:     task.DueDate,
			Completed:   task.Completed,
			UserId:      task.UserID,
			Status:      task.Status,
		},
	}

	return grpcResponse, nil
}

func (r *TaskRepository) UpdateTask(
	ctx context.Context, in *grpc.UpdateTaskRequest,
) (*grpc.UpdateTaskResponse, error) {
	var task Task
	if err := r.dbClient.Conn(ctx).First(&task, in.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	task = Task{
		Title:       *in.Title,
		Description: *in.Description,
		DueDate:     *in.DueDate,
		Completed:   *in.Completed,
		Status:      *in.Status,
		UserID:      *in.UserId,
	}

	if err := r.dbClient.Conn(ctx).Save(&task).Error; err != nil {
		return nil, status.Error(codes.Internal, "failed to update user")
	}

	grpcResponse := &grpc.UpdateTaskResponse{
		Task: &grpc.Task{
			Id:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			DueDate:     task.DueDate,
			Completed:   task.Completed,
			UserId:      task.UserID,
			Status:      task.Status,
		},
	}

	return grpcResponse, nil
}

func (r *TaskRepository) DeleteTask(ctx context.Context, in *grpc.DeleteTaskRequest) (*grpc.DeleteTaskResponse, error) {
	var task Task
	if err := r.dbClient.Conn(ctx).First(&task, in.Id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, status.Error(codes.NotFound, "user not found")
		}
		return nil, status.Error(codes.Internal, "failed to get user")
	}

	if err := r.dbClient.Conn(ctx).Delete(&task).Error; err != nil {
		return nil, status.Error(codes.Internal, "failed to delete user")
	}

	return &grpc.DeleteTaskResponse{}, nil
}
