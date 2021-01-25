package core

import (
	"context"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/auth"
	"github.com/spf13/cast"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	securityv1beta1 "istio.io/api/security/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func (m *Manager) CreateAPI(ctx context.Context, app *meshpaaspb.APIInput) (*meshpaaspb.API, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	usrNamespace := cast.ToString(usr[m.namespaceClaim])
	namespace, err := m.kclient.Namespaces().Get(ctx, usrNamespace, v1.GetOptions{})
	if err != nil {
		namespace, err = m.kclient.Namespaces().Create(ctx, m.toNamespace(usr), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	dep, err := m.toDeployment(usr, app)
	if err != nil {
		return nil, err
	}
	deployment, err := m.kclient.Deployments(usrNamespace).Create(ctx, dep, v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	svc, err := m.kclient.Services(usrNamespace).Create(ctx, m.toService(usr, app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	vsvc, err := m.iclient.VirtualServices(usrNamespace).Create(ctx, m.toVirtualService(usr, app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	a, err := m.iclient.RequestAuthentications(usrNamespace).Create(ctx, m.toRequestAuthentication(usr, app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	authz, err := m.iclient.AuthorizationPolicies(usrNamespace).Create(ctx, m.toAuthorizationPolicy(usr, app), v1.CreateOptions{})
	if err != nil {
		return nil, err
	}
	s, err := m.getStatus(ctx, usrNamespace, app.Name)
	if err != nil {
		return nil, err
	}
	kapp := &k8sAPI{
		namespace:      namespace,
		deployment:     deployment,
		svc:            svc,
		vsvc:           vsvc,
		authentication: a,
		authorization:  authz,
	}
	ap := kapp.toAPI()
	ap.Status = s
	return ap, nil
}

func (m *Manager) UpdateAPI(ctx context.Context, app *meshpaaspb.APIInput) (*meshpaaspb.API, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	usrNamespace := cast.ToString(usr[m.namespaceClaim])
	kapp := &k8sAPI{}
	namespace, err := m.kclient.Namespaces().Get(ctx, usrNamespace, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.namespace = namespace
	deployment, err := m.kclient.Deployments(usrNamespace).Get(ctx, app.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	deployment, err = overwriteDeployment(deployment, app)
	if err != nil {
		return nil, err
	}
	deployment, err = m.kclient.Deployments(usrNamespace).Update(ctx, deployment, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.deployment = deployment
	svc, err := m.kclient.Services(usrNamespace).Get(ctx, app.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	svc = m.overwriteService(svc, app)
	svc, err = m.kclient.Services(usrNamespace).Update(ctx, svc, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.svc = svc

	vsvc, err := m.iclient.VirtualServices(usrNamespace).Get(ctx, app.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	vsvc = m.overwriteVirtualService(usr, vsvc, app)
	vsvc, err = m.iclient.VirtualServices(usrNamespace).Update(ctx, vsvc, v1.UpdateOptions{})
	if err != nil {
		return nil, err
	}
	kapp.vsvc = vsvc
	ath, _ := m.iclient.RequestAuthentications(usrNamespace).Get(ctx, app.Name, v1.GetOptions{})
	if ath != nil {
		var rules []*securityv1beta1.JWTRule
		for _, r := range app.Authentication.Rules {
			rules = append(rules, &securityv1beta1.JWTRule{
				Issuer:               r.Issuer,
				Audiences:            r.Audience,
				JwksUri:              r.JwksUri,
				ForwardOriginalToken: true,
			})
		}
		ath.Spec.JwtRules = rules
		ath, err = m.iclient.RequestAuthentications(usrNamespace).Update(ctx, ath, v1.UpdateOptions{})
		if err != nil {
			return nil, err
		}
		kapp.authentication = ath
	}

	authz, _ := m.iclient.AuthorizationPolicies(usrNamespace).Get(ctx, app.Name, v1.GetOptions{})
	if authz != nil {
		authz, err := m.iclient.AuthorizationPolicies(usrNamespace).Update(ctx, m.toAuthorizationPolicy(usr, app), v1.UpdateOptions{})
		if err != nil {
			return nil, err
		}
		kapp.authorization = authz
	}
	stat, err := m.getStatus(ctx, usrNamespace, app.Name)
	if err != nil {
		return nil, err
	}
	a := kapp.toAPI()
	a.Status = stat
	return a, nil
}

func (m *Manager) GetAPI(ctx context.Context, ref *meshpaaspb.Ref) (*meshpaaspb.API, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	usrNamespace := cast.ToString(usr[m.namespaceClaim])
	kapp := &k8sAPI{}
	namespace, err := m.kclient.Namespaces().Get(ctx, usrNamespace, v1.GetOptions{})
	if err != nil {
		namespace, err = m.kclient.Namespaces().Create(ctx, m.toNamespace(usr), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	kapp.namespace = namespace
	deployment, err := m.kclient.Deployments(usrNamespace).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.deployment = deployment
	svc, err := m.kclient.Services(usrNamespace).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	kapp.svc = svc
	vsvc, err := m.iclient.VirtualServices(usrNamespace).Get(ctx, ref.Name, v1.GetOptions{})
	if err != nil {
		return nil, err
	}
	ath, _ := m.iclient.RequestAuthentications(usrNamespace).Get(ctx, ref.Name, v1.GetOptions{})
	kapp.authentication = ath
	authz, _ := m.iclient.AuthorizationPolicies(usrNamespace).Get(ctx, ref.Name, v1.GetOptions{})
	kapp.authorization = authz
	kapp.vsvc = vsvc
	stat, err := m.getStatus(ctx, usrNamespace, ref.Name)
	if err != nil {
		return nil, err
	}
	a := kapp.toAPI()
	a.Status = stat
	return a, nil
}

func (m *Manager) DeleteAPI(ctx context.Context, ref *meshpaaspb.Ref) error {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	usrNamespace := cast.ToString(usr[m.namespaceClaim])
	_, err := m.kclient.Namespaces().Get(ctx, usrNamespace, v1.GetOptions{})
	if err != nil {
		return err
	}
	m.iclient.RequestAuthentications(usrNamespace).Delete(ctx, ref.Name, v1.DeleteOptions{})
	m.iclient.AuthorizationPolicies(usrNamespace).Delete(ctx, ref.Name, v1.DeleteOptions{})
	if err := m.kclient.Services(usrNamespace).Delete(ctx, ref.Name, v1.DeleteOptions{}); err != nil {
		return err
	}
	if err := m.kclient.Deployments(usrNamespace).Delete(ctx, ref.Name, v1.DeleteOptions{}); err != nil {
		return err
	}
	return nil
}

func (m *Manager) ListAPIs(ctx context.Context) (*meshpaaspb.APIs, error) {
	usr, ok := auth.UserContext(ctx)
	if !ok {
		return nil, status.Error(codes.Unauthenticated, "failed to get logged in user")
	}
	usrNamespace := cast.ToString(usr[m.namespaceClaim])
	var kapps = &meshpaaspb.APIs{}

	namespace, err := m.kclient.Namespaces().Get(ctx, usrNamespace, v1.GetOptions{})
	if err != nil {
		namespace, err = m.kclient.Namespaces().Create(ctx, m.toNamespace(usr), v1.CreateOptions{})
		if err != nil {
			return nil, err
		}
	}
	deployments, err := m.kclient.Deployments(usrNamespace).List(ctx, v1.ListOptions{
		TypeMeta:      v1.TypeMeta{},
		LabelSelector: labelSelector,
	})
	if err != nil {
		return nil, err
	}
	for _, deployment := range deployments.Items {
		svc, err := m.kclient.Services(usrNamespace).Get(ctx, deployment.Name, v1.GetOptions{})
		if err != nil {
			return nil, err
		}
		vsvc, err := m.iclient.VirtualServices(usrNamespace).Get(ctx, deployment.Name, v1.GetOptions{})
		if err != nil {
			return nil, err
		}
		kapp := &k8sAPI{
			namespace:      namespace,
			deployment:     &deployment,
			svc:            svc,
			vsvc:           vsvc,
			authentication: nil,
			authorization:  nil,
		}
		athn, _ := m.iclient.RequestAuthentications(usrNamespace).Get(ctx, deployment.Name, v1.GetOptions{})
		kapp.authentication = athn
		athz, _ := m.iclient.AuthorizationPolicies(usrNamespace).Get(ctx, deployment.Name, v1.GetOptions{})
		kapp.authorization = athz
		a := kapp.toAPI()
		stat, err := m.getStatus(ctx, usrNamespace, deployment.Name)
		if err != nil {
			return nil, err
		}
		a.Status = stat
		kapps.Apis = append(kapps.Apis, a)
	}
	return kapps, nil
}
