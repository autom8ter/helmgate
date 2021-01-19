package client_test

import (
	"context"
	"github.com/autom8ter/kubego"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/client"
	"github.com/autom8ter/meshpaas/internal/logger"
	"testing"
)

func TestTasks(t *testing.T) {
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
		nil,
		"https://openidconnect.googleapis.com/v1/userinfo",
		nil,
	)
	namespaces, err := cli.ListNamespaces(context.Background())
	if err != nil {
		t.Fatal(err.Error())
	}
	for _, n := range namespaces.GetNamespaces() {
		t.Log(n)
	}
	tsk, err := cli.CreateTask(context.Background(), &meshpaaspb.TaskInput{
		Name:        "echo-date",
		Namespace:   "colemanw",
		Image:       "busybox",
		Args:        []string{"/bin/sh", "-c", "date; echo Hello from the Kubernetes cluster"},
		Env:         nil,
		Schedule:    "*/1 * * * *",
		Completions: 0,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	t.Logf("created task: %s", tsk.String())
}
