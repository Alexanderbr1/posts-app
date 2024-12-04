build:
	docker-compose build posts-app

run:
	docker-compose up posts-app

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up

swag:
	swag init -g cmd/main.go