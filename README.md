# Payment Service

[![SonarCloud](https://sonarcloud.io/images/project_badges/sonarcloud-white.svg)](https://sonarcloud.io/summary/new_code?id=gogovan_ggx-kr-Payment-service)

## Overview
Payment Service 
## Prerequisites
- [Go] (https://golang.org/dl/) 1.19 or later
- [Docker] (https://docs.docker.com/get-docker/) 20.10 or later

## Getting Started

### Running the service
1. Clone the repository
2. Run `make run` to start the service
3. Run `make test` to run the tests.
4. Run `make migrate` to create migration.

We have 3 config folder for different environment. You can change the config file in `make run` to run the service in different environment.
Default configuration is set to run on port 5001 for http, 5002 for grpc. you can change it in the `config.yaml` file respectively.

### Docker compose
```bash
$ make compose
```
This will use docker-compose.local.yaml to build up environment

### Build
```bash
$ make build
```
### Docker build
```bash
$ make docker-build
```

## Tools
### Open Telemetry - Jaeger
Distributed tracing system
(https://www.jaegertracing.io/)
### Prometheus
Prometheus exporter for hardware and OS metrics exposed by *NIX kernels, written in Go with pluggable metric collectors.
(https://github.com/prometheus)
### Grafana
Grafana is a multi-platform open source analytics and interactive visualization web application. It provides charts, graphs, and alerts for the web when connected to supported data sources
(https://grafana.com/)
## API
### Swagger
```bash
$ make swagger
```
### Swagger UI
```link
http://localhost:5001/swagger/index.html
```


### Deployment checklist
|                      | Content                                                                                                    | From Infra Teams |
|----------------------|------------------------------------------------------------------------------------------------------------|------------------|
| Routing              | /payment -> /                                                                                              |                  |
| Ports                | HTTP :5000 /api/v1<br/> GRPC :5001 /Payment.PaymentService.*                                               |                  |
| Health Check         | :6001/live <br> :6001/ready                                                                                |                  |
| Metrics              | :8001/metrics                                                                                              |                  |
| Environment Variable | APP_ENV=dev or prod<br/>VAULT_TOKEN<br/>AWS_ACCESS_KEY_ID<br/>AWS_DEFAULT_REGION<br/>AWS_SECRET_ACCESS_KEY | [x]              |
| Note                 | Need to setup authenticate for http request same as User service or Order Service                          |                  |
