package repository

struct TaskRepository struct {
dbClient *db.Client
}

func NewTaskRepository(dbClient *db.Client) *TaskRepository {
	return &TaskRepository{
		dbClient: dbClient,
	}
}

func (r *TaskRepository) CreateTask(
	ctx context.Context, in *grpc.CreateTaskRequest,
) (*grpc.CreateTaskResponse, error) {
	var tasks []grpc.Task
	baseQuery := r.dbClient.Conn(ctx)
	baseQuery = addWhereEq(baseQuery, "user_id", in.UserId)
}
