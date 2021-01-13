package app

import (
	"context"
	"github.com/autom8ter/kdeploy/gen/gql/go/model"
	"github.com/autom8ter/kdeploy/logger"
	"github.com/graphikDB/generic"
	"github.com/graphikDB/kubego"
	"go.uber.org/zap"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"time"
)

type authCtx string

const (
	userInfo authCtx = "userinfo"
)

type Manager struct {
	client   *kubego.Client
	jwtCache *generic.Cache
	logger   *logger.Logger
}

func New(client *kubego.Client, logger *logger.Logger) *Manager {
	return &Manager{
		client:   client,
		jwtCache: generic.NewCache(1 * time.Minute),
		logger:   logger,
	}
}

func (m *Manager) Create(ctx context.Context, app model.AppInput) (*model.App, error) {
	kapp := &k8sApp{}
	namespace, err := m.client.Namespaces().Create(ctx, toNamespace(app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	if app.State != nil {
		ss, err := toStatefulSet(app)
		if err != nil {
			return nil, err
		}
		statefulset, err := m.client.StatefulSets(app.Namespace).Create(ctx, ss, v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
		kapp.statefulset = statefulset
	} else {
		dep, err := toDeployment(app)
		if err != nil {
			return nil, err
		}
		deployment, err := m.client.Deployments(app.Namespace).Create(ctx, dep, v1.CreateOptions{})
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
		ss, err := toStatefulSet(app)
		if err != nil {
			return nil, err
		}
		statefulset, err := m.client.StatefulSets(app.Namespace).Update(ctx, ss, v1.UpdateOptions{})
		if err != nil {
			return nil, err
		}
		kapp.statefulset = statefulset
	} else {
		dep, err := toDeployment(app)
		if err != nil {
			return nil, err
		}
		deployment, err := m.client.Deployments(app.Namespace).Update(ctx, dep, v1.UpdateOptions{})
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

func (m *Manager) Delete(ctx context.Context, name, namespace string) error {
	if err := m.client.Services(namespace).Delete(ctx, name, v1.DeleteOptions{}); err != nil {
		m.logger.Error("failed to delete service",
			zap.Error(err),
			zap.String("name", name),
			zap.String("namespace", namespace),
		)

	}
	if err := m.client.StatefulSets(namespace).Delete(ctx, name, v1.DeleteOptions{}); err != nil {
		if err := m.client.Deployments(namespace).Delete(ctx, name, v1.DeleteOptions{}); err != nil {
			m.logger.Error("failed to delete deployment",
				zap.Error(err),
				zap.String("name", name),
				zap.String("namespace", namespace),
			)
		}
	}
	if err := m.client.Namespaces().Delete(ctx, namespace, v1.DeleteOptions{}); err != nil {
		m.logger.Error("failed to delete namespace",
			zap.Error(err),
			zap.String("name", name),
			zap.String("namespace", namespace),
		)
	}
	return nil
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
