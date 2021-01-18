package client

import (
	"fmt"
	kdeploypb "github.com/autom8ter/kdeploy/gen/grpc/go"
	"github.com/autom8ter/kdeploy/internal/helpers"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"istio.io/api/meta/v1alpha1"
	"istio.io/api/networking/v1alpha3"
	networking "istio.io/client-go/pkg/apis/networking/v1alpha3"
	apps "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const Always = "Always"
const OnFailure = "OnFailure"
const RWO = "ReadWriteOnce"

func appLabels(app *kdeploypb.AppConstructor) map[string]string {
	return map[string]string{
		"kdeploy.app": app.Name,
		"kdeploy":     "true",
	}
}

func taskLabels(app *kdeploypb.TaskConstructor) map[string]string {
	return map[string]string{
		"kdeploy.task": app.Name,
		"kdeploy":      "true",
	}
}

func namespaceLabels() map[string]string {
	return map[string]string{
		"kdeploy": "true",
	}
}

func deploymentLabels(dep *apps.Deployment) map[string]string {
	return map[string]string{
		"kdeploy.app": dep.Name,
		"kdeploy":     "true",
	}
}

func appContainers(app *kdeploypb.AppConstructor) ([]v12.Container, error) {
	ports := []v12.ContainerPort{}
	for name, p := range app.Ports {
		ports = append(ports, v12.ContainerPort{
			Name:          name,
			ContainerPort: cast.ToInt32(p),
		})
	}
	env := []v12.EnvVar{}
	for name, val := range app.Env {
		env = append(env, v12.EnvVar{
			Name:  name,
			Value: cast.ToString(val),
		})
	}
	return []v12.Container{
		{
			Name:            app.Name,
			Image:           app.Image,
			Ports:           ports,
			Env:             env,
			Args:            app.Args,
			Resources:       v12.ResourceRequirements{},
			ImagePullPolicy: Always,
		},
	}, nil
}

func taskContainers(app *kdeploypb.TaskConstructor) ([]v12.Container, error) {
	env := []v12.EnvVar{}
	for name, val := range app.Env {
		env = append(env, v12.EnvVar{
			Name:  name,
			Value: cast.ToString(val),
		})
	}
	return []v12.Container{
		{
			Name:            app.Name,
			Image:           app.Image,
			Env:             env,
			Args:            app.Args,
			Resources:       v12.ResourceRequirements{},
			ImagePullPolicy: Always,
		},
	}, nil
}

func toNamespace(app *kdeploypb.AppConstructor) *v12.Namespace {
	return &v12.Namespace{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Namespace,
			Namespace: app.Namespace,
			Labels:    namespaceLabels(),
		},
		Spec:   v12.NamespaceSpec{},
		Status: v12.NamespaceStatus{},
	}
}

func toTaskNamespace(app *kdeploypb.TaskConstructor) *v12.Namespace {
	return &v12.Namespace{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Namespace,
			Namespace: app.Namespace,
			Labels:    namespaceLabels(),
		},
		Spec:   v12.NamespaceSpec{},
		Status: v12.NamespaceStatus{},
	}
}

func toServicePorts(app *kdeploypb.AppConstructor) []v12.ServicePort {
	var ports []v12.ServicePort
	for name, p := range app.Ports {
		ports = append(ports, v12.ServicePort{
			Name: name,
			Port: cast.ToInt32(p),
		})
	}
	return ports
}

