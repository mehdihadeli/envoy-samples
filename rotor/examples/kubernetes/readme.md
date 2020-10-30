Kubernetes requires a number of RBAC objects to be created before running Rotor. The easiest way to create all these is via the YAML file in this repo:

``` bash
kubectl create -f https://raw.githubusercontent.com/turbinelabs/rotor/master/examples/kubernetes/kubernetes-rotor.yaml
```
Rotor discovers clusters by looking for active pods in Kubernetes and grouping them based on their labels. You will have to add two pieces of information to each pod to have Rotor recognize it:

- A `tbn_cluster: <name>` label to name the service to which the Pod belongs. The label key can be customized.

- An exposed port named http or a customized name. A pod must have both the tbn_cluster label and a port named http to be collected by Rotor.

Rotor will also collect all other labels on the Pod, which can be used for routing.

An example of a pod with labels correctly configured is included here. An example Envoy-simple yaml is also included.