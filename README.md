# sample-golang-app
Sample Golang app.

## Usage

### Build Container Image

```
docker build -t sample-golang-app .
```

### Deploy to Kubernetes with Helm

```
helm install sample-golang-app ./deploy/charts/sample-golang-app
```

### Deploy to Docker

```
# Replace CHANGEME with the tenant name
docker run -dt -p 8080:10000 -e TENANT_NAME=CHANGEME sample-golang-app
```

## Local Development

### Pre-requisites

* Golang v1.19
* Docker

### Build

```
# Install dependencies
go get -d -v ./...

# Install the package
go install -v ./...
```

### Run

```
# Replace CHANGEME with the tenant name
TENANT_NAME=CHANGEME sample-golang-app
```