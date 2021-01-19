package client

import (
	"fmt"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/helpers"
	"github.com/spf13/cast"
	"istio.io/api/meta/v1alpha1"
	"istio.io/api/networking/v1alpha3"
	nv1alpha3 "istio.io/api/networking/v1alpha3"
	networking "istio.io/client-go/pkg/apis/networking/v1alpha3"
	pkgnv1alpha3 "istio.io/client-go/pkg/apis/networking/v1alpha3"
	apps "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

const Always = "Always"
const OnFailure = "OnFailure"
const RWO = "ReadWriteOnce"

func appContainers(app *meshpaaspb.AppInput) ([]v12.Container, error) {
	var containers []v12.Container
	for _, c := range app.Containers {
		ports := []v12.ContainerPort{}
		for name, p := range c.Ports {
			ports = append(ports, v12.ContainerPort{
				Name:          name,
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
			Name:      c.Name,
			Image:     c.Image,
			Args:      c.Args,
			Ports:     ports,
			Env:       env,
			Resources: v12.ResourceRequirements{},
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
		for name, p := range c.Ports {
			ports = append(ports, v12.ContainerPort{
				Name:          name,
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

func toNamespace(project *meshpaaspb.ProjectInput) *v12.Namespace {
	return &v12.Namespace{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      project.Name,
			Namespace: project.Name,
			Labels:    project.Labels,
		},
		Spec:   v12.NamespaceSpec{},
		Status: v12.NamespaceStatus{},
	}
}

func toGwNamespace(gw *meshpaaspb.GatewayInput) *v12.Namespace {
	return &v12.Namespace{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      gw.Project,
			Namespace: gw.Project,
			Labels:    gw.Labels,
		},
		Spec:   v12.NamespaceSpec{},
		Status: v12.NamespaceStatus{},
	}
}

func toSecretNamespace(secret *meshpaaspb.SecretInput) *v12.Namespace {
	return &v12.Namespace{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      secret.Project,
			Namespace: secret.Project,
			Labels:    secret.Labels,
		},
		Spec:   v12.NamespaceSpec{},
		Status: v12.NamespaceStatus{},
	}
}

func toTaskNamespace(app *meshpaaspb.TaskInput) *v12.Namespace {
	return &v12.Namespace{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Project,
			Namespace: app.Project,
		},
		Spec:   v12.NamespaceSpec{},
		Status: v12.NamespaceStatus{},
	}
}

func overwriteService(svc *networking.VirtualService, app *meshpaaspb.AppInput) *networking.VirtualService {
	if svc.Namespace != "" {
		svc.Namespace = app.Project
	}
	if svc.Name != "" {
		svc.Name = app.Name
	}
	if app.GetNetworking().GetGateways() != nil {
		svc.Spec.Gateways = app.GetNetworking().GetGateways()
	}
	if app.GetNetworking().GetHosts() != nil {
		svc.Spec.Hosts = app.GetNetworking().GetHosts()
	}
	if app.GetNetworking().GetExport() {
		svc.Spec.ExportTo = []string{"*"}
	} else {
		svc.Spec.ExportTo = []string{"."}
	}
	if app.GetNetworking().GetHttpRoutes() != nil {
		var (
			routes       []*v1alpha3.HTTPRoute
			origins      []*v1alpha3.StringMatch
			destinations []*v1alpha3.Destination
		)

		for _, h := range app.GetNetworking().GetHttpRoutes() {
			for _, o := range h.AllowOrigins {
				origins = append(origins, &v1alpha3.StringMatch{
					MatchType: &v1alpha3.StringMatch_Exact{Exact: o},
				})
			}
			destinations = append(destinations, &v1alpha3.Destination{
				Host: fmt.Sprintf("%s.%s.svc.cluster.local", app.Name, app.Project),
				Port: &v1alpha3.PortSelector{
					Number: h.Port,
				},
			})
		}
		for _, h := range app.GetNetworking().GetHttpRoutes() {
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
		Namespace: app.Project,
		Labels:    app.Labels,
	},
	Spec: v1.ServiceSpec{
		Ports:    toServicePorts(app),
		Selector: app.Labels,
		Type:     "",
	},
	Status: v1.ServiceStatus{},
*/

func toService(app *meshpaaspb.AppInput) *networking.VirtualService {
	svc := &networking.VirtualService{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Project,
			Labels:    app.Labels,
		},
		Spec:   v1alpha3.VirtualService{},
		Status: v1alpha1.IstioStatus{},
	}
	if svc.Namespace != "" {
		svc.Namespace = app.Project
	}
	if svc.Name != "" {
		svc.Name = app.Name
	}
	if app.GetNetworking().GetGateways() != nil {
		svc.Spec.Gateways = app.GetNetworking().GetGateways()
	}
	if app.GetNetworking().GetHosts() != nil {
		svc.Spec.Hosts = app.GetNetworking().GetHosts()
	}
	if app.GetNetworking().GetExport() {
		svc.Spec.ExportTo = []string{"*"}
	} else {
		svc.Spec.ExportTo = []string{"."}
	}
	if app.GetNetworking().GetHttpRoutes() != nil {
		var (
			routes []*v1alpha3.HTTPRoute
		)
		for _, h := range app.GetNetworking().GetHttpRoutes() {
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
							Host:   fmt.Sprintf("%s.%s.svc.cluster.local", app.Name, app.Project),
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

func toDeployment(app *meshpaaspb.AppInput) (*apps.Deployment, error) {
	var (
		replicas        = int32(app.Replicas)
		containers, err = appContainers(app)
	)
	if err != nil {
		return nil, err
	}
	return &apps.Deployment{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Project,
			Labels:    app.Labels,
		},
		Spec: apps.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: app.Selector,
			},
			Template: v12.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      app.Name,
					Namespace: app.Project,
					Labels:    app.Labels,
				},
				Spec: v12.PodSpec{
					Containers:    containers,
					RestartPolicy: Always,
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

func toTask(app *meshpaaspb.TaskInput) (*v1beta1.CronJob, error) {
	var (
		containers, err = taskContainers(app)
	)
	if err != nil {
		return nil, err
	}
	return &v1beta1.CronJob{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Project,
			Labels:    app.Labels,
		},
		Spec: v1beta1.CronJobSpec{
			Schedule: app.Schedule,
			JobTemplate: v1beta1.JobTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{},
				Spec: batchv1.JobSpec{
					Completions: helpers.Int32Pointer(app.Completions),
					Selector: &metav1.LabelSelector{
						MatchLabels: app.Selector,
					},
					Template: v12.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name:      app.Name,
							Namespace: app.Project,
							Labels:    app.Labels,
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

func overwriteDeployment(deployment *apps.Deployment, app *meshpaaspb.AppInput) (*apps.Deployment, error) {
	replicas := int32(app.Replicas)
	if replicas != *deployment.Spec.Replicas {
		deployment.Spec.Replicas = &replicas
	}
	if app.Labels != nil {
		deployment.Labels = app.Labels
	}
	if app.Selector != nil {
		deployment.Spec.Selector = &metav1.LabelSelector{MatchLabels: app.Selector}
	}
	containers, err := appContainers(app)
	if err != nil {
		return nil, err
	}
	deployment.Spec.Template.Spec.Containers = containers
	return deployment, nil
}

func overwriteGateway(gateway *pkgnv1alpha3.Gateway, gw *meshpaaspb.GatewayInput) *pkgnv1alpha3.Gateway {
	gateway.Name = gw.Name
	gateway.Namespace = gw.Project
	gateway.Labels = gw.Labels
	gateway.Spec.Selector = gw.Selector
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
				CredentialName:        l.TlsConfig.CredentialName,
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

func toGateway(gateway *meshpaaspb.GatewayInput) *pkgnv1alpha3.Gateway {
	return &pkgnv1alpha3.Gateway{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      gateway.GetName(),
			Namespace: gateway.GetProject(),
			Labels:    gateway.GetLabels(),
		},
		Spec: nv1alpha3.Gateway{
			Servers:  toServers(gateway),
			Selector: gateway.GetSelector(),
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
	cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers = containers
	return cronJob, nil
}

type k8sApp struct {
	namespace  *v12.Namespace
	deployment *apps.Deployment
	service    *networking.VirtualService
}

func (k *k8sApp) toApp() *meshpaaspb.App {
	a := &meshpaaspb.App{
		Name:       k.deployment.Name,
		Project:    k.deployment.Namespace,
		Containers: nil,
		Replicas:   uint32(*k.deployment.Spec.Replicas),
		Labels:     k.deployment.Labels,
		Selector:   k.deployment.Spec.Selector.MatchLabels,
		Networking: &meshpaaspb.Networking{},
		Status:     nil,
	}
	a.Networking.Gateways = k.service.Spec.Gateways
	a.Networking.Hosts = k.service.Spec.Hosts
	if len(k.service.Spec.ExportTo) > 0 {
		a.Networking.Export = k.service.Spec.ExportTo[0] == "*"
	}
	for _, r := range k.service.Spec.Http {
		var origins []string
		var prefix string
		if len(r.Match) > 0 {
			prefix = r.Match[0].Uri.GetPrefix()
		}
		a.Networking.HttpRoutes = append(a.Networking.HttpRoutes, &meshpaaspb.HTTPRoute{
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
	for _, c := range k.deployment.Spec.Template.Spec.Containers {
		var env = map[string]string{}
		for _, e := range c.Env {
			env[e.Name] = e.Value
		}
		var ports = map[string]uint32{}
		for _, p := range c.Ports {
			ports[p.Name] = cast.ToUint32(p.ContainerPort)
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

type k8sTask struct {
	namespace *v12.Namespace
	cronJob   *v1beta1.CronJob
}

func (k *k8sTask) toTask() *meshpaaspb.Task {
	a := &meshpaaspb.Task{
		Name:    k.cronJob.Name,
		Project: k.cronJob.Namespace,
	}
	for _, c := range k.cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		var env = map[string]string{}
		for _, e := range c.Env {
			env[e.Name] = e.Value
		}
		a.Schedule = k.cronJob.Spec.Schedule
		if k.cronJob.Spec.JobTemplate.Spec.Completions != nil {
			a.Completions = uint32(*k.cronJob.Spec.JobTemplate.Spec.Completions)
		}
		var ports = map[string]uint32{}
		for _, p := range c.Ports {
			ports[p.Name] = cast.ToUint32(p.ContainerPort)
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
			Protocol: func() meshpaaspb.Protocol {
				switch l.GetPort().GetProtocol() {
				case "GRPC":
					return meshpaaspb.Protocol_GRPC
				case "HTTP":
					return meshpaaspb.Protocol_HTTP
				case "HTTP2":
					return meshpaaspb.Protocol_HTTP2
				case "HTTPS":
					return meshpaaspb.Protocol_HTTPS
				case "MONGO":
					return meshpaaspb.Protocol_MONGO
				case "TLS":
					return meshpaaspb.Protocol_TLS
				default:
					return meshpaaspb.Protocol_TCP
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
				CredentialName:        l.GetTls().GetCredentialName(),
				SubjectAltNames:       l.GetTls().GetSubjectAltNames(),
				VerifyCertificateSpki: l.GetTls().GetVerifyCertificateSpki(),
				VerifyCertificateHash: l.GetTls().GetVerifyCertificateHash(),
				CipherSuites:          l.GetTls().GetCipherSuites(),
			},
		})
	}
	return &meshpaaspb.Gateway{
		Name:      k.gateway.ObjectMeta.Name,
		Project:   k.gateway.ObjectMeta.Namespace,
		Listeners: listeners,
		Labels:    k.gateway.ObjectMeta.Labels,
		Selector:  k.gateway.Spec.Selector,
	}
}

type k8sSecret struct {
	namespace *v12.Namespace
	secret    *v1.Secret
}

func (k *k8sSecret) toSecret() *meshpaaspb.Secret {
	return &meshpaaspb.Secret{
		Name:    k.secret.Name,
		Project: k.secret.Namespace,
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
		Labels: k.secret.Labels,
	}
}

func toSecret(secret *meshpaaspb.SecretInput) *v1.Secret {
	return &v1.Secret{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      secret.GetName(),
			Namespace: secret.GetProject(),
			Labels:    secret.GetLabels(),
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
	secret.Namespace = update.Project
	secret.Name = update.Name
	secret.Labels = update.Labels
	secret.Immutable = &update.Immutable
	secret.StringData = update.Data
	return secret
}
