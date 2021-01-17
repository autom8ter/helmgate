package client_test

import (
	"context"
	kdeploypb "github.com/autom8ter/kdeploy/gen/grpc/go"
	"github.com/autom8ter/kdeploy/internal/client"
	"github.com/autom8ter/kdeploy/internal/logger"
	"github.com/autom8ter/kubego"
	"testing"
)

func TestTasks(t *testing.T) {
	kclient, err := kubego.NewOutOfClusterClient()
	if err != nil {
		t.Fatal(err.Error())
	}
	lgger := logger.New(true)
	cli := client.New(
		kclient,
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
	tsk, err := cli.CreateTask(context.Background(), &kdeploypb.TaskConstructor{
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
