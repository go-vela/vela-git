# Copyright (c) 2020 Target Brands, Inc. All rights reserved.
#
# Use of this source code is governed by the LICENSE file in this repository.

.PHONY: build
build: binary-build

.PHONY: run
run: build docker-build docker-run

.PHONY: test
test: build docker-build docker-example

#################################
######      Go clean       ######
#################################

.PHONY: clean
clean:

	@go mod tidy
	@go vet ./...
	@go fmt ./...
	@echo "I'm kind of the only name in clean energy right now"

#################################
######    Build Binary     ######
#################################

.PHONY: binary-build
binary-build:

	GOOS=linux CGO_ENABLED=0 go build -o release/vela-git github.com/go-vela/vela-git/cmd/vela-git

#################################
######    Docker Build     ######
#################################

.PHONY: docker-build
docker-build:

	docker build --no-cache -t vela-git:local .

#################################
######     Docker Run      ######
#################################

.PHONY: docker-run
docker-run:

	docker run --rm \
		-e PARAMETER_REMOTE \
		-e PARAMETER_PATH \
		-e PARAMETER_REF \
		-e PARAMETER_SHA \
		-e PARAMETER_TAGS \
		-e PARAMETER_SUBMODULES \
		-e VELA_NETRC_MACHINE \
		-e VELA_NETRC_USERNAME \
		-e VELA_NETRC_PASSWORD \
		vela-git:local

.PHONY: docker-example
docker-example:

	docker run --rm \
		-e PARAMETER_REMOTE=https://github.com/octocat/hello-world.git \
		-e PARAMETER_PATH=home/octocat_hello-world_1 \
		-e PARAMETER_REF=refs/heads/master \
		-e PARAMETER_SHA=7fd1a60b01f91b314f59955a4e4d4e80d8edf11d \
		-e PARAMETER_TAGS \
		-e PARAMETER_SUBMODULES \
		-e VELA_NETRC_MACHINE=github.com \
		-e VELA_NETRC_USERNAME=octocat \
		-e VELA_NETRC_PASSWORD=superSecretPassword \
		vela-git:local
