package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/AI1411/go-psql_grpc_gql/graph/generated"
	"github.com/AI1411/go-psql_grpc_gql/graph/model"
	"github.com/AI1411/go-psql_grpc_gql/grpc"
)

// CreateTest is the resolver for the createTest field.
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

// UpdateTest is the resolver for the updateTest field.
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

// DeleteTest is the resolver for the deleteTest field.
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

// CreateUser is the resolver for the createUser field.
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

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	user, err := r.UserServer.UpdateUser(ctx, &grpc.UpdateUserRequest{
		Id:       uint32(input.ID),
		Name:     *input.Name,
		Email:    *input.Email,
		Password: *input.Password,
	})

	if err != nil {
		return nil, err
	}

	response := &model.User{
		ID:       int(user.User.Id),
		Name:     user.User.Name,
		Email:    user.User.Email,
		Password: user.User.Password,
	}

	return response, nil
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*model.User, error) {
	user, err := r.UserServer.DeleteUser(ctx, &grpc.DeleteUserRequest{
		Id: uint32(id),
	})

	if err != nil {
		return nil, err
	}

	response := &model.User{
		ID: int(user.User.Id),
	}

	return response, nil
}

// ChangePassword is the resolver for the changePassword field.
func (r *mutationResolver) ChangePassword(ctx context.Context, input model.ChangePasswordInput) (*model.ChangePasswordResponse, error) {
	password, err := r.UserServer.ChangePassword(ctx, &grpc.ChangePasswordRequest{
		Id:          uint32(input.ID),
		OldPassword: input.OldPassword,
		NewPassword: input.NewPassword,
	})

	if err != nil {
		return nil, err
	}

	response := &model.ChangePasswordResponse{
		Password: password.NewPassword,
	}

	return response, nil
}

// CreateTask is the resolver for the createTask field.
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
