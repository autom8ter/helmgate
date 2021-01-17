package kdeploy_client_go

import (
	"context"
	"fmt"
	kdeploypb "github.com/autom8ter/kdeploy/gen/grpc/go"
	"github.com/autom8ter/kdeploy/logger"
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

type Options struct {
	tokenSource oauth2.TokenSource
	metrics     bool
	logging     bool
	logPayload  bool
	creds       credentials.TransportCredentials
}

type Opt func(o *Options)

func WithTransportCreds(creds credentials.TransportCredentials) Opt {
	return func(o *Options) {
		o.creds = creds
	}
}

func WithTokenSource(tokenSource oauth2.TokenSource) Opt {
	return func(o *Options) {
		o.tokenSource = tokenSource
	}
}

func WithMetrics(metrics bool) Opt {
	return func(o *Options) {
		o.metrics = metrics
	}
}

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

// NewClient creates a new gRPC kdeploy client
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
		client: kdeploypb.NewKdeployServiceClient(conn),
	}, nil
}

// Client is a kdeploy gRPC client
type Client struct {
	client kdeploypb.KdeployServiceClient
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
func (c *Client) CreateApp(ctx context.Context, app *kdeploypb.AppConstructor) (*kdeploypb.App, error) {
	return c.client.CreateApp(ctx, app)
}

// UpdateApp updates an application - it performs a patch, not a full replace
func (c *Client) UpdateApp(ctx context.Context, app *kdeploypb.AppUpdate) (*kdeploypb.App, error) {
	return c.client.UpdateApp(ctx, app)
}

// DeleteApp deletes an application by reference(name/namespace)
func (c *Client) DeleteApp(ctx context.Context, ref *kdeploypb.AppRef) error {
	_, err := c.client.DeleteApp(ctx, ref)
	return err
}

// GetApp get an application by reference(name/namespace)
func (c *Client) GetApp(ctx context.Context, ref *kdeploypb.AppRef) (*kdeploypb.App, error) {
	return c.client.GetApp(ctx, ref)
}

// Logs streams logs from an application until the context cancelled or the function(fn) return false
func (c *Client) Logs(ctx context.Context, ref *kdeploypb.AppRef, fn func(l *kdeploypb.Log) bool) error {
	stream, err := c.client.Logs(ctx, ref)
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
