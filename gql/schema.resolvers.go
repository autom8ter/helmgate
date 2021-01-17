package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/autom8ter/kdeploy/gen/gql/go/generated"
	"github.com/autom8ter/kdeploy/gen/gql/go/model"
	kdeploypb "github.com/autom8ter/kdeploy/gen/grpc/go"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateApp(ctx context.Context, input model.AppConstructor) (*model.App, error) {
	app, err := r.client.CreateApp(ctx, toAppC(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromApp(app), nil
}

func (r *mutationResolver) UpdateApp(ctx context.Context, input model.AppUpdate) (*model.App, error) {
	app, err := r.client.UpdateApp(ctx, toAppU(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromApp(app), nil
}

func (r *mutationResolver) DelApp(ctx context.Context, input model.AppRef) (*string, error) {
	_, err := r.client.DeleteApp(ctx, &kdeploypb.AppRef{
		Name:      input.Name,
		Namespace: input.Namespace,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *mutationResolver) DelAll(ctx context.Context, input model.Namespace) (*string, error) {
	_, err := r.client.DeleteAll(ctx, &kdeploypb.Namespace{Namespace: input.Namespace})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *queryResolver) GetApp(ctx context.Context, input model.AppRef) (*model.App, error) {
	app, err := r.client.GetApp(ctx, &kdeploypb.AppRef{
		Name:      input.Name,
		Namespace: input.Namespace,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromApp(app), nil
}

func (r *queryResolver) ListApps(ctx context.Context, input model.Namespace) ([]*model.App, error) {
	apps, err := r.client.ListApps(ctx, &kdeploypb.Namespace{Namespace: input.Namespace})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var toReturn []*model.App
	for _, a := range apps.GetApplications() {
		toReturn = append(toReturn, fromApp(a))
	}
	return toReturn, nil
}

func (r *queryResolver) ListNamespaces(ctx context.Context, input *string) (*model.Namespaces, error) {
	namespaces, err := r.client.ListNamespaces(ctx, &empty.Empty{})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var toReturn = &model.Namespaces{}
	for _, n := range namespaces.GetNamespaces() {
		toReturn.Namespaces = append(toReturn.Namespaces, n)
	}
	return toReturn, nil
}

func (r *subscriptionResolver) Logs(ctx context.Context, input model.AppRef) (<-chan string, error) {
	stream, err := r.client.Logs(ctx, &kdeploypb.AppRef{
		Name:      input.Name,
		Namespace: input.Namespace,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	ch := make(chan string)
	go func() {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		for {
			select {
			case <-ctx.Done():
				close(ch)
				return
			default:
				msg, err := stream.Recv()
				if err != nil {
					ch <- err.Error()
				}
				ch <- msg.GetMessage()
			}
		}
	}()
	return ch, nil
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
