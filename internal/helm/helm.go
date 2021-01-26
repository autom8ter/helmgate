package helm

import (
	"context"
	"github.com/autom8ter/kubego"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/logger"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"helm.sh/helm/v3/pkg/repo"
)

type Helm struct {
	logger *logger.Logger
	client *kubego.Helm
}

func NewHelm(logger *logger.Logger, repos []*repo.Entry) (*Helm, error) {
	client, err := kubego.NewHelm()
	if err != nil {
		return nil, err
	}
	for _, r := range repos {
		if err := client.AddRepo(r); err != nil {
			return nil, err
		}
	}
	if err := client.UpdateRepos(); err != nil {
		return nil, err
	}
	return &Helm{client: client, logger: logger}, nil
}

func (h Helm) ListProjects(ctx context.Context, empty *empty.Empty) (*meshpaaspb.ProjectRefs, error) {
	return nil, status.Error(codes.Unimplemented, "unimplemented")
}

func (h Helm) GetApp(ctx context.Context, ref *meshpaaspb.AppRef) (*meshpaaspb.App, error) {
	release, err := h.client.Get(ref.Project, ref.Name)
	if err != nil {
		return nil, err
	}
	return h.toApp(release)
}

func (h Helm) ListApps(ctx context.Context, ref *meshpaaspb.ProjectRef) (*meshpaaspb.Apps, error) {
	releases, err := h.client.ListReleases(ref.Name)
	if err != nil {
		return nil, err
	}
	return h.toApps(releases)
}

func (h Helm) UninstallApp(ctx context.Context, ref *meshpaaspb.AppRef) (*empty.Empty, error) {
	_, err := h.client.Uninstall(ref.Project, ref.Name)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (h Helm) RollbackApp(ctx context.Context, ref *meshpaaspb.AppRef) (*meshpaaspb.App, error) {
	if err := h.client.Rollback(ref.Project, ref.Name); err != nil {
		return nil, err
	}
	return h.GetApp(ctx, ref)
}

func (h Helm) CreateApp(ctx context.Context, input *meshpaaspb.AppInput) (*meshpaaspb.App, error) {
	release, err := h.client.Install(input.Project, input.TemplateName, input.AppName, true, input.Config)
	if err != nil {
		return nil, err
	}
	return h.toApp(release)
}

func (h Helm) UpdateApp(ctx context.Context, input *meshpaaspb.AppInput) (*meshpaaspb.App, error) {
	release, err := h.client.Upgrade(input.Project, input.TemplateName, input.AppName, true, input.Config)
	if err != nil {
		return nil, err
	}
	return h.toApp(release)
}

func (h Helm) SearchAppTemplates(ctx context.Context, filter *meshpaaspb.Filter) (*meshpaaspb.AppTemplates, error) {
	charts, err := h.client.SearchCharts(filter.Term, filter.Regex)
	if err != nil {
		return nil, err
	}
	return h.toTempalates(charts), nil
}
