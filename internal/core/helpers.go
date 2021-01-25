package core

import (
	"fmt"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/helpers"
	"github.com/spf13/cast"
	"istio.io/api/meta/v1alpha1"
	"istio.io/api/networking/v1alpha3"
	nv1alpha3 "istio.io/api/networking/v1alpha3"
	securityv1beta1 "istio.io/api/security/v1beta1"
	networking "istio.io/client-go/pkg/apis/networking/v1alpha3"
	pkgnv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	security "istio.io/client-go/pkg/apis/security/v1beta1"
	apps "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

const (
	labelSelector = "meshpaas = true"
	Always        = "Always"
	OnFailure     = "OnFailure"
	meshpaasApp   = "meshpaas.app"
	meshpaas      = "meshpaas"
)

func appContainers(app *meshpaaspb.APIInput) ([]v12.Container, error) {
	var containers []v12.Container
	for _, c := range app.Containers {
		ports := []v12.ContainerPort{}
		for _, p := range c.Ports {
			ports = append(ports, v12.ContainerPort{
				Name:          p.Name,
				ContainerPort: cast.ToInt32(p),
			})
		}
		env := []v12.EnvVar{}
		for name, val := range c.Env {
			env = append(env, v12.EnvVar{
				Name:  name,
				Value: cast.ToString(val),
			})
		}
		containers = append(containers, v12.Container{
			Name:            c.Name,
			Image:           c.Image,
			Args:            c.Args,
			Ports:           ports,
			Env:             env,
			ImagePullPolicy: Always,
			Resources:       v12.ResourceRequirements{},
		})
	}

	return containers, nil
}

func taskContainers(app *meshpaaspb.TaskInput) ([]v12.Container, error) {
	var containers []v12.Container
	for _, c := range app.Containers {
		env := []v12.EnvVar{}
		for name, val := range c.Env {
			env = append(env, v12.EnvVar{
				Name:  name,
				Value: cast.ToString(val),
			})
		}
		ports := []v12.ContainerPort{}
		for _, p := range c.Ports {
			ports = append(ports, v12.ContainerPort{
				Name:          p.Name,
				ContainerPort: cast.ToInt32(p),
			})
		}
		containers = append(containers, v12.Container{
			Name:            c.Name,
			Image:           c.Image,
			Args:            c.Args,
			Ports:           ports,
			Env:             env,
			Resources:       v12.ResourceRequirements{},
			ImagePullPolicy: Always,
		})

	}

	return containers, nil
}

func (m *Manager) toNamespace(usr map[string]interface{}) *v12.Namespace {
	return &v12.Namespace{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      cast.ToString(usr[m.namespaceClaim]),
			Namespace: cast.ToString(usr[m.namespaceClaim]),
		},
		Spec:   v12.NamespaceSpec{},
		Status: v12.NamespaceStatus{},
	}
}

func (m *Manager) overwriteService(svc *v1.Service, app *meshpaaspb.APIInput) *v1.Service {
	var ports []v1.ServicePort

	for _, c := range app.Containers {
		for _, p := range c.Ports {
			ports = append(ports, v1.ServicePort{
				Name: p.Name,
				Port: int32(p.GetNumber()),
			})
		}
	}
	svc.Spec.Ports = ports
	return svc
}

