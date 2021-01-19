package client

import (
	"context"
	kdeploypb "github.com/autom8ter/kdeploy/gen/grpc/go"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m *Manager) CreateGateway(ctx context.Context, gateway *kdeploypb.Gateway) (*kdeploypb.Gateway, error) {
	kapp := &k8sGateway{}
	namespace, err := m.kclient.Namespaces().Get(ctx, gateway.Namespace, v1.GetOptions{})
	if err != nil {
		namespace, err = m.kclient.Namespaces().Create(ctx, toGwNamespace(gateway), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	kapp.namespace = namespace
	resp, err := m.iclient.Gateways(gateway.Namespace).Create(ctx, toGateway(gateway), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.gateway = resp

	return kapp.toGateway(), nil
}

func (m *Manager) UpdateGateway(ctx context.Context, gateway *kdeploypb.Gateway) (*kdeploypb.Gateway, error) {
	kapp := &k8sGateway{}
	namespace, err := m.kclient.Namespaces().Get(ctx, gateway.Namespace, v1.GetOptions{})
	if err != nil {
		namespace, err = m.kclient.Namespaces().Create(ctx, toGwNamespace(gateway), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	kapp.namespace = namespace
	gw, err := m.iclient.Gateways(gateway.Namespace).Get(ctx, gateway.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	gw = overwriteGateway(gw, gateway)
	gw, err = m.iclient.Gateways(gateway.Namespace).Update(ctx, gw, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.gateway = gw
	return kapp.toGateway(), nil
}

func (m *Manager) DeleteGateway(ctx context.Context, ref *kdeploypb.Ref) error {
	err := m.iclient.Gateways(ref.GetNamespace()).Delete(ctx, ref.Name, v1.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}

func (m *Manager) GetGateway(ctx context.Context, ref *kdeploypb.Ref) (*kdeploypb.Gateway, error) {
	kapp := &k8sGateway{}
	namespace, err := m.kclient.Namespaces().Get(ctx, ref.Namespace, v1.GetOptions{})
	if err != nil {
		return nil, err

	}
	kapp.namespace = namespace
	gw, err := m.iclient.Gateways(ref.Namespace).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.gateway = gw
	return kapp.toGateway(), nil
}
