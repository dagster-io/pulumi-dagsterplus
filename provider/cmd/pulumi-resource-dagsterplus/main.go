package main

import (
	"context"
	_ "embed"

	pf "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/pf/tfbridge"

	dagsterplus "github.com/dagster-io/pulumi-dagsterplus/provider"
)

//go:embed schema.json
var schema []byte

func main() {
	meta := pf.ProviderMetadata{PackageSchema: schema}
	pf.Main(context.Background(), "dagsterplus", dagsterplus.Provider(), meta)
}
