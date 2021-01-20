package client

import (
	"bytes"
	"context"
	"fmt"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"io"
	securityv1beta1 "istio.io/api/security/v1beta1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sync"
)

func (m *Manager) CreateApp(ctx context.Context, app *meshpaaspb.AppInput) (*meshpaaspb.App, error) {
	kapp := &k8sApp{}
	namespace, err := m.kclient.Namespaces().Get(ctx, app.Project, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	dep, err := toDeployment(app)
	if err != nil {
		return nil, err
	}
	deployment, err := m.kclient.Deployments(app.Project).Create(ctx, dep, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.deployment = deployment
	svc, err := m.iclient.VirtualServices(app.Project).Create(ctx, toService(app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.service = svc
	auth, err := m.iclient.RequestAuthentications(app.Project).Create(ctx, toRequestAuthentication(app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.authentication = auth
	authz, err := m.iclient.AuthorizationPolicies(app.Project).Create(ctx, toAuthorizationPolicy(app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.authorization = authz
	status, err := m.getStatus(ctx, app.Project, app.Name)
	if err != nil {
		return nil, err
	}
	a := kapp.toApp()
	a.Status = status
	return a, nil
}

func (m *Manager) UpdateApp(ctx context.Context, app *meshpaaspb.AppInput) (*meshpaaspb.App, error) {
	kapp := &k8sApp{}
	namespace, err := m.kclient.Namespaces().Get(ctx, app.Project, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	deployment, err := m.kclient.Deployments(app.Project).Get(ctx, app.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	deployment, err = overwriteDeployment(deployment, app)
	if err != nil {
		return nil, err
	}
	deployment, err = m.kclient.Deployments(app.Project).Update(ctx, deployment, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.deployment = deployment
	svc, err := m.iclient.VirtualServices(app.Project).Get(ctx, app.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	svc = overwriteService(svc, app)
	svc, err = m.iclient.VirtualServices(app.Project).Update(ctx, svc, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.service = svc
	auth, err := m.iclient.RequestAuthentications(app.Project).Get(ctx, app.Name, v1.GetOptions{})
	if err != nil {
		auth, err = m.iclient.RequestAuthentications(app.Project).Create(ctx, toRequestAuthentication(app), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	var rules []*securityv1beta1.JWTRule
	for _, r := range app.Authentication.Rules {
		rules = append(rules, &securityv1beta1.JWTRule{
			Issuer:                r.Issuer,
			Audiences:             r.Audience,
			JwksUri:               r.JwksUri,
			OutputPayloadToHeader: r.OuputPayloadHeader,
			ForwardOriginalToken:  false,
		})
	}
	auth.Spec.JwtRules = rules
	auth, err = m.iclient.RequestAuthentications(app.Project).Update(ctx, auth, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.authentication = auth
	_, err = m.iclient.AuthorizationPolicies(app.Project).Get(ctx, app.Name, v1.GetOptions{})
	if err != nil {
		_, err = m.iclient.AuthorizationPolicies(app.Project).Create(ctx, toAuthorizationPolicy(app), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	authz, err := m.iclient.AuthorizationPolicies(app.Project).Update(ctx, toAuthorizationPolicy(app), v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.authorization = authz
	status, err := m.getStatus(ctx, app.Project, app.Name)
	if err != nil {
		return nil, err
	}
	a := kapp.toApp()
	a.Status = status
	return a, nil
}

func (m *Manager) GetApp(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.App, error) {
	kapp := &k8sApp{}

	ns, err := m.kclient.Namespaces().Get(ctx, ref.Project, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = ns
	deployment, err := m.kclient.Deployments(ref.Project).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.deployment = deployment
	svc, err := m.iclient.VirtualServices(ref.Project).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	auth, _ := m.iclient.RequestAuthentications(ref.Project).Get(ctx, ref.Name, v1.GetOptions{})
	kapp.authentication = auth
	authz, _ := m.iclient.AuthorizationPolicies(ref.Project).Get(ctx, ref.Name, v1.GetOptions{})
	kapp.authorization = authz
	kapp.service = svc
	status, err := m.getStatus(ctx, ref.Project, ref.Name)
	if err != nil {
		return nil, err
	}
	a := kapp.toApp()
	a.Status = status
	return a, nil
}

func (m *Manager) DeleteApp(ctx context.Context, ref *meshpaaspb.Ref) error {
	m.iclient.RequestAuthentications(ref.Project).Delete(ctx, ref.Name, v1.DeleteOptions{})
	m.iclient.AuthorizationPolicies(ref.Project).Delete(ctx, ref.Name, v1.DeleteOptions{})
	if err := m.kclient.Services(ref.Project).Delete(ctx, ref.Name, v1.DeleteOptions{}); err != nil {
		m.logger.Error("failed to delete service",
			zap.Error(err),
			zap.String("name", ref.Name),
			zap.String("namespace", ref.Project),
		)
	}
	if err := m.kclient.Deployments(ref.Project).Delete(ctx, ref.Name, v1.DeleteOptions{}); err != nil {
		m.logger.Error("failed to delete deployment",
			zap.Error(err),
			zap.String("name", ref.Name),
			zap.String("namespace", ref.Project),
		)
	}
	return nil
}

func (m *Manager) ListApps(ctx context.Context, namespace *meshpaaspb.ProjectRef) (*meshpaaspb.Apps, error) {
	var kapps = &meshpaaspb.Apps{}

	ns, err := m.kclient.Namespaces().Get(ctx, namespace.GetName(), v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	deployments, err := m.kclient.Deployments(namespace.GetName()).List(ctx, v1.ListOptions{
		TypeMeta:      v1.TypeMeta{},
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, err
	}
	for _, deployment := range deployments.Items {
		svc, err := m.iclient.VirtualServices(namespace.GetName()).Get(ctx, deployment.Name, v1.GetOptions{})
		if err != nil {
			return nil, err
		}
		kapp := &k8sApp{
			namespace:  ns,
			deployment: &deployment,
			service:    svc,
		}
		auth, _ := m.iclient.RequestAuthentications(namespace.GetName()).Get(ctx, deployment.Name, v1.GetOptions{})
		kapp.authentication = auth
		a := kapp.toApp()
		status, err := m.getStatus(ctx, namespace.GetName(), deployment.Name)
		if err != nil {
			return nil, err
		}
		a.Status = status
		kapps.Applications = append(kapps.Applications, a)
	}
	return kapps, nil
}

func (m *Manager) StreamLogs(ctx context.Context, ref *meshpaaspb.Ref) (chan string, error) {
	pods, err := m.kclient.Pods(ref.Project).List(ctx, v1.ListOptions{
		TypeMeta:      v1.TypeMeta{},
		Watch:         false,
		LabelSelector: fmt.Sprintf("app = %s", ref.Name),
	})
	if err != nil {
		return nil, err
	}
	if len(pods.Items) == 0 {
		return nil, errors.New("zero pods")
	}
	logChan := make(chan string)
	var streamMu = sync.RWMutex{}
	for _, pod := range pods.Items {
		go func(p corev1.Pod) {
			closer, err := m.kclient.GetLogs(context.Background(), p.Name, p.Namespace, &corev1.PodLogOptions{
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
					zap.String("namespace", ref.Project),
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
