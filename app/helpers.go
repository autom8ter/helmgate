package app

import (
	"github.com/autom8ter/kdeploy/gen/gql/go/model"
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
		"app":     app.Name,
		"kdeploy": "true",
	}
}

func appContainers(app model.AppInput) ([]v12.Container, error) {
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
	var volumes []v12.VolumeMount
	if app.State != nil {
		volumes = append(volumes, v12.VolumeMount{
			Name:      app.Name,
			ReadOnly:  false,
			MountPath: app.State.StoragePath,
		})
	}
	return []v12.Container{
		{
			Name:            app.Name,
			Image:           app.Image,
			Ports:           ports,
			Env:             env,
			Resources:       v12.ResourceRequirements{},
			VolumeMounts:    volumes,
			ImagePullPolicy: Always,
		},
	}, nil
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
	for name, p := range app.Ports {
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

func toDeployment(app model.AppInput) (*apps.Deployment, error) {
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
	}, nil
}

func toStatefulSet(app model.AppInput) (*apps.StatefulSet, error) {
	var (
		replicas        = int32(app.Replicas)
		containers, err = appContainers(app)
	)
	if err != nil {
		return nil, err
	}
	strg, err := resource.ParseQuantity(app.State.StorageSize)
	if err != nil {
		return nil, err
	}
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
							Requests: v1.ResourceList{
								v1.ResourceRequestsStorage: strg,
							},
						},
						VolumeName: app.Name,
					},
				},
			},
			ServiceName: app.Name,
		},
		Status: apps.StatefulSetStatus{},
	}, nil
}

type k8sApp struct {
	namespace   *v12.Namespace
	deployment  *apps.Deployment
	statefulset *apps.StatefulSet
	service     *v1.Service
}

func (k *k8sApp) toApp() *model.App {
	a := &model.App{
		Name:      k.service.Name,
		Namespace: k.service.Namespace,
		State:     nil,
		Status:    nil,
	}
	var state *model.State
	a.Status = &model.Status{}
	if k.deployment != nil {
		a.Replicas = int(*k.deployment.Spec.Replicas)
		a.Image = k.deployment.Spec.Template.Spec.Containers[0].Image
		var env = map[string]interface{}{}
		for _, e := range k.deployment.Spec.Template.Spec.Containers[0].Env {
			env[e.Name] = e.Value
		}
		a.Env = env
		var ports = map[string]interface{}{}
		for _, p := range k.deployment.Spec.Template.Spec.Containers[0].Ports {
			ports[p.Name] = p.ContainerPort
		}
		a.Ports = ports
		a.Status.Deployment = k.deployment.Status.String()
	}
	if k.statefulset != nil {
		a.Replicas = int(*k.statefulset.Spec.Replicas)
		a.Image = k.statefulset.Spec.Template.Spec.Containers[0].Image
		var env = map[string]interface{}{}
		for _, e := range k.statefulset.Spec.Template.Spec.Containers[0].Env {
			env[e.Name] = e.Value
		}
		a.Env = env
		var ports = map[string]interface{}{}
		for _, p := range k.statefulset.Spec.Template.Spec.Containers[0].Ports {
			ports[p.Name] = p.ContainerPort
		}
		a.Ports = ports
		a.Status.Deployment = k.statefulset.Status.String()
		state = &model.State{
			Statefulset: true,
			StoragePath: k.statefulset.Spec.Template.Spec.Containers[0].VolumeMounts[0].MountPath,
			StorageSize: k.statefulset.Spec.VolumeClaimTemplates[0].Spec.Resources.Requests.Storage().String(),
		}
	}
	if k.service != nil {
		a.Status.LoadBalancer = k.service.Status.String()
	}
	if k.namespace != nil {
		a.Status.Namespace = k.namespace.Status.String()
	}
	a.State = state
	return a
}
