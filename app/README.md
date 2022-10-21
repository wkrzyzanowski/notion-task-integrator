# Development

During development activities it is possible to run backend and frontend part of application separately.

&nbsp;

## Backend

Application will run on **http://localhost:8080/**

```sh
cd app-backend/ && \
 go run main.go
```

&nbsp;

## Frontend (not valid yet)

Application will run on **http://localhost:8081/**

```sh
cd app-frontend/ && \
 npm install && \
 npm run serve
```

When application is running on dev-server all requests prefixed with **/api** are redirected to the default backend which is set to **http://localhost:8080/**

&nbsp;

# Docker

## Build image

In order to build image let's execute following command:

```sh
make docker
```

It will build image with complete application which is runnable via:

```sh
docker run -d -p 8080:8080 notion-task-integrator:1.0.0
```

&nbsp;

# Kubernetes (Helm)

Feture will be available soon...
