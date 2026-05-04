PROVIDER_NAME := dagsterplus
TFGEN         := pulumi-tfgen-$(PROVIDER_NAME)
PROVIDER      := pulumi-resource-$(PROVIDER_NAME)
SCHEMA_OUT    := provider/cmd/pulumi-resource-$(PROVIDER_NAME)

.PHONY: tidy tfgen schema provider install sdk_nodejs all

tidy:
	go mod tidy

tfgen:
	go build -o $(TFGEN) ./provider/cmd/pulumi-tfgen-$(PROVIDER_NAME)

schema: tfgen
	./$(TFGEN) schema --out $(SCHEMA_OUT)

provider: schema
	go build -o $(PROVIDER) ./provider/cmd/pulumi-resource-$(PROVIDER_NAME)

install: provider
	mkdir -p ~/.pulumi/plugins/resource-$(PROVIDER_NAME)-v0.0.0
	cp $(PROVIDER) ~/.pulumi/plugins/resource-$(PROVIDER_NAME)-v0.0.0/$(PROVIDER)

sdk_python: tfgen
	./$(TFGEN) python --out sdk/python

test_provider:
	go test -v ./provider/...

test_examples: provider
	cd examples && go test -v -timeout 60m ./...

sdk_nodejs: tfgen
	./$(TFGEN) nodejs --out sdk/nodejs

all: tidy provider
