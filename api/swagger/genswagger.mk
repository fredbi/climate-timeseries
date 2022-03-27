TARGET=./pkg/restapi
SPEC_DIR=api/swagger
SPEC_MODELS=$(SPEC_DIR)/timeseries-models.yaml

.PHONY: generate-models
## Generate API data model from swagger spec
generate-models:
	cd $(git rev-parse --show-toplevel)
	swagger generate model -f $(SPEC_MODELS) --accept-definitions-only --target $(TARGET)
