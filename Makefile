PROJECT_NAME=go-telemetry-api
DOCKER_COMPOSE_FILE=docker-compose.yml


.PHONY: run
run: 
	@echo "Starting services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up --build -d

.PHONY: clean
clean:
	@echo "Stopping and removing services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

.PHONY: restart
restart: clean run

.PHONY: logs
logs: 
	@echo "Displaying logs..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

.PHONY: ps
ps: 
	@echo "Displaying service status..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) ps

.PHONY: build
build:
	@echo "Building services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

.PHONY: help
help:
	@echo "Usage: make [target]"
	@echo
	@echo "Targets:"
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-20s %s\n", $$1, $$2}'
