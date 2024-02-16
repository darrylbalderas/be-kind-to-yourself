# be-kind-to-yourself
Repo getting started using KIND on Mac


## Steps

```bash
kind create cluster --config kind-config.yaml --name development

kind get clusters

kind delete cluster --name development

kubectl cluster-info --context kind-development

kind load docker-image local_kind_app:latest --name development

k apply -f manifests/deployment.yaml

k logs app-5f69685f7-8x7c9

kubectly delete -f manifests/deployment.yaml
```