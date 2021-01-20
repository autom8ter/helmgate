# meshpaas

an opinionated graphQL/gRPC API for easily deploying applications & jobs on Istio service mesh

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


