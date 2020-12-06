## Prerequisites

### Minikube

## Install helm 3.x

```shell
# TIP: If you are using OS X you can install it with the brew install command: brew install helm.
curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get-helm-3 > get_helm.sh
chmod 700 get_helm.sh
./get_helm.sh
```

### Minikube

Minikube is a tool that makes it easy to run Kubernetes locally. Minikube runs a single-node Kubernetes cluster inside a Virtual Machine (VM) on your laptop.

- Distro specific install instructions [here](https://kubernetes.io/docs/tasks/tools/install-minikube/), choose the `Virtualbox` driver option


## 🚀 Launch Minikube / Orchestrator

```bash
# spin up minikube, 
make start

# ensure services was deployed
kubectl get all

# launch the kashti brigade dashboard
make kashti
```