#!/usr/bin/env bash

CONTAINER_ID_FILE=.db-container-id

main() {
  if [[ ! -f "${CONTAINER_ID_FILE}" ]]; then
    exit
  fi

  DB_CONTAINER_ID=$(cat ${CONTAINER_ID_FILE})

  # stop the container
  docker kill "${DB_CONTAINER_ID}" 1>/dev/null

  # cleanup the container id file
  rm -f ${CONTAINER_ID_FILE}
}

main
