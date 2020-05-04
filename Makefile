# this Makefile depends on
# scripts from https://github.com/jcosentino11/scripts

DOCKER_NETWORK_NAME := meals-local
DB_HOST_NAME := db
DB_PORT := 8000
LOCAL_TEMPLATE := template-local.yml
GENERATED_TEMPLATE := .aws-sam/build/template.yaml
DEBUG_PORT := 3001

clean: db-kill network-kill

run-local: clean network-create db
	sam build \
		--template $(LOCAL_TEMPLATE)
	sam local start-api \
		--debug \
		--profile local \
		--docker-network $(DOCKER_NETWORK_NAME) \
		--parameter-overrides 'DbEndpoint=http://$(DB_HOST_NAME):$(DB_PORT)' \
		--template $(GENERATED_TEMPLATE)

debug-local: clean network-create db
	sam build \
		--template $(LOCAL_TEMPLATE)
	sam local start-api \
		--debug \
		--debug-port $(DEBUG_PORT) \
		--profile local \
		--docker-network $(DOCKER_NETWORK_NAME) \
		--parameter-overrides 'DbEndpoint=http://$(DB_HOST_NAME):$(DB_PORT)' \
		--template $(GENERATED_TEMPLATE)

db: login-local
	@./scripts/start_db.sh $(DOCKER_NETWORK_NAME) $(DB_HOST_NAME) $(DB_PORT)

db-kill:
	@./scripts/kill_db.sh

network-create:
	@docker network create $(DOCKER_NETWORK_NAME)

network-kill:
	@docker network rm $(DOCKER_NETWORK_NAME) || true

login-local:
	@aws_login local

# logout of aws-cli
logout:
	@aws_logout
