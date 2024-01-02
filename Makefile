include .env

FILES_DIR := files
PROJECT_NAME := banner-writer
PROJECT_BINARY := $(FILES_DIR)/$(PROJECT_NAME)

DOCKER := docker run --rm -it -v ${PWD}:${PWD} --env-file .env -w ${PWD} ${PROJECT_NAME}

.PHONY: build_dir
.PHONY: build_project
.PHONY: build_docker_image
.PHONY: run docker/run

build_dir:
	mkdir -p ${FILES_DIR}

build_project:
	go build -o $(PROJECT_BINARY) .

build_docker_image:
	docker build -t ${PROJECT_NAME} .

run: build_dir build_project
	./${PROJECT_BINARY} 

docker/run: build_docker_image
	${DOCKER} make run