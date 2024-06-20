include deploy/app.env

############################################################
# Переменные
############################################################

CONTNAME_LOCAL_PSQL=local-psql

DIR_LOCAL_DEPLOY=./deploy

ENV_LOCAL_FILE=${DIR_LOCAL_DEPLOY}/app.env

DEPLOY_LOCAL_DOCKER_COMPOSE_FILE=${DIR_LOCAL_DEPLOY}/docker-compose.yaml

############################################################
# Команды
############################################################

# run
run:
	go run cmd/main.go

# Сборка контейнера local_postgres
up_local: down_local
	@echo "Building (when required) and starting docker images..."
	docker compose --env-file=${ENV_LOCAL_FILE} -f ${DEPLOY_LOCAL_DOCKER_COMPOSE_FILE} up --build -d

# Сброс контейнера BD(локально)
down_local:
	@echo "Stopping docker images (if running...)"
	docker compose --env-file=${ENV_LOCAL_FILE} -f ${DEPLOY_LOCAL_DOCKER_COMPOSE_FILE} down

# Пересоздать локальную базу данных
psql_local_recreate:
	@echo "Dropping local Postgres database..."
	docker exec -id ${CONTNAME_LOCAL_PSQL} dropdb --username=$(PSQL_USER) $(PSQL_DB)
	@echo "Creating local Postgres database..."
	docker exec -id ${CONTNAME_LOCAL_PSQL} createdb --username=$(PSQL_USER) --owner=$(PSQL_USER) $(PSQL_DB)
