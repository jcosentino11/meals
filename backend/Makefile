
BIN_DIR := ./bin

API_CMD := main.go

GOCACHE := $(shell realpath ${HOME}/.cache) # makes sure path is normalized
GOENV := GOCACHE=$(GOCACHE)

clean:
	rm -rf $(BIN_DIR)

init-modules:
	$(GOENV) go mod init $(shell dirname $(API_CMD))

build: clean
	$(GOENV) go build -o $(BIN_DIR)/main $(API_CMD)

build-image: clean
	docker build -t meals-api .

run:
	$(GOENV) go run $(API_CMD)
