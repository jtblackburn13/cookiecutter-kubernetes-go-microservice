.PHONY: build run test clean docker-build docker-run deploy swagger-validate

APP_NAME={{cookiecutter.service_name}}
DOCKER_IMAGE={{cookiecutter.docker_image_name}}:{{cookiecutter.docker_image_tag}}
PORT={{cookiecutter.server_port}}

build:
	go build -o bin/$(APP_NAME) ./cmd/server

run:
	go run ./cmd/server

test:
	go test ./...

clean:
	rm -rf bin/

docker-build:
	docker build -t $(DOCKER_IMAGE) .

docker-run:
	docker run -p $(PORT):$(PORT) $(DOCKER_IMAGE)

deploy:
	kubectl apply -f deployment/kubernetes/

deploy-delete:
	kubectl delete -f deployment/kubernetes/

swagger-validate:
	@if command -v swagger >/dev/null 2>&1; then \
		swagger validate docs/swagger.yaml; \
	else \
		echo "Swagger CLI not installed. Install with: go install github.com/go-swagger/go-swagger/cmd/swagger@latest"; \
	fi

swagger-serve:
	@if command -v swagger >/dev/null 2>&1; then \
		swagger serve docs/swagger.yaml --port=8081; \
	else \
		echo "Swagger CLI not installed. Use the built-in docs at http://localhost:$(PORT)/docs/"; \
	fi

.DEFAULT_GOAL := build