func (m *Manager) overwriteVirtualService(usr map[string]interface{}, svc *networking.VirtualService, app *meshpaaspb.APIInput) *networking.VirtualService {
	if svc.Name != "" {
		svc.Name = app.Name
	}
	if app.GetRouting().GetGateway() != "" {
		svc.Spec.Gateways = []string{app.GetRouting().GetGateway()}
	}
	if app.GetRouting().GetHosts() != nil {
		svc.Spec.Hosts = app.GetRouting().GetHosts()
	}

	if app.GetRouting().GetHttpRoutes() != nil {
		var (
			routes       []*v1alpha3.HTTPRoute
			origins      []*v1alpha3.StringMatch
			destinations []*v1alpha3.Destination
		)

		for _, h := range app.GetRouting().GetHttpRoutes() {
			for _, o := range h.AllowOrigins {
				origins = append(origins, &v1alpha3.StringMatch{
					MatchType: &v1alpha3.StringMatch_Exact{Exact: o},
				})
			}
			destinations = append(destinations, &v1alpha3.Destination{
				Host: fmt.Sprintf("%s.%s.svc.cluster.local", app.Name, cast.ToString(usr[m.namespaceClaim])),
				Port: &v1alpha3.PortSelector{
					Number: h.Port,
				},
			})
		}
		for _, h := range app.GetRouting().GetHttpRoutes() {
			routes = append(routes, &v1alpha3.HTTPRoute{
				Name:  h.Name,
				Match: nil,
				Route: []*v1alpha3.HTTPRouteDestination{
					{
						Destination: &v1alpha3.Destination{
							Host:   app.Name,
							Subset: "",
							Port: &v1alpha3.PortSelector{
								Number: h.Port,
							},
						},
					},
				},
				Rewrite: &v1alpha3.HTTPRewrite{
					Uri:       h.RewriteUri,
					Authority: "",
				},
				Timeout:          nil,
				Retries:          nil,
				Fault:            nil,
				Mirror:           nil,
				MirrorPercent:    nil,
				MirrorPercentage: nil,
				CorsPolicy: &v1alpha3.CorsPolicy{
					AllowOrigins:  origins,
					AllowMethods:  h.AllowMethods,
					AllowHeaders:  h.AllowHeaders,
					ExposeHeaders: h.ExposeHeaders,
				},
				Headers: nil,
			})
		}
	}

	return svc
}

/*

	TypeMeta: metav1.TypeMeta{},
	ObjectMeta: metav1.ObjectMeta{
		Name:      app.Name,
		Namespace: cast.ToString(usr[m.namespaceClaim]),
		Labels:    app.Labels,
	},
	Spec: v1.ServiceSpec{
		Ports:    toVirtualServicePorts(app),
		Selector: app.Labels,
		Type:     "",
	},
	Status: v1.ServiceStatus{},
*/

func (m *Manager) toVirtualService(usr map[string]interface{}, app *meshpaaspb.APIInput) *networking.VirtualService {
	svc := &networking.VirtualService{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: cast.ToString(usr[m.namespaceClaim]),
			Labels: map[string]string{
				meshpaasApp: app.Name,
				meshpaas:    "true",
			},
		},
		Spec:   v1alpha3.VirtualService{},
		Status: v1alpha1.IstioStatus{},
	}
	if svc.Namespace != "" {
		svc.Namespace = cast.ToString(usr[m.namespaceClaim])
	}
	if svc.Name != "" {
		svc.Name = app.Name
	}
	if app.GetRouting().GetGateway() != "" {
		svc.Spec.Gateways = []string{app.GetRouting().GetGateway()}
	}
	if app.GetRouting().GetHosts() != nil {
		svc.Spec.Hosts = app.GetRouting().GetHosts()
	}

	if app.GetRouting().GetHttpRoutes() != nil {
		var (
			routes []*v1alpha3.HTTPRoute
		)
		for _, h := range app.GetRouting().GetHttpRoutes() {
			var (
				origins []*v1alpha3.StringMatch
			)
			for _, o := range h.AllowOrigins {
				origins = append(origins, &v1alpha3.StringMatch{
					MatchType: &v1alpha3.StringMatch_Exact{Exact: o},
				})
			}
			route := &v1alpha3.HTTPRoute{
				Name: h.Name,
				Route: []*v1alpha3.HTTPRouteDestination{
					{
						Destination: &v1alpha3.Destination{
							Host:   fmt.Sprintf("%s.%s.svc.cluster.local", app.Name, cast.ToString(usr[m.namespaceClaim])),
							Subset: "",
							Port: &v1alpha3.PortSelector{
								Number: h.Port,
							},
						},
					},
				},
				Timeout:          nil,
				Retries:          nil,
				Fault:            nil,
				Mirror:           nil,
				MirrorPercent:    nil,
				MirrorPercentage: nil,
				CorsPolicy: &v1alpha3.CorsPolicy{
					AllowOrigins:  origins,
					AllowMethods:  h.AllowMethods,
					AllowHeaders:  h.AllowHeaders,
					ExposeHeaders: h.ExposeHeaders,
				},
				Headers: nil,
			}
			if h.PathPrefix != "" {
				route.Match = append(route.Match, &v1alpha3.HTTPMatchRequest{
					Uri: &v1alpha3.StringMatch{
						MatchType: &v1alpha3.StringMatch_Prefix{
							Prefix: h.PathPrefix,
						},
					},
					Port: h.Port,
				})
			}
			if h.RewriteUri != "" {
				route.Rewrite = &v1alpha3.HTTPRewrite{
					Uri:       h.RewriteUri,
					Authority: "",
				}
			}
			routes = append(routes, route)
		}
		svc.Spec.Http = routes
	}

	return svc
}

