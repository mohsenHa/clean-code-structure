ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))

lint: ## Lint project files
	which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0)
	golangci-lint run --modules-download-mode=vendor --config=$(ROOT)/.golangci.yml $(ROOT)/...


format: ## Format project files
	@which gofumpt || (go install mvdan.cc/gofumpt@latest)
	@gofumpt -l -w $(ROOT)
	@which gci || (go install github.com/daixiang0/gci@latest)
	@gci write --skip-vendor $(ROOT)
	@which golangci-lint || (go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0)
	@golangci-lint run --fix --modules-download-mode=vendor --config=$(ROOT)/.golangci.yml $(ROOT)/...

help:  ## Display this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[.a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)
