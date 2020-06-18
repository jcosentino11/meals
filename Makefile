
BIN_DIR := ./bin

API_CMD := cmd/meals-api/main.go
API_OUTPUT = $(BIN_DIR)/meals-api

GOCACHE := $(shell realpath ${HOME}/.cache) # makes sure path is normalized
GOENV := GOCACHE=$(GOCACHE)

clean:
	rm -rf $(BIN_DIR)

init-modules:
	$(GOENV) go mod init $(shell dirname $(API_CMD))

build-api: clean
	$(GOENV) go build -o $(API_OUTPUT) $(API_CMD)

run-api:
	$(GOENV) go run $(API_CMD)