package gql

import (
	"github.com/autom8ter/kdeploy/gen/gql/go/model"
	kdeploypb "github.com/autom8ter/kdeploy/gen/grpc/go"
	"github.com/spf13/cast"
)

func toAppC(input model.AppConstructor) *kdeploypb.AppConstructor {
	var env map[string]string
	var ports map[string]uint32
	if input.Env != nil {
		env = cast.ToStringMapString(input.Env)
	}
	if input.Ports != nil {
		ports = map[string]uint32{}
		for k, v := range input.Ports {
			ports[k] = cast.ToUint32(v)
		}
	}
	return &kdeploypb.AppConstructor{
		Name:      input.Name,
		Namespace: input.Namespace,
		Image:     input.Image,
		Env:       env,
		Ports:     ports,
		Replicas:  uint32(input.Replicas),
		Args:      input.Args,
	}
}

func toTaskC(input model.TaskConstructor) *kdeploypb.TaskConstructor {
	var env map[string]string
	if input.Env != nil {
		env = cast.ToStringMapString(input.Env)
	}
	return &kdeploypb.TaskConstructor{
		Name:      input.Name,
		Namespace: input.Namespace,
		Image:     input.Image,
		Env:       env,
		Args:      input.Args,
		Schedule:  input.Schedule,
	}
}

func toAppU(input model.AppUpdate) *kdeploypb.AppUpdate {
	var (
		env      map[string]string
		ports    map[string]uint32
		image    string
		replicas uint32
	)
	if input.Env != nil {
		env = cast.ToStringMapString(input.Env)
	}
	if input.Ports != nil {
		ports = map[string]uint32{}
		for k, v := range input.Ports {
			ports[k] = cast.ToUint32(v)
		}
	}
	if input.Image != nil {
		image = *input.Image
	}
	if input.Replicas != nil {
		replicas = uint32(*input.Replicas)
	}
	return &kdeploypb.AppUpdate{
		Name:      input.Name,
		Namespace: input.Namespace,
		Image:     image,
		Env:       env,
		Ports:     ports,
		Replicas:  replicas,
		Args:      input.Args,
	}
}

func toTaskU(input model.TaskUpdate) *kdeploypb.TaskUpdate {
	var (
		env      map[string]string
		schedule string
		image    string
	)
	if input.Env != nil {
		env = cast.ToStringMapString(input.Env)
	}
	if input.Image != nil {
		image = *input.Image
	}
	if input.Schedule != nil {
		schedule = *input.Schedule
	}
	return &kdeploypb.TaskUpdate{
		Name:      input.Name,
		Namespace: input.Namespace,
		Image:     image,
		Env:       env,
		Schedule:  schedule,
		Args:      input.Args,
	}
}

func fromApp(app *kdeploypb.App) *model.App {
	var (
		env    map[string]interface{}
		ports  map[string]interface{}
		status = &model.Status{}
	)
	if app.Env != nil {
		env = map[string]interface{}{}
		for k, v := range app.Env {
			env[k] = v
		}
	}
	if app.Ports != nil {
		ports = map[string]interface{}{}
		for k, v := range app.Ports {
			ports[k] = v
		}
	}
	for _, r := range app.Status.Replicas {
		status.Replicas = append(status.Replicas, &model.Replica{
			Phase:     r.Phase,
			Condition: r.Condition,
			Reason:    r.Reason,
		})
	}
	return &model.App{
		Name:      app.Name,
		Namespace: app.Namespace,
		Image:     app.Image,
		Env:       env,
		Ports:     ports,
		Replicas:  int(app.Replicas),
		Status:    status,
		Args:      app.Args,
	}
}

func fromTask(app *kdeploypb.Task) *model.Task {
	var (
		env map[string]interface{}
	)
	if app.Env != nil {
		env = map[string]interface{}{}
		for k, v := range app.Env {
			env[k] = v
		}
	}
	return &model.Task{
		Name:      app.Name,
		Namespace: app.Namespace,
		Image:     app.Image,
		Args:      app.Args,
		Env:       env,
		Schedule:  app.Schedule,
	}
}

func toApp(input *model.App) *kdeploypb.App {
	var (
		env   map[string]string
		ports map[string]uint32
	)
	if input.Env != nil {
		env = cast.ToStringMapString(input.Env)
	}
	if input.Ports != nil {
		ports = map[string]uint32{}
		for k, v := range input.Ports {
			ports[k] = cast.ToUint32(v)
		}
	}
	var status = &kdeploypb.AppStatus{}
	for _, replica := range input.Status.Replicas {
		status.Replicas = append(status.Replicas, &kdeploypb.Replica{
			Phase:     replica.Phase,
			Condition: replica.Condition,
			Reason:    replica.Reason,
		})
	}
	return &kdeploypb.App{
		Name:      input.Name,
		Namespace: input.Namespace,
		Image:     input.Image,
		Args:      input.Args,
		Env:       env,
		Ports:     ports,
		Replicas:  uint32(input.Replicas),
		Status:    status,
	}
}
