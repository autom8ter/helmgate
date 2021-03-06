package service

import (
	helmgatepb "github.com/autom8ter/helmgate/gen/grpc/go"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/structpb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"helm.sh/helm/v3/cmd/helm/search"
	"helm.sh/helm/v3/pkg/release"
)

func (h Helm) toApp(release *release.Release) (*helmgatepb.App, error) {
	config, err := structpb.NewStruct(release.Config)
	if err != nil {
		h.logger.Error("bad config values", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "bad config values")
	}
	app := &helmgatepb.App{
		Name:      release.Name,
		Namespace: release.Namespace,
		Release: &helmgatepb.Release{
			Version:     uint32(release.Version),
			Config:      config,
			Notes:       release.Info.Notes,
			Description: release.Info.Description,
			Status:      release.Info.Status.String(),
			Timestamps: &helmgatepb.Timestamps{
				Created: timestamppb.New(release.Info.FirstDeployed.Time),
				Updated: timestamppb.New(release.Info.LastDeployed.Time),
				Deleted: timestamppb.New(release.Info.Deleted.Time),
			},
		},
		Chart: &helmgatepb.Chart{
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
		app.Chart.Maintainers = append(app.Chart.Maintainers, &helmgatepb.Maintainer{
			Name:  m.Name,
			Email: m.Email,
		})
	}
	for _, d := range release.Chart.Metadata.Dependencies {
		app.Chart.Dependencies = append(app.Chart.Dependencies, &helmgatepb.Dependency{
			Chart:      d.Name,
			Version:    d.Version,
			Repository: d.Repository,
		})
	}

	return app, nil
}

func (h Helm) toApps(releases []*release.Release) (*helmgatepb.Apps, error) {
	apps := &helmgatepb.Apps{}
	for _, r := range releases {
		a, err := h.toApp(r)
		if err != nil {
			return nil, err
		}
		apps.Apps = append(apps.Apps, a)
	}
	return apps, nil
}

func (h Helm) toTempalates(results []*search.Result) *helmgatepb.Charts {
	t := &helmgatepb.Charts{}
	for _, r := range results {
		tmpl := &helmgatepb.Chart{
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
			tmpl.Maintainers = append(tmpl.Maintainers, &helmgatepb.Maintainer{
				Name:  m.Name,
				Email: m.Email,
			})
		}

		for _, d := range r.Chart.Metadata.Dependencies {
			tmpl.Dependencies = append(tmpl.Dependencies, &helmgatepb.Dependency{
				Chart:      d.Name,
				Version:    d.Version,
				Repository: d.Repository,
			})
		}
		t.Charts = append(t.Charts, tmpl)
	}
	return t
}
