# Variables
BINARY_NAME=myapp
DOCKER_IMAGE=mydockerhubusername/myapp:latest
K8S_NAMESPACE=default
K8S_DEPLOYMENT=myapp-deployment
HEALTH_CHECK_URL=http://localhost:8080/health

# Build Go binary
build:
	go build -o $(BINARY_NAME) ./cmd/main.go

# Run unit tests
test:
	go test ./... -v

# Run linting using golangci-lint
lint:
	golangci-lint run ./...

# Build and push Docker image using a multi-stage build
docker-build:
	docker build -t $(DOCKER_IMAGE) .
	docker push $(DOCKER_IMAGE)

# Deploy to Kubernetes
deploy:
	kubectl apply -f k8s/deployment.yaml -n $(K8S_NAMESPACE)

# Check service health using curl
health-check:
	@echo "Checking health of the application..."
	curl --fail $(HEALTH_CHECK_URL) || (echo "Health check failed!" && exit 1)
	@echo "Health check passed!"

# Clean build artifacts
clean:
	rm -f $(BINARY_NAME)

# Run all tasks in sequence
all: clean build test lint docker-build deploy health-check