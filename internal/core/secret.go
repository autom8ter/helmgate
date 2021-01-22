package core

import (
	"context"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/auth"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	apiv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m *Manager) CreateSecret(ctx context.Context, secret *meshpaaspb.SecretInput) (*meshpaaspb.Secret, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	k8s := &k8sSecret{}
	namespace, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr[m.namespaceClaim]), apiv1.GetOptions{})
	if err != nil {
		return nil, err
	}
	k8s.namespace = namespace
	resp, err := m.kclient.Secrets(cast.ToString(usr[m.namespaceClaim])).Create(ctx, m.toSecret(usr, secret), apiv1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	k8s.secret = resp
	return k8s.toSecret(), nil
}

func (m *Manager) DeleteSecret(ctx context.Context, ref *meshpaaspb.Ref) error {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	return m.kclient.Secrets(cast.ToString(usr[m.namespaceClaim])).Delete(ctx, ref.GetName(), apiv1.DeleteOptions{})
}

func (m *Manager) GetSecret(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.Secret, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	k8s := &k8sSecret{}
	namespace, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr[m.namespaceClaim]), apiv1.GetOptions{})
	if err != nil {
		return nil, err
	}
	k8s.namespace = namespace
	resp, err := m.kclient.Secrets(cast.ToString(usr[m.namespaceClaim])).Get(ctx, ref.Name, apiv1.GetOptions{})
	if err != nil {
		return nil, err
	}
	k8s.secret = resp
	return k8s.toSecret(), nil
}

func (m *Manager) UpdateSecret(ctx context.Context, secret *meshpaaspb.SecretInput) (*meshpaaspb.Secret, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	kapp := &k8sSecret{}
	namespace, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr[m.namespaceClaim]), apiv1.GetOptions{})
	if err != nil {
		return nil, err
	}

	kapp.namespace = namespace
	ksecret, err := m.kclient.Secrets(cast.ToString(usr[m.namespaceClaim])).Get(ctx, secret.Name, apiv1.GetOptions{})
	if err != nil {
		return nil, err
	}
	ksecret = overwriteSecret(ksecret, secret)
	ksecret, err = m.kclient.Secrets(cast.ToString(usr[m.namespaceClaim])).Update(ctx, ksecret, apiv1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.secret = ksecret
	return kapp.toSecret(), nil
}
