# k8s-in-cluster-config
Sample to show K8s in-cluster config

## How to use
Apply the configuration in `manifest.yaml` to create the resources.
This creates a namespace `test-in-cluster-config` and creates a Role and RoleBinding which can list pods in that namespace.
The pod runs an image built from `main.go` which attempts to create an in-cluster config and then use that to list the pods in the namespace where it is running.

```
kubectl apply -f manifest.yaml
kubectl logs -n test-in-cluster-config list-pods
```
