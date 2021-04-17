setup:
	make build
	docker-compose up -d

build:
	docker-compose build

destroy:
	docker-compose down