func (m *Manager) toRequestAuthentication(usr map[string]interface{}, input *meshpaaspb.APIInput) *security.RequestAuthentication {
	if input.Authentication == nil {
		return nil
	}
	var auth = &security.RequestAuthentication{
		ObjectMeta: metav1.ObjectMeta{
			Name:      input.Name,
			Namespace: cast.ToString(usr[m.namespaceClaim]),
		},
	}
	auth.Spec.Selector.MatchLabels = map[string]string{
		meshpaasApp: input.Name,
		meshpaas:    "true",
	}
	for _, r := range input.Authentication.Rules {
		auth.Spec.JwtRules = append(auth.Spec.JwtRules, &securityv1beta1.JWTRule{
			Issuer:               r.Issuer,
			Audiences:            r.Audience,
			JwksUri:              r.JwksUri,
			ForwardOriginalToken: true,
		})
	}
	return auth
}

func (m *Manager) toAuthorizationPolicy(usr map[string]interface{}, input *meshpaaspb.APIInput) *security.AuthorizationPolicy {
	var auth = &security.AuthorizationPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name:      input.Name,
			Namespace: cast.ToString(usr[m.namespaceClaim]),
		},
	}
	auth.Spec.Selector.MatchLabels = map[string]string{
		meshpaasApp: input.Name,
		meshpaas:    "true",
	}
	auth.Spec.Action = securityv1beta1.AuthorizationPolicy_ALLOW
	return auth
}

func (m *Manager) overwriteAuthorizationPolicy(auth *security.AuthorizationPolicy, input *meshpaaspb.APIInput) *security.AuthorizationPolicy {
	return auth
}

func (m *Manager) toDeployment(usr map[string]interface{}, app *meshpaaspb.APIInput) (*apps.Deployment, error) {
	var (
		replicas        = int32(app.Replicas)
		containers, err = appContainers(app)
		imagePullSecret []v12.LocalObjectReference
	)
	if err != nil {
		return nil, err
	}
	if app.ImagePullSecret != "" {
		imagePullSecret = append(imagePullSecret, v12.LocalObjectReference{
			Name: app.ImagePullSecret,
		})
	}
	var labels = map[string]string{
		meshpaasApp: app.Name,
		meshpaas:    "true",
	}

	return &apps.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: cast.ToString(usr[m.namespaceClaim]),
			Labels:    labels,
		},
		Spec: apps.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: labels,
			},
			Template: v12.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      app.Name,
					Namespace: cast.ToString(usr[m.namespaceClaim]),
					Labels:    labels,
				},
				Spec: v12.PodSpec{
					ImagePullSecrets: imagePullSecret,
					Containers:       containers,
					RestartPolicy:    Always,
				},
			},
			Strategy:                apps.DeploymentStrategy{},
			MinReadySeconds:         0,
			RevisionHistoryLimit:    nil,
			Paused:                  false,
			ProgressDeadlineSeconds: nil,
		},
		Status: apps.DeploymentStatus{},
	}, nil
}

