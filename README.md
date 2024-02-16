# be-kind-to-yourself
Repo getting started using KIND on Mac


## Steps

```bash
kind create cluster --config kind-config.yaml --name development

kind get clusters

kind delete cluster --name development

kubectl cluster-info --context kind-development
```