ENV := $(PWD)/.env
include $(ENV)
export

run-api:
	docker-compose up -d
	go run cmd/icash/main.go

run-migrate:
	go run migrate/main.go

run-build:
	go build cmd/icash/main.go

run-dependecies:
	docker-compose up -d