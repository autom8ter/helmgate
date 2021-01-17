package service

import (
	"context"
	"github.com/autom8ter/kdeploy/client"
	kdeploypb "github.com/autom8ter/kdeploy/gen/grpc/go"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type KdeployService struct {
	client *client.Manager
}

func NewKdeployService(client *client.Manager) *KdeployService {
	return &KdeployService{client: client}
}

func (k KdeployService) CreateApp(ctx context.Context, constructor *kdeploypb.AppConstructor) (*kdeploypb.App, error) {
	return k.client.Create(ctx, constructor)
}

func (k KdeployService) UpdateApp(ctx context.Context, update *kdeploypb.AppUpdate) (*kdeploypb.App, error) {
	return k.client.Update(ctx, update)
}

func (k KdeployService) DeleteApp(ctx context.Context, ref *kdeploypb.AppRef) (*empty.Empty, error) {
	if err := k.client.Delete(ctx, ref); err != nil {
		k.client.L().Error("failed to delete app", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete app")
	}
	return &empty.Empty{}, nil
}

func (k KdeployService) GetApp(ctx context.Context, ref *kdeploypb.AppRef) (*kdeploypb.App, error) {
	return k.client.Get(ctx, ref)
}

func (k KdeployService) ListNamespaces(ctx context.Context, _ *empty.Empty) (*kdeploypb.Namespaces, error) {
	return k.client.ListNamespaces(ctx)
}

func (k KdeployService) Logs(ref *kdeploypb.AppRef, server kdeploypb.KdeployService_LogsServer) error {
	stream, err := k.client.StreamLogs(server.Context(), ref)
	if err != nil {
		k.client.L().Error("failed to stream logs", zap.Error(err))
		return status.Error(codes.Internal, "failed to stream logs")
	}
	for {
		select {
		case <-server.Context().Done():
			close(stream)

		case log := <-stream:
			if err := server.Send(&kdeploypb.Log{Message: log}); err != nil {
				k.client.L().Error("failed to stream log", zap.Error(err))
				return status.Error(codes.Internal, "failed to stream log")
			}
		}
	}
}
