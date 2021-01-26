package gql

import (
	"github.com/autom8ter/meshpaas/gen/gql/go/model"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/helpers"
)

func toAppRef(ref model.AppRef) *meshpaaspb.AppRef {
	return &meshpaaspb.AppRef{
		Project: ref.Project,
		Name:    ref.Name,
	}
}

func toAppInput(ref model.AppInput) *meshpaaspb.AppInput {
	return &meshpaaspb.AppInput{
		Project:      ref.Project,
		TemplateName: ref.TemplateName,
		AppName:      ref.AppName,
		Config:       helpers.ConvertMapS(ref.Config),
	}
}

func gqlAppTemplate(app *meshpaaspb.AppTemplate) *model.AppTemplate {
	return &model.AppTemplate{
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

func gqlApp(app *meshpaaspb.App) *model.App {
	return &model.App{
		Name:     app.Name,
		Project:  app.Project,
		Release:  gqlRelease(app.Release),
		Template: gqlTemplate(app.Template),
	}
}

func gqlRelease(release *meshpaaspb.Release) *model.Release {
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

func gqlTemplate(template *meshpaaspb.AppTemplate) *model.AppTemplate {
	return &model.AppTemplate{
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

func gqlMaintainers(maintainer []*meshpaaspb.Maintainer) []*model.Maintainer {
	var maintainers []*model.Maintainer
	for _, m := range maintainer {
		maintainers = append(maintainers, &model.Maintainer{
			Name:  m.Name,
			Email: m.Email,
		})
	}
	return maintainers
}

func gqlDependencies(maintainer []*meshpaaspb.Dependency) []*model.Dependency {
	var deps []*model.Dependency
	for _, m := range maintainer {
		deps = append(deps, &model.Dependency{
			TemplateName: m.GetTemplateName(),
			Version:      m.GetVersion(),
			Repository:   m.GetRepository(),
		})
	}
	return deps
}
