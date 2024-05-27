APP_NAME = arvancloud-challenge-app
DOCKER_IMAGE = saniyar-dev/$(APP_NAME)
POSTGRES_CONTAINER = arvancloud-challenge-app-postgres
POSTGRES_PORT = 5432
POSTGRES_USER = saniyar
POSTGRES_PASSWORD = localtest
POSTGRES_DB = arvancloud-challenge-appdb

.PHONY: all build run docker-build docker-run docker-push postgres-up postgres-down migrate

all: build

build:
	@echo "Building the application..."
	go build -o $(APP_NAME) ./cmd

run: postgres-up
	@echo "Running the application after 1min"
	sleep 60
	./$(APP_NAME)

docker-build:
	@echo "Building Docker image..."
	docker build -t $(DOCKER_IMAGE) .

docker-run: postgres-up
	@echo "Running Docker container..."
	docker run --rm -d -p 8080:8080 --name $(APP_NAME) --network host $(DOCKER_IMAGE)

docker-push:
	@echo "Pushing Docker image to registry..."
	docker push $(DOCKER_IMAGE)

postgres-up:
	@echo "Starting PostgreSQL container..."
	docker run --rm -d \
		--name $(POSTGRES_CONTAINER) \
		-p $(POSTGRES_PORT):5432 \
		-e POSTGRES_USER=$(POSTGRES_USER) \
		-e POSTGRES_PASSWORD=$(POSTGRES_PASSWORD) \
		-e POSTGRES_DB=$(POSTGRES_DB) \
		postgres:13

postgres-down:
	@echo "Stopping PostgreSQL container..."
	docker stop $(POSTGRES_CONTAINER)

