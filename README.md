# kdeploy

![create-redis](assets/create-redis.png)
![get-redis](assets/get-redis.png)

an OAuth-protected graphQL API for deploying "serverless" kubernetes applications

[graphQL Documentation](https://autom8ter.github.io/kdeploy/)

## Features

- [x] [gRPC API](kdeploy.proto)
    - [x] [golang client sdk](gen/grpc/go)
    - [x] [javascript client sdk](gen/grpc/js)
    - [x] [java client sdk](gen/grpc/java)
    - [x] [php client sdk](gen/grpc/php)
    - [x] [python client sdk](gen/grpc/python)
    - [x] [ruby client sdk](gen/grpc/ruby)
    - [x] [csharp client sdk](gen/grpc/csharp)
- [x] [graphQL API](schema.graphql)
- [x] SSO/Oauth protected graphQL playground(autocomplete/schema-documentation/query-console)
- [x] Run in cluster
- [x] Run out of cluster
- [x] OpenID Connect based Authentication
- [x] Create Application
- [x] Update Application
- [x] Get Application
- [x] Destroy Application
- [ ] Search Applications
- [x] Stream Application Logs
- [ ] Expression based Application "authorizers"(execute against open-id profile)


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
curl https://raw.githubusercontent.com/autom8ter/kdeploy/master/k8s.yaml >> k8s.yaml
```

inspect the manifest and add/adjust environmental variables in the deployment spec(see flags for supported environmental variables)

then run:

    kubectl apply -f k8s.yaml

to view pods as they spin up, run:

    kubectl get pods -n kdeploy -w

Kdeploy is intended to be deployed behind an SSL ingress/proxy and doesn't handle TLS termination.
