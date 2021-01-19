package client

import (
	"context"
	kdeploypb "github.com/autom8ter/kdeploy/gen/grpc/go"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m *Manager) CreateGateway(ctx context.Context, gateway *kdeploypb.Gateway) (*kdeploypb.Gateway, error) {
	resp, err := m.iclient.Gateways(gateway.Namespace).Create(ctx, toGateway(gateway), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return &kdeploypb.Gateway{
		Name:      resp.Name,
		Namespace: resp.Namespace,
		Listeners: gateway.Listeners,
		Labels:    resp.Spec.Selector,
	}, nil
}