func (m *Manager) toTask(usr map[string]interface{}, app *meshpaaspb.TaskInput) (*v1beta1.CronJob, error) {
	var (
		containers, err = taskContainers(app)
	)
	if err != nil {
		return nil, err
	}
	var labels = map[string]string{
		meshpaasApp: app.Name,
		meshpaas:    "true",
	}
	return &v1beta1.CronJob{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: cast.ToString(usr[m.namespaceClaim]),
			Labels:    labels,
		},
		Spec: v1beta1.CronJobSpec{
			Schedule: app.Schedule,
			JobTemplate: v1beta1.JobTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{},
				Spec: batchv1.JobSpec{
					Completions: helpers.Int32Pointer(app.Completions),
					Selector: &metav1.LabelSelector{
						MatchLabels: labels,
					},
					Template: v12.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name:      app.Name,
							Namespace: cast.ToString(usr[m.namespaceClaim]),
							Labels:    labels,
						},
						Spec: v12.PodSpec{
							Containers:    containers,
							RestartPolicy: OnFailure,
						},
					},
					TTLSecondsAfterFinished: nil,
				},
			},
		},
	}, nil
}

func overwriteDeployment(deployment *apps.Deployment, app *meshpaaspb.APIInput) (*apps.Deployment, error) {
	replicas := int32(app.Replicas)
	if replicas != *deployment.Spec.Replicas {
		deployment.Spec.Replicas = &replicas
	}
	containers, err := appContainers(app)
	if err != nil {
		return nil, err
	}
	if app.ImagePullSecret != "" {
		deployment.Spec.Template.Spec.ImagePullSecrets = []v12.LocalObjectReference{
			{
				Name: app.ImagePullSecret,
			},
		}
	}
	deployment.Spec.Template.Spec.Containers = containers
	return deployment, nil
}

func overwriteGateway(gateway *pkgnv1alpha3.Gateway, gw *meshpaaspb.GatewayInput) *pkgnv1alpha3.Gateway {
	gateway.Spec.Servers = toServers(gw)
	return gateway
}

func toServers(gateway *meshpaaspb.GatewayInput) []*nv1alpha3.Server {
	var servers []*nv1alpha3.Server
	for _, l := range gateway.GetListeners() {
		var tls *nv1alpha3.ServerTLSSettings
		if l.TlsConfig != nil {
			tls = &nv1alpha3.ServerTLSSettings{
				HttpsRedirect: l.TlsConfig.HttpsRedirect,
				Mode: func() nv1alpha3.ServerTLSSettings_TLSmode {
					switch l.TlsConfig.Mode {
					case meshpaaspb.TLSmode_SIMPLE:
						return nv1alpha3.ServerTLSSettings_SIMPLE
					case meshpaaspb.TLSmode_AUTO_PASSTHROUGH:
						return nv1alpha3.ServerTLSSettings_AUTO_PASSTHROUGH
					case meshpaaspb.TLSmode_PASSTHROUGH:
						return nv1alpha3.ServerTLSSettings_PASSTHROUGH
					case meshpaaspb.TLSmode_ISTIO_MUTUAL:
						return nv1alpha3.ServerTLSSettings_ISTIO_MUTUAL
					case meshpaaspb.TLSmode_MUTUAL:
						return nv1alpha3.ServerTLSSettings_MUTUAL
					default:
						return nv1alpha3.ServerTLSSettings_SIMPLE
					}
				}(),
				CredentialName:        l.TlsConfig.SecretName,
				SubjectAltNames:       l.TlsConfig.SubjectAltNames,
				VerifyCertificateSpki: l.TlsConfig.VerifyCertificateSpki,
				VerifyCertificateHash: l.TlsConfig.VerifyCertificateHash,
				MinProtocolVersion:    nv1alpha3.ServerTLSSettings_TLS_AUTO,
				MaxProtocolVersion:    nv1alpha3.ServerTLSSettings_TLS_AUTO,
				CipherSuites:          l.TlsConfig.CipherSuites,
			}
		}
		servers = append(servers, &nv1alpha3.Server{
			Port: &nv1alpha3.Port{
				Number:   l.Port,
				Protocol: l.Protocol.String(),
				Name:     strings.ToLower(l.Protocol.String()),
			},
			Hosts: l.GetHosts(),
			Tls:   tls,
			Name:  l.GetName(),
		})
	}
	return servers
}

