TARGET=./pkg/restapi
CLIENT_TARGET=./clients/go
SPEC_DIR=api/swagger
HANDLERS=handlers
MODELS_PKG=github.com/fredbi/climate-timeseries/pkg/restapi/models
CLIENT_PKG=public
SPEC_MODELS=$(SPEC_DIR)/timeseries-models.yaml
SPEC_API=$(SPEC_DIR)/timeseries.yaml

.PHONY: ensure-codegen-deps
## Install API codegen tools
ensure-codegen-deps:
	go install github.com/go-swagger/go-swagger@master

.PHONY: generate-models
## Generate API data model from swagger spec
generate-models:
	cd $(git rev-parse --show-toplevel)
	swagger generate model -f $(SPEC_MODELS) \
		--accept-definitions-only \
		--target $(TARGET)

.PHONY: generate-handlers
## Generate API handlers
generate-handlers:
	cd $(git rev-parse --show-toplevel)
	swagger generate server -f $(SPEC_API) \
		--default-scheme https \
		--name 'climate change API' \
		--strict-responders \
		--flag-strategy pflag \
		--server-package $(HANDLERS) \
		--exclude-main \
		--skip-models \
		--existing-models $(MODELS_PKG) \
		--target $(TARGET)

.PHONY: generate-go-client
## Generate go client
generate-go-client:
	cd $(git rev-parse --show-toplevel)
	swagger generate client -f $(SPEC_API) \
		--default-scheme https \
		--name 'climate change API' \
		--strict-responders \
		--client-package $(CLIENT_PKG) \
		--skip-models \
		--existing-models $(MODELS_PKG) \
		--target $(CLIENT_TARGET)
