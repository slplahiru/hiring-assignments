# Build and install dummy-pdf-or-png project in Kubernetes

#TODO: intro 

## Prerequisites
* Linux VM
* Docker installed
* k3s installed
* kubectl installed
* helm CLI installed

## Build and push Docker image
Docker image needs to be built locally and push into the Docker Hub public repository

```bash
cd dummy-pdf-or-png
docker build -t slplahiru/dummy-test:1.0 .
docker push slplahiru/dummy-test:1.0
```

## Update helm chart values file
#TODO: write

## Install helm chart

```bash
# helm install -n <namesapce> <chart-name> <chart-directory/source> --create-namespace
helm install -n dummy dummy dummy --create-namespace
```   
