package client

import (
	"context"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m *Manager) CreateProject(ctx context.Context, project *meshpaaspb.ProjectInput) (*meshpaaspb.Project, error) {
	namespace, err := m.kclient.Namespaces().Create(ctx, toNamespace(project), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	return &meshpaaspb.Project{
		Name: namespace.Name,
	}, nil
}

func (m *Manager) GetProject(ctx context.Context, project *meshpaaspb.ProjectRef) (*meshpaaspb.Project, error) {
	namespace, err := m.kclient.Namespaces().Get(ctx, project.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	return &meshpaaspb.Project{
		Name: namespace.Name,
	}, nil
}

func (m *Manager) UpdateProject(ctx context.Context, project *meshpaaspb.ProjectInput) (*meshpaaspb.Project, error) {
	namespace, err := m.kclient.Namespaces().Get(ctx, project.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	namespace, err = m.kclient.Namespaces().Update(ctx, namespace, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	return &meshpaaspb.Project{
		Name: namespace.Name,
	}, nil
}

func (m *Manager) DeleteProject(ctx context.Context, ref *meshpaaspb.ProjectRef) error {
	if err := m.kclient.Namespaces().Delete(ctx, ref.GetName(), v1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func (m *Manager) ListProjects(ctx context.Context) (*meshpaaspb.Projects, error) {
	namespaces, err := m.kclient.Namespaces().List(ctx, v1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, err
	}
	var ns = &meshpaaspb.Projects{}
	for _, n := range namespaces.Items {
		ns.Projects = append(ns.Projects, n.Name)
	}
	return ns, nil
}
