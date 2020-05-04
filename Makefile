SHELL = bash

APP := $(shell basename $(PWD) | tr '[:upper:]' '[:lower:]')

GIT_COMMIT_HASH ?= $(shell git rev-parse HEAD)
GIT_SHORT_COMMIT_HASH := $(shell git rev-parse --short HEAD)

.PHONY: all
all: test

.PHONY: ci_setup
ci_setup: ## Setup the ci environment
	@wget -O flow.tar.gz https://github.com/gomicro/flow/releases/latest/download/flow_linux_amd64.tar.gz
	@tar xvf flow.tar.gz -C /home/travis/.local/bin

.PHONY: clean
clean: ## Cleans out all generated items
	-@rm -f output.txt
	-@rm -rf coverage
	-@rm -f coverage.txt
	-@docker-compose down

.PHONY: coverage
coverage: ## Generates the code coverage from all the tests
	docker run -v $$PWD:/go$${PWD/$$GOPATH} --workdir /go$${PWD/$$GOPATH} gomicro/gocover

.PHONY: dockerize
dockerize: ecr_login  ## Create a docker image of the project
	docker build \
	--build-arg BUILD_PATH=/go$${PWD/$$GOPATH} \
	-t dan9186/superman .

.PHONY: ecr_login
ecr_login:  ## Login to the ECR using local credentials
	@eval $$(flow aws ecr auth)

.PHONY: help
help:  ## Show This Help
	@for line in $$(cat Makefile | grep "##" | grep -v "grep" | sed  "s/:.*##/:/g" | sed "s/\ /!/g"); do verb=$$(echo $$line | cut -d ":" -f 1); desc=$$(echo $$line | cut -d ":" -f 2 | sed "s/!/\ /g"); printf "%-30s--%s\n" "$$verb" "$$desc"; done

.PHONY: linters
linters: ## Run all the linters
	docker run -v $$PWD:/go$${PWD/$$GOPATH} --workdir /go$${PWD/$$GOPATH} gomicro/golinters

.PHONY: pull_dependencies
pull_dependencies: ## Forcibly pull the latest images of other dependent services

.PHONY: run
run: dockerize pull_dependencies ## Run a dockerized version of the app
	docker-compose up -d

.PHONY: test
test: unit_test functional_test ## Run all available tests

.PHONY: unit_test
unit_test: ## Run unit tests
	go test ./...

.PHONY: functional_test
functional_test: ## Runs the functional tests against the running service
	docker run -i -v $$PWD/features:/usr/app/features --network=$(APP)_services gomicro/cucumber cucumber $$CUKE_TAGS
