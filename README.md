# MyApp

MyApp is a simple Go backend application that logs incoming HTTP requests, including the requester's IP address and user agent, to a PostgreSQL database. The application is built using Gin, GORM, and follows clean architecture principles. It also exports some data as Prometheus metrics for monitoring.

## Project Structure

myapp/
├── Makefile
├── Dockerfile
├── go.mod
├── go.sum
├── cmd/
│ └── main.go
├── configs/
│ └── config.yaml
├── internal/
│ ├── domain/
│ │ └── request.go
│ ├── usecase/
│ │ └── request_usecase.go
│ ├── repository/
│ │ └── request_repository.go
│ └── delivery/
│ └── http/
│ └── request_handler.go
├── scripts/
│ ├── check_namespace.sh
│ └── migrate.sh
└── helm/
└── myapp/
├── Chart.yaml
├── values.yaml
└── templates/
├── deployment.yaml
├── service.yaml
└── ingress.yaml

## Getting Started

### Prerequisites

- Go 1.18+
- Docker
- Kubernetes cluster
- ArgoCD
- PostgreSQL database

### Installation

1. Clone the repository:

```sh
git clone https://github.com/saniyar-dev/arvancloud-challenge-app.git
cd arvancloud-challenge-app
```

2. Build the application:

`make build`

3. Run the application locally:
   `make run`
