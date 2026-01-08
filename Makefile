.PHONY: docker-run docker-run-alt docker-stop

APP_NAME := golang-web-app
VERSION := latest
PORT := 8080
ALT_PORT := 8080

# Run Docker container dengan port default
docker-run:
	@echo "🚀 Running Docker container on port $(PORT)..."
	@docker rm -f $(APP_NAME) 2>/dev/null || true
	docker run -d \
		--name $(APP_NAME) \
		-p $(PORT):8080 \
		-e ENV=production \
		$(APP_NAME):$(VERSION) || \
	(echo "❌ Port $(PORT) tidak tersedia, coba: make docker-run-alt" && exit 1)
	@echo "✅ Container started"
	@echo "🔗 Access at: http://localhost:$(PORT)"

# Run dengan port alternatif
docker-run-alt:
	@echo "🚀 Running Docker container on port $(ALT_PORT)..."
	@docker rm -f $(APP_NAME) 2>/dev/null || true
	docker run -d \
		--name $(APP_NAME) \
		-p $(ALT_PORT):8080 \
		-e ENV=production \
		$(APP_NAME):$(VERSION)
	@echo "✅ Container started"
	@echo "🔗 Access at: http://localhost:$(ALT_PORT)"

# Run dengan custom port
docker-run-custom:
	@read -p "Enter port number: " port; \
	echo "🚀 Running Docker container on port $$port..."; \
	docker rm -f $(APP_NAME) 2>/dev/null || true; \
	docker run -d \
		--name $(APP_NAME) \
		-p $$port:8080 \
		-e ENV=production \
		$(APP_NAME):$(VERSION); \
	echo "✅ Container started"; \
	echo "🔗 Access at: http://localhost:$$port"

# Stop container
docker-stop:
	@echo "🛑 Stopping container..."
	@docker stop $(APP_NAME) 2>/dev/null || true
	@docker rm $(APP_NAME) 2>/dev/null || true
	@echo "✅ Container stopped"

# Check available ports
check-ports:
	@echo "🔍 Checking common ports..."
	@for port in 8080 8081 9000 3000 5000; do \
		if ! netstat -ano | findstr :$$port > /dev/null 2>&1; then \
			echo "✅ Port $$port is available"; \
		else \
			echo "❌ Port $$port is in use"; \
		fi; \
	done