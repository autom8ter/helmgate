package core_test

import (
	"context"
	"github.com/autom8ter/kubego"
	"github.com/autom8ter/meshpaas/internal/core"
	"github.com/autom8ter/meshpaas/internal/logger"
	"testing"
)

func TestApps(t *testing.T) {
	kclient, err := kubego.NewOutOfClusterKubeClient()
	if err != nil {
		t.Fatal(err.Error())
	}
	iclient, err := kubego.NewOutOfClusterIstioClient()
	if err != nil {
		t.Fatal(err.Error())
	}
	lgger := logger.New(true)
	cli := core.NewManager(
		kclient,
		iclient,
		lgger,
		"aud",
	)
	cli.CreateSecret(context.Background(), nil)
}
