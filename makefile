ENV := $(PWD)/.env
include $(ENV)
export

setup: migrate
	go run cmd/icash/main.go

migrate: dependecies
	go run migrate/main.go

run-build:
	go build cmd/icash/main.go

dependecies:
	docker-compose up -d