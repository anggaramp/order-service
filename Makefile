run:
	go run main.go -env=local

setup:
	docker-compose -f docker/dev/docker-compose.yaml up -d
	go mod tidy

build-image:
	docker build --tag order-service .