package client

import (
	"context"
	"fmt"
	"github.com/autom8ter/kubego"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/logger"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Manager struct {
	kclient *kubego.Kube
	iclient *kubego.Istio
	logger  *logger.Logger
}

func New(kclient *kubego.Kube, iclient *kubego.Istio, logger *logger.Logger) *Manager {
	return &Manager{
		kclient: kclient,
		iclient: iclient,
		logger:  logger,
	}
}

func (m *Manager) L() *logger.Logger {
	return m.logger
}

func (m *Manager) getStatus(ctx context.Context, namespace, name string) (*meshpaaspb.AppStatus, error) {
	var replicas []*meshpaaspb.Replica
	pods, err := m.kclient.Pods(namespace).List(ctx, v1.ListOptions{
		TypeMeta:      v1.TypeMeta{},
		LabelSelector: fmt.Sprintf("meshpaas.app = %s", name),
	})
	if err != nil {
		return nil, err
	}
	for _, pod := range pods.Items {
		replicas = append(replicas, &meshpaaspb.Replica{
			Phase:     string(pod.Status.Phase),
			Condition: string(pod.Status.Conditions[0].Status),
			Reason:    pod.Status.Reason,
		})
	}
	return &meshpaaspb.AppStatus{Replicas: replicas}, nil
}
