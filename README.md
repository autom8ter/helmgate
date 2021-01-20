# meshpaas

an opinionated graphQL/gRPC API for easily deploying applications & jobs on Istio service mesh

[API Documentation](https://autom8ter.github.io/meshpaas/)

## Command Line

```
Usage of meshpaas:
      --allow-headers strings     cors allow headers (env: MESHPAAS_ALLOW_HEADERS) (default [*])
      --allow-jwt-issuer string   allowed jwt.claim.iss issuer (env: MESHPAAS_ALLOW_ISSUER)
      --allow-methods strings     cors allow methods (env: MESHPAAS_ALLOW_METHODS) (default [HEAD,GET,POST,PUT,PATCH,DELETE])
      --allow-origins strings     cors allow origins (env: MESHPAAS_ALLOW_ORIGINS) (default [*])
      --debug                     enable debug logs (env: MESHPAAS_DEBUG)
      --jwks-uri string           remote json web key set uri for verifying authorization tokens (env: MESHPAAS_JWKS_URI)
      --listen-port int           serve gRPC & graphQL on this port (env: MESHPAAS_LISTEN_PORT) (default 8820)
      --out-of-cluster            enable out of cluster k8s config discovery (env: MESHPAAS_OUT_OF_CLUSTER)
```


