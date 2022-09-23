package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/AI1411/go-psql_grpc_gql/graph/generated"
	"github.com/AI1411/go-psql_grpc_gql/graph/model"
	"github.com/AI1411/go-psql_grpc_gql/grpc"
)

func (r *mutationResolver) CreateTest(ctx context.Context, input model.CreateTest) (*model.Test, error) {
	test, err := r.TestServer.CreateTest(ctx, &grpc.CreateTestRequest{
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}

	response := &model.Test{
		ID:   int(test.Id),
		Name: test.Name,
	}

	return response, nil
}

func (r *mutationResolver) UpdateTest(ctx context.Context, input model.UpdateTest) (*model.Test, error) {
	test, err := r.TestServer.UpdateTest(ctx, &grpc.UpdateTestRequest{
		Id:   uint32(input.ID),
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}

	response := &model.Test{
		ID:   int(test.Id),
		Name: test.Name,
	}

	return response, nil
}

func (r *mutationResolver) DeleteTest(ctx context.Context, id int) (*model.Test, error) {
	test, err := r.TestServer.DeleteTest(ctx, &grpc.DeleteTestRequest{
		Id: uint32(id),
	})

	if err != nil {
		return nil, err
	}

	response := &model.Test{
		ID: int(test.Id),
	}

	return response, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	user, err := r.UserServer.CreateUser(ctx, &grpc.CreateUserRequest{
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
	})
	if err != nil {
		return nil, err
	}

	response := &model.User{
		ID:        int(user.User.Id),
		Name:      user.User.Name,
		Email:     user.User.Email,
		Password:  user.User.Password,
		CreatedAt: user.User.CreatedAt,
		UpdatedAt: user.User.UpdatedAt,
	}

	return response, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	task, err := r.TaskServer.CreateTask(ctx, &grpc.CreateTaskRequest{
		Title:       input.Title,
		Description: *input.Description,
		UserId:      uint32(*input.UserID),
		Status:      input.Status,
		DueDate:     *input.DueDate,
	})

	if err != nil {
		return nil, err
	}

	response := &model.Task{
		ID:          int(task.Task.Id),
		Title:       task.Task.Title,
		Description: task.Task.Description,
		UserID:      int(task.Task.UserId),
		Completed:   task.Task.Completed,
		Status:      task.Task.Status,
		DueDate:     task.Task.DueDate,
		CreatedAt:   task.Task.CreatedAt,
		UpdatedAt:   task.Task.UpdatedAt,
	}

	return response, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
