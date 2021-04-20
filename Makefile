setup:
	make build
	docker-compose up -d

build:
	docker-compose build

test:
	go test -v ./...

destroy:
	docker-compose down