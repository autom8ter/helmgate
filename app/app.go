package app

import (
	"context"
	"github.com/graphikDB/kdeploy/gen/gql/go/model"
	"github.com/graphikDB/kubego"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Manager struct {
	client *kubego.Client
}

func New(client *kubego.Client) *Manager {
	return &Manager{client: client}
}

func (m *Manager) Create(ctx context.Context, app model.AppInput) (*model.App, error) {
	kapp := &k8sApp{}
	namespace, err := m.client.Namespaces().Create(ctx, toNamespace(app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	if app.State != nil {
		statefulset, err := m.client.StatefulSets(app.Namespace).Create(ctx, toStatefulSet(app), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
		kapp.statefulset = statefulset
	} else {
		deployment, err := m.client.Deployments(app.Namespace).Create(ctx, toDeployment(app), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
		kapp.deployment = deployment
	}
	svc, err := m.client.Services(app.Namespace).Create(ctx, toService(app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.service = svc
	return kapp.toApp(), nil
}

func (m *Manager) Update(ctx context.Context, app model.AppInput) (*model.App, error) {
	kapp := &k8sApp{}
	namespace, err := m.client.Namespaces().Update(ctx, toNamespace(app), v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	if app.State != nil {
		statefulset, err := m.client.StatefulSets(app.Namespace).Update(ctx, toStatefulSet(app), v1.UpdateOptions{})
		if err != nil {
			return nil, err
		}
		kapp.statefulset = statefulset
	} else {
		deployment, err := m.client.Deployments(app.Namespace).Update(ctx, toDeployment(app), v1.UpdateOptions{})
		if err != nil {
			return nil, err
		}
		kapp.deployment = deployment
	}
	svc, err := m.client.Services(app.Namespace).Update(ctx, toService(app), v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.service = svc
	return kapp.toApp(), nil
}

func (m *Manager) Get(ctx context.Context, name, namespace string) (*model.App, error) {
	kapp := &k8sApp{}
	ns, err := m.client.Namespaces().Get(ctx, namespace, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = ns
	statefulset, err := m.client.StatefulSets(namespace).Get(ctx, name, v1.GetOptions{})
	if err == nil {
		kapp.statefulset = statefulset
	} else {
		deployment, err := m.client.Deployments(namespace).Get(ctx, name, v1.GetOptions{})
		if err == nil {
			kapp.deployment = deployment
		}
	}

	svc, err := m.client.Services(namespace).Get(ctx, name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.service = svc
	return kapp.toApp(), nil
}
