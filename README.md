# Build and install dummy-pdf-or-png project in Kubernetes

The Purpose of this assignment is to a develop a microservice that takes HTTP GET requests with a random ID and request a document from the microservice that's provided and return that document with the correct MIME type.

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
![Cluster](Kuber.jpg)

## Update helm chart values file
Here, the replica set is defined and directs to the image tthat is in the repository and define a tag for it. 


## Install helm chart

```bash
# helm install -n <namesapce> <chart-name> <chart-directory/source> --create-namespace
helm install -n dummy dummy dummy --create-namespace
```   
