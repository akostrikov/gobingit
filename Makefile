.PHONY: test build
PWD := $(shell pwd)

build:
	docker build . -t test-git

test:
	docker run -v ${PWD}:/opt/git test-git