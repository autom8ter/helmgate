package helm

import (
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/golang/protobuf/ptypes/timestamp"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"helm.sh/helm/v3/cmd/helm/search"
	"helm.sh/helm/v3/pkg/release"
)

func (h Helm) toApp(release *release.Release) (*meshpaaspb.App, error) {
	config, err := structpb.NewStruct(release.Config)
	if err != nil {
		h.logger.Error("bad config values", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "bad config values")
	}
	app := &meshpaaspb.App{
		Name:    release.Name,
		Project: release.Namespace,
		Version: uint32(release.Version),
		Config:  config,
		LifeCycle: &meshpaaspb.LifeCycle{
			Notes:       release.Info.Notes,
			Description: release.Info.Description,
			Status:      release.Info.Status.String(),
			Timestamps: map[string]*timestamp.Timestamp{
				"deleted": timestamppb.New(release.Info.Deleted.Time),
				"created": timestamppb.New(release.Info.FirstDeployed.Time),
				"updated": timestamppb.New(release.Info.LastDeployed.Time),
			},
		},
		Template: &meshpaaspb.AppTemplate{
			Name:        release.Chart.Name(),
			Home:        release.Chart.Metadata.Home,
			Description: release.Chart.Metadata.Description,
			Version:     release.Chart.Metadata.Version,
			Sources:     release.Chart.Metadata.Sources,
			Keywords:    release.Chart.Metadata.Keywords,
			Icon:        release.Chart.Metadata.Icon,
			Deprecated:  release.Chart.Metadata.Deprecated,
			Metadata:    release.Chart.Metadata.Annotations,
		},
	}
	for _, m := range release.Chart.Metadata.Maintainers {
		app.Template.Maintainers = append(app.Template.Maintainers, &meshpaaspb.Maintainer{
			Name:  m.Name,
			Email: m.Email,
		})
	}
	for _, d := range release.Chart.Metadata.Dependencies {
		app.Template.Dependencies = append(app.Template.Dependencies, &meshpaaspb.Dependency{
			TemplateName: d.Name,
			Version:      d.Version,
			Repository:   d.Repository,
		})
	}

	return app, nil
}

func (h Helm) toApps(releases []*release.Release) (*meshpaaspb.Apps, error) {
	apps := &meshpaaspb.Apps{}
	for _, r := range releases {
		a, err := h.toApp(r)
		if err != nil {
			return nil, err
		}
		apps.Apps = append(apps.Apps, a)
	}
	return apps, nil
}

func (h Helm) toTempalates(resultes []*search.Result) *meshpaaspb.AppTemplates {
	t := &meshpaaspb.AppTemplates{}
	for _, r := range resultes {
		tmpl := &meshpaaspb.AppTemplate{
			Name:        r.Name,
			Home:        r.Chart.Home,
			Description: r.Chart.Description,
			Version:     r.Chart.Version,
			Sources:     r.Chart.Sources,
			Keywords:    r.Chart.Keywords,
			Icon:        r.Chart.Icon,
			Deprecated:  r.Chart.Deprecated,
			Metadata:    r.Chart.Annotations,
		}
		for _, m := range r.Chart.Metadata.Maintainers {
			tmpl.Maintainers = append(tmpl.Maintainers, &meshpaaspb.Maintainer{
				Name:  m.Name,
				Email: m.Email,
			})
		}

		for _, d := range r.Chart.Metadata.Dependencies {
			tmpl.Dependencies = append(tmpl.Dependencies, &meshpaaspb.Dependency{
				TemplateName: d.Name,
				Version:      d.Version,
				Repository:   d.Repository,
			})
		}
		t.Templates = append(t.Templates)
	}
	return t
}
