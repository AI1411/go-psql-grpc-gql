package repository

import (
	"context"

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
		}
	}

	grpcResponse := &grpc.ListTasksResponse{
		Tasks: res,
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
