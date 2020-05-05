#!/usr/bin/env bash

CONTAINER_ID_FILE=.db-container-id

DOCKER_NETWORK_NAME=$1
DB_HOSTNAME=$2
DB_PORT=$3

# run local dynamodb in docker
# returns: container id
run_local_dynamodb() {
  docker \
    run \
    -d \
    --rm \
    --network "${DOCKER_NETWORK_NAME}" \
    --name "${DB_HOSTNAME}" \
    -p "${DB_PORT}":"${DB_PORT}" \
    amazon/dynamodb-local:1.12.0 \
    -jar DynamoDBLocal.jar -sharedDb -inMemory
}

main() {
  printf "[start_local_db][main][INFO] spinning up local database...\n"
  run_local_dynamodb > ${CONTAINER_ID_FILE}
  sleep 2

  printf "[start_local_db][main][INFO] database ready: http://%s:%s\n" "${DB_HOSTNAME}" "${DB_PORT}"
}

main
