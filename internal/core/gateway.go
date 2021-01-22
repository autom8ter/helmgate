package core

import (
	"context"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/auth"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m *Manager) CreateGateway(ctx context.Context, gateway *meshpaaspb.GatewayInput) (*meshpaaspb.Gateway, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	kapp := &k8sGateway{}
	namespace, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr[m.namespaceClaim]), v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	resp, err := m.iclient.Gateways(cast.ToString(usr[m.namespaceClaim])).Create(ctx, m.toGateway(usr, gateway), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.gateway = resp

	return kapp.toGateway(), nil
}

func (m *Manager) UpdateGateway(ctx context.Context, gateway *meshpaaspb.GatewayInput) (*meshpaaspb.Gateway, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	kapp := &k8sGateway{}
	namespace, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr[m.namespaceClaim]), v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	gw, err := m.iclient.Gateways(cast.ToString(usr[m.namespaceClaim])).Get(ctx, gateway.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	gw = overwriteGateway(gw, gateway)
	gw, err = m.iclient.Gateways(cast.ToString(usr[m.namespaceClaim])).Update(ctx, gw, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.gateway = gw
	return kapp.toGateway(), nil
}

func (m *Manager) DeleteGateway(ctx context.Context, ref *meshpaaspb.Ref) error {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	_, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr[m.namespaceClaim]), v1.GetOptions{})
	if err != nil {
		return err
	}
	err = m.iclient.Gateways(cast.ToString(usr[m.namespaceClaim])).Delete(ctx, ref.Name, v1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) GetGateway(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.Gateway, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	kapp := &k8sGateway{}
	namespace, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr[m.namespaceClaim]), v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	gw, err := m.iclient.Gateways(cast.ToString(usr[m.namespaceClaim])).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.gateway = gw
	return kapp.toGateway(), nil
}
