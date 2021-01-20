# meshpaas

an opinionated, OAuth-protected graphQL/gRPC API for easily deploying applications & jobs on Istio service mesh

[API Documentation](https://autom8ter.github.io/meshpaas/)

## Command Line

```
Usage of meshpaas:
      --allow-headers strings         cors allow headers (env: KDEPLOY_ALLOW_HEADERS) (default [*])
      --allow-methods strings         cors allow methods (env: KDEPLOY_ALLOW_METHODS) (default [HEAD,GET,POST,PUT,PATCH,DELETE])
      --allow-origins strings         cors allow origins (env: KDEPLOY_ALLOW_ORIGINS) (default [*])
      --debug                         enable debug logs (env: KDEPLOY_DEBUG) (default true)
      --listen-port int               serve gRPC & graphQL on this port (env: KDEPLOY_LISTEN_PORT) (default 8820)
      --out-of-cluster                enable out of cluster k8s config discovery (env: KDEPLOY_OUT_OF_CLUSTER)
```

## Installation

Given a running Kubernetes cluster, run:

```yaml
curl https://raw.githubusercontent.com/autom8ter/meshpaas/master/k8s.yaml >> k8s.yaml
```

inspect the manifest and add/adjust environmental variables in the deployment spec(see flags for supported environmental variables)

then run:

    kubectl apply -f k8s.yaml

to view pods as they spin up, run:

    kubectl get pods -n meshpaas -w

