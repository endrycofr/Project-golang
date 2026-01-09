.PHONY: build push build-push run stop logs remove push-app-all

# Docker Hub namespace
NAMESPACE := endrycofr
REGISTRY  := docker.io
IMAGE_NAME := golang-web-app

VERSION := $(shell git describe --tags --always --dirty)

IMAGE   := $(REGISTRY)/$(NAMESPACE)/$(IMAGE_NAME):$(VERSION)
LATEST  := $(REGISTRY)/$(NAMESPACE)/$(IMAGE_NAME):latest

COMPOSE_FILE := deployments/docker-compose/docker-compose.yaml

build:
	docker build \
		-t $(IMAGE) \
		-t $(LATEST) \
		-f deployments/docker/Dockerfile .

	@echo "Image size:"
	@docker images $(IMAGE) --format "{{.Size}}"

push:
	docker push $(IMAGE)
	docker push $(LATEST)

build-push: build push

run:
	docker compose -f $(COMPOSE_FILE) up -d

stop:
	docker compose -f $(COMPOSE_FILE) down -v

logs:
	docker compose -f $(COMPOSE_FILE) logs -f

remove:
	docker system prune -a --volumes --force

# Push semua tag untuk image ini saja
push-app-all:
	@docker images $(REGISTRY)/$(NAMESPACE)/$(IMAGE_NAME) \
		--format "{{.Repository}}:{{.Tag}}" \
	| grep -v "<none>" \
	| while read image; do \
		echo "👉 Pushing $$image"; \
		docker push $$image; \
	done
