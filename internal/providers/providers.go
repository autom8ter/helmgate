package providers

import (
	"github.com/autom8ter/kubego"
	meshpaaspb "github.com/autom8ter/meshpaas/gen/grpc/go"
	"github.com/autom8ter/meshpaas/internal/providers/helm"
	"github.com/pkg/errors"
)

type Backend string

const (
	K8sHelm Backend = "k8s-helm"
)

func GetBackend(backend Backend) (meshpaaspb.MeshPaasServiceServer, error) {
	switch backend {
	case K8sHelm:
		client, err := kubego.NewHelm()
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get backend: %s", backend)
		}
		return helm.NewHelm(client), nil
	default:
		return nil, errors.Errorf("unsupported backend: %s", backend)
	}
}