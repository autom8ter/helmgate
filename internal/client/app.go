package client

import (
	"bytes"
	"context"
	"fmt"
	kdeploypb "github.com/autom8ter/kdeploy/gen/grpc/go"
	"github.com/autom8ter/kdeploy/internal/logger"
	"github.com/autom8ter/kubego"
	"github.com/graphikDB/generic"
	"github.com/graphikDB/trigger"
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
	client             *kubego.Client
	jwtCache           generic.Cache
	logger             *logger.Logger
	rootUsers          []string
	requestAuthorizers []*trigger.Decision
	userInfoEndpoint   string
}

func New(client *kubego.Client, logger *logger.Logger, rootUsers []string, userInfoEndpoint string, authorizers []*trigger.Decision) *Manager {
	return &Manager{
		client:             client,
		jwtCache:           generic.NewCache(1 * time.Minute),
		logger:             logger,
		rootUsers:          rootUsers,
		userInfoEndpoint:   userInfoEndpoint,
		requestAuthorizers: authorizers,
	}
}

func (m *Manager) L() *logger.Logger {
	return m.logger
}

func (m *Manager) getStatus(ctx context.Context, namespace, name string) (*kdeploypb.Status, error) {
	var replicas []*kdeploypb.Replica
	pods, err := m.client.Pods(namespace).List(ctx, v1.ListOptions{
		TypeMeta:      v1.TypeMeta{},
		LabelSelector: fmt.Sprintf("kdeploy.app = %s", name),
	})
	if err != nil {
		return nil, err
	}
	for _, pod := range pods.Items {
		replicas = append(replicas, &kdeploypb.Replica{
			Phase:     string(pod.Status.Phase),
			Condition: string(pod.Status.Conditions[0].Status),
			Reason:    pod.Status.Reason,
		})
	}
	return &kdeploypb.Status{Replicas: replicas}, nil
}

func (m *Manager) CreateApp(ctx context.Context, app *kdeploypb.AppConstructor) (*kdeploypb.App, error) {
	kapp := &k8sApp{}
	namespace, err := m.client.Namespaces().Get(ctx, app.Namespace, v1.GetOptions{})
	if err != nil {
		namespace, err = m.client.Namespaces().Create(ctx, toNamespace(app), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
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
	status, err := m.getStatus(ctx, app.Namespace, app.Name)
	if err != nil {
		return nil, err
	}
	a := kapp.toApp()
	a.Status = status
	return a, nil
}

func (m *Manager) UpdateApp(ctx context.Context, app *kdeploypb.AppUpdate) (*kdeploypb.App, error) {
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
	status, err := m.getStatus(ctx, app.Namespace, app.Name)
	if err != nil {
		return nil, err
	}
	a := kapp.toApp()
	a.Status = status
	return a, nil
}

func (m *Manager) GetApp(ctx context.Context, ref *kdeploypb.AppRef) (*kdeploypb.App, error) {
	kapp := &k8sApp{}

	ns, err := m.client.Namespaces().Get(ctx, ref.Namespace, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = ns
	deployment, err := m.client.Deployments(ref.Namespace).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.deployment = deployment
	svc, err := m.client.Services(ref.Namespace).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.service = svc
	status, err := m.getStatus(ctx, ref.Namespace, ref.Name)
	if err != nil {
		return nil, err
	}
	a := kapp.toApp()
	a.Status = status
	return a, nil
}

func (m *Manager) DeleteApp(ctx context.Context, ref *kdeploypb.AppRef) error {
	if err := m.client.Services(ref.Namespace).Delete(ctx, ref.Name, v1.DeleteOptions{}); err != nil {
		m.logger.Error("failed to delete service",
			zap.Error(err),
			zap.String("name", ref.Name),
			zap.String("namespace", ref.Namespace),
		)

	}
	if err := m.client.Deployments(ref.Namespace).Delete(ctx, ref.Name, v1.DeleteOptions{}); err != nil {
		m.logger.Error("failed to delete deployment",
			zap.Error(err),
			zap.String("name", ref.Name),
			zap.String("namespace", ref.Namespace),
		)
	}
	if err := m.client.Namespaces().Delete(ctx, ref.Namespace, v1.DeleteOptions{}); err != nil {
		m.logger.Error("failed to delete namespace",
			zap.Error(err),
			zap.String("name", ref.Name),
			zap.String("namespace", ref.Namespace),
		)
	}
	return nil
}

func (m *Manager) ListNamespaces(ctx context.Context) (*kdeploypb.Namespaces, error) {
	namespaces, err := m.client.Namespaces().List(ctx, v1.ListOptions{
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, err
	}
	var ns = &kdeploypb.Namespaces{}
	for _, n := range namespaces.Items {
		ns.Namespaces = append(ns.Namespaces, n.Name)
	}
	return ns, nil
}

func (m *Manager) ListApps(ctx context.Context, namespace *kdeploypb.Namespace) (*kdeploypb.Apps, error) {
	var kapps = &kdeploypb.Apps{}

	ns, err := m.client.Namespaces().Get(ctx, namespace.GetNamespace(), v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	deployments, err := m.client.Deployments(namespace.GetNamespace()).List(ctx, v1.ListOptions{
		TypeMeta:      v1.TypeMeta{},
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, err
	}
	for _, deployment := range deployments.Items {
		svc, err := m.client.Services(namespace.GetNamespace()).Get(ctx, deployment.Name, v1.GetOptions{})
		if err != nil {
			return nil, err
		}
		kapp := &k8sApp{
			namespace:  ns,
			deployment: &deployment,
			service:    svc,
		}
		a := kapp.toApp()
		status, err := m.getStatus(ctx, namespace.GetNamespace(), deployment.Name)
		if err != nil {
			return nil, err
		}
		a.Status = status
		kapps.Applications = append(kapps.Applications, a)
	}
	return kapps, nil
}

func (m *Manager) StreamLogs(ctx context.Context, ref *kdeploypb.AppRef) (chan string, error) {
	pods, err := m.client.Pods(ref.Namespace).List(ctx, v1.ListOptions{
		TypeMeta:      v1.TypeMeta{},
		Watch:         false,
		LabelSelector: fmt.Sprintf("kdeploy.app = %s", ref.Name),
	})
	if err != nil {
		return nil, err
	}
	if len(pods.Items) == 0 {
		return nil, errors.New("zero pods")
	}
	m.logger.Info("setup log stream",
		zap.String("name", ref.Name),
		zap.String("namespace", ref.Namespace),
	)
	logChan := make(chan string)
	var streamMu = sync.RWMutex{}
	for _, pod := range pods.Items {
		go func(p corev1.Pod) {
			closer, err := m.client.GetLogs(context.Background(), p.Name, p.Namespace, &corev1.PodLogOptions{
				TypeMeta:  v1.TypeMeta{},
				Container: ref.Name,
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
					zap.String("name", ref.Name),
					zap.String("namespace", ref.Namespace),
					zap.String("pod", p.Name),
				)
				return
			}
			for {
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
