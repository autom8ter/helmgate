package service

import (
	"context"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/client"
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

func (k KdeployService) CreateApp(ctx context.Context, constructor *meshpaaspb.AppInput) (*meshpaaspb.App, error) {
	return k.client.CreateApp(ctx, constructor)
}

func (k KdeployService) UpdateApp(ctx context.Context, update *meshpaaspb.AppInput) (*meshpaaspb.App, error) {
	return k.client.UpdateApp(ctx, update)
}

func (k KdeployService) DeleteApp(ctx context.Context, ref *meshpaaspb.Ref) (*empty.Empty, error) {
	if err := k.client.DeleteApp(ctx, ref); err != nil {
		k.client.L().Error("failed to delete app", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete app")
	}
	return &empty.Empty{}, nil
}

func (k KdeployService) GetApp(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.App, error) {
	return k.client.GetApp(ctx, ref)
}

func (k KdeployService) ListNamespaces(ctx context.Context, _ *empty.Empty) (*meshpaaspb.Namespaces, error) {
	return k.client.ListNamespaces(ctx)
}

func (k KdeployService) ListApps(ctx context.Context, ns *meshpaaspb.Namespace) (*meshpaaspb.Apps, error) {
	return k.client.ListApps(ctx, ns)
}

func (k KdeployService) DeleteAll(ctx context.Context, ns *meshpaaspb.Namespace) (*empty.Empty, error) {
	if err := k.client.DeleteAll(ctx, ns); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (k KdeployService) StreamLogs(ref *meshpaaspb.Ref, server meshpaaspb.KdeployService_StreamLogsServer) error {
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
			if err := server.Send(&meshpaaspb.Log{Message: log}); err != nil {
				k.client.L().Error("failed to stream log", zap.Error(err))
				return status.Error(codes.Internal, "failed to stream log")
			}
		}
	}
}

func (k KdeployService) CreateTask(ctx context.Context, constructor *meshpaaspb.TaskInput) (*meshpaaspb.Task, error) {
	return k.client.CreateTask(ctx, constructor)
}

func (k KdeployService) UpdateTask(ctx context.Context, update *meshpaaspb.TaskInput) (*meshpaaspb.Task, error) {
	return k.client.UpdateTask(ctx, update)
}

func (k KdeployService) DeleteTask(ctx context.Context, ref *meshpaaspb.Ref) (*empty.Empty, error) {
	return &empty.Empty{}, k.client.DeleteTask(ctx, ref)
}

func (k KdeployService) GetTask(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.Task, error) {
	return k.client.GetTask(ctx, ref)
}

func (k KdeployService) ListTasks(ctx context.Context, ns *meshpaaspb.Namespace) (*meshpaaspb.Tasks, error) {
	return k.client.ListTasks(ctx, ns)
}

func (k KdeployService) CreateGateway(ctx context.Context, gateway *meshpaaspb.GatewayInput) (*meshpaaspb.Gateway, error) {
	return k.client.CreateGateway(ctx, gateway)
}

func (k KdeployService) UpdateGateway(ctx context.Context, gateway *meshpaaspb.GatewayInput) (*meshpaaspb.Gateway, error) {
	return k.client.UpdateGateway(ctx, gateway)
}

func (k KdeployService) DeleteGateway(ctx context.Context, ref *meshpaaspb.Ref) (*empty.Empty, error) {
	return &empty.Empty{}, k.client.DeleteGateway(ctx, ref)
}

func (k KdeployService) GetGateway(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.Gateway, error) {
	return k.client.GetGateway(ctx, ref)

}
