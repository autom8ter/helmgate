package client_test

import (
	"context"
	"github.com/autom8ter/kdeploy/internal/client"
	"github.com/autom8ter/kdeploy/internal/logger"
	"github.com/autom8ter/kubego"
	"testing"
)

func Test(t *testing.T) {
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
}
