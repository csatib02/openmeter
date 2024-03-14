# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

export PATH := $(abspath bin/):${PATH}

.PHONY: up
up: ## Start the dependencies via docker compose
	$(call print-target)
	docker compose --profile ksqldb up -d

.PHONY: down
down: ## Stop the dependencies via docker compose
	$(call print-target)
	docker compose down --remove-orphans --volumes

.PHONY: generate
generate: ## Generate code
	$(call print-target)
	go generate ./...

.PHONY: build-server
build-server: ## Build server binary
	$(call print-target)
	go build -o build/server ./cmd/server

.PHONY: build-sink-worker
build-sink-worker: ## Build binary
	$(call print-target)
	go build -o build/sink-worker ./cmd/sink-worker

config.yaml:
	cp config.example.yaml config.yaml

.PHONY: server
server: ## Run sink-worker
	@ if [ config.yaml -ot config.example.yaml ]; then diff -u config.yaml config.example.yaml || (echo "!!! The configuration example changed. Please update your config.yaml file accordingly (or at least touch it). !!!" && false); fi
	$(call print-target)
	$(AIR_BIN) -c ./cmd/server/.air.toml

.PHONY: sink-worker
sink-worker: ## Run sink-worker
	@ if [ config.yaml -ot config.example.yaml ]; then diff -u config.yaml config.example.yaml || (echo "!!! The configuration example changed. Please update your config.yaml file accordingly (or at least touch it). !!!" && false); fi
	$(call print-target)
	$(AIR_BIN) -c ./cmd/sink-worker/.air.toml

.PHONY: test
test: ## Run tests
	$(call print-target)
	$(DAGGER_BIN) call test

.PHONY: lint
lint: ## Run linters
	$(call print-target)
	$(DAGGER_BIN) call lint all

.PHONY: license-check
license-check: ## Run license check
	$(call print-target)
	$(LICENSEI_BIN) check
	$(LICENSEI_BIN) header

.PHONY: fmt
fmt: ## Format code
	$(call print-target)
	$(GOLANGCI_LINT_BIN) run --fix

.PHONY: mod
mod: ## go mod tidy
	$(call print-target)
	go mod tidy

.PHONY: seed
seed: ## Seed OpenMeter with test data
	$(call print-target)
	$(BENTHOS_BIN) -c etc/seed/seed.yaml
	
.PHONY: deps
deps: bin/air bin/dagger bin/golangci-lint bin/licensei bin/benthos ## Install dependencies

# Dependency versions
AIR_VERSION = 1.51.0
DAGGER_VERSION = 0.9.0
GOLANGCI_VERSION = 1.53.1
LICENSEI_VERSION = 0.8.0
BENTHOS_VERSION = 4.24.0

# Dependency binaries
AIR_BIN := air
DAGGER_BIN := dagger
GOLANGCI_LINT_BIN := golangci-lint
LICENSEI_BIN := licensei
BENTHOS_BIN := benthos

# If we have a "bin" dir, use those binaries instead
ifneq ($(wildcard ./bin/.),)
	AIR_BIN := bin/$(AIR_BIN)
	DAGGER_BIN := bin/$(DAGGER_BIN)
	GOLANGCI_LINT_BIN := bin/$(GOLANGCI_LINT_BIN)
	LICENSEI_BIN := bin/$(LICENSEI_BIN)
	BENTHOS_BIN := bin/$(BENTHOS_BIN)
endif

bin/air:
	@mkdir -p bin
	curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | bash -s -- v${AIR_VERSION}

bin/dagger:
	@mkdir -p bin
	curl -sfL https://raw.githubusercontent.com/dagger/dagger/main/install.sh | bash -s -- v${DAGGER_VERSION}

bin/golangci-lint:
	@mkdir -p bin
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | bash -s -- v${GOLANGCI_VERSION}

bin/licensei:
	@mkdir -p bin
	curl -sfL https://raw.githubusercontent.com/goph/licensei/master/install.sh | bash -s -- v${LICENSEI_VERSION}

bin/benthos:
	@mkdir -p bin
	curl -Lsf https://sh.benthos.dev | bash -s -- ${BENTHOS_VERSION} ${PWD}/bin

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

# Variable outputting/exporting rules
var-%: ; @echo $($*)
varexport-%: ; @echo $*=$($*)

define print-target
    @printf "Executing target: \033[36m$@\033[0m\n"
endef
