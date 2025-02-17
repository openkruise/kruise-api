CURRENT_DIR=$(shell pwd)
# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Run go vet against code
vet:
	go vet ./...

# Generate code
generate: controller-gen
	#$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."
	@hack/generate_client.sh

CONTROLLER_GEN = $(shell pwd)/bin/controller-gen
controller-gen: ## Download controller-gen locally if necessary.
ifeq ("$(shell $(CONTROLLER_GEN) --version 2> /dev/null)", "Version: v0.16.5")
else
	rm -rf $(CONTROLLER_GEN)
	$(call go-get-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen@v0.16.5)
endif

OPENAPI_GEN = $(shell pwd)/bin/openapi-gen
module=$(shell go list -f '{{.Module}}' k8s.io/kube-openapi/cmd/openapi-gen | awk '{print $$1}')
module_version=$(shell go list -m $(module) | awk '{print $$NF}' | head -1)
openapi-gen: ## Download openapi-gen locally if necessary.
ifeq ("$(shell command -v $(OPENAPI_GEN) 2> /dev/null)", "")
	$(call go-get-tool,$(OPENAPI_GEN),k8s.io/kube-openapi/cmd/openapi-gen@$(module_version))
else
	@echo "openapi-gen is already installed."
endif

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
cd $$TMP_DIR ;\
go mod init tmp ;\
echo "Downloading $(2)" ;\
GOBIN=$(PROJECT_DIR)/bin go install $(2) ;\
rm -rf $$TMP_DIR ;\
}
endef

.PHONY: gen-schema-only
gen-schema-only:
	go run cmd/gen-schema/main.go

.PHONY: gen-openapi-schema
gen-openapi-schema: gen-all-openapi
	go run cmd/gen-schema/main.go

.PHONY: gen-all-openapi
	@hack/generate_openapi.sh
