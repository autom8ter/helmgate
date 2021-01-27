package gql

import (
	"github.com/autom8ter/hpaas/gen/gql/go/model"
	hpaaspb "github.com/autom8ter/hpaas/gen/grpc/go"
	"github.com/autom8ter/hpaas/internal/helpers"
)

func toAppRef(ref model.AppRef) *hpaaspb.AppRef {
	return &hpaaspb.AppRef{
		Namespace: ref.Namespace,
		Name:      ref.Name,
	}
}

func toAppInput(ref model.AppInput) *hpaaspb.AppInput {
	return &hpaaspb.AppInput{
		Namespace: ref.Namespace,
		Chart:     ref.Chart,
		AppName:   ref.AppName,
		Config:    helpers.ConvertMapS(ref.Config),
	}
}

func gqlChart(app *hpaaspb.Chart) *model.Chart {
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

func gqlApp(app *hpaaspb.App) *model.App {
	return &model.App{
		Name:      app.Name,
		Namespace: app.Namespace,
		Release:   gqlRelease(app.Release),
		Chart:     gqlTemplate(app.Chart),
	}
}

func gqlRelease(release *hpaaspb.Release) *model.Release {
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

func gqlTemplate(template *hpaaspb.Chart) *model.Chart {
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

func gqlMaintainers(maintainer []*hpaaspb.Maintainer) []*model.Maintainer {
	var maintainers []*model.Maintainer
	for _, m := range maintainer {
		maintainers = append(maintainers, &model.Maintainer{
			Name:  m.Name,
			Email: m.Email,
		})
	}
	return maintainers
}

func gqlDependencies(maintainer []*hpaaspb.Dependency) []*model.Dependency {
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
