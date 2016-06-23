# devnation app

## Demo

```
kubectl run devnation --image=kelseyhightower/devnation:1.0.0 --port=8080
```

```
kubectl scale deployments devnation --replicas=3
```

```
kubectl expose deployments devnation
```

## Build Guide

Edit main.go to update HTML.

### Build the go binary

The following command will output a static binary that will be used with the Dockerfile.

```
GOOS=linux bash build
```

You should now have the devnation binary.

### Build the Docker image

The following command requires that the devnation binary exist first.

```
docker build -t <username>/devnation:1.0.0 .
```
