# Troubleshoot Talk

## Set up

This requires [krew](https://krew.sigs.k8s.io/)

```sh
kubectl krew install support-bundle
```

```sh
kubectl create secret -n default generic mysecret --from-file=secrets.yaml

kubectl apply -f deployment.yaml
```

## Troubleshoot

```sh
kubectl support-bundle ./support-bundle.yaml
```

Slide deck

https://docs.google.com/presentation/d/1MC_-yq1E0Yp8cGCedWrhLTOvTxBEAFHQCzGefYe6C1Y/edit?usp=sharing