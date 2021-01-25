package providers

import (
	"github.com/autom8ter/kubego"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/providers/helm"
)

func GetHelmProvider(client *kubego.Helm) meshpaaspb.MeshPaasServiceServer {
	return helm.NewHelm(client)
}
