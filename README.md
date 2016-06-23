# devnation app


```
kubectl run devnation --image=kelseyhightower/devnation:1.0.0 --port=8080
```

```
kubectl scale deployments devnation --replicas=3
```

```
kubectl expose deployments devnation
```
