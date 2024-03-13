# Define variables
DOCKER_COMPOSE = docker compose
URI = "host=localhost port=5430 user=admin password=password123 dbname=devsoc-24-backend"

# Targets
.PHONY: build up down logs restart clean migrate-up migrate-down

build:
	$(DOCKER_COMPOSE) up --build -d

up:
	$(DOCKER_COMPOSE) up -d

down:
	$(DOCKER_COMPOSE) down

logs:
	$(DOCKER_COMPOSE) logs -f --tail 100

restart:
	$(DOCKER_COMPOSE) restart

clean:
	$(DOCKER_COMPOSE) down -v

migrate-up:
	cd db/migrations && goose postgres $(URI) up && cd ../..

migrate-down:
	cd db/migrations && goose postgres $(URI) down-to 0 && cd ../..

# Help target
help:
	@echo "Usage: make [target]"
	@echo "Targets:"
	@echo "  build      Build Docker containers"
	@echo "  up         Start Docker containers in the background"
	@echo "  down       Stop and remove Docker containers"
	@echo "  logs       View logs of Docker containers"
	@echo "  restart    Restart Docker containers"
	@echo "  clean      Stop, remove containers, and also remove volumes (data)"
	@echo "  help       Display this help message"
