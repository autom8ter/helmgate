package client_test

import (
	"context"
	"github.com/autom8ter/kubego"
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
	)
	cli.CreateSecret(context.Background(), nil)
	if err != nil {
		t.Fatal(err.Error())
	}
	//for _, n := range namespaces.GetProjects() {
	//	t.Log(n)
	//}
	//tsk, err := cli.CreateTask(context.Background(), &meshpaaspb.TaskInput{
	//	Name:    "echo-date",
	//	Project: "colemanw",
	//	Containers: []*meshpaaspb.Container{
	//		{
	//			Image: "busybox",
	//			Args:  []string{"/bin/sh", "-c", "date; echo Hello from the Kubernetes cluster"},
	//			Env:   nil,
	//		},
	//	},
	//	Schedule:    "*/1 * * * *",
	//	Completions: 0,
	//})
	//if err != nil {
	//	t.Fatal(err.Error())
	//}
	//t.Logf("created task: %s", tsk.String())
}
