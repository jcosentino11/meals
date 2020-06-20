API_CMD := cmd/meals-api/main.go

GOCACHE := $(shell realpath ${HOME}/.cache) # makes sure path is normalized
GOENV := GOCACHE=$(GOCACHE)

clean:
	cd backend && make clean

build-api: clean
	cd backend && make build

build-api-image: clean
	cd backend && make build-image

run-api:
	cd backend && make run

docker:
	docker-compose up --build
