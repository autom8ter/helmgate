# meshpaas

an opinionated graphQL/gRPC API for easily deploying applications & jobs on Istio service mesh

- [API Documentation](https://autom8ter.github.io/meshpaas/)

## Helpful Links
- [Istio Docs](https://istio.io/latest/)
- [Free Online GraphQL Console](graphqlbin.com)

## Command Line

```
Usage of meshpaas:
      --allow-jwt-issuer string   allowed jwt.claim.iss issuer (required) (env: MESHPAAS_ALLOW_ISSUER)
      --debug                     enable debug logs (env: MESHPAAS_DEBUG)
      --jwks-uri string           remote json web key set uri for verifying authorization tokens (required) (env: MESHPAAS_JWKS_URI)
      --listen-port int           serve gRPC & graphQL on this port (env: MESHPAAS_LISTEN_PORT) (default 8820)
      --namespace-claim string    the jwt attribute on the id token that returns the namespace the user is allowed access to (required) (env: MESHPAAS_NAMESPACE_CLAIM) (default "aud")
      --out-of-cluster            enable out of cluster k8s config discovery (env: MESHPAAS_OUT_OF_CLUSTER)

```


## Notes

- the meshpaas API manages resources within the namespace matching the `jwt.${namespace-claim}` claim(see flags) of the identity token of the user making the request
    - this namespace is automatically created

## Concepts

### Gateway
A Gateway is a public Ingress or API-Gateway for traffic coming from the public internet.

Components:
- [istio gateway](https://istio.io/latest/docs/reference/config/networking/gateway/)

### Application

An Application

Components: 
- [k8s namespace](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)
- [k8s deployment](https://kubernetes.io/docs/concepts/workloads/controllers/deployment/)
- [k8s service](https://kubernetes.io/docs/concepts/services-networking/service/)
- [istio virtual service](https://istio.io/latest/docs/reference/config/networking/virtual-service/)
- [istio request authentication policy](https://istio.io/latest/docs/reference/config/security/request_authentication/)
- [istio authorization policy](https://istio.io/latest/docs/reference/config/security/authorization-policy/) 

### Task

Components: 
- [k8s namespace](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)
- [k8s cronjob](https://kubernetes.io/docs/concepts/workloads/controllers/cron-jobs/)

### Secret

Components: 
- [k8s namespace](https://kubernetes.io/docs/concepts/overview/working-with-objects/namespaces/)
- [k8s secret](https://kubernetes.io/docs/concepts/configuration/secret/)