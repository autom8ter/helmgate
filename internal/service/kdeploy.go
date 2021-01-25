package service

import (
	"context"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/core"
	"github.com/golang/protobuf/ptypes/empty"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MeshPaasService struct {
	client *core.Manager
}

func NewMeshPaasService(client *core.Manager) *MeshPaasService {
	return &MeshPaasService{client: client}
}

func (k MeshPaasService) CreateAPI(ctx context.Context, constructor *meshpaaspb.APIInput) (*meshpaaspb.API, error) {
	return k.client.CreateAPI(ctx, constructor)
}

func (k MeshPaasService) UpdateAPI(ctx context.Context, update *meshpaaspb.APIInput) (*meshpaaspb.API, error) {
	return k.client.UpdateAPI(ctx, update)
}

func (k MeshPaasService) DeleteAPI(ctx context.Context, ref *meshpaaspb.Ref) (*empty.Empty, error) {
	if err := k.client.DeleteAPI(ctx, ref); err != nil {
		k.client.L().Error("failed to delete app", zap.Error(err))
		return nil, status.Error(codes.Internal, "failed to delete app")
	}
	return &empty.Empty{}, nil
}

func (k MeshPaasService) GetAPI(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.API, error) {
	return k.client.GetAPI(ctx, ref)
}

func (k MeshPaasService) ListAPIs(ctx context.Context, e *empty.Empty) (*meshpaaspb.APIs, error) {
	return k.client.ListAPIs(ctx)
}

func (k MeshPaasService) ListSecrets(ctx context.Context, e *empty.Empty) (*meshpaaspb.Secrets, error) {
	return k.client.ListSecrets(ctx)
}

func (k MeshPaasService) ListGateways(ctx context.Context, e *empty.Empty) (*meshpaaspb.Gateways, error) {
	return k.client.ListGateways(ctx)
}

func (k MeshPaasService) StreamLogs(opts *meshpaaspb.LogOpts, server meshpaaspb.MeshPaasService_StreamLogsServer) error {
	stream, err := k.client.StreamLogs(server.Context(), opts)
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

func (k MeshPaasService) CreateTask(ctx context.Context, constructor *meshpaaspb.TaskInput) (*meshpaaspb.Task, error) {
	return k.client.CreateTask(ctx, constructor)
}

func (k MeshPaasService) UpdateTask(ctx context.Context, update *meshpaaspb.TaskInput) (*meshpaaspb.Task, error) {
	return k.client.UpdateTask(ctx, update)
}

func (k MeshPaasService) DeleteTask(ctx context.Context, ref *meshpaaspb.Ref) (*empty.Empty, error) {
	return &empty.Empty{}, k.client.DeleteTask(ctx, ref)
}

func (k MeshPaasService) GetTask(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.Task, error) {
	return k.client.GetTask(ctx, ref)
}

func (k MeshPaasService) ListTasks(ctx context.Context, _ *empty.Empty) (*meshpaaspb.Tasks, error) {
	return k.client.ListTasks(ctx)
}

func (k MeshPaasService) CreateGateway(ctx context.Context, gateway *meshpaaspb.GatewayInput) (*meshpaaspb.Gateway, error) {
	return k.client.CreateGateway(ctx, gateway)
}

func (k MeshPaasService) UpdateGateway(ctx context.Context, gateway *meshpaaspb.GatewayInput) (*meshpaaspb.Gateway, error) {
	return k.client.UpdateGateway(ctx, gateway)
}

func (k MeshPaasService) DeleteGateway(ctx context.Context, ref *meshpaaspb.Ref) (*empty.Empty, error) {
	return &empty.Empty{}, k.client.DeleteGateway(ctx, ref)
}

func (k MeshPaasService) GetGateway(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.Gateway, error) {
	return k.client.GetGateway(ctx, ref)
}

func (k MeshPaasService) CreateSecret(ctx context.Context, input *meshpaaspb.SecretInput) (*meshpaaspb.Secret, error) {
	return k.client.CreateSecret(ctx, input)
}

func (k MeshPaasService) UpdateSecret(ctx context.Context, input *meshpaaspb.SecretInput) (*meshpaaspb.Secret, error) {
	return k.client.UpdateSecret(ctx, input)
}

func (k MeshPaasService) DeleteSecret(ctx context.Context, ref *meshpaaspb.Ref) (*empty.Empty, error) {
	return &empty.Empty{}, k.client.DeleteSecret(ctx, ref)
}

func (k MeshPaasService) GetSecret(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.Secret, error) {
	return k.client.GetSecret(ctx, ref)
}
