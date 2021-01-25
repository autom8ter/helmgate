package gql

import (
	"github.com/autom8ter/meshpaas/gen/gql/go/model"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/helpers"
)

func toApp(input model.AppInput) *meshpaaspb.AppInput {
	var networking = &meshpaaspb.Routing{}
	var containers []*meshpaaspb.Container
	for _, c := range input.Containers {
		ports := []*meshpaaspb.ContainerPort{}
		for _, v := range c.Ports {
			ports = append(ports, &meshpaaspb.ContainerPort{
				Name:   v.Name,
				Number: uint32(v.Number),
				Expose: v.Expose,
			})
		}
		containers = append(containers, &meshpaaspb.Container{
			Name:  c.Name,
			Image: c.Image,
			Args:  c.Args,
			Env:   helpers.ConvertMapS(c.Env),
			Ports: ports,
		})
	}
	for _, r := range input.Routing.HTTPRoutes {
		var (
			pathPrefix       string
			rewriteUri       string
			allowCredentials bool
		)

		networking.HttpRoutes = append(networking.HttpRoutes, &meshpaaspb.HTTPRoute{
			Name:             r.Name,
			Port:             uint32(r.Port),
			PathPrefix:       pathPrefix,
			RewriteUri:       rewriteUri,
			AllowOrigins:     r.AllowOrigins,
			AllowMethods:     r.AllowMethods,
			AllowHeaders:     r.AllowHeaders,
			ExposeHeaders:    r.ExposeHeaders,
			AllowCredentials: allowCredentials,
		})
	}
	return &meshpaaspb.AppInput{
		Name:       input.Name,
		Containers: containers,
		Replicas:   uint32(input.Replicas),
		Routing:    networking,
	}
}

func toTask(input model.TaskInput) *meshpaaspb.TaskInput {
	var completions uint32
	var containers []*meshpaaspb.Container
	for _, c := range input.Containers {
		var ports []*meshpaaspb.ContainerPort
		for _, v := range c.Ports {
			ports = append(ports, &meshpaaspb.ContainerPort{
				Name:   v.Name,
				Number: uint32(v.Number),
				Expose: v.Expose,
			})
		}
		containers = append(containers, &meshpaaspb.Container{
			Name:  c.Name,
			Image: c.Image,
			Args:  c.Args,
			Env:   helpers.ConvertMapS(c.Env),
			Ports: ports,
		})
	}
	if input.Completions != nil {
		completions = uint32(*input.Completions)
	}
	return &meshpaaspb.TaskInput{
		Name:        input.Name,
		Containers:  containers,
		Schedule:    input.Schedule,
		Completions: completions,
	}
}

func fromApp(app *meshpaaspb.App) *model.App {
	var (
		status     = &model.AppStatus{}
		containers []*model.Container
	)
	for _, c := range app.Containers {
		var (
			env   map[string]interface{}
			ports []*model.ContainerPort
		)
		if c.Env != nil {
			env = map[string]interface{}{}
			for k, v := range c.Env {
				env[k] = v
			}
		}
		if c.Ports != nil {
			for _, v := range c.Ports {
				ports = append(ports, &model.ContainerPort{
					Name:   v.Name,
					Number: int(v.Number),
					Expose: v.Expose,
				})
			}
		}
		containers = append(containers, &model.Container{
			Name:  c.Name,
			Image: c.Image,
			Args:  c.Args,
			Env:   env,
			Ports: ports,
		})
	}

	for _, r := range app.Status.Replicas {
		status.Replicas = append(status.Replicas, &model.Replica{
			Phase:     r.Phase,
			Condition: r.Condition,
			Reason:    r.Reason,
		})
	}
	return &model.App{
		Name:       app.Name,
		Containers: containers,
		Replicas:   int(app.Replicas),
		Routing:    fromRouting(app.GetRouting()),
		Status:     status,
	}
}

func fromRouting(networking *meshpaaspb.Routing) *model.Routing {
	var routes []*model.HTTPRoute
	for _, r := range networking.GetHttpRoutes() {
		route := &model.HTTPRoute{
			Name:             r.Name,
			Port:             int(r.Port),
			PathPrefix:       helpers.StringPointer(r.PathPrefix),
			RewriteURI:       helpers.StringPointer(r.RewriteUri),
			AllowOrigins:     r.AllowOrigins,
			AllowMethods:     r.AllowMethods,
			AllowHeaders:     r.AllowHeaders,
			ExposeHeaders:    r.ExposeHeaders,
			AllowCredentials: helpers.BoolPointer(r.AllowCredentials),
		}
		if r.Name != "" {
			route.Name = r.Name
		}
		if r.Port != 0 {
			route.Port = int(r.Port)
		}
		routes = append(routes, route)
	}
	r := &model.Routing{
		Hosts:      networking.GetHosts(),
		HTTPRoutes: routes,
	}
	if networking.Gateway != "" {
		r.Gateway = &networking.Gateway
	}
	return r
}

func fromTask(app *meshpaaspb.Task) *model.Task {
	var containers []*model.Container
	var completions int
	if app.Completions > 0 {
		completions = int(app.Completions)
	}
	for _, c := range app.Containers {
		var (
			env   map[string]interface{}
			ports []*model.ContainerPort
		)
		if c.Env != nil {
			env = map[string]interface{}{}
			for k, v := range c.Env {
				env[k] = v
			}
		}
		if c.Ports != nil {
			var ports []*model.ContainerPort
			for _, v := range c.Ports {
				ports = append(ports, &model.ContainerPort{
					Name:   v.Name,
					Number: int(v.Number),
					Expose: v.Expose,
				})
			}
		}
		containers = append(containers, &model.Container{
			Name:  c.Name,
			Image: c.Image,
			Args:  c.Args,
			Env:   helpers.ConvertMap(c.Env),
			Ports: ports,
		})
	}

	return &model.Task{
		Name:        app.Name,
		Containers:  containers,
		Schedule:    app.Schedule,
		Completions: &completions,
	}
}
