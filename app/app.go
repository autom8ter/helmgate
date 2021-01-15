package app

import (
	"bytes"
	"context"
	"github.com/autom8ter/kdeploy/gen/gql/go/model"
	"github.com/autom8ter/kdeploy/logger"
	"github.com/autom8ter/kubego"
	"github.com/graphikDB/generic"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sync"
	"time"
)

type authCtx string

const (
	userInfo authCtx = "userinfo"
)

type Manager struct {
	client   *kubego.Client
	jwtCache generic.Cache
	logger   *logger.Logger
}

func New(client *kubego.Client, logger *logger.Logger) *Manager {
	return &Manager{
		client:   client,
		jwtCache: generic.NewCache(1 * time.Minute),
		logger:   logger,
	}
}

func (m *Manager) getStatus(ctx context.Context, namespace string) (*model.Status, error) {
	var replicas []*model.Replica
	pods, err := m.client.Pods(namespace).List(ctx, v1.ListOptions{})
	if err != nil {
		return nil, err
	}
	for _, pod := range pods.Items {
		replicas = append(replicas, &model.Replica{
			Phase:     string(pod.Status.Phase),
			Condition: string(pod.Status.Conditions[0].Status),
			Reason:    pod.Status.Reason,
		})
	}
	return &model.Status{Replicas: replicas}, nil
}

func (m *Manager) Create(ctx context.Context, app model.AppInput) (*model.App, error) {
	kapp := &k8sApp{}
	namespace, err := m.client.Namespaces().Create(ctx, toNamespace(app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	dep, err := toDeployment(app)
	if err != nil {
		return nil, err
	}
	deployment, err := m.client.Deployments(app.Namespace).Create(ctx, dep, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.deployment = deployment
	svc, err := m.client.Services(app.Namespace).Create(ctx, toService(app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.service = svc
	status, err := m.getStatus(ctx, app.Namespace)
	if err != nil {
		return nil, err
	}
	a := kapp.toApp()
	a.Status = status
	return a, nil
}

func (m *Manager) Update(ctx context.Context, app model.AppUpdate) (*model.App, error) {
	kapp := &k8sApp{}
	namespace, err := m.client.Namespaces().Get(ctx, app.Namespace, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	deployment, err := m.client.Deployments(app.Namespace).Get(ctx, app.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	deployment, err = overwriteDeployment(deployment, app)
	if err != nil {
		return nil, err
	}
	deployment, err = m.client.Deployments(app.Namespace).Update(ctx, deployment, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.deployment = deployment
	svc, err := m.client.Services(app.Namespace).Get(ctx, app.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	svc = overwriteService(svc, app)
	svc, err = m.client.Services(app.Namespace).Update(ctx, svc, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.service = svc
	status, err := m.getStatus(ctx, app.Namespace)
	if err != nil {
		return nil, err
	}
	a := kapp.toApp()
	a.Status = status
	return a, nil
}

func (m *Manager) Get(ctx context.Context, name, namespace string) (*model.App, error) {
	kapp := &k8sApp{}

	ns, err := m.client.Namespaces().Get(ctx, namespace, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = ns
	deployment, err := m.client.Deployments(namespace).Get(ctx, name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.deployment = deployment
	svc, err := m.client.Services(namespace).Get(ctx, name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.service = svc
	status, err := m.getStatus(ctx, namespace)
	if err != nil {
		return nil, err
	}
	a := kapp.toApp()
	a.Status = status
	return a, nil
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

func (m *Manager) StreamLogs(ctx context.Context, name, namespace string) (<-chan string, error) {
	pods, err := m.client.Pods(namespace).List(ctx, v1.ListOptions{
		TypeMeta: v1.TypeMeta{},
		Watch:    false,
	})
	if err != nil {
		return nil, err
	}
	if len(pods.Items) == 0 {
		return nil, errors.New("zero pods")
	}
	m.logger.Info("setup log stream",
		zap.String("name", name),
		zap.String("namespace", namespace),
	)
	logChan := make(chan string)
	var streamMu = sync.RWMutex{}
	//mach.Go(func(routine machine.Routine) {
	//	for {
	//		select {
	//		case <-routine.Context().Done():
	//			close(logs)
	//			return
	//		}
	//	}
	//})
	for _, pod := range pods.Items {
		go func(p corev1.Pod) {
			m.logger.Debug("setup log stream",
				zap.String("name", name),
				zap.String("namespace", namespace),
				zap.String("pod", p.Name),
			)
			closer, err := m.client.GetLogs(context.Background(), p.Name, p.Namespace, &corev1.PodLogOptions{
				TypeMeta:  v1.TypeMeta{},
				Container: name,
				//Container: name,
				Follow:                       true,
				Previous:                     false,
				Timestamps:                   true,
				InsecureSkipTLSVerifyBackend: true,
			})
			defer closer.Close()
			if err != nil {
				m.logger.Error("failed to stream pod logs",
					zap.Error(err),
					zap.String("name", name),
					zap.String("namespace", namespace),
					zap.String("pod", p.Name),
				)
				return
			}
			for {
				m.logger.Debug("streaming log",
					zap.String("name", name),
					zap.String("namespace", namespace),
					zap.String("pod", p.Name),
				)
				buf := make([]byte, 1024)
				_, err := closer.Read(buf)
				if err != nil {
					if err == io.EOF {
						return
					}
				}
				streamMu.Lock()
				logChan <- string(bytes.Trim(buf, "\x00"))
				streamMu.Unlock()
			}
		}(pod)
	}
	return logChan, nil
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
