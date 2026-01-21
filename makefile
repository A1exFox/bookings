include .env

ARG ?=

COMPOSE = docker compose

up:
	@$(COMPOSE) down
	@$(COMPOSE) up -d

build:
	@$(COMPOSE) down
	@$(COMPOSE) build

down:
	@$(COMPOSE) down


remove:
	@echo "Stopping:"
	@docker stop adminer postgres
	@echo ""
	@echo "Removing:"
	@docker rm adminer postgres
	@docker network rm localnet

# make soda-create ARG="CreateUserTable"
soda-create:
	soda generate fizz $(ARG)

migrate:
	soda migrate

migrate-down:
	soda migrate down
