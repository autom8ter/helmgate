package gql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/autom8ter/meshpaas/gen/gql/go/generated"
	"github.com/autom8ter/meshpaas/gen/gql/go/model"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func (r *mutationResolver) CreateAPI(ctx context.Context, input model.APIInput) (*model.API, error) {
	app, err := r.client.CreateAPI(ctx, toAPI(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromAPI(app), nil
}

func (r *mutationResolver) UpdateAPI(ctx context.Context, input model.APIInput) (*model.API, error) {
	app, err := r.client.UpdateAPI(ctx, toAPI(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromAPI(app), nil
}

func (r *mutationResolver) DelAPI(ctx context.Context, input model.Ref) (*string, error) {
	_, err := r.client.DeleteAPI(ctx, &meshpaaspb.Ref{
		Name: input.Name,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *mutationResolver) CreateTask(ctx context.Context, input model.TaskInput) (*model.Task, error) {
	task, err := r.client.CreateTask(ctx, toTask(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromTask(task), nil
}

func (r *mutationResolver) UpdateTask(ctx context.Context, input model.TaskInput) (*model.Task, error) {
	task, err := r.client.UpdateTask(ctx, toTask(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromTask(task), nil
}

func (r *mutationResolver) DelTask(ctx context.Context, input model.Ref) (*string, error) {
	_, err := r.client.DeleteAPI(ctx, &meshpaaspb.Ref{
		Name: input.Name,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *mutationResolver) CreateGateway(ctx context.Context, input model.GatewayInput) (*model.Gateway, error) {
	gw, err := r.client.CreateGateway(ctx, toGateway(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromGateway(gw), nil
}

func (r *mutationResolver) UpdateGateway(ctx context.Context, input model.GatewayInput) (*model.Gateway, error) {
	gw, err := r.client.UpdateGateway(ctx, toGateway(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromGateway(gw), nil
}

func (r *mutationResolver) DelGateway(ctx context.Context, input model.Ref) (*string, error) {
	_, err := r.client.DeleteGateway(ctx, &meshpaaspb.Ref{
		Name: input.Name,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *mutationResolver) CreateSecret(ctx context.Context, input model.SecretInput) (*model.Secret, error) {
	s, err := r.client.CreateSecret(ctx, toSecret(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromSecret(s), nil
}

func (r *mutationResolver) UpdateSecret(ctx context.Context, input model.SecretInput) (*model.Secret, error) {
	s, err := r.client.UpdateSecret(ctx, toSecret(input))
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromSecret(s), nil
}

func (r *mutationResolver) DelSecret(ctx context.Context, input model.Ref) (*string, error) {
	_, err := r.client.DeleteSecret(ctx, &meshpaaspb.Ref{
		Name: input.Name,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return nil, nil
}

func (r *queryResolver) GetAPI(ctx context.Context, input model.Ref) (*model.API, error) {
	app, err := r.client.GetAPI(ctx, &meshpaaspb.Ref{
		Name: input.Name,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromAPI(app), nil
}

func (r *queryResolver) ListAPIs(ctx context.Context, input *string) ([]*model.API, error) {
	apps, err := r.client.ListAPIs(ctx, &empty.Empty{})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var toReturn []*model.API
	for _, a := range apps.GetApis() {
		toReturn = append(toReturn, fromAPI(a))
	}
	return toReturn, nil
}

func (r *queryResolver) GetTask(ctx context.Context, input model.Ref) (*model.Task, error) {
	app, err := r.client.GetTask(ctx, &meshpaaspb.Ref{
		Name: input.Name,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromTask(app), nil
}

func (r *queryResolver) ListTasks(ctx context.Context, input *string) ([]*model.Task, error) {
	apps, err := r.client.ListTasks(ctx, &empty.Empty{})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var toReturn []*model.Task
	for _, a := range apps.GetTasks() {
		toReturn = append(toReturn, fromTask(a))
	}
	return toReturn, nil
}

func (r *queryResolver) GetGateway(ctx context.Context, input model.Ref) (*model.Gateway, error) {
	gw, err := r.client.GetGateway(ctx, &meshpaaspb.Ref{
		Name: input.Name,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromGateway(gw), nil
}

func (r *queryResolver) ListGateways(ctx context.Context, input *string) ([]*model.Gateway, error) {
	gws, err := r.client.ListGateways(ctx, &empty.Empty{})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var toReturn []*model.Gateway
	for _, a := range gws.GetGateways() {
		toReturn = append(toReturn, fromGateway(a))
	}
	return toReturn, nil
}

func (r *queryResolver) GetSecret(ctx context.Context, input model.Ref) (*model.Secret, error) {
	s, err := r.client.GetSecret(ctx, &meshpaaspb.Ref{
		Name: input.Name,
	})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	return fromSecret(s), nil
}

func (r *queryResolver) ListSecrets(ctx context.Context, input *string) ([]*model.Secret, error) {
	ss, err := r.client.ListSecrets(ctx, &empty.Empty{})
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Path:    graphql.GetPath(ctx),
		}
	}
	var toReturn []*model.Secret
	for _, a := range ss.GetSecrets() {
		toReturn = append(toReturn, fromSecret(a))
	}
	return toReturn, nil
}

func (r *subscriptionResolver) StreamLogs(ctx context.Context, input model.LogOpts) (<-chan string, error) {
	var (
		since        int64
		tail         int64
		previous     bool
		shouldStream bool
	)
	if input.SinceSeconds != nil {
		since = int64(*input.SinceSeconds)
	}
	if input.TailLines != nil {
		tail = int64(*input.TailLines)
	}
	if input.Previous != nil {
		previous = *input.Previous
	}
	if input.Stream != nil {
		shouldStream = *input.Stream
	}
	stream, err := r.client.StreamLogs(ctx, &meshpaaspb.LogOpts{
		Name:         input.Name,
		Container:    input.Container,
		SinceSeconds: since,
		TailLines:    tail,
		Previous:     previous,
		Stream:       shouldStream,
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