func overwriteService(svc *networking.VirtualService, app *kdeploypb.AppUpdate) *networking.VirtualService {
	if svc.Namespace != "" {
		svc.Namespace = app.Namespace
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
	if app.GetNetworking().GetRoutes() != nil {
		var (
			routes       []*v1alpha3.HTTPRoute
			origins      []*v1alpha3.StringMatch
			destinations []*v1alpha3.Destination
		)

		for _, h := range app.GetNetworking().GetRoutes() {
			for _, o := range h.AllowOrigins {
				origins = append(origins, &v1alpha3.StringMatch{
					MatchType: &v1alpha3.StringMatch_Exact{Exact: o},
				})
			}
			destinations = append(destinations, &v1alpha3.Destination{
				Host: fmt.Sprintf("%s.%s.svc.cluster.local", app.Name, app.Namespace),
				Port: &v1alpha3.PortSelector{
					Number: h.Port,
				},
			})
		}
		for _, h := range app.GetNetworking().GetRoutes() {
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
		Namespace: app.Namespace,
		Labels:    appLabels(app),
	},
	Spec: v1.ServiceSpec{
		Ports:    toServicePorts(app),
		Selector: appLabels(app),
		Type:     "",
	},
	Status: v1.ServiceStatus{},
*/

func toService(app *kdeploypb.AppConstructor) *networking.VirtualService {
	svc := &networking.VirtualService{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Namespace,
			Labels:    appLabels(app),
		},
		Spec:   v1alpha3.VirtualService{},
		Status: v1alpha1.IstioStatus{},
	}
	return overwriteService(svc, &kdeploypb.AppUpdate{
		Name:       app.Name,
		Namespace:  app.Namespace,
		Image:      app.Image,
		Args:       app.Args,
		Env:        app.Env,
		Ports:      app.Ports,
		Replicas:   app.Replicas,
		Networking: app.Networking,
	})
}

func toDeployment(app *kdeploypb.AppConstructor) (*apps.Deployment, error) {
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
			Namespace: app.Namespace,
			Labels:    appLabels(app),
		},
		Spec: apps.DeploymentSpec{
			Replicas: &replicas,
			Selector: &metav1.LabelSelector{
				MatchLabels: appLabels(app),
			},
			Template: v12.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      app.Name,
					Namespace: app.Namespace,
					Labels:    appLabels(app),
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

func toTask(app *kdeploypb.TaskConstructor) (*v1beta1.CronJob, error) {
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
			Namespace: app.Namespace,
			Labels:    taskLabels(app),
		},
		Spec: v1beta1.CronJobSpec{
			Schedule: app.Schedule,
			JobTemplate: v1beta1.JobTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{},
				Spec: batchv1.JobSpec{
					Completions: helpers.Int32Pointer(app.Completions),
					Template: v12.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name:      app.Name,
							Namespace: app.Namespace,
							Labels:    taskLabels(app),
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

func overwriteDeployment(deployment *apps.Deployment, app *kdeploypb.AppUpdate) (*apps.Deployment, error) {
	var container *v1.Container
	for _, c := range deployment.Spec.Template.Spec.Containers {
		if c.Name == app.Name {
			container = &c
		}
	}
	if container == nil {
		return nil, errors.Errorf("failed to find container: %s", app.Name)
	}
	replicas := int32(app.Replicas)
	if replicas != *deployment.Spec.Replicas {
		deployment.Spec.Replicas = &replicas
	}
	if app.Image != "" {
		container.Image = app.Image
	}
	if app.Args != nil {
		container.Args = app.Args
	}
	if app.Ports != nil {
		ports := []v12.ContainerPort{}
		for name, p := range app.Ports {
			ports = append(ports, v12.ContainerPort{
				Name:          name,
				ContainerPort: cast.ToInt32(p),
			})
		}
		container.Ports = ports
	}
	if app.Env != nil {
		env := []v12.EnvVar{}
		for name, val := range app.Env {
			env = append(env, v12.EnvVar{
				Name:  name,
				Value: cast.ToString(val),
			})
		}
		container.Env = env
	}
	var containers = []v1.Container{*container}
	for _, c := range deployment.Spec.Template.Spec.Containers {
		if c.Name != app.Name {
			containers = append(containers, c)
		}
	}
	deployment.Spec.Template.Spec.Containers = containers
	return deployment, nil
}

func overwriteTask(cronJob *v1beta1.CronJob, task *kdeploypb.TaskUpdate) (*v1beta1.CronJob, error) {
	var container *v12.Container
	for _, c := range cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if c.Name == cronJob.Name {
			container = &c
		}
	}
	if container == nil {
		return nil, errors.Errorf("failed to find container: %s", task.Name)
	}
	if task.Schedule != "" {
		cronJob.Spec.Schedule = task.Schedule
	}
	if task.Completions != 0 {
		cronJob.Spec.JobTemplate.Spec.Completions = helpers.Int32Pointer(task.Completions)
	}
	if task.Image != "" {
		container.Image = task.Image
	}
	if task.Args != nil {
		container.Args = task.Args
	}
	if task.Env != nil {
		env := []v12.EnvVar{}
		for name, val := range task.Env {
			env = append(env, v12.EnvVar{
				Name:  name,
				Value: cast.ToString(val),
			})
		}
		container.Env = env
	}
	var containers = []v1.Container{*container}
	for _, c := range cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if c.Name != task.Name {
			containers = append(containers, c)
		}
	}
	cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers = containers
	return cronJob, nil
}

type k8sApp struct {
	namespace  *v12.Namespace
	deployment *apps.Deployment
	service    *networking.VirtualService
}

func (k *k8sApp) toApp() *kdeploypb.App {
	a := &kdeploypb.App{
		Name:      k.deployment.Name,
		Namespace: k.deployment.Namespace,
	}
	a.Replicas = uint32(*k.deployment.Spec.Replicas)
	a.Image = k.deployment.Spec.Template.Spec.Containers[0].Image
	var container *v1.Container
	for _, c := range k.deployment.Spec.Template.Spec.Containers {
		if c.Name == a.Name {
			container = &c
		}
	}
	if container == nil {
		panic("failed to find container")
	}
	var env = map[string]string{}
	for _, e := range container.Env {
		env[e.Name] = e.Value
	}
	a.Env = env
	var ports = map[string]uint32{}
	for _, p := range container.Ports {
		ports[p.Name] = cast.ToUint32(p.ContainerPort)
	}
	a.Ports = ports
	return a
}

type k8sTask struct {
	namespace *v12.Namespace
	cronJob   *v1beta1.CronJob
}

func (k *k8sTask) toTask() *kdeploypb.Task {
	a := &kdeploypb.Task{
		Name:      k.cronJob.Name,
		Namespace: k.cronJob.Namespace,
	}
	var container v12.Container
	for _, c := range k.cronJob.Spec.JobTemplate.Spec.Template.Spec.Containers {
		if c.Name == k.cronJob.Name {
			container = c
		}
	}
	a.Image = container.Image
	var env = map[string]string{}
	for _, e := range container.Env {
		env[e.Name] = e.Value
	}
	a.Schedule = k.cronJob.Spec.Schedule
	if k.cronJob.Spec.JobTemplate.Spec.Completions != nil {
		a.Completions = uint32(*k.cronJob.Spec.JobTemplate.Spec.Completions)
	}
	a.Env = env
	var ports = map[string]uint32{}
	for _, p := range container.Ports {
		ports[p.Name] = cast.ToUint32(p.ContainerPort)
	}
	return a
}
