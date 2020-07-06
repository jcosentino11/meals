
BIN_DIR := ./bin

API_CMD := *.go

GOCACHE := $(shell realpath ${HOME}/.cache) # makes sure path is normalized
GOENV := GOCACHE=$(GOCACHE)

clean:
	rm -rf $(BIN_DIR)

go-clean:
	$(GOENV) go clean -cache -modcache -i -r

build: clean
	$(GOENV) go build -o $(BIN_DIR)/main $(API_CMD)

build-image: clean
	docker build -t meals-api .

run:
	$(GOENV) go run $(API_CMD)
