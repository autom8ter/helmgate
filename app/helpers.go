package app

import (
	"fmt"
	"github.com/graphikDB/kdeploy/gen/gql/go/model"
	"github.com/spf13/cast"
	apps "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	v12 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const Always = "Always"
const RWO = "ReadWriteOnce"

func appLabels(app model.AppInput) map[string]string {
	return map[string]string{
		"kdeploy": fmt.Sprintf("%s.%s", app.Namespace, app.Name),
	}
}

func appContainers(app model.AppInput) []v12.Container {
	containers := []v12.Container{}
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
			Name:            c.Name,
			Image:           c.Image,
			Ports:           ports,
			Env:             env,
			Resources:       v12.ResourceRequirements{},
			ImagePullPolicy: Always,
		})
	}
	return containers
}

func toNamespace(app model.AppInput) *v12.Namespace {
	return &v12.Namespace{
		TypeMeta: metav1.TypeMeta{},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Namespace,
			Namespace: app.Namespace,
			Labels:    appLabels(app),
		},
		Spec:   v12.NamespaceSpec{},
		Status: v12.NamespaceStatus{},
	}
}

func toServicePorts(app model.AppInput) []v12.ServicePort {
	var ports []v12.ServicePort
	for name, p := range app.ExposePorts {
		ports = append(ports, v12.ServicePort{
			Name: name,
			Port: cast.ToInt32(p),
		})
	}
	return ports
}

func toService(app model.AppInput) *v1.Service {
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

func toDeployment(app model.AppInput) *apps.Deployment {
	var (
		replicas   = int32(app.Replicas)
		labels     = appLabels(app)
		containers = appContainers(app)
	)
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
			Selector: nil,
			Template: v12.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      app.Name,
					Namespace: app.Namespace,
					Labels:    labels,
				},
				Spec: v12.PodSpec{
					Volumes:       nil,
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
	}
}

func toStatefulSet(app model.AppInput) *apps.StatefulSet {
	var (
		replicas   = int32(app.Replicas)
		labels     = appLabels(app)
		containers = appContainers(app)
	)
	return &apps.StatefulSet{
		TypeMeta: metav1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Namespace,
			Labels:    appLabels(app),
		},
		Spec: apps.StatefulSetSpec{
			Replicas: &replicas,
			Selector: nil,
			Template: v12.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Name:      app.Name,
					Namespace: app.Namespace,
					Labels:    labels,
				},
				Spec: v12.PodSpec{
					Volumes:       nil,
					Containers:    containers,
					RestartPolicy: Always,
				},
			},
			VolumeClaimTemplates: []v1.PersistentVolumeClaim{
				{
					TypeMeta: metav1.TypeMeta{},
					ObjectMeta: metav1.ObjectMeta{
						Name:      app.Name,
						Namespace: app.Namespace,
						Labels:    appLabels(app),
					},
					Spec: v1.PersistentVolumeClaimSpec{
						AccessModes: []v1.PersistentVolumeAccessMode{RWO},
						Selector: &metav1.LabelSelector{
							MatchLabels: appLabels(app),
						},
						Resources: v1.ResourceRequirements{
							Limits: v1.ResourceList{},
							Requests: v1.ResourceList{
								v1.ResourceRequestsStorage: resource.MustParse(app.State.StorageSize),
							},
						},
						VolumeName: app.Name,
					},
				},
			},
			ServiceName: app.Name,
		},
		Status: apps.StatefulSetStatus{},
	}
}

type k8sApp struct {
	input       model.AppInput
	namespace   *v12.Namespace
	deployment  *apps.Deployment
	statefulset *apps.StatefulSet
	service     *v1.Service
}

func (k *k8sApp) toApp() *model.App {
	a := &model.App{
		Name:        k.input.Name,
		Namespace:   k.input.Namespace,
		ExposePorts: k.input.ExposePorts,
		Replicas:    &k.input.Replicas,
	}
	var containers []*model.Container
	var state *model.State
	if k.input.State != nil {
		state = &model.State{
			Statefulset: a.State.Statefulset,
			StoragePath: a.State.StoragePath,
			StorageSize: a.State.StorageSize,
		}
	}
	for _, c := range k.input.Containers {
		containers = append(containers, &model.Container{
			Name:   c.Name,
			Image:  c.Image,
			Env:    c.Env,
			Ports:  c.Ports,
			Memory: c.Memory,
		})
	}
	a.Status = map[string]interface{}{}
	if k.deployment != nil {
		a.Status["deployment"] = k.deployment.Status.String()
	}
	if k.statefulset != nil {
		a.Status["statefulset"] = k.statefulset.Status.String()
	}
	if k.service != nil {
		a.Status["service"] = k.service.Status.String()
	}
	if k.namespace != nil {
		a.Status["namespace"] = k.namespace.Status.String()
	}
	a.Containers = containers
	a.State = state
	return a
}
