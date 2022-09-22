package repository

import "github.com/AI1411/go-psql_grpc_gql/db"

type TaskRepository struct {
	dbClient *db.Client
}

func NewTaskRepository(dbClient *db.Client) *TaskRepository {
	return &TaskRepository{
		dbClient: dbClient,
	}
}

func (r *TaskRepository) ListTasks() {
}
