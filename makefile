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

dump:
	@docker exec postgres pg_dump \
	-U $(POSTGRES_USER) \
	$(POSTGRES_DB) > ./db_backup/$(POSTGRES_DB).sql
