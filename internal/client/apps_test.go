package client_test

import (
	"context"
	"github.com/autom8ter/kubego"
	"github.com/autom8ter/meshpaas/internal/client"
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
	cli := client.New(
		kclient,
		iclient,
		lgger,
	)
	cli.CreateSecret(context.Background(), nil)
}
