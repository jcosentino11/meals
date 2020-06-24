
CWD := $(shell cygpath -w $(shell pwd))

docker:
	docker run \
		-it \
		--rm \
		-p '3000:3000' \
		-w /app \
		-v $(CWD):/app \
		--user $(shell id -u):$(shell id -g) \
		--entrypoint /bin/bash \
		node:14
