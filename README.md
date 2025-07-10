# Cookiecutter Kubernetes Go Microservice

A comprehensive Cookiecutter template for creating production-ready Go microservices with Kubernetes deployment support using Chi router.

## Features
- Clean architecture with internal packages
- Chi router for HTTP routing
- HTTP server with JSON API
- Middleware support (logging, CORS)
- Interactive Swagger/OpenAPI documentation
- Graceful shutdown
- Health checks
- Metrics endpoint
- Docker multi-stage builds
- Kubernetes manifests
- Comprehensive documentation
- Makefile for common tasks

## Usage

1. Install Cookiecutter:
```bash
pip install cookiecutter
```

2. Generate project:

```bash
cookiecutter https://github.com/yourorg/cookiecutter-kubernetes-go-microservice
```

3. Configure your service by answering the prompts
4. Navigate to your new project and start developing:

```bash
cd your-service-name
go mod tidy
make run
```

## Template Variables

service_name: Name of your microservice
service_description: Brief description of the service
go_module_name: Go module path
author_name: Your name
author_email: Your email
server_port: HTTP server port
kubernetes_namespace: Kubernetes namespace
docker_image_name: Docker image name
docker_image_tag: Docker image tag
go_version: Go version

## Generated Structure
The template creates a well-organized Go microservice with:

cmd/server/: Application entry point
internal/: Private packages (config, handlers, middleware)
deployment/: Kubernetes and Docker configurations
docs/: Complete documentation with Swagger
Makefile: Common development tasks

## Quick Start After Generation
```bash
# Install dependencies
go mod tidy

# Run the service
make run

# Visit interactive API docs
open http://localhost:8080/docs/

# Test endpoints
curl http://localhost:8080/health
```