func (m *Manager) toGateway(usr map[string]interface{}, gateway *meshpaaspb.GatewayInput) *pkgnv1alpha3.Gateway {
	return &pkgnv1alpha3.Gateway{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      gateway.GetName(),
			Namespace: cast.ToString(usr[m.namespaceClaim]),
		},
		Spec: nv1alpha3.Gateway{
			Servers: toServers(gateway),
		},
		Status: v1alpha1.IstioStatus{},
	}
}

func overwriteTask(cronJob *v1beta1.CronJob, task *meshpaaspb.TaskInput) (*v1beta1.CronJob, error) {
	containers, err := taskContainers(task)
	if err != nil {
		return nil, err
	}
	if task.Schedule != "" {
		cronJob.Spec.Schedule = task.Schedule
	}
	if task.Completions != 0 {
		cronJob.Spec.JobTemplate.Spec.Completions = helpers.Int32Pointer(task.Completions)
	}
	if task.ImagePullSecret != "" {
		cronJob.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets = []v12.LocalObjectReference{
			{
				Name: task.ImagePullSecret,
			},
		}
	}
	cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers = containers
	return cronJob, nil
}

func (m *Manager) toService(usr map[string]interface{}, input *meshpaaspb.APIInput) *v1.Service {
	svc := &v1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:      input.Name,
			Namespace: cast.ToString(usr[m.namespaceClaim]),
			Labels: map[string]string{
				meshpaasApp: input.Name,
				meshpaas:    "true",
			},
		},
	}
	for _, c := range input.Containers {
		for _, p := range c.GetPorts() {
			if p.Expose {
				svc.Spec.Ports = append(svc.Spec.Ports, v1.ServicePort{
					Name: p.Name,
					Port: int32(p.GetNumber()),
				})
			}
		}
	}
	return svc
}

type k8sAPI struct {
	namespace      *v12.Namespace
	deployment     *apps.Deployment
	svc            *v1.Service
	vsvc           *networking.VirtualService
	authentication *security.RequestAuthentication
	authorization  *security.AuthorizationPolicy
}

