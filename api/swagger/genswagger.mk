REPODIR=$(shell git rev-parse --show-toplevel)

CLIENT_BASE=./clients

CLIENT_GO_TARGET=$(CLIENT_BASE)/go
CLIENT_PY_TARGET=$(CLIENT_BASE)/python
CLIENT_R_TARGET=$(CLIENT_BASE)/R
CLIENT_TSNODE_TARGET=$(CLIENT_BASE)/typescript-node

TARGET=./pkg/restapi
HANDLERS=handlers

MODELS_PKG=github.com/fredbi/climate-timeseries/pkg/restapi/models
CLIENT_PKG=climate

SPEC_DIR=api/swagger
SPEC_MODELS=$(SPEC_DIR)/timeseries-models.yaml
SPEC_API=$(SPEC_DIR)/timeseries.yaml

.PHONY: ensure-codegen-deps
## Install API codegen tools (pin and follow closely version, to avoid spurious incontrolled changes and so many quirks with codegen)
ensure-codegen-deps:
	go install github.com/go-swagger/go-swagger/cmd/swagger@a543a92947790ebcf87dd555a44deb78fba6b932
	docker pull swaggerapi/swagger-codegen-cli:2.4.26

.PHONY: generate-models
## Generate API data model from swagger spec
generate-models:
	cd $(REPODIR) || exit 1
	rm -rf $(TARGET)/models
	swagger generate model -f $(SPEC_MODELS) \
		--accept-definitions-only \
		--target $(TARGET)

.PHONY: generate-handlers
## Generate API handlers
generate-handlers:
	cd $(REPODIR) || exit 1
	rm -rf $(TARGET)/$(HANDLERS)
	swagger generate server -f $(SPEC_API) \
		--default-scheme https \
		--name 'climate change API' \
		--strict-responders \
		--flag-strategy pflag \
		--server-package $(HANDLERS) \
		--exclude-main \
		--skip-models \
		--existing-models $(MODELS_PKG) \
		--principal-is-interface \
		--principal github.com/fredbi/climate-timeseries/pkg/auth.Principal \
		--target $(TARGET)

.PHONY: generate-go-client
## Generate go client
generate-go-client:
	cd $(REPODIR) || exit 1
	rm -rf $(CLIENT_GO_TARGET)/$(CLIENT_PKG)
	swagger generate client -f $(SPEC_API) \
		--default-scheme https \
		--name 'climate change API' \
		--strict-responders \
		--client-package $(CLIENT_PKG) \
		--skip-models \
		--existing-models $(MODELS_PKG) \
		--principal-is-interface \
		--principal github.com/fredbi/climate-timeseries/pkg/auth.Principal \
		--target $(CLIENT_GO_TARGET)

SWAGGERCODEGEN=docker run --rm -v ${REPODIR}/${SPEC_DIR}:/local swaggerapi/swagger-codegen-cli generate -i /local/$(shell basename ${SPEC_API})

.PHONY: generate-python-client
## Generate a python client for this API
generate-python-client:
generate-python-client: export 
	cd $(REPODIR) || exit 1
	rm -rf $(CLIENT_PY_TARGET)/$(CLIENT_PKG)
	$(SWAGGERCODEGEN) -l python -o /local/$(CLIENT_PY_TARGET)/$(CLIENT_PKG)

.PHONY: generate-r-client
## Generate a R client for this API
generate-r-client:
	cd $(REPODIR) || exit 1
	rm -rf $(CLIENT_R_TARGET)/$(CLIENT_PKG)
	$(SWAGGERCODEGEN) -l r -o /local/$(CLIENT_R_TARGET)/$(CLIENT_PKG)

.PHONY: generate-ts-node-client
## Generate a typescript-node client for this API
generate-ts-node-client:
	cd $(REPODIR) || exit 1
	rm -rf $(CLIENT_TSNODE_TARGET)/$(CLIENT_PKG)
	$(SWAGGERCODEGEN) -l typescript-node -o /local/$(CLIENT_TSNODE_TARGET)/$(CLIENT_PKG)
