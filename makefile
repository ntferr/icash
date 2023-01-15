ENV := $(PWD)/.env
include $(ENV)
export

run-migrate:
	go run migrate/main.go 