# Standardization of Cluster resource
Proposal of a CRD for scs-cluster standardization. In this repo we started to move some parameters out of the param-discussion document into a CRD (Custom Resource Definition). This has the benefit that this is a machine-readable format with simple validation (or even arbitrary validation via Webhooks).

You can find the CRD in `scs-cluster-crd.yaml` and an example custom object in `scs-example-cluster.yaml`. You can  apply the CRD with 

```
$ kubectl apply -f scs-cluster-crd.yaml
customresourcedefinition.apiextensions.k8s.io/clusters.scs.community created
```
Afterwards you can create custom objects of the type clusters.scs.community, k8s will validate against the spec and apply the resource if the validation was successful. The contents of an example custom object currently looks like this(totally wip, this wont work):

```yaml
apiVersion: scs.community/v1alpha1
kind: Cluster
metadata:
  name: example-cluster
spec:
  addons:
    nginxIngress:
      enabled: true
    certManager:
      enabled: true
    metricsServer:
      enabled: true
    harbor:
      enabled: false
    flux:
      enabled: false
  kubernetes:
    version: v1.24.0-rc.1
  provider:
    providerType: hcloud
  workers:
    - name: wg1
      flavor: cpx31
      count: 3
```

An apply and a get will look like this:

```
$ kubectl apply -f scs-example-cluster.yaml 
cluster.scs.community/example-cluster created
$ kubectl get clusters.scs.community
NAME              K8S-VERSION    PROVIDER   INGRESS
example-cluster   v1.24.0-rc.1   hcloud     true
```
