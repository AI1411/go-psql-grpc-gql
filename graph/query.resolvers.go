package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/AI1411/go-psql_grpc_gql/graph/generated"
	"github.com/AI1411/go-psql_grpc_gql/graph/model"
	"github.com/AI1411/go-psql_grpc_gql/grpc"
)

// Test is the resolver for the test field.
func (r *queryResolver) Test(ctx context.Context, input int) (*model.Test, error) {
	test, err := r.TestServer.GetTest(ctx, &grpc.GetTestRequest{
		Id: uint32(input),
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

// Tests is the resolver for the tests field.
func (r *queryResolver) Tests(ctx context.Context) ([]*model.Test, error) {
	tests, err := r.TestServer.ListTest(ctx, &grpc.ListTestRequest{})
	if err != nil {
		return nil, err
	}

	response := make([]*model.Test, len(tests.Tests))
	for i, test := range tests.Tests {
		response[i] = &model.Test{
			ID:   int(test.Id),
			Name: test.Name,
		}
	}
	return response, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	users, err := r.UserServer.ListUsers(ctx, &grpc.ListUsersRequest{})
	if err != nil {
		return nil, err
	}

	response := make([]*model.User, len(users.Users))
	for i, user := range users.Users {
		response[i] = &model.User{
			ID:        int(user.Id),
			Name:      user.Name,
			Email:     user.Email,
			CreatedAt: user.CreatedAt,
		}
	}
	return response, nil
}

// User is the resolver for the user field.
func (r *queryResolver) User(ctx context.Context, id int) (*model.User, error) {
	user, err := r.UserServer.GetUser(ctx, &grpc.GetUserRequest{
		Id: uint32(id),
	})
	if err != nil {
		return nil, err
	}

	response := &model.User{
		ID:        int(user.User.Id),
		Name:      user.User.Name,
		Email:     user.User.Email,
		CreatedAt: user.User.CreatedAt,
	}

	return response, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }