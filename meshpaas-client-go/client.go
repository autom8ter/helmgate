package meshpaas_client_go

import (
	"context"
	"fmt"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/logger"
	grpc_zap "github.com/grpc-ecosystem/go-grpc-middleware/logging/zap"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// Options holds configuration options
type Options struct {
	tokenSource oauth2.TokenSource
	metrics     bool
	logging     bool
	logPayload  bool
	creds       credentials.TransportCredentials
}

// Opt is a single configuration option
type Opt func(o *Options)

// WithTransportCreds adds transport credentials to the client
func WithTransportCreds(creds credentials.TransportCredentials) Opt {
	return func(o *Options) {
		o.creds = creds
	}
}

// WithTokenSource uses oauth token add an authorization header to every outbound request
func WithTokenSource(tokenSource oauth2.TokenSource) Opt {
	return func(o *Options) {
		o.tokenSource = tokenSource
	}
}

// WithMetrics registers prometheus metrics
func WithMetrics(metrics bool) Opt {
	return func(o *Options) {
		o.metrics = metrics
	}
}

// WithLogging registers a logging middleware
func WithLogging(logging, logPayload bool) Opt {
	return func(o *Options) {
		o.logging = logging
		o.logPayload = logPayload
	}
}

func unaryAuth(tokenSource oauth2.TokenSource) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		ctx, err := toContext(ctx, tokenSource)
		if err != nil {
			return err
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func streamAuth(tokenSource oauth2.TokenSource) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		ctx, err := toContext(ctx, tokenSource)
		if err != nil {
			return nil, err
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}

// NewClient creates a new gRPC meshpaas client
func NewClient(ctx context.Context, target string, opts ...Opt) (*Client, error) {
	if target == "" {
		return nil, errors.New("empty target")
	}
	dialopts := []grpc.DialOption{}
	var uinterceptors []grpc.UnaryClientInterceptor
	var sinterceptors []grpc.StreamClientInterceptor
	options := &Options{}
	for _, o := range opts {
		o(options)
	}
	if options.creds == nil {
		dialopts = append(dialopts, grpc.WithInsecure())
	} else {
		dialopts = append(dialopts, grpc.WithTransportCredentials(options.creds))
	}
	uinterceptors = append(uinterceptors, grpc_validator.UnaryClientInterceptor())

	if options.metrics {
		uinterceptors = append(uinterceptors, grpc_prometheus.UnaryClientInterceptor)
		sinterceptors = append(sinterceptors, grpc_prometheus.StreamClientInterceptor)
	}

	if options.tokenSource != nil {
		uinterceptors = append(uinterceptors, unaryAuth(options.tokenSource))
		sinterceptors = append(sinterceptors, streamAuth(options.tokenSource))
	}
	if options.logging {
		lgger := logger.New(true, zap.Bool("client", true))
		uinterceptors = append(uinterceptors, grpc_zap.UnaryClientInterceptor(lgger.Zap()))
		sinterceptors = append(sinterceptors, grpc_zap.StreamClientInterceptor(lgger.Zap()))

		if options.logPayload {
			uinterceptors = append(uinterceptors, grpc_zap.PayloadUnaryClientInterceptor(lgger.Zap(), func(ctx context.Context, fullMethodName string) bool {
				return true
			}))
			sinterceptors = append(sinterceptors, grpc_zap.PayloadStreamClientInterceptor(lgger.Zap(), func(ctx context.Context, fullMethodName string) bool {
				return true
			}))
		}
	}
	dialopts = append(dialopts,
		grpc.WithChainUnaryInterceptor(uinterceptors...),
		grpc.WithChainStreamInterceptor(sinterceptors...),
		grpc.WithBlock(),
	)
	conn, err := grpc.DialContext(ctx, target, dialopts...)
	if err != nil {
		return nil, err
	}
	return &Client{
		client: meshpaaspb.NewKdeployServiceClient(conn),
	}, nil
}

// Client is a meshpaas gRPC client
type Client struct {
	client meshpaaspb.KdeployServiceClient
}

func toContext(ctx context.Context, tokenSource oauth2.TokenSource) (context.Context, error) {
	token, err := tokenSource.Token()
	if err != nil {
		return ctx, errors.Wrap(err, "failed to get token")
	}
	return metadata.AppendToOutgoingContext(
		ctx,
		"Authorization", fmt.Sprintf("Bearer %v", token.AccessToken),
	), nil
}

// CreateApp creates a new application
func (c *Client) CreateApp(ctx context.Context, app *meshpaaspb.AppInput) (*meshpaaspb.App, error) {
	return c.client.CreateApp(ctx, app)
}

// UpdateApp updates an application - it performs a full replace
func (c *Client) UpdateApp(ctx context.Context, app *meshpaaspb.AppInput) (*meshpaaspb.App, error) {
	return c.client.UpdateApp(ctx, app)
}

// DeleteApp deletes an application by reference(name/namespace)
func (c *Client) DeleteApp(ctx context.Context, ref *meshpaaspb.Ref) error {
	_, err := c.client.DeleteApp(ctx, ref)
	return err
}

// GetApp get an application by reference(name/namespace)
func (c *Client) GetApp(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.App, error) {
	return c.client.GetApp(ctx, ref)
}

// DeleteApp deletes all tasks/applications in the namespace
func (c *Client) DeleteAll(ctx context.Context, namespace *meshpaaspb.Namespace) error {
	_, err := c.client.DeleteAll(ctx, namespace)
	return err
}

// CreateTask creates a new task
func (c *Client) CreateTask(ctx context.Context, app *meshpaaspb.TaskInput) (*meshpaaspb.Task, error) {
	return c.client.CreateTask(ctx, app)
}

// UpdateTask updates a task - it performs a full replace
func (c *Client) UpdateTask(ctx context.Context, app *meshpaaspb.TaskInput) (*meshpaaspb.Task, error) {
	return c.client.UpdateTask(ctx, app)
}

// DeleteTask deletes a task by reference(name/namespace)
func (c *Client) DeleteTask(ctx context.Context, ref *meshpaaspb.Ref) error {
	_, err := c.client.DeleteTask(ctx, ref)
	return err
}

// GetTask gets a task by reference(name/namespace)
func (c *Client) GetTask(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.Task, error) {
	return c.client.GetTask(ctx, ref)
}

// StreamLogs streams logs from an application until the context cancelled or the function(fn) return false
func (c *Client) StreamLogs(ctx context.Context, ref *meshpaaspb.Ref, fn func(l *meshpaaspb.Log) bool) error {
	stream, err := c.client.StreamLogs(ctx, ref)
	if err != nil {
		return err
	}
	for {
		select {
		case <-ctx.Done():
			return nil
		default:
			msg, err := stream.Recv()
			if err != nil {
				return err
			}
			if !fn(msg) {
				return nil
			}
		}
	}
}


// CreateGateway creates a new gateway
func (c *Client) CreateGateway(ctx context.Context, app *meshpaaspb.GatewayInput) (*meshpaaspb.Gateway, error) {
	return c.client.CreateGateway(ctx, app)
}

// UpdateGateway updates a gateway - it performs a full replace
func (c *Client) UpdateGateway(ctx context.Context, app *meshpaaspb.GatewayInput) (*meshpaaspb.Gateway, error) {
	return c.client.UpdateGateway(ctx, app)
}

// DeleteGateway deletes a gateway by reference(name/namespace)
func (c *Client) DeleteGateway(ctx context.Context, ref *meshpaaspb.Ref) error {
	_, err := c.client.DeleteGateway(ctx, ref)
	return err
}

// GetGateway gets a gateway by reference(name/namespace)
func (c *Client) GetGateway(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.Gateway, error) {
	return c.client.GetGateway(ctx, ref)
}