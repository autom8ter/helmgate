package gql

import (
	"github.com/autom8ter/kdeploy/gen/gql/go/model"
	kdeploypb "github.com/autom8ter/kdeploy/gen/grpc/go"
	"github.com/autom8ter/kdeploy/internal/helpers"
	"github.com/spf13/cast"
)

func toAppC(input model.AppInput) *kdeploypb.AppInput {
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
	return &kdeploypb.AppInput{
		Name:      input.Name,
		Namespace: input.Namespace,
		Image:     input.Image,
		Env:       env,
		Ports:     ports,
		Replicas:  uint32(input.Replicas),
		Args:      input.Args,
	}
}

func toTaskC(input model.TaskInput) *kdeploypb.TaskInput {
	var env map[string]string
	var completions uint32
	if input.Env != nil {
		env = cast.ToStringMapString(input.Env)
	}
	if input.Completions != nil {
		completions = uint32(*input.Completions)
	}
	return &kdeploypb.TaskInput{
		Name:        input.Name,
		Namespace:   input.Namespace,
		Image:       input.Image,
		Args:        input.Args,
		Env:         env,
		Schedule:    input.Schedule,
		Completions: completions,
	}
}

func toAppU(input model.AppInput) *kdeploypb.AppInput {
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
	return &kdeploypb.AppInput{
		Name:       input.Name,
		Namespace:  input.Namespace,
		Image:      input.Image,
		Env:        env,
		Ports:      ports,
		Replicas:   uint32(input.Replicas),
		Args:       input.Args,
		Networking: toNetworking(input.Networking),
	}
}

func toTaskU(input model.TaskInput) *kdeploypb.TaskInput {
	var (
		env         map[string]string
		completions int
	)
	if input.Env != nil {
		env = cast.ToStringMapString(input.Env)
	}
	if input.Completions != nil {
		completions = int(*input.Completions)
	}
	return &kdeploypb.TaskInput{
		Name:        input.Name,
		Namespace:   input.Namespace,
		Image:       input.Image,
		Args:        input.Args,
		Env:         env,
		Schedule:    input.Schedule,
		Completions: uint32(completions),
	}
}

func fromApp(app *kdeploypb.App) *model.App {
	var (
		env    map[string]interface{}
		ports  map[string]interface{}
		status = &model.AppStatus{}
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
		Name:       app.Name,
		Namespace:  app.Namespace,
		Image:      app.Image,
		Args:       app.Args,
		Env:        env,
		Ports:      ports,
		Replicas:   int(app.Replicas),
		Networking: fromNetworking(app.GetNetworking()),
		Status:     status,
	}
}

func toNetworking(input *model.NetworkingInput) *kdeploypb.Networking {
	if input == nil {
		return nil
	}
	var routes []*kdeploypb.HTTPRoute
	for _, r := range input.HTTPRoutes {
		route := &kdeploypb.HTTPRoute{
			AllowOrigins:  r.AllowOrigins,
			AllowMethods:  r.AllowMethods,
			AllowHeaders:  r.AllowHeaders,
			ExposeHeaders: r.ExposeHeaders,
		}
		if r.Name != nil {
			route.Name = *r.Name
		}
		if r.Port != nil {
			p := uint32(*r.Port)
			route.Port = p
		}
		if r.RewriteURI != nil {
			route.RewriteUri = *r.RewriteURI
		}
		if r.PathPrefix != nil {
			route.PathPrefix = *r.PathPrefix
		}
		if r.AllowCredentials != nil {
			route.AllowCredentials = *r.AllowCredentials
		}
		routes = append(routes, route)
	}
	n := &kdeploypb.Networking{
		Gateways:   input.Gateways,
		Hosts:      input.Hosts,
		HttpRoutes: routes,
	}
	if input.Export != nil {
		n.Export = *input.Export
	}
	return n
}

func fromNetworking(networking *kdeploypb.Networking) *model.Networking {
	var routes []*model.HTTPRoute
	for _, r := range networking.GetHttpRoutes() {
		route := &model.HTTPRoute{
			PathPrefix:       helpers.StringPointer(r.PathPrefix),
			RewriteURI:       helpers.StringPointer(r.RewriteUri),
			AllowOrigins:     r.AllowOrigins,
			AllowMethods:     r.AllowMethods,
			AllowHeaders:     r.AllowHeaders,
			ExposeHeaders:    r.ExposeHeaders,
			AllowCredentials: helpers.BoolPointer(r.AllowCredentials),
		}
		if r.Name != "" {
			route.Name = &r.Name
		}
		if r.Port != 0 {
			p := int(r.Port)
			route.Port = &p
		}
		routes = append(routes, route)
	}
	return &model.Networking{
		Gateways:   networking.GetGateways(),
		Hosts:      networking.GetHosts(),
		Export:     &networking.Export,
		HTTPRoutes: routes,
	}
}

func fromTask(app *kdeploypb.Task) *model.Task {
	var (
		env         map[string]interface{}
		completions int
	)
	if app.Env != nil {
		env = map[string]interface{}{}
		for k, v := range app.Env {
			env[k] = v
		}
	}
	if app.Completions > 0 {
		completions = int(app.Completions)
	}
	return &model.Task{
		Name:        app.Name,
		Namespace:   app.Namespace,
		Image:       app.Image,
		Args:        app.Args,
		Env:         env,
		Schedule:    app.Schedule,
		Completions: &completions,
	}
}
