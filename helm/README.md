## Prerequisites

### Minikube

Minikube is a tool that makes it easy to run Kubernetes locally. Minikube runs a single-node Kubernetes cluster inside a Virtual Machine (VM) on your laptop.

- Distro specific install instructions [here](https://kubernetes.io/docs/tasks/tools/install-minikube/), choose the `Virtualbox` driver option


### Brig cli 

Brig is the Brigade command line client. You can use brig to create/update/delete new brigade Projects, run Builds, etc. To get `brig`, navigate to the [Releases](https://github.com/brigadecore/brigade/releases/) page and then download the appropriate client for your platform. For example, if you’re using Linux or WSL, you can get the 1.4.0 version in this way:

```sh
# Note the k8s client used in brig < 1.4.0 is not compatible with k8s >= 1.18
wget -O brig https://github.com/brigadecore/brigade/releases/download/v1.4.0/brig-linux-amd64
chmod +x brig
sudo mv brig /usr/local/bin/
```


### Helm 3.x

Helm helps you manage Kubernetes applications — Helm Charts help you define, install, and upgrade even the most complex Kubernetes application.

```shell
# TIP: If you are using OS X you can install it with the brew install command: brew install helm.
curl https://raw.githubusercontent.com/kubernetes/helm/master/scripts/get-helm-3 > get_helm.sh
chmod 700 get_helm.sh
./get_helm.sh
```
## Setup
- Add [GHCR Auth](https://docs.github.com/en/packages/guides/pushing-and-pulling-docker-images#authenticating-to-github-container-registry)
- Request to join the [@orchestrator](https://github.com/orgs/knowship-io/teams/orchestrator) team

## 🚀 Launch Minikube / Orchestrator

```bash
# install helm dependencies 
make install-deps

# spin up minikube, 
make start

# ensure services was deployed
watch kubectl get all

# launch the kashti brigade dashboard
make kashti
```

## Launch a brigade build

```sh
curl --header "Content-Type: application/json" \
  --request POST \
  --data '{
    "key1": "value1",
    "key2": "value2"
}' \
  http://localhost:8081/simpleevents/v1/brigade-b47f9114e066cb0d78f91eff72e6813dc2ee897c1aeeafbc88fff0/genericsecret
```