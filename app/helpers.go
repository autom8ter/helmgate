package app

import (
	"fmt"
	"github.com/graphikDB/kdeploy/gen/gql/go/model"
	"github.com/graphikDB/kdeploy/version"
	"github.com/spf13/cast"
	apps "k8s.io/api/apps/v1"
	v12 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func appLabels(app model.AppInput) map[string]string {
	return map[string]string{
		"kdeploy":         fmt.Sprintf("%s.%s", app.Namespace, app.Name),
		"kdeploy_version": version.Version,
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
			ImagePullPolicy: "Always",
		})
	}
	return containers
}

func toDeployment(app model.AppInput) *apps.Deployment {
	var (
		replicas   = int32(1)
		labels     = appLabels(app)
		containers = appContainers(app)
	)
	if app.Replicas != nil {
		replicas = int32(*app.Replicas)
	}
	return &apps.Deployment{
		TypeMeta: v1.TypeMeta{
			Kind:       "",
			APIVersion: "",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      app.Name,
			Namespace: app.Namespace,
			Labels:    appLabels(app),
		},
		Spec: apps.DeploymentSpec{
			Replicas: &replicas,
			Selector: nil,
			Template: v12.PodTemplateSpec{
				ObjectMeta: v1.ObjectMeta{
					Name:      app.Name,
					Namespace: app.Namespace,
					Labels:    labels,
				},
				Spec: v12.PodSpec{
					Volumes:       nil,
					Containers:    containers,
					RestartPolicy: "Always",
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
