package client

import (
	"bytes"
	"context"
	"fmt"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/auth"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	securityv1beta1 "istio.io/api/security/v1beta1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sync"
)

func (m *Manager) CreateApp(ctx context.Context, app *meshpaaspb.AppInput) (*meshpaaspb.App, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	kapp := &k8sApp{}
	namespace, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr["aud"]), v1.GetOptions{})
	if err != nil {
		namespace, err = m.kclient.Namespaces().Create(ctx, toNamespace(usr), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	kapp.namespace = namespace
	dep, err := toDeployment(usr, app)
	if err != nil {
		return nil, err
	}
	deployment, err := m.kclient.Deployments(cast.ToString(usr["aud"])).Create(ctx, dep, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.deployment = deployment
	svc, err := m.iclient.VirtualServices(cast.ToString(usr["aud"])).Create(ctx, toService(usr, app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.service = svc
	a, err := m.iclient.RequestAuthentications(cast.ToString(usr["aud"])).Create(ctx, toRequestAuthentication(usr, app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.authentication = a
	authz, err := m.iclient.AuthorizationPolicies(cast.ToString(usr["aud"])).Create(ctx, toAuthorizationPolicy(usr, app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.authorization = authz
	s, err := m.getStatus(ctx, cast.ToString(usr["aud"]), app.Name)
	if err != nil {
		return nil, err
	}
	ap := kapp.toApp()
	ap.Status = s
	return ap, nil
}

func (m *Manager) UpdateApp(ctx context.Context, app *meshpaaspb.AppInput) (*meshpaaspb.App, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	kapp := &k8sApp{}
	namespace, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr["aud"]), v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	deployment, err := m.kclient.Deployments(cast.ToString(usr["aud"])).Get(ctx, app.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	deployment, err = overwriteDeployment(deployment, app)
	if err != nil {
		return nil, err
	}
	deployment, err = m.kclient.Deployments(cast.ToString(usr["aud"])).Update(ctx, deployment, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.deployment = deployment
	svc, err := m.iclient.VirtualServices(cast.ToString(usr["aud"])).Get(ctx, app.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	svc = overwriteService(usr, svc, app)
	svc, err = m.iclient.VirtualServices(cast.ToString(usr["aud"])).Update(ctx, svc, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.service = svc
	ath, err := m.iclient.RequestAuthentications(cast.ToString(usr["aud"])).Get(ctx, app.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
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
	ath.Spec.JwtRules = rules
	ath, err = m.iclient.RequestAuthentications(cast.ToString(usr["aud"])).Update(ctx, ath, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.authentication = ath
	_, err = m.iclient.AuthorizationPolicies(cast.ToString(usr["aud"])).Get(ctx, app.Name, v1.GetOptions{})
	if err != nil {
		_, err = m.iclient.AuthorizationPolicies(cast.ToString(usr["aud"])).Create(ctx, toAuthorizationPolicy(usr, app), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	authz, err := m.iclient.AuthorizationPolicies(cast.ToString(usr["aud"])).Update(ctx, toAuthorizationPolicy(usr, app), v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.authorization = authz
	stat, err := m.getStatus(ctx, cast.ToString(usr["aud"]), app.Name)
	if err != nil {
		return nil, err
	}
	a := kapp.toApp()
	a.Status = stat
	return a, nil
}

func (m *Manager) GetApp(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.App, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	kapp := &k8sApp{}
	namespace, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr["aud"]), v1.GetOptions{})
	if err != nil {
		namespace, err = m.kclient.Namespaces().Create(ctx, toNamespace(usr), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	kapp.namespace = namespace
	deployment, err := m.kclient.Deployments(cast.ToString(usr["aud"])).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.deployment = deployment
	svc, err := m.iclient.VirtualServices(cast.ToString(usr["aud"])).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	ath, _ := m.iclient.RequestAuthentications(cast.ToString(usr["aud"])).Get(ctx, ref.Name, v1.GetOptions{})
	kapp.authentication = ath
	authz, _ := m.iclient.AuthorizationPolicies(cast.ToString(usr["aud"])).Get(ctx, ref.Name, v1.GetOptions{})
	kapp.authorization = authz
	kapp.service = svc
	stat, err := m.getStatus(ctx, cast.ToString(usr["aud"]), ref.Name)
	if err != nil {
		return nil, err
	}
	a := kapp.toApp()
	a.Status = stat
	return a, nil
}

func (m *Manager) DeleteApp(ctx context.Context, ref *meshpaaspb.Ref) error {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	_, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr["aud"]), v1.GetOptions{})
	if err != nil {
		return err
	}
	m.iclient.RequestAuthentications(cast.ToString(usr["aud"])).Delete(ctx, ref.Name, v1.DeleteOptions{})
	m.iclient.AuthorizationPolicies(cast.ToString(usr["aud"])).Delete(ctx, ref.Name, v1.DeleteOptions{})
	if err := m.kclient.Services(cast.ToString(usr["aud"])).Delete(ctx, ref.Name, v1.DeleteOptions{}); err != nil {
		m.logger.Error("failed to delete service",
			zap.Error(err),
			zap.String("name", ref.Name),
			zap.String("namespace", cast.ToString(usr["aud"])),
		)
	}
	if err := m.kclient.Deployments(cast.ToString(usr["aud"])).Delete(ctx, ref.Name, v1.DeleteOptions{}); err != nil {
		m.logger.Error("failed to delete deployment",
			zap.Error(err),
			zap.String("name", ref.Name),
			zap.String("namespace", cast.ToString(usr["aud"])),
		)
	}
	return nil
}

func (m *Manager) ListApps(ctx context.Context) (*meshpaaspb.Apps, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	var kapps = &meshpaaspb.Apps{}

	namespace, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr["aud"]), v1.GetOptions{})
	if err != nil {
		namespace, err = m.kclient.Namespaces().Create(ctx, toNamespace(usr), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	deployments, err := m.kclient.Deployments(cast.ToString(usr["aud"])).List(ctx, v1.ListOptions{
		TypeMeta:      v1.TypeMeta{},
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, err
	}
	for _, deployment := range deployments.Items {
		svc, err := m.iclient.VirtualServices(cast.ToString(usr["aud"])).Get(ctx, deployment.Name, v1.GetOptions{})
		if err != nil {
			return nil, err
		}
		kapp := &k8sApp{
			namespace:  namespace,
			deployment: &deployment,
			service:    svc,
		}
		ath, _ := m.iclient.RequestAuthentications(cast.ToString(usr["aud"])).Get(ctx, deployment.Name, v1.GetOptions{})
		kapp.authentication = ath
		a := kapp.toApp()
		stat, err := m.getStatus(ctx, cast.ToString(usr["aud"]), deployment.Name)
		if err != nil {
			return nil, err
		}
		a.Status = stat
		kapps.Applications = append(kapps.Applications, a)
	}
	return kapps, nil
}

func (m *Manager) StreamLogs(ctx context.Context, ref *meshpaaspb.Ref) (chan string, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	_, err := m.kclient.Namespaces().Get(ctx, cast.ToString(usr["aud"]), v1.GetOptions{})
	if err != nil {
		_, err = m.kclient.Namespaces().Create(ctx, toNamespace(usr), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}

	pods, err := m.kclient.Pods(cast.ToString(usr["aud"])).List(ctx, v1.ListOptions{
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
					zap.String("namespace", cast.ToString(usr["aud"])),
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
