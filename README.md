# Standardization of Cluster resource
Proposal of a CRD for scs-cluster standardization. In this repo we started to move some parameters out of the param-discussion document into a CRD (Custom Resource Definition). This has the benefit that this is a machine-readable format with simple validation (or even arbitrary validation via Webhooks).

You can find the CRD in `scs-cluster-crd.yaml` and an example custom object in `scs-example-cluster.yaml`. You can  apply the CRD with `kubectl apply -f scs-cluster-crd.yaml`. 
