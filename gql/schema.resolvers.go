package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/autom8ter/kdeploy/gen/gql/go/generated"
	"github.com/autom8ter/kdeploy/gen/gql/go/model"
)

func (r *mutationResolver) CreateApp(ctx context.Context, input model.AppConstructor) (*model.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateApp(ctx context.Context, input model.AppUpdate) (*model.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelApp(ctx context.Context, input model.Ref) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateTask(ctx context.Context, input model.TaskConstructor) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UpdateTask(ctx context.Context, input model.TaskUpdate) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelTask(ctx context.Context, input model.Ref) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DelAll(ctx context.Context, input model.Namespace) (*string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetApp(ctx context.Context, input model.Ref) (*model.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListApps(ctx context.Context, input model.Namespace) ([]*model.App, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTask(ctx context.Context, input model.Ref) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListTasks(ctx context.Context, input model.Namespace) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) ListNamespaces(ctx context.Context, input *string) (*model.Namespaces, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *subscriptionResolver) StreamLogs(ctx context.Context, input model.Ref) (<-chan string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Subscription returns generated.SubscriptionResolver implementation.
func (r *Resolver) Subscription() generated.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
