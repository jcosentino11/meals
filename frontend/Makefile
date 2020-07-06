
CWD := $(shell cygpath -w $(shell pwd))

clean:
	rm -rf ./build

# open a shell in docker container,
# with all the dependencies needed to run the 
# application.
#
# Mapping to /frontend is a silly hack to get
# vscode debugging to correctly map source files.
docker:
	docker run \
		-it \
		--rm \
		-p '3000:3000' \
		-p '5000:5000' \
		-w /frontend \
		-v $(CWD):/frontend \
		--entrypoint /bin/bash \
		node:14
# --user $(shell id -u):$(shell id -g) \

# start development server, port 3000
start:
	yarn start

# create optimized build
build:
	yarn build

# serve optimized build files locally, port 5000.
serve:
	yarn global add serve
	serve -s build

# refresh build files and serve
serve-fresh: clean build serve
