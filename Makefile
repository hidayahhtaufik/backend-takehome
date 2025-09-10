.PHONY: swag build up down logs restart

# Install Swaggo
install-swag:
	cd ./app&& go install github.com/swaggo/swag/cmd/swag@v1.8.12

# Generate Swagger
swag:
	cd ./app && swag init -g main.go -o ./docs

# Build image
build:
	docker compose build

# Start container
up:
	docker compose up -d

#  Stop & remove container
down:
	docker compose down

# Check Logs
logs:
	docker compose logs -f app

# Restart app container
restart:
	docker compose restart app
