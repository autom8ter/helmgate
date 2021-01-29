package testing

import (
	"context"
	helmProxypb "github.com/autom8ter/helmProxy/gen/grpc/go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"testing"
)

func Test(t *testing.T) {
	conn, err := grpc.Dial("localhost:8820", grpc.WithInsecure())
	if err != nil {
		t.Fatal(err.Error())
	}
	client := helmProxypb.NewHelmProxyServiceClient(conn)
	ctx := context.Background()
	ctx = metadata.AppendToOutgoingContext(ctx,
		"Authorization",
		"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c")

	charts, err := client.SearchCharts(ctx, &helmProxypb.ChartFilter{
		Term:  "redis",
		Regex: true,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	if len(charts.GetCharts()) == 0 {
		t.Fatal("expected at least one chart")
	}
	for _, c := range charts.GetCharts() {
		t.Log(c.Name)
	}
}
