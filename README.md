# kdeploy

## Command Line

```
Usage of kdeploy:
      --allow-headers strings        cors allow headers (env: KDEPLOY_ALLOW_HEADERS) (default [*])
      --allow-methods strings        cors allow methods (env: KDEPLOY_ALLOW_METHODS) (default [HEAD,GET,POST,PUT,PATCH,DELETE])
      --allow-origins strings        cors allow origins (env: KDEPLOY_ALLOW_ORIGINS) (default [*])
      --debug                        enable debug logs (env: KDEPLOY_DEBUG)
      --listen-port int              serve gRPC & graphQL on this port (env: KDEPLOY_LISTEN_PORT) (default 8820)
      --oauth-client-id string       playground oauth client id (env: KDEPLOY_OAUTH_CLIENT_ID)
      --oauth-client-secret string   playground oauth client secret (env: KDEPLOY_OAUTH_CLIENT_SECRET)
      --oauth-redirect string        playground oauth redirect (env: KDEPLOY_OAUTH_REDIRECT)
      --open-id string               open id connect discovery uri ex: https://accounts.google.com/.well-known/openid-configuration (env: KDEPLOY_OPEN_ID) (required)
      --out-of-cluster               enable out of cluster k8s config discovery (env: KDEPLOY_OUT_OF_CLUSTER)
```

## Installation

Given a running Kubernetes cluster, run:

```yaml
curl https://raw.githubusercontent.com/graphikDB/kdeploy/master/k8s.yaml >> k8s.yaml
```

inspect the manifest and add/adjust environmental variables in the deployment spec(see flags for supported environmental variables)

then run:

    kubectl apply -f k8s.yaml

to view pods as they spin up, run:

    kubectl get pods -n graphik-system -w

Kdeploy is intended to be deployed behind an SSL ingress/proxy and doesn't handle TLS termination. 
Please see [gproxy](https://github.com/graphikDB/gproxy) if an easy to use ingress/proxy for Kubernetes is needed.