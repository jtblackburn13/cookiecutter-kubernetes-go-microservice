# {{cookiecutter.service_name}}

{{cookiecutter.service_description}}

## Features
- HTTP server with JSON API using Chi router
- Interactive Swagger/OpenAPI documentation
- Graceful shutdown
- Health checks (liveness and readiness)
- Metrics endpoint
- Version information
- CORS middleware
- Request logging
- Built-in Chi middleware (RequestID, Logger, Recoverer, Timeout)
- Kubernetes ready
- Docker support

## Quick Start

### Local Development
```bash
# Install dependencies
go mod tidy

# Run the service
go run ./cmd/server
```

Visit `http://localhost:{{cookiecutter.server_port}}/docs/` for interactive API documentation.

### Docker
```bash
# Build image
docker build -t {{cookiecutter.docker_image_name}}:{{cookiecutter.docker_image_tag}} .

# Run container
docker run -p {{cookiecutter.server_port}}:{{cookiecutter.server_port}} {{cookiecutter.docker_image_name}}:{{cookiecutter.docker_image_tag}}
```

### Kubernetes
```bash
# Deploy to cluster
kubectl apply -f deployment/kubernetes/

# Port forward for local testing
kubectl port-forward svc/{{cookiecutter.service_name}}-service {{cookiecutter.server_port}}:80
```

## API Endpoints
- `GET /` - Welcome message
- `GET /health` - Health check
- `GET /ready` - Readiness check
- `GET /metrics` - Service metrics
- `GET /version` - Version information
- `GET /docs/` - Interactive API documentation
- `GET /swagger.yaml` - OpenAPI specification

## Documentation
- **Interactive API Docs**: `http://localhost:{{cookiecutter.server_port}}/docs/`
- [API Documentation](docs/api.md)
- [Deployment Guide](docs/deployment.md)
- [Development Guide](docs/development.md)

## Configuration
Set environment variables:
- `PORT`: Server port (default: {{cookiecutter.server_port}})
- `ENVIRONMENT`: Environment name (default: development)

## Testing
```bash
# Test all endpoints
curl http://localhost:{{cookiecutter.server_port}}/
curl http://localhost:{{cookiecutter.server_port}}/health
curl http://localhost:{{cookiecutter.server_port}}/ready
curl http://localhost:{{cookiecutter.server_port}}/metrics
curl http://localhost:{{cookiecutter.server_port}}/version
```

## Chi Router Features
This service uses Chi router which provides:
- Clean, idiomatic route definitions
- Powerful middleware ecosystem
- URL parameter handling
- Route grouping and mounting
- Built-in middleware for common tasks

## Author
{{cookiecutter.author_name}} ({{cookiecutter.author_email}})