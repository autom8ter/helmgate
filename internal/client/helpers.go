package client

import (
	kdeploypb "github.com/autom8ter/kdeploy/gen/grpc/go"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	apps "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1"
	"k8s.io/api/batch/v1beta1"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const Always = "Always"
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

func overwriteService(svc *v1.Service, app *kdeploypb.AppUpdate) *v1.Service {
	if app.Ports != nil {
		var ports []v12.ServicePort
		for name, p := range app.Ports {
			ports = append(ports, v12.ServicePort{
				Name: name,
				Port: cast.ToInt32(p),
			})
		}
		svc.Spec.Ports = ports
	}
	return svc
}

func toService(app *kdeploypb.AppConstructor) *v1.Service {
	return &v1.Service{
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
	}
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

			JobTemplate: v1beta1.JobTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{},
				Spec: batchv1.JobSpec{
					Parallelism:           nil,
					Completions:           int32Pointer(app.Completions),
					ActiveDeadlineSeconds: nil,
					BackoffLimit:          nil,
					Selector: &metav1.LabelSelector{
						MatchLabels: taskLabels(app),
					},
					ManualSelector: nil,
					Template: v12.PodTemplateSpec{
						ObjectMeta: metav1.ObjectMeta{
							Name:      app.Name,
							Namespace: app.Namespace,
							Labels:    taskLabels(app),
						},
						Spec: v12.PodSpec{
							Containers:    containers,
							RestartPolicy: Always,
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
		cronJob.Spec.JobTemplate.Spec.Completions = int32Pointer(task.Completions)
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
	service    *v1.Service
}

func (k *k8sApp) toApp() *kdeploypb.App {
	a := &kdeploypb.App{
		Name:      k.service.Name,
		Namespace: k.service.Namespace,
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

func int32Pointer(value uint32) *int32 {
	if value != 0 {
		i := int32(value)
		return &i
	}
	return nil
}
