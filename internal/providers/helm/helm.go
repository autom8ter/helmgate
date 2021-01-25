package helm

import (
	"context"
	"github.com/autom8ter/kubego"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/golang/protobuf/ptypes/empty"
)

type Helm struct {
	client *kubego.Helm
}

func (h Helm) ListProjects(ctx context.Context, empty *empty.Empty) (*meshpaaspb.ProjectRefs, error) {
	panic("implement me")
}

func (h Helm) GetApp(ctx context.Context, ref *meshpaaspb.AppRef) (*meshpaaspb.App, error) {
	panic("implement me")
}

func (h Helm) ListApps(ctx context.Context, ref *meshpaaspb.ProjectRef) (*meshpaaspb.Apps, error) {
	panic("implement me")
}

func (h Helm) UninstallApp(ctx context.Context, ref *meshpaaspb.AppRef) (*meshpaaspb.Apps, error) {
	panic("implement me")
}

func (h Helm) RollbackApp(ctx context.Context, ref *meshpaaspb.AppRef) (*meshpaaspb.App, error) {
	panic("implement me")
}

func (h Helm) CreateApp(ctx context.Context, input *meshpaaspb.AppInput) (*meshpaaspb.App, error) {
	panic("implement me")
}

func (h Helm) UpdateApp(ctx context.Context, input *meshpaaspb.AppInput) (*meshpaaspb.App, error) {
	panic("implement me")
}

func (h Helm) SearchAppTemplates(ctx context.Context, filter *meshpaaspb.Filter) (*meshpaaspb.AppTemplates, error) {
	panic("implement me")
}


