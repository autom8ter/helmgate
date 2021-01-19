package client

import (
	"context"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	apiv1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m *Manager) CreateSecret(ctx context.Context, secret *meshpaaspb.SecretInput) (*meshpaaspb.Secret, error) {
	k8s := &k8sSecret{}
	namespace, err := m.kclient.Namespaces().Get(ctx, secret.Project, apiv1.GetOptions{})
	if err != nil {
		namespace, err = m.kclient.Namespaces().Create(ctx, toSecretNamespace(secret), apiv1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	k8s.namespace = namespace
	resp, err := m.kclient.Secrets(secret.GetProject()).Create(ctx, toSecret(secret), apiv1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	k8s.secret = resp
	return k8s.toSecret(), nil
}

func (m *Manager) DeleteSecret(ctx context.Context, ref *meshpaaspb.Ref) error {
	return m.kclient.Secrets(ref.GetProject()).Delete(ctx, ref.GetName(), apiv1.DeleteOptions{})
}

func (m *Manager) GetSecret(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.Secret, error) {
	k8s := &k8sSecret{}
	namespace, err := m.kclient.Namespaces().Get(ctx, ref.Project, apiv1.GetOptions{})
	if err != nil {
		return nil, err
	}
	k8s.namespace = namespace
	resp, err := m.kclient.Secrets(ref.GetProject()).Get(ctx, ref.Name, apiv1.GetOptions{})
	if err != nil {
		return nil, err
	}
	k8s.secret = resp
	return k8s.toSecret(), nil
}

func (m *Manager) UpdateSecret(ctx context.Context, secret *meshpaaspb.SecretInput) (*meshpaaspb.Secret, error) {
	kapp := &k8sSecret{}
	namespace, err := m.kclient.Namespaces().Get(ctx, secret.Project, apiv1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	ksecret, err := m.kclient.Secrets(secret.Project).Get(ctx, secret.Name, apiv1.GetOptions{})
	if err != nil {
		return nil, err
	}
	ksecret = overwriteSecret(ksecret, secret)
	ksecret, err = m.kclient.Secrets(secret.Project).Update(ctx, ksecret, apiv1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.secret = ksecret
	return kapp.toSecret(), nil
}
