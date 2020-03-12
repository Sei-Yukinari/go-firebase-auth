# Variables
PROJECT = go-firebase-auth
COMPOSE_PATH := ./docker-compose.yml
COMPOSE_COMMAND := docker-compose -f $(COMPOSE_PATH) -p $(PROJECT)
SERVICE_API := api
SERVIVE_DB := db

### Start selected environment
.PHONY: start
start:
	${COMPOSE_COMMAND} up -d

### Build and Start selected environment
.PHONY: restart
restart:
	${COMPOSE_COMMAND} kill \
 && ${COMPOSE_COMMAND} rm -f \
 && ${COMPOSE_COMMAND} up -d --build

### Shut down selected environment
.PHONY: stop
stop:
	${COMPOSE_COMMAND} kill \
 && ${COMPOSE_COMMAND} rm -f \

### Dell selected environment
.PHONY: down
down:
	${COMPOSE_COMMAND} down -v

 ## List containers
.PHONY: ps
ps:
	${COMPOSE_COMMAND} ps

### Show logs from selected environment
.PHONY: logs
ifeq (logs,$(firstword $(MAKECMDGOALS)))
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  $(eval $(RUN_ARGS):;@:)
endif
logs:
	${COMPOSE_COMMAND} logs -f $(RUN_ARGS)

### Enter Conteiner
.PHONY: exec
exec:
	${COMPOSE_COMMAND} exec ${SERVICE_API} ash

.PHONY: exec-db
exec-db:
	${COMPOSE_COMMAND} exec ${SERVIVE_DB} bash
