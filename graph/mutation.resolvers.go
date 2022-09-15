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
	panic(fmt.Errorf("not implemented: UpdateTest - updateTest"))
}

// DeleteTest is the resolver for the deleteTest field.
func (r *mutationResolver) DeleteTest(ctx context.Context, id int) (int, error) {
	panic(fmt.Errorf("not implemented: DeleteTest - deleteTest"))
}

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input model.CreateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: CreateUser - createUser"))
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUserInput) (*model.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (*model.User, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
