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

	_, err := m.client.Deployments(app.Namespace).Create(ctx, toDeployment(app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *Manager) Update(ctx context.Context, app model.AppInput) (*model.App, error) {
	_, err := m.client.Deployments(app.Namespace).Update(ctx, toDeployment(app), v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *Manager) Get(ctx context.Context, namespace string) (*model.App, error) {

	return nil, nil
}
