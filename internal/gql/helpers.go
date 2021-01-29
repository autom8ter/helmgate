package gql

import (
	"github.com/autom8ter/helmProxy/gen/gql/go/model"
	helmProxypb "github.com/autom8ter/helmProxy/gen/grpc/go"
	"github.com/autom8ter/helmProxy/internal/helpers"
)

func toAppRef(ref model.AppRef) *helmProxypb.AppRef {
	return &helmProxypb.AppRef{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
}

func toAppInput(ref model.AppInput) *helmProxypb.AppInput {
	return &helmProxypb.AppInput{
		Namespace: ref.Namespace,
		Chart:     ref.Chart,
		AppName:   ref.AppName,
		Config:    helpers.ConvertMapS(ref.Config),
	}
}

func gqlChart(app *helmProxypb.Chart) *model.Chart {
	return &model.Chart{
		Name:         app.Name,
		Home:         helpers.ToStringPointer(app.Home),
		Icon:         helpers.ToStringPointer(app.Icon),
		Version:      helpers.ToStringPointer(app.Version),
		Description:  helpers.ToStringPointer(app.Description),
		Sources:      app.Sources,
		Keywords:     app.Keywords,
		Deprecated:   helpers.ToBoolPointer(app.Deprecated),
		Metadata:     helpers.ConvertMap(app.Metadata),
		Maintainers:  gqlMaintainers(app.GetMaintainers()),
		Dependencies: gqlDependencies(app.GetDependencies()),
	}
}

func gqlApp(app *helmProxypb.App) *model.App {
	return &model.App{
		Name:      app.Name,
		Namespace: app.Namespace,
		Release:   gqlRelease(app.Release),
		Chart:     gqlTemplate(app.Chart),
	}
}

func gqlRelease(release *helmProxypb.Release) *model.Release {
	return &model.Release{
		Version:     int(release.GetVersion()),
		Config:      release.Config.AsMap(),
		Notes:       helpers.ToStringPointer(release.GetNotes()),
		Description: helpers.ToStringPointer(release.GetDescription()),
		Status:      helpers.ToStringPointer(release.GetStatus()),
		Timestamps: &model.Timestamps{
			Created: helpers.ToTimePointer(release.GetTimestamps().GetCreated().AsTime()),
			Updated: helpers.ToTimePointer(release.GetTimestamps().GetUpdated().AsTime()),
			Deleted: helpers.ToTimePointer(release.GetTimestamps().GetDeleted().AsTime()),
		},
	}
}

func gqlTemplate(template *helmProxypb.Chart) *model.Chart {
	return &model.Chart{
		Name:         template.GetName(),
		Home:         helpers.ToStringPointer(template.GetHome()),
		Icon:         helpers.ToStringPointer(template.GetIcon()),
		Version:      helpers.ToStringPointer(template.GetVersion()),
		Description:  helpers.ToStringPointer(template.GetDescription()),
		Sources:      template.GetSources(),
		Keywords:     template.GetKeywords(),
		Deprecated:   helpers.ToBoolPointer(template.GetDeprecated()),
		Metadata:     helpers.ConvertMap(template.GetMetadata()),
		Maintainers:  gqlMaintainers(template.GetMaintainers()),
		Dependencies: gqlDependencies(template.GetDependencies()),
	}
}

func gqlMaintainers(maintainer []*helmProxypb.Maintainer) []*model.Maintainer {
	var maintainers []*model.Maintainer
	for _, m := range maintainer {
		maintainers = append(maintainers, &model.Maintainer{
			Name:  m.Name,
			Email: m.Email,
		})
	}
	return maintainers
}

func gqlDependencies(maintainer []*helmProxypb.Dependency) []*model.Dependency {
	var deps []*model.Dependency
	for _, m := range maintainer {
		deps = append(deps, &model.Dependency{
			Chart:      m.GetChart(),
			Version:    m.GetVersion(),
			Repository: m.GetRepository(),
		})
	}
	return deps
}
