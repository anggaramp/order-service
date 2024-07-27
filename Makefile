run:
	go run main.go -env=local

setup:
	docker-compose -f docker/dev/docker-compose.yaml up -d

build-image:
	docker build --tag order-service .