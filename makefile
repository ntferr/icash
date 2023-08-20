ENV := $(PWD)/.env
include $(ENV)
export

run-api: run-migrate
	go run cmd/icash/main.go

run-migrate: run-dependecies
	go run migrate/main.go

run-build:
	go build cmd/icash/main.go

run-dependecies:
	docker-compose up -d