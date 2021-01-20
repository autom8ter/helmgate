package client

import (
	"context"
	"fmt"
	"github.com/autom8ter/kubego"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/logger"
	"github.com/graphikDB/generic"
	"github.com/graphikDB/trigger"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

type authCtx string

const (
	userInfo authCtx = "userinfo"
)

type Manager struct {
	kclient            *kubego.Kube
	iclient            *kubego.Istio
	jwtCache           generic.Cache
	logger             *logger.Logger
	rootUsers          []string
	requestAuthorizers []*trigger.Decision
	userInfoEndpoint   string
}

func New(kclient *kubego.Kube, iclient *kubego.Istio, logger *logger.Logger) *Manager {
	return &Manager{
		kclient:  kclient,
		iclient:  iclient,
		jwtCache: generic.NewCache(1 * time.Minute),
		logger:   logger,
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

func (r *Manager) GetUserInfo(ctx context.Context) map[string]interface{} {
	if ctx.Value(userInfo) == nil {
		return map[string]interface{}{}
	}
	return ctx.Value(userInfo).(map[string]interface{})
}

func (r *Manager) SetUserInfo(ctx context.Context, userInfoData map[string]interface{}) context.Context {
	return context.WithValue(ctx, userInfo, userInfoData)
}

func (r *Manager) GetJWTHash(hash string) (interface{}, bool) {
	return r.jwtCache.Get(hash)
}

func (r *Manager) SetJWTHash(hash string, userInfo map[string]interface{}) {
	r.jwtCache.Set(hash, userInfo, 1*time.Hour)
}