func (k *k8sAPI) toAPI() *meshpaaspb.API {
	a := &meshpaaspb.API{
		Name:           k.deployment.Name,
		Containers:     nil,
		Replicas:       uint32(*k.deployment.Spec.Replicas),
		Routing:        &meshpaaspb.Routing{},
		Authentication: &meshpaaspb.Authn{},
		Status:         nil,
	}
	if len(k.vsvc.Spec.Gateways) > 0 {
		a.Routing.Gateway = k.vsvc.Spec.Gateways[0]
	}
	a.Routing.Hosts = k.vsvc.Spec.Hosts
	for _, r := range k.vsvc.Spec.Http {
		var origins []string
		var prefix string
		if len(r.Match) > 0 {
			prefix = r.Match[0].Uri.GetPrefix()
		}
		a.Routing.HttpRoutes = append(a.Routing.HttpRoutes, &meshpaaspb.HTTPRoute{
			Name:             r.Name,
			Port:             r.Route[0].Destination.Port.Number,
			PathPrefix:       prefix,
			RewriteUri:       r.Rewrite.GetUri(),
			AllowOrigins:     origins,
			AllowMethods:     r.CorsPolicy.AllowMethods,
			AllowHeaders:     r.CorsPolicy.AllowHeaders,
			ExposeHeaders:    r.CorsPolicy.ExposeHeaders,
			AllowCredentials: r.CorsPolicy.AllowCredentials != nil && r.CorsPolicy.AllowCredentials.Value,
		})
	}
	if len(k.deployment.Spec.Template.Spec.ImagePullSecrets) > 0 {
		a.ImagePullSecret = k.deployment.Spec.Template.Spec.ImagePullSecrets[0].Name
	}
	for _, c := range k.deployment.Spec.Template.Spec.Containers {
		var env = map[string]string{}
		for _, e := range c.Env {
			env[e.Name] = e.Value
		}
		var ports = []*meshpaaspb.ContainerPort{}
		for _, p := range c.Ports {
			port := &meshpaaspb.ContainerPort{
				Name:   p.Name,
				Number: uint32(p.ContainerPort),
				Expose: false,
			}
			for _, svcP := range k.svc.Spec.Ports {
				if svcP.Name == p.Name && svcP.Port == p.ContainerPort {
					port.Expose = true
				}
			}
			ports = append(ports, port)
		}
		a.Containers = append(a.Containers, &meshpaaspb.Container{
			Name:  c.Name,
			Image: c.Image,
			Args:  c.Args,
			Env:   env,
			Ports: ports,
		})
	}
	for _, r := range k.authentication.Spec.JwtRules {
		a.Authentication.Rules = append(a.Authentication.Rules, &meshpaaspb.AuthnRule{
			JwksUri:  r.JwksUri,
			Issuer:   r.Issuer,
			Audience: r.Audiences,
		})
	}

	return a
}

type k8sTask struct {
	namespace *v12.Namespace
	cronJob   *v1beta1.CronJob
}

func (k *k8sTask) toTask() *meshpaaspb.Task {
	a := &meshpaaspb.Task{
		Name: k.cronJob.Name,
	}
	a.Schedule = k.cronJob.Spec.Schedule
	if len(k.cronJob.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets) > 0 {
		a.ImagePullSecret = k.cronJob.Spec.JobTemplate.Spec.Template.Spec.ImagePullSecrets[0].Name
	}
	for _, c := range k.cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		var env = map[string]string{}
		for _, e := range c.Env {
			env[e.Name] = e.Value
		}
		if k.cronJob.Spec.JobTemplate.Spec.Completions != nil {
			a.Completions = uint32(*k.cronJob.Spec.JobTemplate.Spec.Completions)
		}
		var ports []*meshpaaspb.ContainerPort
		for _, p := range c.Ports {
			ports = append(ports, &meshpaaspb.ContainerPort{
				Name:   p.Name,
				Number: uint32(p.ContainerPort),
				Expose: false,
			})
		}
		a.Containers = append(a.Containers, &meshpaaspb.Container{
			Name:  c.Name,
			Image: c.Image,
			Args:  c.Args,
			Env:   env,
			Ports: ports,
		})
	}
	return a
}

type k8sGateway struct {
	namespace *v12.Namespace
	gateway   *pkgnv1alpha3.Gateway
}

