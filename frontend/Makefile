
CWD := $(shell cygpath -w $(shell pwd))

docker:
	docker run \
		-it \
		--rm \
		-p '3000:3000' \
		-w /app \
		-v $(CWD):/app \
		--entrypoint /bin/bash \
		node:14
