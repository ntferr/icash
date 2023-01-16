ENV := $(PWD)/.env
include $(ENV)
export

run-api:
	docker-compose up -d
	go run cmd/main.go

run-migrate:
	go run migrate/main.go