func (k *k8sGateway) toGateway() *meshpaaspb.Gateway {
	var listeners []*meshpaaspb.GatewayListener
	for _, l := range k.gateway.Spec.GetServers() {
		listeners = append(listeners, &meshpaaspb.GatewayListener{
			Port: l.GetPort().GetNumber(),
			Name: l.GetName(),
			Protocol: func() meshpaaspb.TransportProtocol {
				switch l.GetPort().GetProtocol() {
				case "GRPC":
					return meshpaaspb.TransportProtocol_GRPC
				case "HTTP":
					return meshpaaspb.TransportProtocol_HTTP
				case "HTTP2":
					return meshpaaspb.TransportProtocol_HTTP2
				case "HTTPS":
					return meshpaaspb.TransportProtocol_HTTPS
				case "MONGO":
					return meshpaaspb.TransportProtocol_MONGO
				case "TLS":
					return meshpaaspb.TransportProtocol_TLS
				default:
					return meshpaaspb.TransportProtocol_TCP
				}
			}(),
			Hosts: l.GetHosts(),
			TlsConfig: &meshpaaspb.ServerTLSSettings{
				HttpsRedirect: l.GetTls().GetHttpsRedirect(),
				Mode: func() meshpaaspb.TLSmode {
					switch l.GetTls().GetMode() {
					case nv1alpha3.ServerTLSSettings_SIMPLE:
						return meshpaaspb.TLSmode_SIMPLE
					case nv1alpha3.ServerTLSSettings_MUTUAL:
						return meshpaaspb.TLSmode_MUTUAL
					case nv1alpha3.ServerTLSSettings_ISTIO_MUTUAL:
						return meshpaaspb.TLSmode_ISTIO_MUTUAL
					case nv1alpha3.ServerTLSSettings_AUTO_PASSTHROUGH:
						return meshpaaspb.TLSmode_AUTO_PASSTHROUGH
					case nv1alpha3.ServerTLSSettings_PASSTHROUGH:
						return meshpaaspb.TLSmode_PASSTHROUGH
					default:
						return meshpaaspb.TLSmode_SIMPLE
					}
				}(),
				SecretName:            l.GetTls().GetCredentialName(),
				SubjectAltNames:       l.GetTls().GetSubjectAltNames(),
				VerifyCertificateSpki: l.GetTls().GetVerifyCertificateSpki(),
				VerifyCertificateHash: l.GetTls().GetVerifyCertificateHash(),
				CipherSuites:          l.GetTls().GetCipherSuites(),
			},
		})
	}
	return &meshpaaspb.Gateway{
		Name:      k.gateway.ObjectMeta.Name,
		Listeners: listeners,
	}
}

type k8sSecret struct {
	namespace *v12.Namespace
	secret    *v1.Secret
}

func (k *k8sSecret) toSecret() *meshpaaspb.Secret {
	return &meshpaaspb.Secret{
		Name: k.secret.Name,
		Type: func() meshpaaspb.SecretType {
			switch k.secret.Type {
			case v1.DockerConfigKey:
				return meshpaaspb.SecretType_DOCKER_CONFIG
			case v1.SecretTypeTLS:
				return meshpaaspb.SecretType_TLS_CERT_KEY
			default:
				return meshpaaspb.SecretType_OPAQUE
			}
		}(),
		Immutable: *k.secret.Immutable,
		Data: func() map[string]string {
			var data = map[string]string{}
			for k, v := range k.secret.Data {
				data[k] = cast.ToString(v)
			}
			return data
		}(),
	}
}

func (m *Manager) toSecret(usr map[string]interface{}, secret *meshpaaspb.SecretInput) *v1.Secret {
	return &v1.Secret{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      secret.GetName(),
			Namespace: cast.ToString(usr[m.namespaceClaim]),
		},
		Immutable:  &secret.Immutable,
		Data:       nil,
		StringData: secret.Data,
		Type: func() v1.SecretType {
			switch secret.Type {
			case meshpaaspb.SecretType_DOCKER_CONFIG:
				return v1.SecretTypeDockercfg
			case meshpaaspb.SecretType_TLS_CERT_KEY:
				return v1.SecretTypeTLS
			default:
				return v1.SecretTypeOpaque
			}
		}(),
	}
}

func overwriteSecret(secret *v1.Secret, update *meshpaaspb.SecretInput) *v1.Secret {
	secret.Name = update.Name
	secret.Immutable = &update.Immutable
	secret.StringData = update.Data
	return secret
}
