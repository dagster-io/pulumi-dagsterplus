PROVIDER_NAME := dagsterplus
TFGEN         := pulumi-tfgen-$(PROVIDER_NAME)
PROVIDER      := pulumi-resource-$(PROVIDER_NAME)
SCHEMA_OUT    := provider/cmd/pulumi-resource-$(PROVIDER_NAME)
VERSION       ?= 0.0.0
LDFLAGS       := -X github.com/dagster-io/pulumi-dagsterplus/provider.Version=$(VERSION)

.PHONY: tidy tfgen schema provider install all

tidy:
	go mod tidy

tfgen:
	go build -ldflags "$(LDFLAGS)" -o $(TFGEN) ./provider/cmd/pulumi-tfgen-$(PROVIDER_NAME)

schema: tfgen
	./$(TFGEN) schema --out $(SCHEMA_OUT)
	go run ./provider/cmd/strip-schema-languages $(SCHEMA_OUT)/schema.json

provider: schema
	go build -ldflags "$(LDFLAGS)" -o $(PROVIDER) ./provider/cmd/pulumi-resource-$(PROVIDER_NAME)

install: provider
	mkdir -p ~/.pulumi/plugins/resource-$(PROVIDER_NAME)-v0.0.0
	cp $(PROVIDER) ~/.pulumi/plugins/resource-$(PROVIDER_NAME)-v0.0.0/$(PROVIDER)

sdk_python: tfgen
	./$(TFGEN) python --out sdk/python

test_provider:
	go test -v ./provider/...

test_examples: provider
	cd examples && go test -v -timeout 60m ./...

all: tidy provider
